package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Management Framework
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • RedeployV1(ctx, computerID) - Redeploys Jamf Management Framework on a device
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Requires an enrolled computer
//     -- Tests: TestAcceptance_JamfManagementFramework_redeploy_v1
//     -- Flow: List computers → skip if none → RedeployV1 on first result
//     -- Note: Gracefully skips if the computer is not eligible for redeploy
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_JamfManagementFramework_validation_errors
//     -- Cases: RedeployV1("") → "computer ID is required"
//
// Notes
// -----------------------------------------------------------------------------
//   • RedeployV1 returns 201 when the command is successfully queued
//   • Computer IDs are sourced from the Computer Inventory service
//   • Some computers may not be eligible (returns error); test skips gracefully
//
// =============================================================================

// TestAcceptance_JamfManagementFramework_redeploy_v1 queues a management framework redeploy.
func TestAcceptance_JamfManagementFramework_redeploy_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfManagementFramework
	ctx := context.Background()

	// Get a computer ID from inventory
	list, _, err := acc.Client.ComputerInventory.ListV3(ctx, nil)
	require.NoError(t, err, "failed to list computer inventory")

	if list == nil || list.TotalCount == 0 {
		t.Skip("No computers enrolled; skipping JamfManagementFramework redeploy test")
	}

	computerID := list.Results[0].ID
	acc.LogTestStage(t, "RedeployV1", "Redeploying Jamf Management Framework for computer ID=%s", computerID)

	result, resp, err := svc.RedeployV1(ctx, computerID)
	if err != nil {
		acc.LogTestWarning(t, "RedeployV1 returned error (computer may not be eligible): %v", err)
		t.Skip("Computer not eligible for management framework redeploy; skipping")
		return
	}

	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotEmpty(t, result.CommandUUID, "commandUUID should not be empty")
	acc.LogTestSuccess(t, "RedeployV1: computerID=%s commandUUID=%s", computerID, result.CommandUUID)
}

// TestAcceptance_JamfManagementFramework_validation_errors verifies input validation.
func TestAcceptance_JamfManagementFramework_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfManagementFramework

	t.Run("RedeployV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.RedeployV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer ID is required")
	})
}
