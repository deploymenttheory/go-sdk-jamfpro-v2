package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Inventory Information
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Returns statistics about managed/unmanaged devices and computers
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service is read-only; returns inventory statistics
//     -- Tests: TestAcceptance_InventoryInformation_get_v1
//     -- Flow: Get → Verify counts are non-negative
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Read operations (GetV1)
//   ✓ Field validation (all inventory counts are non-negative integers)
//
// Notes
// -----------------------------------------------------------------------------
//   • Returns counts for managed/unmanaged computers and mobile devices
//   • All counts should be >= 0 in any environment
//   • No CRUD or filtering operations available
//
// =============================================================================

// TestAcceptance_InventoryInformation_get_v1 verifies the inventory information
// endpoint returns valid non-negative counts.
func TestAcceptance_InventoryInformation_get_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.InventoryInformation
	ctx := context.Background()

	acc.LogTestStage(t, "Get", "Getting inventory information")

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	assert.GreaterOrEqual(t, result.ManagedComputers, 0)
	assert.GreaterOrEqual(t, result.UnmanagedComputers, 0)
	assert.GreaterOrEqual(t, result.ManagedDevices, 0)
	assert.GreaterOrEqual(t, result.UnmanagedDevices, 0)

	acc.LogTestSuccess(t, "GetV1: managedComputers=%d unmanagedComputers=%d managedDevices=%d unmanagedDevices=%d",
		result.ManagedComputers, result.UnmanagedComputers, result.ManagedDevices, result.UnmanagedDevices)
}
