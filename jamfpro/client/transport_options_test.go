package client

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestWithBaseURL(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, WithBaseURL("https://custom.example.com"))
	require.NoError(t, err)
	assert.Equal(t, "https://custom.example.com", tr.BaseURL)
}

func TestWithLogger(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	logger := zap.NewNop()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, WithLogger(logger))
	require.NoError(t, err)
	assert.Same(t, logger, tr.GetLogger())
}

func TestWithLogger_Nil(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, WithLogger(nil))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "logger")
}

func TestWithTimeout(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, WithTimeout(30*time.Second))
	require.NoError(t, err)
}

func TestWithRetryCount_WithRetryWaitTime_WithRetryMaxWaitTime(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, WithRetryCount(2), WithRetryWaitTime(time.Second), WithRetryMaxWaitTime(10*time.Second))
	require.NoError(t, err)
}

func TestWithUserAgent_WithGlobalHeader_WithGlobalHeaders(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, WithUserAgent("custom/1.0"), WithGlobalHeader("A", "a"), WithGlobalHeaders(map[string]string{"B": "b"}))
	require.NoError(t, err)
	assert.Equal(t, "custom/1.0", tr.userAgent)
	assert.Equal(t, "a", tr.globalHeaders["A"])
	assert.Equal(t, "b", tr.globalHeaders["B"])
}

func TestWithProxy_WithTLSClientConfig_WithTransport_WithInsecureSkipVerify(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, WithProxy("http://proxy:8080"), WithTLSClientConfig(&tls.Config{}), WithTransport(http.DefaultTransport), WithInsecureSkipVerify())
	require.NoError(t, err)
}

func TestWithMaxConcurrentRequests_WithMandatoryRequestDelay_WithTotalRetryDuration(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, WithMaxConcurrentRequests(2), WithMandatoryRequestDelay(time.Millisecond), WithTotalRetryDuration(10*time.Second))
	require.NoError(t, err)
	assert.NotNil(t, tr.sem)
	assert.Greater(t, tr.requestDelay, time.Duration(0))
	assert.Greater(t, tr.totalRetryDuration, time.Duration(0))
}

func TestWithMaxConcurrentRequests_Zero(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, WithMaxConcurrentRequests(0))
	require.NoError(t, err)
	assert.Nil(t, tr.sem)
}

func TestWithDebug(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, WithDebug())
	require.NoError(t, err)
}
