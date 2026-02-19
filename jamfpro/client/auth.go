package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"resty.dev/v3"
)

// AuthMethod constants for the Jamf Pro authentication methods.
const (
	AuthMethodOAuth2 = "oauth2"
	AuthMethodBasic  = "basic"
)

// AuthConfig holds authentication configuration for the Jamf Pro API.
//
// Two authentication flows are supported:
//   - OAuth2 client credentials (recommended): POST /api/v1/oauth/token
//   - Basic auth to bearer token exchange:     POST /api/v1/auth/token
//
// See: https://developer.jamf.com/jamf-pro/docs/classic-api-authentication-changes
type AuthConfig struct {
	// InstanceDomain is the Jamf Pro instance base URL (e.g. https://example.jamfcloud.com).
	InstanceDomain string

	// AuthMethod selects the authentication flow: "oauth2" or "basic".
	AuthMethod string

	// OAuth2 credentials (required when AuthMethod == "oauth2").
	ClientID     string
	ClientSecret string

	// Basic auth credentials (required when AuthMethod == "basic").
	Username string
	Password string

	// TokenRefreshBufferPeriod is how far before expiry to proactively refresh
	// the token. Defaults to 5 minutes if zero.
	TokenRefreshBufferPeriod time.Duration

	// HideSensitiveData suppresses bearer token values in log output.
	// Enable in production to prevent tokens from appearing in log files.
	HideSensitiveData bool
}

// Validate checks the auth configuration for required fields.
func (a *AuthConfig) Validate() error {
	if a.InstanceDomain == "" {
		return fmt.Errorf("instance domain is required")
	}
	if a.AuthMethod != AuthMethodOAuth2 && a.AuthMethod != AuthMethodBasic {
		return fmt.Errorf("auth method must be %q or %q", AuthMethodOAuth2, AuthMethodBasic)
	}
	if a.AuthMethod == AuthMethodOAuth2 {
		if a.ClientID == "" || a.ClientSecret == "" {
			return fmt.Errorf("client_id and client_secret are required for oauth2")
		}
	}
	if a.AuthMethod == AuthMethodBasic {
		if a.Username == "" || a.Password == "" {
			return fmt.Errorf("username and password are required for basic auth")
		}
	}
	return nil
}

// tokenHolder holds a bearer token and refreshes it automatically before expiry.
// All exported methods are safe for concurrent use.
type tokenHolder struct {
	mu                sync.Mutex
	token             string
	expiry            time.Time
	buffer            time.Duration
	auth              *AuthConfig
	logger            *zap.Logger
	http              *http.Client
	baseURL           string
	hideSensitiveData bool
	fetchFn           func() (string, time.Time, error)
}

// logToken returns the token string for logging, redacted when HideSensitiveData is set.
func (h *tokenHolder) logToken() string {
	if h.hideSensitiveData {
		return "[REDACTED]"
	}
	return h.token
}

// getToken returns the current bearer token, refreshing it when expired or
// within the buffer period before expiry.
func (h *tokenHolder) getToken() (string, error) {
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
func (h *tokenHolder) invalidate() error {
	h.mu.Lock()
	currentToken := h.token
	h.mu.Unlock()

	if currentToken == "" {
		return nil
	}

	endpoint := strings.TrimSuffix(h.baseURL, "/") + invalidateTokenEndpoint
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return fmt.Errorf("invalidate token: build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+currentToken)

	resp, err := h.http.Do(req)
	if err != nil {
		return fmt.Errorf("invalidate token: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("invalidate token: unexpected status %d: %s", resp.StatusCode, string(body))
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
func (h *tokenHolder) keepAlive() error {
	h.mu.Lock()
	currentToken := h.token
	h.mu.Unlock()

	if currentToken == "" {
		return fmt.Errorf("keep-alive: no active token")
	}

	endpoint := strings.TrimSuffix(h.baseURL, "/") + keepAliveTokenEndpoint
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return fmt.Errorf("keep-alive: build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+currentToken)

	resp, err := h.http.Do(req)
	if err != nil {
		return fmt.Errorf("keep-alive: request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("keep-alive: read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("keep-alive: unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var out struct {
		Token   string    `json:"token"`
		Expires time.Time `json:"expires"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		return fmt.Errorf("keep-alive: decode response: %w", err)
	}

	h.mu.Lock()
	h.token = out.Token
	h.expiry = out.Expires
	h.mu.Unlock()

	h.logger.Info("Bearer token keep-alive successful", zap.Time("new_expiry", out.Expires))
	return nil
}

func (h *tokenHolder) fetchOAuth2() (string, time.Time, error) {
	endpoint := strings.TrimSuffix(h.baseURL, "/") + oAuthTokenEndpoint

	data := url.Values{}
	data.Set("client_id", h.auth.ClientID)
	data.Set("client_secret", h.auth.ClientSecret)
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", time.Time{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := h.http.Do(req)
	if err != nil {
		return "", time.Time{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", time.Time{}, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", time.Time{}, fmt.Errorf("oauth2 token request failed: %d %s", resp.StatusCode, string(body))
	}

	var out struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		return "", time.Time{}, fmt.Errorf("decode oauth2 response: %w", err)
	}
	if out.AccessToken == "" {
		return "", time.Time{}, fmt.Errorf("empty access_token in oauth2 response")
	}

	expiry := time.Now().Add(time.Duration(out.ExpiresIn) * time.Second)
	h.logger.Info("OAuth2 bearer token obtained",
		zap.Time("expiry", expiry),
		zap.String("token", h.logToken()),
	)
	return out.AccessToken, expiry, nil
}

func (h *tokenHolder) fetchBasic() (string, time.Time, error) {
	endpoint := strings.TrimSuffix(h.baseURL, "/") + bearerTokenEndpoint

	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return "", time.Time{}, err
	}
	req.SetBasicAuth(h.auth.Username, h.auth.Password)

	resp, err := h.http.Do(req)
	if err != nil {
		return "", time.Time{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", time.Time{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", time.Time{}, fmt.Errorf("basic auth token request failed: %d %s", resp.StatusCode, string(body))
	}

	var out struct {
		Token   string    `json:"token"`
		Expires time.Time `json:"expires"`
	}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&out); err != nil {
		return "", time.Time{}, fmt.Errorf("decode basic auth response: %w", err)
	}
	if out.Token == "" {
		return "", time.Time{}, fmt.Errorf("empty token in basic auth response")
	}

	h.logger.Info("Basic auth bearer token obtained",
		zap.Time("expiry", out.Expires),
		zap.String("token", h.logToken()),
	)
	return out.Token, out.Expires, nil
}

// SetupAuthentication configures the resty client with Jamf Pro bearer token
// authentication. A token is fetched immediately to surface misconfiguration
// at startup. Subsequent requests refresh the token automatically via middleware.
//
// Returns the tokenHolder so the Transport can expose InvalidateToken and
// KeepAliveToken.
//
// See: https://developer.jamf.com/jamf-pro/docs/classic-api-authentication-changes
func SetupAuthentication(restyClient *resty.Client, authConfig *AuthConfig, logger *zap.Logger) (*tokenHolder, error) {
	if err := authConfig.Validate(); err != nil {
		return nil, fmt.Errorf("authentication configuration invalid: %w", err)
	}

	baseURL := strings.TrimSuffix(authConfig.InstanceDomain, "/")
	buffer := authConfig.TokenRefreshBufferPeriod
	if buffer == 0 {
		buffer = 5 * time.Minute
	}

	holder := &tokenHolder{
		auth:              authConfig,
		logger:            logger,
		http:              &http.Client{Timeout: 30 * time.Second},
		baseURL:           baseURL,
		buffer:            buffer,
		hideSensitiveData: authConfig.HideSensitiveData,
	}

	switch authConfig.AuthMethod {
	case AuthMethodOAuth2:
		holder.fetchFn = holder.fetchOAuth2
	case AuthMethodBasic:
		holder.fetchFn = holder.fetchBasic
	default:
		return nil, fmt.Errorf("unsupported auth method: %q", authConfig.AuthMethod)
	}

	if _, err := holder.getToken(); err != nil {
		return nil, fmt.Errorf("initial token fetch failed: %w", err)
	}

	restyClient.AddRequestMiddleware(func(_ *resty.Client, r *resty.Request) error {
		token, err := holder.getToken()
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
	return holder, nil
}
