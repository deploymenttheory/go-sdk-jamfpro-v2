package client

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthConfigFromEnv(t *testing.T) {
	clear := setEnvMap(map[string]string{
		"INSTANCE_DOMAIN": "https://example.jamfcloud.com",
		"AUTH_METHOD":     AuthMethodOAuth2,
		"CLIENT_ID":       "cid",
		"CLIENT_SECRET":   "secret",
	})
	defer clear()

	cfg := AuthConfigFromEnv()
	require.NotNil(t, cfg)
	assert.Equal(t, "https://example.jamfcloud.com", cfg.InstanceDomain)
	assert.Equal(t, AuthMethodOAuth2, cfg.AuthMethod)
	assert.Equal(t, "cid", cfg.ClientID)
	assert.Equal(t, "secret", cfg.ClientSecret)
	assert.Equal(t, 300*time.Second, cfg.TokenRefreshBufferPeriod)
}

func TestAuthConfigFromEnv_Defaults(t *testing.T) {
	clear := setEnvMap(map[string]string{
		"INSTANCE_DOMAIN":       "https://x.com",
		"AUTH_METHOD":           "basic",
		"BASIC_AUTH_USERNAME":   "u",
		"BASIC_AUTH_PASSWORD":   "p",
	})
	defer clear()

	cfg := AuthConfigFromEnv()
	require.NotNil(t, cfg)
	assert.Equal(t, 300*time.Second, cfg.TokenRefreshBufferPeriod)
	assert.False(t, cfg.HideSensitiveData)
}

func TestAuthConfigFromEnv_OptionalInt(t *testing.T) {
	clear := setEnvMap(map[string]string{
		"INSTANCE_DOMAIN":                 "https://x.com",
		"AUTH_METHOD":                     "basic",
		"BASIC_AUTH_USERNAME":             "u",
		"BASIC_AUTH_PASSWORD":             "p",
		"TOKEN_REFRESH_BUFFER_SECONDS":    "60",
	})
	defer clear()

	cfg := AuthConfigFromEnv()
	require.NotNil(t, cfg)
	assert.Equal(t, 60*time.Second, cfg.TokenRefreshBufferPeriod)
}

func TestAuthConfigFromEnv_OptionalBool(t *testing.T) {
	clear := setEnvMap(map[string]string{
		"INSTANCE_DOMAIN":      "https://x.com",
		"AUTH_METHOD":          "basic",
		"BASIC_AUTH_USERNAME":  "u",
		"BASIC_AUTH_PASSWORD":  "p",
		"HIDE_SENSITIVE_DATA":  "true",
	})
	defer clear()

	cfg := AuthConfigFromEnv()
	require.NotNil(t, cfg)
	assert.True(t, cfg.HideSensitiveData)
}

func setEnvMap(m map[string]string) (clear func()) {
	orig := make(map[string]string)
	for k, v := range m {
		orig[k] = os.Getenv(k)
		_ = os.Setenv(k, v)
	}
	return func() {
		for k, v := range orig {
			if v == "" {
				_ = os.Unsetenv(k)
			} else {
				_ = os.Setenv(k, v)
			}
		}
	}
}
