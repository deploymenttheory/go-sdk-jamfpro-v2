package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_OAuth2SessionTokens_get_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.Oauth2SessionTokens
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	if err != nil && resp != nil && resp.StatusCode() == 400 {
		t.Skip("OAuth token data not available in session (requires OAuth authentication)")
	}
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.AccessToken)
	assert.NotEmpty(t, result.IDToken)
}
