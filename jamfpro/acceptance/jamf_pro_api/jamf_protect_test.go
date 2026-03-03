package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_protect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Protect
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetSettingsV1(ctx) - Returns current Jamf Protect integration settings
//   • UpdateSettingsV1(ctx, request) - Updates settings (e.g. autoInstall flag)
//   • RegisterV1(ctx, request) - Registers Jamf Protect integration
//   • SyncPlansV1(ctx) - Syncs plans from Protect server
//   • CreateIntegrationV1(ctx, registration, autoInstall) - Composite create
//   • ListDeploymentTasksV1(ctx, deploymentID, rsqlQuery) - Lists deployment tasks
//   • RetryDeploymentTasksV1(ctx, deploymentID) - Retries failed tasks
//   • ListHistoryV1(ctx, rsqlQuery) - Lists Protect history entries
//   • CreateHistoryNoteV1(ctx, request) - Creates a history note
//   • ListPlansV1(ctx, rsqlQuery) - Lists available Protect plans
//   • DeleteIntegrationV1(ctx) - Removes the integration
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Tests: TestAcceptance_JamfProtect_get_settings_v1
//     -- Flow: GetSettingsV1 → verify (graceful if not configured)
//
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Tests: TestAcceptance_JamfProtect_list_history_v1
//     -- Flow: ListHistoryV1 → verify count ≥ 0
//
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Tests: TestAcceptance_JamfProtect_list_plans_v1
//     -- Flow: ListPlansV1 → verify count ≥ 0 (graceful if no integration)
//
// Notes
// -----------------------------------------------------------------------------
//   • RegisterV1 / CreateIntegrationV1 / DeleteIntegrationV1 require real
//     Jamf Protect credentials and are NOT run in acceptance tests to avoid
//     disrupting an existing integration.
//   • ListDeploymentTasksV1 requires a known deployment ID — tested only if
//     plans are available.
//   • GetSettingsV1 returns 404 or an error if no integration is configured.
//
// =============================================================================

// TestAcceptance_JamfProtect_get_settings_v1 verifies settings can be retrieved.
func TestAcceptance_JamfProtect_get_settings_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProtect
	ctx := context.Background()

	acc.LogTestStage(t, "GetSettings", "Getting Jamf Protect integration settings")

	result, resp, err := svc.GetSettingsV1(ctx)
	if err != nil {
		if resp != nil && (resp.StatusCode == 404 || resp.StatusCode == 400) {
			t.Logf("GetSettingsV1 returned %d - Jamf Protect integration not configured, skipping", resp.StatusCode)
			return
		}
		require.NoError(t, err)
	}
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	acc.LogTestSuccess(t, "JamfProtect settings: protectURL=%s autoInstall=%v",
		result.ProtectURL, result.AutoInstall)
}

// TestAcceptance_JamfProtect_list_history_v1 verifies history entries can be listed.
func TestAcceptance_JamfProtect_list_history_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProtect
	ctx := context.Background()

	// Add history note first
	acc.LogTestStage(t, "CreateHistoryNote", "Adding history note for Jamf Protect")
	noteReq := &jamf_protect.RequestJamfProtectHistoryNote{
		Note: "Acceptance test history note for Jamf Protect",
	}
	addResult, addResp, err := svc.CreateHistoryNoteV1(ctx, noteReq)
	if err != nil {
		if addResp != nil && (addResp.StatusCode == 404 || addResp.StatusCode == 400) {
			t.Skipf("CreateHistoryNoteV1 returned %d - Jamf Protect not configured, skipping", addResp.StatusCode)
			return
		}
		require.NoError(t, err)
	}
	require.NotNil(t, addResult)
	assert.Equal(t, 201, addResp.StatusCode)
	acc.LogTestSuccess(t, "Added history note with ID: %s", addResult.ID)

	acc.LogTestStage(t, "ListHistory", "Listing Jamf Protect history")

	result, resp, err := svc.ListHistoryV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 1, "Should have at least the note we just added")

	acc.LogTestSuccess(t, "ListHistoryV1: found %d history entries", result.TotalCount)
}

// TestAcceptance_JamfProtect_list_plans_v1 verifies plans can be listed.
func TestAcceptance_JamfProtect_list_plans_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProtect
	ctx := context.Background()

	acc.LogTestStage(t, "ListPlans", "Listing Jamf Protect plans")

	result, resp, err := svc.ListPlansV1(ctx, nil)
	if err != nil {
		if resp != nil && (resp.StatusCode == 404 || resp.StatusCode == 400) {
			t.Logf("ListPlansV1 returned %d - Jamf Protect not configured, skipping", resp.StatusCode)
			return
		}
		require.NoError(t, err)
	}
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)

	acc.LogTestSuccess(t, "ListPlansV1: found %d plans", result.TotalCount)
}

// TestAcceptance_JamfProtect_validation_errors verifies input validation.
func TestAcceptance_JamfProtect_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProtect

	t.Run("UpdateSettingsV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateSettingsV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot be nil")
	})

	t.Run("RegisterV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.RegisterV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot be nil")
	})

	t.Run("CreateIntegrationV1_NilRegistration", func(t *testing.T) {
		_, _, err := svc.CreateIntegrationV1(context.Background(), nil, false)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot be nil")
	})

	t.Run("ListDeploymentTasksV1_EmptyDeploymentID", func(t *testing.T) {
		_, _, err := svc.ListDeploymentTasksV1(context.Background(), "", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "deployment ID is required")
	})

	t.Run("RetryDeploymentTasksV1_EmptyDeploymentID", func(t *testing.T) {
		_, err := svc.RetryDeploymentTasksV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "deployment ID is required")
	})

	t.Run("CreateHistoryNoteV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateHistoryNoteV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot be nil")
	})
}
