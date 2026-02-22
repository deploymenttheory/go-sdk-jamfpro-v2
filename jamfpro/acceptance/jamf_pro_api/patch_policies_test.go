package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Patch Policies
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   Patch Policy Read Operations (V2 API - Read Only):
//   • ListV2(ctx) - Lists all patch policies with automatic pagination
//   • GetByIDV2(ctx, id) - Retrieves a patch policy by ID (helper using ListV2)
//   • GetByNameV2(ctx, name) - Retrieves a patch policy by name (helper using ListV2)
//   • GetDashboardStatusV2(ctx, id) - Checks if a patch policy is on the dashboard
//   • AddToDashboardV2(ctx, id) - Adds a patch policy to the dashboard
//   • RemoveFromDashboardV2(ctx, id) - Removes a patch policy from the dashboard
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Operations
//     -- Reason: v2 API is read-only for patch policies. Create/Update/Delete
//        operations require the Classic API (XML).
//     -- Tests: TestAcceptance_PatchPolicies_ListV2,
//               TestAcceptance_PatchPolicies_DashboardOperations
//     -- Flow: List → Get (if available) → Dashboard Operations
//
//   Note: RSQL Filter Testing NOT applicable
//     -- ListV2 uses automatic pagination, not RSQL filtering
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ List operations (all patch policies with pagination)
//   ✓ Read operations (GetByID, GetByName if policies exist)
//   ✓ Dashboard operations (check status, add, remove)
//   ✗ Create/Update/Delete operations (not available in v2 API)
//   ✗ Input validation and error handling (not yet tested)
//
// Notes
// -----------------------------------------------------------------------------
//   • Patch policies define automated software update deployments
//   • The v2 API is READ-ONLY - create/update/delete require Classic API
//   • GetByID and GetByName are helper methods that use ListV2 for lookup
//   • Dashboard operations allow adding/removing policies from the dashboard view
//   • Tests may skip if no patch policies exist in the environment
//   • All dashboard operations are tested and cleaned up (policy removed from dashboard)
//   • TODO: Add validation error tests for empty IDs/names
//
// =============================================================================

func TestAcceptance_PatchPolicies_ListV2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.PatchPolicies
	ctx := context.Background()

	result, resp, err := svc.ListV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
	assert.NotNil(t, result.Results)

	if result.TotalCount > 0 {
		acc.LogTestSuccess(t, "Found %d patch policies", result.TotalCount)

		// Verify first policy has expected fields
		policy := result.Results[0]
		assert.NotEmpty(t, policy.ID)
		assert.NotEmpty(t, policy.Name)
		assert.NotEmpty(t, policy.SoftwareTitleId)
		acc.LogTestSuccess(t, "Sample policy: ID=%s, Name=%s, Enabled=%v",
			policy.ID, policy.Name, policy.Enabled)

		// Test GetByID
		acc.LogTestStage(t, "Read", "Getting patch policy by ID")
		byID, resp, err := svc.GetByIDV2(ctx, policy.ID)
		require.NoError(t, err)
		require.NotNil(t, byID)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, policy.ID, byID.ID)
		assert.Equal(t, policy.Name, byID.Name)
		acc.LogTestSuccess(t, "Retrieved patch policy by ID: %s", byID.Name)

		// Test GetByName
		acc.LogTestStage(t, "Read", "Getting patch policy by name")
		byName, resp, err := svc.GetByNameV2(ctx, policy.Name)
		require.NoError(t, err)
		require.NotNil(t, byName)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, policy.ID, byName.ID)
		assert.Equal(t, policy.Name, byName.Name)
		acc.LogTestSuccess(t, "Retrieved patch policy by name: %s", byName.Name)
	} else {
		acc.LogTestSuccess(t, "No patch policies found (empty list OK)")
	}
}

func TestAcceptance_PatchPolicies_DashboardOperations(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.PatchPolicies
	ctx := context.Background()

	// First, get a list of patch policies
	acc.LogTestStage(t, "Setup", "Getting list of patch policies")
	list, resp, err := svc.ListV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, resp.StatusCode)

	if list.TotalCount == 0 {
		t.Skip("No patch policies available - skipping dashboard operations test")
		return
	}

	// Use the first policy for dashboard operations
	policyID := list.Results[0].ID
	policyName := list.Results[0].Name
	acc.LogTestSuccess(t, "Using policy ID=%s, Name=%s for dashboard operations", policyID, policyName)

	// Check initial dashboard status
	acc.LogTestStage(t, "Check Status", "Getting initial dashboard status")
	initialStatus, resp, err := svc.GetDashboardStatusV2(ctx, policyID)
	require.NoError(t, err)
	require.NotNil(t, initialStatus)
	assert.Equal(t, 200, resp.StatusCode)
	acc.LogTestSuccess(t, "Initial dashboard status - OnDashboard: %v", initialStatus.OnDashboard)

	// If already on dashboard, remove it first for clean test
	if initialStatus.OnDashboard {
		acc.LogTestStage(t, "Setup", "Removing policy from dashboard for clean test")
		resp, err := svc.RemoveFromDashboardV2(ctx, policyID)
		require.NoError(t, err)
		assert.Contains(t, []int{200, 204}, resp.StatusCode)
		acc.LogTestSuccess(t, "Removed policy from dashboard")
	}

	// Add to dashboard
	acc.LogTestStage(t, "Add", "Adding patch policy to dashboard")
	addResp, err := svc.AddToDashboardV2(ctx, policyID)
	require.NoError(t, err)
	assert.Contains(t, []int{200, 201, 204}, addResp.StatusCode)
	acc.LogTestSuccess(t, "Added patch policy to dashboard")

	// Verify it's on the dashboard
	acc.LogTestStage(t, "Verify", "Verifying policy is on dashboard")
	statusAfterAdd, resp, err := svc.GetDashboardStatusV2(ctx, policyID)
	require.NoError(t, err)
	require.NotNil(t, statusAfterAdd)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, statusAfterAdd.OnDashboard, "Policy should be on dashboard after adding")
	acc.LogTestSuccess(t, "Verified policy is on dashboard")

	// Remove from dashboard
	acc.LogTestStage(t, "Remove", "Removing patch policy from dashboard")
	removeResp, err := svc.RemoveFromDashboardV2(ctx, policyID)
	require.NoError(t, err)
	assert.Contains(t, []int{200, 204}, removeResp.StatusCode)
	acc.LogTestSuccess(t, "Removed patch policy from dashboard")

	// Verify it's removed from the dashboard
	acc.LogTestStage(t, "Verify Removal", "Verifying policy is removed from dashboard")
	statusAfterRemove, resp, err := svc.GetDashboardStatusV2(ctx, policyID)
	require.NoError(t, err)
	require.NotNil(t, statusAfterRemove)
	assert.Equal(t, 200, resp.StatusCode)
	assert.False(t, statusAfterRemove.OnDashboard, "Policy should not be on dashboard after removal")
	acc.LogTestSuccess(t, "Verified policy is removed from dashboard")

	// Restore to initial state if needed
	if initialStatus.OnDashboard {
		acc.LogTestStage(t, "Cleanup", "Restoring initial dashboard state")
		_, _ = svc.AddToDashboardV2(ctx, policyID)
		acc.LogTestSuccess(t, "Restored policy to dashboard")
	}
}
