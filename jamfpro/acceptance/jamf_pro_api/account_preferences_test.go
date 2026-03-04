package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_AccountPreferences_get_v3(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.AccountPreferences
	ctx := context.Background()

	result, resp, err := svc.GetV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.Language)
}

func TestAcceptance_AccountPreferences_update_v2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.AccountPreferences
	ctx := context.Background()

	current, _, err := svc.GetV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	request := *current
	request.DisablePageLeaveCheck = !request.DisablePageLeaveCheck
	updated, resp, err := svc.UpdateV3(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.True(t, resp.StatusCode() == 200 || resp.StatusCode() == 204, "expected 200 or 204, got %d", resp.StatusCode())
	_ = updated

	// Restore original
	request.DisablePageLeaveCheck = !request.DisablePageLeaveCheck
	_, _, _ = svc.UpdateV3(ctx, &request)
}
