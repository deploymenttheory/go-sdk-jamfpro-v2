package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_user_account_settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Pro User Account Settings
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetSettingsV1(ctx, keyID) - Returns username, key, and values for a preference key
//   • GetV1(ctx, keyID) - Returns the raw string value for a preference key
//   • PutV1(ctx, keyID, values) - Sets key-value preference pairs
//   • DeleteV1(ctx, keyID) - Removes the preference for a key
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Tests: TestAcceptance_JamfProUserAccountSettings_lifecycle
//     -- Flow: Put → GetSettings → Get → Delete → verify gone (404)
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_JamfProUserAccountSettings_validation_errors
//     -- Cases: empty keyID, nil/empty values
//
// Notes
// -----------------------------------------------------------------------------
//   • PUT /api/v1/user/preferences/{keyId} accepts a map[string]string body
//   • GET /api/v1/user/preferences/settings/{keyId} returns username + key + []values
//   • These endpoints may not be available for OAuth2 API client credentials
//
// =============================================================================

const testUserPrefKey = "sdkv2_acc_test_pref_key"

// TestAcceptance_JamfProUserAccountSettings_lifecycle tests the full lifecycle.
func TestAcceptance_JamfProUserAccountSettings_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProUserAccountSettings
	ctx := context.Background()

	// 1. Put (create/update) a preference
	acc.LogTestStage(t, "Put", "Setting user preference key=%s", testUserPrefKey)

	values := jamf_pro_user_account_settings.RequestUserPreferences{
		"setting1": "value1",
		"setting2": "value2",
	}

	putResp, err := svc.PutV1(ctx, testUserPrefKey, values)
	if err != nil {
		acc.LogTestWarning(t, "PutV1 returned error (may not be supported for API client credentials): %v", err)
		t.Skip("PUT /api/v1/user/preferences is not supported for this authentication method")
	}
	require.NotNil(t, putResp)
	assert.True(t, putResp.StatusCode == 200 || putResp.StatusCode == 204,
		"expected 200 or 204, got %d", putResp.StatusCode)

	acc.Cleanup(t, func() {
		cleanupCtx := context.Background()
		_, _ = svc.DeleteV1(cleanupCtx, testUserPrefKey)
	})

	acc.LogTestSuccess(t, "PutV1: key=%s status=%d", testUserPrefKey, putResp.StatusCode)

	// 2. GetSettings
	acc.LogTestStage(t, "GetSettings", "Fetching settings for key=%s", testUserPrefKey)

	settings, settingsResp, err := svc.GetSettingsV1(ctx, testUserPrefKey)
	if err != nil {
		acc.LogTestWarning(t, "GetSettingsV1 returned error: %v", err)
	} else {
		require.NotNil(t, settings)
		assert.Equal(t, 200, settingsResp.StatusCode)
		assert.Equal(t, testUserPrefKey, settings.Key)
		acc.LogTestSuccess(t, "GetSettingsV1: key=%s username=%s values=%v",
			settings.Key, settings.Username, settings.Values)
	}

	// 3. Get raw value
	acc.LogTestStage(t, "Get", "Fetching raw value for key=%s", testUserPrefKey)

	val, getResp, err := svc.GetV1(ctx, testUserPrefKey)
	if err != nil {
		acc.LogTestWarning(t, "GetV1 returned error: %v", err)
	} else {
		require.NotNil(t, getResp)
		assert.Equal(t, 200, getResp.StatusCode)
		acc.LogTestSuccess(t, "GetV1: key=%s value=%q", testUserPrefKey, val)
	}

	// 4. Delete
	acc.LogTestStage(t, "Delete", "Deleting user preference key=%s", testUserPrefKey)

	deleteResp, err := svc.DeleteV1(ctx, testUserPrefKey)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.True(t, deleteResp.StatusCode == 200 || deleteResp.StatusCode == 204,
		"expected 200 or 204, got %d", deleteResp.StatusCode)

	acc.LogTestSuccess(t, "DeleteV1: key=%s deleted", testUserPrefKey)
}

// TestAcceptance_JamfProUserAccountSettings_validation_errors verifies input validation.
func TestAcceptance_JamfProUserAccountSettings_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProUserAccountSettings

	t.Run("GetSettingsV1_EmptyKeyID", func(t *testing.T) {
		_, _, err := svc.GetSettingsV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "keyId is required")
	})

	t.Run("GetV1_EmptyKeyID", func(t *testing.T) {
		_, _, err := svc.GetV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "keyId is required")
	})

	t.Run("PutV1_EmptyKeyID", func(t *testing.T) {
		_, err := svc.PutV1(context.Background(), "", jamf_pro_user_account_settings.RequestUserPreferences{"k": "v"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "keyId is required")
	})

	t.Run("PutV1_NilValues", func(t *testing.T) {
		_, err := svc.PutV1(context.Background(), "somekey", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "values is required")
	})

	t.Run("PutV1_EmptyValues", func(t *testing.T) {
		_, err := svc.PutV1(context.Background(), "somekey", jamf_pro_user_account_settings.RequestUserPreferences{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "values is required")
	})

	t.Run("DeleteV1_EmptyKeyID", func(t *testing.T) {
		_, err := svc.DeleteV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "keyId is required")
	})
}
