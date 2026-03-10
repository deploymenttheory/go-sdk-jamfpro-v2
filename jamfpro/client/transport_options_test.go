package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/config"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// These tests verify that NewTransport correctly reads and applies each
// TransportSettings field. Options are constructed inline rather than via
// the With* helpers (which live in the root jamfpro package) so that this
// package remains cycle-free.

func TestNewTransport_BaseURL(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.BaseURL = "https://custom.example.com"
		return nil
	})
	require.NoError(t, err)
	assert.Equal(t, "https://custom.example.com", tr.BaseURL)
}

func TestNewTransport_CustomLogger(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	logger := zap.NewNop()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.Logger = logger
		return nil
	})
	require.NoError(t, err)
	assert.Same(t, logger, tr.GetLogger())
}

func TestNewTransport_OptionError(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, func(s *TransportSettings) error {
		return fmt.Errorf("logger cannot be nil")
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "logger")
}

func TestNewTransport_Timeout(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.Timeout = 30 * time.Second
		return nil
	})
	require.NoError(t, err)
}

func TestNewTransport_RetrySettings(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.RetryCount = 2
		s.RetryWaitTime = time.Second
		s.RetryMaxWaitTime = 10 * time.Second
		return nil
	})
	require.NoError(t, err)
}

func TestNewTransport_UserAgent_GlobalHeaders(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.UserAgent = "custom/1.0"
		s.GlobalHeaders = map[string]string{"A": "a", "B": "b"}
		return nil
	})
	require.NoError(t, err)
	assert.Equal(t, "custom/1.0", tr.userAgent)
	assert.Equal(t, "a", tr.globalHeaders["A"])
	assert.Equal(t, "b", tr.globalHeaders["B"])
}

func TestNewTransport_TLSAndProxyOptions(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.ProxyURL = "http://proxy:8080"
		s.TLSClientConfig = &tls.Config{}
		s.HTTPTransport = http.DefaultTransport
		s.InsecureSkipVerify = true
		return nil
	})
	require.NoError(t, err)
}

func TestNewTransport_ConcurrencyAndDelays(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.MaxConcurrentRequests = 2
		s.MandatoryRequestDelay = time.Millisecond
		s.TotalRetryDuration = 10 * time.Second
		return nil
	})
	require.NoError(t, err)
	assert.NotNil(t, tr.sem)
	assert.Greater(t, tr.requestDelay, time.Duration(0))
	assert.Greater(t, tr.totalRetryDuration, time.Duration(0))
}

func TestNewTransport_MaxConcurrentRequests_Zero(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.MaxConcurrentRequests = 0
		return nil
	})
	require.NoError(t, err)
	assert.Nil(t, tr.sem)
}

func TestNewTransport_Debug(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.Debug = true
		return nil
	})
	require.NoError(t, err)
}
