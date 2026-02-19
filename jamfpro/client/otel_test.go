package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultOTelConfig(t *testing.T) {
	cfg := DefaultOTelConfig()
	require.NotNil(t, cfg)
	assert.Equal(t, "jamfpro-client", cfg.ServiceName)
	assert.NotNil(t, cfg.TracerProvider)
	assert.NotNil(t, cfg.Propagators)
}

func TestTransport_EnableTracing(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)

	err = tr.EnableTracing(nil)
	require.NoError(t, err)

	err = tr.EnableTracing(DefaultOTelConfig())
	require.NoError(t, err)
}
