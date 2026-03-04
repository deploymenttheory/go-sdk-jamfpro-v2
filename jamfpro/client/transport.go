package client

import (
	"context"
	"fmt"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// Transport is the HTTP transport layer for the Jamf Pro API.
// It wraps a resty.Client with Jamf-specific behaviour: bearer token auth,
// idempotent-only retries with exponential backoff, adaptive response-time
// throttling, sticky-session cookie jar, optional concurrency limiting,
// and structured logging.
type Transport struct {
	client        *resty.Client
	logger        *zap.Logger
	authConfig    *AuthConfig
	holder        *tokenHolder
	BaseURL       string
	globalHeaders map[string]string
	userAgent     string

	// Optional throttles — nil / zero means disabled.
	sem                *semaphore
	requestDelay       time.Duration
	totalRetryDuration time.Duration

	// responseTracker measures per-request latency and derives an adaptive
	// inter-request delay when the server begins responding slowly.
	responseTracker *responseTimeTracker
}

// NewTransport creates and fully configures a Jamf Pro API transport.
//
// Behaviour applied at construction time (resty native where possible):
//   - Bearer token authentication with automatic refresh
//   - Idempotent-only retry (GET/PUT/DELETE) with exponential backoff
//   - Sticky-session cookie jar (handles jpro-ingress, APBALANCEID, JSESSIONID)
//   - Deprecation header warning logged on every response
//   - Adaptive inter-request delay derived from response-time EMA tracking
//
// Jamf Pro does not emit rate-limit HTTP headers. Throttling is inferred
// from observed response times per Jamf scalability best practices.
func NewTransport(authConfig *AuthConfig, options ...ClientOption) (*Transport, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	if authConfig == nil {
		return nil, fmt.Errorf("auth config is required")
	}
	baseURL := authConfig.InstanceDomain
	if baseURL == "" {
		return nil, fmt.Errorf("instance domain is required")
	}
	baseURL = trimTrailingSlash(baseURL)
	userAgent := fmt.Sprintf("%s/%s", UserAgentBase, Version)

	// Resty creates a cookie jar by default, which enables sticky sessions automatically.
	// Jamf Cloud sets jpro-ingress / APBALANCEID / JSESSIONID in Set-Cookie
	// headers; resty resends them on subsequent requests without extra logic.
	// See: https://developer.jamf.com/jamf-pro/docs/sticky-sessions-for-jamf-cloud
	restyClient := resty.New()
	restyClient.SetBaseURL(baseURL)
	restyClient.SetTimeout(DefaultTimeout)
	restyClient.SetRetryCount(MaxRetries)
	restyClient.SetRetryWaitTime(RetryWaitTime)
	restyClient.SetRetryMaxWaitTime(RetryMaxWaitTime)
	restyClient.SetHeader("User-Agent", userAgent)

	// Only retry idempotent methods on transient server errors.
	// Resty's built-in exponential backoff handles the wait between retries.
	// See: https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-scalability-best-practices
	restyClient.AddRetryConditions(retryCondition)

	transport := &Transport{
		client:          restyClient,
		logger:          logger,
		authConfig:      authConfig,
		BaseURL:         baseURL,
		globalHeaders:   make(map[string]string),
		userAgent:       userAgent,
		responseTracker: newResponseTimeTracker(),
	}

	// Apply caller-supplied options before auth so WithLogger takes effect first.
	for _, opt := range options {
		if err := opt(transport); err != nil {
			return nil, fmt.Errorf("failed to apply client option: %w", err)
		}
	}

	// Log deprecated endpoint warnings and cookie usage via resty response middleware.
	restyClient.AddResponseMiddleware(func(_ *resty.Client, r *resty.Response) error {
		if dep := r.Header().Get("Deprecation"); dep != "" {
			transport.logger.Warn("Jamf Pro API endpoint is deprecated",
				zap.String("endpoint", r.Request.URL),
				zap.String("deprecation", dep),
				zap.String("sunset", r.Header().Get("Sunset")),
			)
		}

		if r.Request != nil && r.Request.Header != nil {
			cookieHeader := r.Request.Header.Get("Cookie")
			transport.logger.Info("Request cookie status",
				zap.String("method", r.Request.Method),
				zap.String("path", r.Request.URL),
				zap.String("cookie_sent", cookieHeader),
				zap.Bool("has_cookie", cookieHeader != ""),
			)
		}

		return nil
	})

	// Wire authentication — returns tokenHolder for InvalidateToken/KeepAliveToken.
	// The OAuth token request (first request) captures the sticky session cookie automatically.
	// IMPORTANT: This must happen AFTER SetBaseURL so the cookie jar associates the cookie
	// with the correct domain for all subsequent requests.
	holder, err := SetupAuthentication(restyClient, authConfig, transport.logger)
	if err != nil {
		return nil, fmt.Errorf("failed to setup authentication: %w", err)
	}
	transport.holder = holder

	transport.logger.Info("Jamf Pro API transport created",
		zap.String("base_url", transport.BaseURL),
		zap.String("auth_method", authConfig.AuthMethod),
	)
	return transport, nil
}

func trimTrailingSlash(s string) string {
	if len(s) > 0 && s[len(s)-1] == '/' {
		return s[:len(s)-1]
	}
	return s
}

// GetHTTPClient returns the underlying resty client for advanced use.
func (t *Transport) GetHTTPClient() *resty.Client {
	return t.client
}

// GetLogger returns the configured logger.
func (t *Transport) GetLogger() *zap.Logger {
	return t.logger
}

// RSQLBuilder returns a new RSQL filter expression builder.
// Pass the Build() result as rsqlQuery["filter"] to filter endpoint results.
func (t *Transport) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return NewRSQLFilterBuilder()
}

// InvalidateToken revokes the current bearer token at the Jamf Pro API and
// clears the local cache. The next request triggers a full re-authentication.
func (t *Transport) InvalidateToken() error {
	return t.holder.invalidate()
}

// KeepAliveToken extends the current bearer token lifetime without re-auth.
// Use before long-running operations to prevent mid-operation token expiry.
func (t *Transport) KeepAliveToken() error {
	return t.holder.keepAlive()
}

// executeRequest is the central request executor used by all HTTP verb methods.
// It applies the concurrency semaphore, total-retry deadline, mandatory
// per-request delay, and adaptive response-time throttling.
func (t *Transport) executeRequest(req *resty.Request, method, path string) (*resty.Response, error) {
	ctx := req.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// Wrap in a deadline for the total allowed retry window if configured and
	// the caller has not already set a more restrictive deadline.
	if t.totalRetryDuration > 0 {
		if _, hasDeadline := ctx.Deadline(); !hasDeadline {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, t.totalRetryDuration)
			defer cancel()
			req.SetContext(ctx)
		}
	}

	// Acquire concurrency slot — blocks until available or context cancelled.
	if t.sem != nil {
		if err := t.sem.acquire(ctx); err != nil {
			return nil, fmt.Errorf("concurrency limit: %w", err)
		}
		defer t.sem.release()
	}

	t.logger.Debug("Executing API request", zap.String("method", method), zap.String("path", path))

	resp, execErr := req.Execute(method, path)

	if execErr != nil {
		t.logger.Error("Request failed",
			zap.String("method", method),
			zap.String("path", path),
			zap.Error(execErr),
		)
		return resp, fmt.Errorf("request failed: %w", execErr)
	}

	if err := t.validateResponse(resp, method, path); err != nil {
		return resp, err
	}

	if resp.IsError() {
		return resp, ParseErrorResponse(
			[]byte(resp.String()),
			resp.StatusCode(),
			resp.Status(),
			method,
			path,
			t.logger,
		)
	}

	duration := resp.Duration()

	var stickySessionCookie string
	var allCookies []string
	if resp.RawResponse != nil {
		for _, cookie := range resp.Cookies() {
			allCookies = append(allCookies, cookie.Name)
			if cookie.Name == "jpro-ingress" || cookie.Name == "APBALANCEID" || cookie.Name == "JSESSIONID" {
				stickySessionCookie = fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
			}
		}
	}

	t.logger.Info("Request completed",
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status_code", resp.StatusCode()),
		zap.Duration("duration", duration),
		zap.String("sticky_session_cookie", stickySessionCookie),
		zap.Strings("all_response_cookies", allCookies),
	)

	// Mandatory fixed delay (user-configured for bulk operations).
	if t.requestDelay > 0 {
		time.Sleep(t.requestDelay)
	}

	// Adaptive delay: when the server is responding more slowly than its own
	// EMA baseline, pause proportionally before the next request.
	// This implements Jamf's guidance to "measure response times and dynamically
	// adjust time between requests accordingly."
	if adaptive := t.responseTracker.record(duration); adaptive > 0 {
		t.logger.Debug("Adaptive delay applied due to elevated response time",
			zap.Duration("response_time", duration),
			zap.Duration("adaptive_delay", adaptive),
		)
		time.Sleep(adaptive)
	}

	return resp, nil
}
