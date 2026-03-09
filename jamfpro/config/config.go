package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
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

// authConfigFile is the JSON shape for LoadAuthConfigFromFile.
type authConfigFile struct {
	InstanceDomain            string `json:"instance_domain"`
	AuthMethod                string `json:"auth_method"`
	ClientID                  string `json:"client_id"`
	ClientSecret              string `json:"client_secret"`
	Username                  string `json:"basic_auth_username"`
	Password                  string `json:"basic_auth_password"`
	TokenRefreshBufferSeconds int    `json:"token_refresh_buffer_period_seconds"`
	HideSensitiveData         bool   `json:"hide_sensitive_data"`
}

// LoadAuthConfigFromFile loads AuthConfig from a JSON file.
// Expected keys: instance_domain, auth_method; for oauth2: client_id, client_secret; for basic: basic_auth_username, basic_auth_password.
// Optional: token_refresh_buffer_period_seconds (default 300).
func LoadAuthConfigFromFile(path string) (*AuthConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config file: %w", err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}
	var c authConfigFile
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}
	buffer := time.Duration(c.TokenRefreshBufferSeconds) * time.Second
	if buffer == 0 {
		buffer = 5 * time.Minute
	}
	return &AuthConfig{
		InstanceDomain:           c.InstanceDomain,
		AuthMethod:               c.AuthMethod,
		ClientID:                 c.ClientID,
		ClientSecret:             c.ClientSecret,
		Username:                 c.Username,
		Password:                 c.Password,
		TokenRefreshBufferPeriod: buffer,
		HideSensitiveData:        c.HideSensitiveData,
	}, nil
}

// AuthConfigFromEnv builds AuthConfig from environment variables.
// Required: INSTANCE_DOMAIN, AUTH_METHOD; for oauth2: CLIENT_ID, CLIENT_SECRET; for basic: BASIC_AUTH_USERNAME, BASIC_AUTH_PASSWORD.
// Optional: TOKEN_REFRESH_BUFFER_SECONDS (default 300), HIDE_SENSITIVE_DATA (default false).
func AuthConfigFromEnv() *AuthConfig {
	bufferSec := getEnvAsInt("TOKEN_REFRESH_BUFFER_SECONDS", 300)
	return &AuthConfig{
		InstanceDomain:           getEnv("INSTANCE_DOMAIN", ""),
		AuthMethod:               getEnv("AUTH_METHOD", ""),
		ClientID:                 getEnv("CLIENT_ID", ""),
		ClientSecret:             getEnv("CLIENT_SECRET", ""),
		Username:                 getEnv("BASIC_AUTH_USERNAME", ""),
		Password:                 getEnv("BASIC_AUTH_PASSWORD", ""),
		TokenRefreshBufferPeriod: time.Duration(bufferSec) * time.Second,
		HideSensitiveData:        getEnvAsBool("HIDE_SENSITIVE_DATA", false),
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func getEnvAsInt(key string, def int) int {
	v := getEnv(key, "")
	if v == "" {
		return def
	}
	i, _ := strconv.Atoi(v)
	return i
}

func getEnvAsBool(key string, def bool) bool {
	v := getEnv(key, "")
	if v == "" {
		return def
	}
	b, _ := strconv.ParseBool(v)
	return b
}
