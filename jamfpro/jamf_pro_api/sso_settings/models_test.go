package sso_settings

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnit_SsoSettings_BrokerUpdateSerialization(t *testing.T) {
	body, err := json.Marshal(RequestOidcBrokerConfig{Enabled: false})
	require.NoError(t, err)
	var fields map[string]any
	require.NoError(t, json.Unmarshal(body, &fields))
	require.Equal(t, false, fields["enabled"])
	require.NotContains(t, fields, "clientSecret")
	require.NotContains(t, fields, "privateKeyJwt")
	require.NotContains(t, fields, "capabilities")
	require.NotContains(t, fields, "redirectUris")
}
