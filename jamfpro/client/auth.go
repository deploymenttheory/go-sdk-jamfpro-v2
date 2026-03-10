package client

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/config"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// bearerTokenManager manages the full lifecycle of a Jamf Pro bearer token:
// automatic refresh before expiry, keep-alive extension, and revocation.
// All methods are safe for concurrent use.
//
// authClient is a dedicated resty client used exclusively for auth operations
// (token fetch, keep-alive, invalidate). It is intentionally separate from
// the main Transport client to prevent a deadlock: the Transport client carries
// an auth middleware that calls getToken(), which holds the mutex and calls
// fetchFn(). If fetchFn() made requests through the same client, the middleware
// would fire again and attempt to re-acquire the mutex — deadlocking.
type bearerTokenManager struct {
	mu                sync.Mutex
	token             string
	expiry            time.Time
	buffer            time.Duration
	auth              *config.AuthConfig
	logger            *zap.Logger
	authClient        *resty.Client
	baseURL           string
	hideSensitiveData bool
	fetchFn           func() (string, time.Time, error)
}

// logToken returns the token string for logging, redacted when HideSensitiveData is set.
func (h *bearerTokenManager) logToken() string {
	if h.hideSensitiveData {
		return "[REDACTED]"
	}
	return h.token
}

// getToken returns the current bearer token, refreshing it when expired or
// within the buffer period before expiry.
func (h *bearerTokenManager) getToken() (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	now := time.Now()
	if h.token != "" && now.Before(h.expiry) && now.Add(h.buffer).Before(h.expiry) {
		return h.token, nil
	}

	token, expiry, err := h.fetchFn()
	if err != nil {
		return "", err
	}
	h.token = token
	h.expiry = expiry
	return token, nil
}

// invalidate revokes the current bearer token at the Jamf Pro API and clears
// the local cache so the next request forces a full re-authentication.
func (h *bearerTokenManager) invalidate() error {
	h.mu.Lock()
	currentToken := h.token
	h.mu.Unlock()

	if currentToken == "" {
		return nil
	}

	endpoint := strings.TrimSuffix(h.baseURL, "/") + constants.EndpointInvalidateToken

	resp, err := h.authClient.R().
		SetAuthToken(currentToken).
		Post(endpoint)

	if err != nil {
		return fmt.Errorf("invalidate token: request failed: %w", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		return fmt.Errorf("invalidate token: unexpected status %d: %s", resp.StatusCode(), resp.String())
	}

	h.mu.Lock()
	h.token = ""
	h.expiry = time.Time{}
	h.mu.Unlock()

	h.logger.Info("Bearer token invalidated")
	return nil
}

// keepAlive extends the current token lifetime via the Jamf Pro API and updates
// the cached token and expiry time from the response.
func (h *bearerTokenManager) keepAlive() error {
	h.mu.Lock()
	currentToken := h.token
	h.mu.Unlock()

	if currentToken == "" {
		return fmt.Errorf("keep-alive: no active token")
	}

	endpoint := strings.TrimSuffix(h.baseURL, "/") + constants.EndpointKeepAliveToken

	var result struct {
		Token   string    `json:"token"`
		Expires time.Time `json:"expires"`
	}

	resp, err := h.authClient.R().
		SetAuthToken(currentToken).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return fmt.Errorf("keep-alive: request failed: %w", err)
	}

	if resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		return fmt.Errorf("keep-alive: unexpected status %d: %s", resp.StatusCode(), resp.String())
	}

	h.mu.Lock()
	h.token = result.Token
	h.expiry = result.Expires
	h.mu.Unlock()

	h.logger.Info("Bearer token keep-alive successful", zap.Time("new_expiry", result.Expires))
	return nil
}

func (h *bearerTokenManager) fetchOAuth2() (string, time.Time, error) {
	endpoint := strings.TrimSuffix(h.baseURL, "/") + constants.EndpointOAuthToken

	data := url.Values{}
	data.Set("client_id", h.auth.ClientID)
	data.Set("client_secret", h.auth.ClientSecret)
	data.Set("grant_type", "client_credentials")

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	resp, err := h.authClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormDataFromValues(data).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return "", time.Time{}, err
	}

	if resp.StatusCode() < 200 || resp.StatusCode() > 299 {
		return "", time.Time{}, fmt.Errorf("oauth2 token request failed: %d %s", resp.StatusCode(), resp.String())
	}

	if result.AccessToken == "" {
		return "", time.Time{}, fmt.Errorf("empty access_token in oauth2 response")
	}

	expiry := time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)
	
	var stickySessionCookie string
	var allCookies []string
	var cookieDetails []string
	for _, cookie := range resp.Cookies() {
		allCookies = append(allCookies, cookie.Name)
		cookieDetails = append(cookieDetails, fmt.Sprintf("%s=%s (Path:%s, Domain:%s, Secure:%v, HttpOnly:%v)", 
			cookie.Name, cookie.Value, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly))
		if cookie.Name == "jpro-ingress" || cookie.Name == "APBALANCEID" || cookie.Name == "JSESSIONID" {
			stickySessionCookie = fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
		}
	}
	
	h.logger.Info("OAuth2 bearer token obtained",
		zap.Time("expiry", expiry),
		zap.String("token", h.logToken()),
		zap.String("sticky_session_cookie", stickySessionCookie),
		zap.Strings("all_cookies_from_auth", allCookies),
		zap.Strings("cookie_details", cookieDetails),
	)
	return result.AccessToken, expiry, nil
}

func (h *bearerTokenManager) fetchBasic() (string, time.Time, error) {
	endpoint := strings.TrimSuffix(h.baseURL, "/") + constants.EndpointBearerToken

	var result struct {
		Token   string    `json:"token"`
		Expires time.Time `json:"expires"`
	}

	resp, err := h.authClient.R().
		SetBasicAuth(h.auth.Username, h.auth.Password).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return "", time.Time{}, err
	}

	if resp.StatusCode() != 200 {
		return "", time.Time{}, fmt.Errorf("basic auth token request failed: %d %s", resp.StatusCode(), resp.String())
	}

	if result.Token == "" {
		return "", time.Time{}, fmt.Errorf("empty token in basic auth response")
	}

	var stickySessionCookie string
	var allCookies []string
	var cookieDetails []string
	for _, cookie := range resp.Cookies() {
		allCookies = append(allCookies, cookie.Name)
		cookieDetails = append(cookieDetails, fmt.Sprintf("%s=%s (Path:%s, Domain:%s, Secure:%v, HttpOnly:%v)", 
			cookie.Name, cookie.Value, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly))
		if cookie.Name == "jpro-ingress" || cookie.Name == "APBALANCEID" || cookie.Name == "JSESSIONID" {
			stickySessionCookie = fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
		}
	}

	h.logger.Info("Basic auth bearer token obtained",
		zap.Time("expiry", result.Expires),
		zap.String("token", h.logToken()),
		zap.String("sticky_session_cookie", stickySessionCookie),
		zap.Strings("all_cookies_from_auth", allCookies),
		zap.Strings("cookie_details", cookieDetails),
	)
	return result.Token, result.Expires, nil
}

// SetupAuthentication configures the resty client with Jamf Pro bearer token
// authentication. A token is fetched immediately to surface misconfiguration
// at startup. Subsequent requests refresh the token automatically via middleware.
//
// A dedicated authClient is created for token operations (fetch, keep-alive,
// invalidate). It carries the same TLS/proxy/transport settings as the main
// client but has no auth middleware, preventing the deadlock that would occur
// if token refresh fired the middleware re-entrantly while holding the mutex.
// The passed restyClient is only used to install the request middleware.
//
// Returns the bearerTokenManager so the Transport can expose InvalidateToken
// and KeepAliveToken.
//
// See: https://developer.jamf.com/jamf-pro/docs/classic-api-authentication-changes
func SetupAuthentication(restyClient *resty.Client, authConfig *config.AuthConfig, logger *zap.Logger, settings *TransportSettings) (*bearerTokenManager, error) {
	if err := authConfig.Validate(); err != nil {
		return nil, fmt.Errorf("authentication configuration invalid: %w", err)
	}

	baseURL := strings.TrimSuffix(authConfig.InstanceDomain, "/")
	buffer := authConfig.TokenRefreshBufferPeriod
	if buffer == 0 {
		buffer = 5 * time.Minute
	}

	// Dedicated client for auth requests only — no middleware, no retry logic.
	// Mirrors the TLS/proxy/transport settings of the main client so that
	// environments requiring custom TLS (e.g. self-signed certs, corporate
	// proxies) work correctly for auth calls too.
	authClient := resty.New()
	if settings != nil {
		if settings.InsecureSkipVerify {
			authClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //nolint:gosec
		} else if settings.TLSClientConfig != nil {
			authClient.SetTLSClientConfig(settings.TLSClientConfig)
		}
		if settings.ProxyURL != "" {
			authClient.SetProxy(settings.ProxyURL)
		}
		if settings.HTTPTransport != nil {
			authClient.SetTransport(settings.HTTPTransport)
		}
	}

	tokenManager := &bearerTokenManager{
		auth:              authConfig,
		logger:            logger,
		authClient:        authClient,
		baseURL:           baseURL,
		buffer:            buffer,
		hideSensitiveData: authConfig.HideSensitiveData,
	}

	switch authConfig.AuthMethod {
	case constants.AuthMethodOAuth2:
		tokenManager.fetchFn = tokenManager.fetchOAuth2
	case constants.AuthMethodBasic:
		tokenManager.fetchFn = tokenManager.fetchBasic
	default:
		return nil, fmt.Errorf("unsupported auth method: %q", authConfig.AuthMethod)
	}

	if _, err := tokenManager.getToken(); err != nil {
		return nil, fmt.Errorf("initial token fetch failed: %w", err)
	}

	restyClient.AddRequestMiddleware(func(_ *resty.Client, r *resty.Request) error {
		token, err := tokenManager.getToken()
		if err != nil {
			return err
		}
		r.SetAuthToken(token)
		return nil
	})

	logger.Info("Jamf Pro API authentication configured",
		zap.String("auth_method", authConfig.AuthMethod),
		zap.String("instance", baseURL),
	)
	return tokenManager, nil
}
