package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Impact Alert Notification Settings
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Retrieves impact alert notification settings
//   • UpdateV1(ctx, request) - Updates impact alert notification settings (returns 204)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration Testing
//     -- Reason: Impact Alert Notification Settings is a singleton settings object
//     -- Tests: TestAcceptance_ImpactAlertNotificationSettings_get_and_update
//     -- Flow: Get current → Update → Get (verify) → Restore original
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_ImpactAlertNotificationSettings_validation_errors
//     -- Cases: Nil request validation
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get operations (retrieve current settings)
//   ✓ Update operations (modify settings via PUT, returns 204)
//   ✓ Input validation and error handling
//   ✓ Settings restoration (restore original config after tests)
//
// Notes
// -----------------------------------------------------------------------------
//   • Impact Alert Notification Settings is a settings object, not a CRUD resource
//   • Tests preserve original settings and restore them after modifications
//   • Update endpoint returns 204 No Content (no response body)
//   • All tests use proper cleanup to restore state
//
// =============================================================================

func TestAcceptance_ImpactAlertNotificationSettings_get_and_update(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ImpactAlertNotificationSettings
	ctx := context.Background()

	// Get current settings to preserve them
	original, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, original)
	assert.Equal(t, 200, resp.StatusCode)

	// Register cleanup to restore original settings
	t.Cleanup(func() {
		_, err := svc.UpdateV1(ctx, original)
		if err != nil {
			t.Logf("Warning: Failed to restore original impact alert notification settings: %v", err)
		}
	})

	// Update with modified settings (toggle all boolean values as a test)
	modified := *original
	modified.ScopeableObjectsAlertEnabled = !original.ScopeableObjectsAlertEnabled
	modified.ScopeableObjectsConfirmationCodeEnabled = !original.ScopeableObjectsConfirmationCodeEnabled
	modified.DeployableObjectsAlertEnabled = !original.DeployableObjectsAlertEnabled
	modified.DeployableObjectsConfirmationCodeEnabled = !original.DeployableObjectsConfirmationCodeEnabled

	updateResp, err := svc.UpdateV1(ctx, &modified)
	require.NoError(t, err)
	require.NotNil(t, updateResp)

	// Update returns 204 No Content
	assert.Equal(t, 204, updateResp.StatusCode)

	// Get again to verify persistence
	current, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, modified.ScopeableObjectsAlertEnabled, current.ScopeableObjectsAlertEnabled)
	assert.Equal(t, modified.ScopeableObjectsConfirmationCodeEnabled, current.ScopeableObjectsConfirmationCodeEnabled)
	assert.Equal(t, modified.DeployableObjectsAlertEnabled, current.DeployableObjectsAlertEnabled)
	assert.Equal(t, modified.DeployableObjectsConfirmationCodeEnabled, current.DeployableObjectsConfirmationCodeEnabled)
}

func TestAcceptance_ImpactAlertNotificationSettings_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ImpactAlertNotificationSettings
	ctx := context.Background()

	// Test nil request validation
	resp, err := svc.UpdateV1(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}
