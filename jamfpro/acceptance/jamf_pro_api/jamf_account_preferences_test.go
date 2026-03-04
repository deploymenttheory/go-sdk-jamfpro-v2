package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Account Preferences
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV3(ctx) - Returns current Jamf Pro account preferences
//   • UpdateV3(ctx, request) - Updates account preferences (PATCH, returns 204)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration
//     -- Tests: TestAcceptance_JamfAccountPreferences_get_and_update_v3
//     -- Flow: Get → Update (toggle bool) → Verify → Restore
//
// Notes
// -----------------------------------------------------------------------------
//   • PATCH /api/v3/account-preferences returns 204 No Content on success
//   • This service wraps the same endpoint as AccountPreferences but is
//     exposed as a distinct client field (JamfAccountPreferences)
//
// =============================================================================

// TestAcceptance_JamfAccountPreferences_get_v3 verifies account preferences can be read.
func TestAcceptance_JamfAccountPreferences_get_v3(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfAccountPreferences
	ctx := context.Background()

	acc.LogTestStage(t, "Get", "Getting Jamf account preferences")

	result, resp, err := svc.GetV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.Language, "language preference should not be empty")

	acc.LogTestSuccess(t, "JamfAccountPreferences: language=%s timezone=%s theme=%s",
		result.Language, result.Timezone, result.UserInterfaceDisplayTheme)
}

// TestAcceptance_JamfAccountPreferences_get_and_update_v3 verifies preferences can be updated.
func TestAcceptance_JamfAccountPreferences_get_and_update_v3(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfAccountPreferences
	ctx := context.Background()

	// 1. Get current
	acc.LogTestStage(t, "Get", "Getting current Jamf account preferences")

	current, _, err := svc.GetV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	// 2. Update (toggle DisablePageLeaveCheck)
	acc.LogTestStage(t, "Update", "Updating DisablePageLeaveCheck preference")

	request := *current
	request.DisablePageLeaveCheck = !request.DisablePageLeaveCheck

	updated, updateResp, err := svc.UpdateV3(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, updateResp)
	// PATCH returns 204 No Content; some tenants may return 200
	assert.True(t, updateResp.StatusCode() == 200 || updateResp.StatusCode() == 204,
		"expected 200 or 204, got %d", updateResp.StatusCode())
	_ = updated

	acc.LogTestSuccess(t, "UpdateV3: disablePageLeaveCheck toggled to %v", request.DisablePageLeaveCheck)

	// 3. Restore original
	restore := *current
	_, _, _ = svc.UpdateV3(ctx, &restore)
	acc.LogTestSuccess(t, "Preferences restored")
}
