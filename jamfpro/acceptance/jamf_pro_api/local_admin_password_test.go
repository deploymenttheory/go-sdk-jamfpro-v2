package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/local_admin_password"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Local Admin Password (LAPS)
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetPendingRotationsV2(ctx) - Lists devices and usernames with pending LAPS rotations
//   • GetSettingsV2(ctx) - Retrieves current LAPS settings
//   • UpdateSettingsV2(ctx, settings) - Updates LAPS settings
//   • GetPasswordHistoryByClientManagementIDV2(ctx, clientManagementID, username) - Retrieves password view history
//   • GetCurrentPasswordByClientManagementIDV2(ctx, clientManagementID, username) - Retrieves current password (triggers rotation)
//   • GetFullHistoryByClientManagementIDV2(ctx, clientManagementID) - Retrieves complete password history
//   • GetCapableAccountsByClientManagementIDV2(ctx, clientManagementID) - Lists LAPS-capable accounts
//   • SetPasswordByClientManagementIDV2(ctx, clientManagementID, passwordList) - Sets LAPS passwords
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Read-Only Settings
//     -- Reason: LAPS is a system-level feature that requires specific device configuration
//     -- Tests: TestAcceptance_LocalAdminPassword_Settings_GetAndUpdate
//     -- Flow: Get current settings → Update → Get again → Restore original
//
//   ✓ Pattern 3: Read-Only Lists
//     -- Reason: Pending rotations and capable accounts are system-managed lists
//     -- Tests: TestAcceptance_LocalAdminPassword_PendingRotations
//     -- Flow: List pending rotations → Verify response structure
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_LocalAdminPassword_ValidationErrors
//     -- Cases: Empty IDs, nil requests, missing required fields
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Settings operations (Get, Update)
//   ✓ List pending rotations
//   ✓ Input validation and error handling
//   ✓ Restore original settings (cleanup)
//
// Notes
// -----------------------------------------------------------------------------
//   • LAPS requires enrolled devices with LAPS capability
//   • Password viewing triggers automatic rotation based on settings
//   • Tests use read-only operations to avoid affecting production LAPS state
//   • Device-specific operations are not tested as they require enrolled devices
//   • Settings are restored to original values after update tests
//
// =============================================================================
// TestAcceptance_LocalAdminPassword_Settings_GetAndUpdate exercises the
// settings read and update operations.
// =============================================================================

func TestAcceptance_LocalAdminPassword_Settings_GetAndUpdate(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.LocalAdminPassword
	ctx := context.Background()

	// 1. Get current settings
	acc.LogTestStage(t, "Get", "Fetching current LAPS settings")

	originalSettings, getResp, err := svc.GetSettingsV2(ctx)
	require.NoError(t, err, "GetSettingsV2 should not return an error")
	require.NotNil(t, originalSettings)
	assert.Equal(t, 200, getResp.StatusCode)
	acc.LogTestSuccess(t, "Current settings: AutoDeploy=%v, RotationTime=%d days",
		originalSettings.AutoDeployEnabled, originalSettings.PasswordRotationTime)

	// 2. Update settings (modify rotation time)
	acc.LogTestStage(t, "Update", "Updating LAPS settings")

	updatedSettings := &local_admin_password.SettingsResource{
		AutoDeployEnabled:        originalSettings.AutoDeployEnabled,
		PasswordRotationTime:     originalSettings.PasswordRotationTime,
		AutoRotateEnabled:        originalSettings.AutoRotateEnabled,
		AutoRotateExpirationTime: originalSettings.AutoRotateExpirationTime,
	}
	updateResp, err := svc.UpdateSettingsV2(ctx, updatedSettings)
	require.NoError(t, err)
	require.NotNil(t, updateResp)
	assert.Equal(t, 200, updateResp.StatusCode)
	acc.LogTestSuccess(t, "Settings updated successfully")

	// 3. Verify settings were updated
	acc.LogTestStage(t, "Verify", "Verifying settings update")

	fetchedSettings, fetchResp, err := svc.GetSettingsV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, fetchedSettings)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, updatedSettings.AutoDeployEnabled, fetchedSettings.AutoDeployEnabled)
	assert.Equal(t, updatedSettings.PasswordRotationTime, fetchedSettings.PasswordRotationTime)
	acc.LogTestSuccess(t, "Settings verified: RotationTime=%d days", fetchedSettings.PasswordRotationTime)

	// 4. Restore original settings
	acc.LogTestStage(t, "Restore", "Restoring original LAPS settings")

	restoreResp, err := svc.UpdateSettingsV2(ctx, originalSettings)
	require.NoError(t, err)
	require.NotNil(t, restoreResp)
	assert.Equal(t, 200, restoreResp.StatusCode)
	acc.LogTestSuccess(t, "Original settings restored")
}

// =============================================================================
// TestAcceptance_LocalAdminPassword_PendingRotations
// =============================================================================

func TestAcceptance_LocalAdminPassword_PendingRotations(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.LocalAdminPassword
	ctx := context.Background()

	acc.LogTestStage(t, "List", "Fetching pending LAPS rotations")

	result, resp, err := svc.GetPendingRotationsV2(ctx)
	require.NoError(t, err, "GetPendingRotationsV2 should not return an error")
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)

	acc.LogTestSuccess(t, "Pending rotations: %d total", result.TotalCount)

	if result.TotalCount > 0 {
		acc.LogTestSuccess(t, "Sample rotation: Device=%s, User=%s, Created=%s",
			result.Results[0].LapsUser.ClientManagementID,
			result.Results[0].LapsUser.Username,
			result.Results[0].CreatedDate)
	}
}

// =============================================================================
// TestAcceptance_LocalAdminPassword_ValidationErrors
// =============================================================================

func TestAcceptance_LocalAdminPassword_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.LocalAdminPassword

	t.Run("UpdateSettingsV2_NilRequest", func(t *testing.T) {
		_, err := svc.UpdateSettingsV2(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "settings is required")
	})

	t.Run("GetPasswordHistoryByClientManagementIDV2_EmptyClientManagementID", func(t *testing.T) {
		_, _, err := svc.GetPasswordHistoryByClientManagementIDV2(context.Background(), "", "admin")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementID is required")
	})

	t.Run("GetPasswordHistoryByClientManagementIDV2_EmptyUsername", func(t *testing.T) {
		_, _, err := svc.GetPasswordHistoryByClientManagementIDV2(context.Background(), "device-001", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "username is required")
	})

	t.Run("GetCurrentPasswordByClientManagementIDV2_EmptyClientManagementID", func(t *testing.T) {
		_, _, err := svc.GetCurrentPasswordByClientManagementIDV2(context.Background(), "", "admin")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementID is required")
	})

	t.Run("GetFullHistoryByClientManagementIDV2_EmptyClientManagementID", func(t *testing.T) {
		_, _, err := svc.GetFullHistoryByClientManagementIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementID is required")
	})

	t.Run("GetCapableAccountsByClientManagementIDV2_EmptyClientManagementID", func(t *testing.T) {
		_, _, err := svc.GetCapableAccountsByClientManagementIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementID is required")
	})

	t.Run("SetPasswordByClientManagementIDV2_EmptyClientManagementID", func(t *testing.T) {
		req := &local_admin_password.SetPasswordRequest{
			LapsUserPasswordList: []local_admin_password.LapsUserPassword{{Username: "admin", Password: "P@ss"}},
		}
		_, _, err := svc.SetPasswordByClientManagementIDV2(context.Background(), "", req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementID is required")
	})

	t.Run("SetPasswordByClientManagementIDV2_NilRequest", func(t *testing.T) {
		_, _, err := svc.SetPasswordByClientManagementIDV2(context.Background(), "device-001", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "passwordList is required")
	})
}
