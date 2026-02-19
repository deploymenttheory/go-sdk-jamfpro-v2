package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"resty.dev/v3"
)

func TestAuthConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *AuthConfig
		wantErr string
	}{
		{"nil", nil, ""},
		{"missing instance", &AuthConfig{InstanceDomain: "", AuthMethod: AuthMethodOAuth2}, "instance domain"},
		{"invalid method", &AuthConfig{InstanceDomain: "https://x.com", AuthMethod: "invalid"}, "auth method must"},
		{"oauth2 missing client_id", &AuthConfig{InstanceDomain: "https://x.com", AuthMethod: AuthMethodOAuth2, ClientSecret: "s"}, "client_id and client_secret"},
		{"oauth2 missing client_secret", &AuthConfig{InstanceDomain: "https://x.com", AuthMethod: AuthMethodOAuth2, ClientID: "c"}, "client_id and client_secret"},
		{"oauth2 ok", &AuthConfig{InstanceDomain: "https://x.com", AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}, ""},
		{"basic missing username", &AuthConfig{InstanceDomain: "https://x.com", AuthMethod: AuthMethodBasic, Password: "p"}, "username and password"},
		{"basic missing password", &AuthConfig{InstanceDomain: "https://x.com", AuthMethod: AuthMethodBasic, Username: "u"}, "username and password"},
		{"basic ok", &AuthConfig{InstanceDomain: "https://x.com", AuthMethod: AuthMethodBasic, Username: "u", Password: "p"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.cfg != nil {
				err = tt.cfg.Validate()
			}
			if tt.wantErr == "" {
				if tt.cfg != nil {
					assert.NoError(t, err)
				}
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestSetupAuthentication_OAuth2(t *testing.T) {
	oauthBody := func() []byte {
		b, _ := json.Marshal(map[string]any{
			"access_token": "tok-oauth",
			"expires_in":   3600,
		})
		return b
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(oauthBody())
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	cfg := &AuthConfig{
		InstanceDomain:   srv.URL,
		AuthMethod:       AuthMethodOAuth2,
		ClientID:         "cid",
		ClientSecret:     "secret",
		HideSensitiveData: true,
	}
	logger := zap.NewNop()
	restyClient := resty.New()
	restyClient.SetBaseURL(srv.URL)

	holder, err := SetupAuthentication(restyClient, cfg, logger)
	require.NoError(t, err)
	require.NotNil(t, holder)
}

func TestSetupAuthentication_Basic(t *testing.T) {
	expires := time.Now().Add(24 * time.Hour)
	basicBody := func() []byte {
		b, _ := json.Marshal(map[string]any{
			"token":   "tok-basic",
			"expires": expires.Format(time.RFC3339),
		})
		return b
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/auth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(basicBody())
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	cfg := &AuthConfig{
		InstanceDomain: srv.URL,
		AuthMethod:     AuthMethodBasic,
		Username:       "u",
		Password:       "p",
	}
	logger := zap.NewNop()
	restyClient := resty.New()
	restyClient.SetBaseURL(srv.URL)

	holder, err := SetupAuthentication(restyClient, cfg, logger)
	require.NoError(t, err)
	require.NotNil(t, holder)
}

func TestSetupAuthentication_InvalidMethod(t *testing.T) {
	cfg := &AuthConfig{InstanceDomain: "https://x.com", AuthMethod: "other", ClientID: "c", ClientSecret: "s"}
	logger := zap.NewNop()
	restyClient := resty.New()
	_, err := SetupAuthentication(restyClient, cfg, logger)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "auth method")
}
