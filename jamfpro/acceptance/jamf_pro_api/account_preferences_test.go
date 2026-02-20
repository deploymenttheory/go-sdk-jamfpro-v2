package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_AccountPreferences_GetV2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.AccountPreferences
	ctx := context.Background()

	result, resp, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.Language)
}

func TestAcceptance_AccountPreferences_UpdateV2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.AccountPreferences
	ctx := context.Background()

	current, _, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	request := *current
	request.DisablePageLeaveCheck = !request.DisablePageLeaveCheck
	updated, resp, err := svc.UpdateV2(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, updated)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Restore original
	request.DisablePageLeaveCheck = !request.DisablePageLeaveCheck
	_, _, _ = svc.UpdateV2(ctx, &request)
}
