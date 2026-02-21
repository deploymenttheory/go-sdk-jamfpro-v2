package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory_collection_settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Computer Inventory Collection Settings
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   Settings Management:
//   • GetV2(ctx) - Retrieves computer inventory collection settings
//   • UpdateV2(ctx, request) - Updates settings using merge-patch semantics (PATCH)
//
//   Custom Application Paths:
//   • CreateCustomPathV2(ctx, request) - Creates a custom application path for inventory
//   • DeleteCustomPathByIDV2(ctx, id) - Deletes a custom application path by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration
//     -- Reason: Singleton settings that cannot be created or deleted, only updated
//     -- Tests: TestComputerInventoryCollectionSettings_GetAndUpdate
//     -- Flow: Get original → Update → Verify → Restore → (Verify restoration implicit)
//
//   ✓ Partial CRUD for Custom Paths
//     -- Reason: Custom paths support create and delete (no update or list operations)
//     -- Tests: TestComputerInventoryCollectionSettings_CustomPath
//     -- Flow: Create path → Get settings (verify path exists) → Delete path
//
// Test Coverage
// -----------------------------------------------------------------------------
//   Settings:
//   ✓ Get current settings
//   ✓ Update settings (PATCH merge-patch semantics)
//   ✓ Verify updated settings
//   ✓ Restore original settings
//   ✗ Verify restoration (implicit - not explicitly checked)
//
//   Custom Paths:
//   ✓ Create custom application path
//   ✓ Verify path appears in settings (via GetV2)
//   ✓ Delete custom path
//   ✓ Cleanup and resource management
//
// Notes
// -----------------------------------------------------------------------------
//   • Main settings are singleton configuration - no create/delete operations
//   • UpdateV2 uses PATCH with merge-patch semantics (partial updates)
//   • Custom paths extend default application inventory locations
//   • Scope options for custom paths: USER_LIBRARY, COMPUTER_LIBRARY, etc.
//   • Custom path creation may not be supported in all environments (test handles gracefully)
//   • Settings affect all computer inventory collection for the tenant
//   • MonitorApplicationUsage is toggled in test as a safe change to verify updates
//   • Test properly implements Pattern 2 restoration cycle
//   • TODO: Add explicit verification step after restoration in settings test
//
// =============================================================================

func TestComputerInventoryCollectionSettings_GetAndUpdate(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.ComputerInventoryCollectionSettings

	original, _, err := svc.GetV2(ctx)
	require.NoError(t, err)
	assert.NotNil(t, original)

	updateReq := &computer_inventory_collection_settings.ResourceComputerInventoryCollectionSettings{
		ComputerInventoryCollectionPreferences: original.ComputerInventoryCollectionPreferences,
	}
	updateReq.ComputerInventoryCollectionPreferences.MonitorApplicationUsage = !original.ComputerInventoryCollectionPreferences.MonitorApplicationUsage

	_, err = svc.UpdateV2(ctx, updateReq)
	require.NoError(t, err)

	updated, _, err := svc.GetV2(ctx)
	require.NoError(t, err)
	assert.Equal(t, !original.ComputerInventoryCollectionPreferences.MonitorApplicationUsage, updated.ComputerInventoryCollectionPreferences.MonitorApplicationUsage)

	restoreReq := &computer_inventory_collection_settings.ResourceComputerInventoryCollectionSettings{
		ComputerInventoryCollectionPreferences: original.ComputerInventoryCollectionPreferences,
	}
	_, err = svc.UpdateV2(ctx, restoreReq)
	require.NoError(t, err)
}

func TestComputerInventoryCollectionSettings_CustomPath(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.ComputerInventoryCollectionSettings

	createReq := &computer_inventory_collection_settings.CustomPathRequest{
		Scope: "USER_LIBRARY",
		Path:  "/Library/TestPath",
	}

	created, _, err := svc.CreateCustomPathV2(ctx, createReq)
	if err != nil {
		t.Skipf("Failed to create custom path (may not be supported): %v", err)
		return
	}
	require.NotEmpty(t, created.ID)

	acc.Cleanup(t, func() {
		svc.DeleteCustomPathByIDV2(ctx, created.ID)
	})

	settings, _, err := svc.GetV2(ctx)
	require.NoError(t, err)

	found := false
	for _, path := range settings.ApplicationPaths {
		if path.ID == created.ID {
			found = true
			assert.Equal(t, "/Library/TestPath", path.Path)
			break
		}
	}
	assert.True(t, found, "Created custom path should appear in settings")

	_, err = svc.DeleteCustomPathByIDV2(ctx, created.ID)
	require.NoError(t, err)
}
