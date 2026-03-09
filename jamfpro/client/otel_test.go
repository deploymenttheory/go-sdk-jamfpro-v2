package client

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/config"
	"github.com/stretchr/testify/require"
)

func TestOpenTelemetryAlwaysEnabled(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}

	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	require.NotNil(t, tr)

	httpClient := tr.client.Client()
	require.NotNil(t, httpClient)
	require.NotNil(t, httpClient.Transport)
}
