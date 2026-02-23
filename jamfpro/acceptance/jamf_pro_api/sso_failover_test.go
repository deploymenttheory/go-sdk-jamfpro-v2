package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_SsoFailover_GetV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoFailover
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.FailoverURL)
}

func TestAcceptance_SsoFailover_RegenerateV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoFailover
	ctx := context.Background()

	// Get current failover URL
	before, _, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, before)

	// Regenerate the failover URL
	result, resp, err := svc.RegenerateV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.FailoverURL)
	assert.NotEqual(t, before.FailoverURL, result.FailoverURL, "Failover URL should change after regeneration")

	// Get updated failover URL to verify
	after, _, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, after)
	assert.Equal(t, result.FailoverURL, after.FailoverURL)
}
