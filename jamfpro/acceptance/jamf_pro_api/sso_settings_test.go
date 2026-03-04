package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_settings"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_SsoSettings_get_v3(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoSettings
	ctx := context.Background()

	result, resp, err := svc.GetV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestAcceptance_SsoSettings_get_enrollment_customization_dependencies(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoSettings
	ctx := context.Background()

	result, resp, err := svc.GetEnrollmentCustomizationDependenciesV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestAcceptance_SsoSettings_update_v3(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoSettings
	ctx := context.Background()

	current, _, err := svc.GetV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	request := *current
	request.SsoBypassAllowed = !request.SsoBypassAllowed
	updated, resp, err := svc.UpdateV3(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, updated)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	// Restore
	request.SsoBypassAllowed = current.SsoBypassAllowed
	_, _, _ = svc.UpdateV3(ctx, &request)
}

func TestAcceptance_SsoSettings_get_history_v3(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoSettings
	ctx := context.Background()

	// Add history note first
	noteReq := &shared.SharedHistoryNoteRequest{
		Note: "Acceptance test history note for SSO settings",
	}
	addResult, addResp, err := svc.AddHistoryNoteV3(ctx, noteReq)
	require.NoError(t, err)
	require.NotNil(t, addResult)
	assert.Equal(t, 201, addResp.StatusCode())
	t.Logf("Added history note with ID: %s", addResult.ID)

	result, resp, err := svc.GetHistoryV3(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.GreaterOrEqual(t, result.TotalCount, 1, "Should have at least the note we just added")
}

func TestAcceptance_SsoSettings_add_history_note_v3(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoSettings
	ctx := context.Background()

	noteReq := &sso_settings.AddHistoryNoteRequest{
		Note: "Test history note from acceptance test",
	}

	result, resp, err := svc.AddHistoryNoteV3(ctx, noteReq)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	t.Logf("Added history note successfully")
}
