package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_software_title_configurations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Patch Software Title Configurations
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   Configuration CRUD (V2 API):
//   • ListV2(ctx) - Lists patch software title configurations
//   • GetByIDV2(ctx, id) - Retrieves a patch software title configuration by ID
//   • GetByNameV2(ctx, name) - Retrieves a patch software title configuration by display name (helper)
//   • CreateV2(ctx, request) - Creates a new patch software title configuration
//   • UpdateByIDV2(ctx, id, request) - Updates an existing patch software title configuration
//   • UpdateByNameV2(ctx, name, request) - Updates a patch software title configuration by display name (helper)
//   • DeleteByIDV2(ctx, id) - Deletes a patch software title configuration by ID
//   • DeleteByNameV2(ctx, name) - Deletes a patch software title configuration by display name (helper)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_PatchSoftwareTitleConfigurations_Lifecycle
//     -- Flow: Create → GetByID → GetByName → Update → Delete
//
//   ✓ List Operations
//     -- Tests: TestAcceptance_PatchSoftwareTitleConfigurations_ListV2
//     -- Flow: List all configurations → Verify response structure
//
//   Note: RSQL Filter Testing NOT applicable
//     -- ListV2 uses pagination, not RSQL filtering
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (configuration creation with comprehensive settings)
//   ✓ Read operations (GetByID, GetByName, List)
//   ✓ Update operations (full resource update)
//   ✓ Delete operations (single delete by ID)
//   ✓ Cleanup and resource management
//   ✗ Input validation and error handling (not yet tested)
//   ✗ Update by name operations (not yet tested)
//   ✗ Delete by name operations (not yet tested)
//
// Notes
// -----------------------------------------------------------------------------
//   • Patch software title configurations define patching settings for software titles
//   • Requires a valid softwareTitleId - test assumes at least one exists in the environment
//   • Configurations link packages to patch versions for specific software titles
//   • ExtensionAttributes can be used to control patching behavior
//   • UINotifications and EmailNotifications control user/admin notification settings
//   • GetByName, UpdateByName, DeleteByName are helper methods (use ListV2 for lookup)
//   • All tests register cleanup handlers to remove test configurations
//   • Tests use acc.UniqueName() to avoid conflicts in shared test environments
//   • TODO: Add validation error tests for empty IDs, nil requests, etc.
//   • TODO: Add tests for UpdateByNameV2 and DeleteByNameV2 operations
//
// =============================================================================

func TestAcceptance_PatchSoftwareTitleConfigurations_ListV2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.PatchSoftwareTitleConfigurations
	ctx := context.Background()

	result, resp, err := svc.ListV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, len(*result), 0)
}

func TestAcceptance_PatchSoftwareTitleConfigurations_Lifecycle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.PatchSoftwareTitleConfigurations
	ctx := context.Background()
	name := acc.UniqueName("acc-patch-config")

	// Get existing configurations to find a valid software title ID
	acc.LogTestStage(t, "Pre-check", "Finding valid software title ID")
	existingConfigs, _, err := svc.ListV2(ctx)
	require.NoError(t, err)

	var softwareTitleID string
	if len(*existingConfigs) > 0 {
		softwareTitleID = (*existingConfigs)[0].SoftwareTitleID
		acc.LogTestStage(t, "Pre-check", "Using software title ID: %s", softwareTitleID)
	} else {
		t.Skip("No existing patch software title configurations found - cannot determine valid software title ID")
	}

	// Create
	acc.LogTestStage(t, "Create", "Creating patch software title configuration")
	created, resp, err := svc.CreateV2(ctx, &patch_software_title_configurations.ResourcePatchSoftwareTitleConfiguration{
		DisplayName:        name,
		SoftwareTitleID:    softwareTitleID,
		UINotifications:    true,
		EmailNotifications: false,
		ExtensionAttributes: []patch_software_title_configurations.SubsetExtensionAttribute{
			{
				Accepted: true,
				EAID:     "1",
			},
		},
	})
	require.NoError(t, err, "Failed to create patch software title configuration")
	require.NotNil(t, created)
	require.NotEmpty(t, created.ID)
	assert.Contains(t, created.Href, created.ID)
	assert.Equal(t, 200, resp.StatusCode)
	acc.Cleanup(t, func() {
		acc.LogTestStage(t, "Cleanup", "Deleting patch software title configuration ID: %s", created.ID)
		_, _ = svc.DeleteByIDV2(ctx, created.ID)
	})

	// GetByID
	acc.LogTestStage(t, "Read", "Getting patch software title configuration by ID: %s", created.ID)
	retrieved, resp, err := svc.GetByIDV2(ctx, created.ID)
	require.NoError(t, err, "Failed to get patch software title configuration by ID")
	require.NotNil(t, retrieved)
	assert.Equal(t, created.ID, retrieved.ID)
	assert.Equal(t, name, retrieved.DisplayName)
	assert.Equal(t, softwareTitleID, retrieved.SoftwareTitleID)
	assert.True(t, retrieved.UINotifications)
	assert.False(t, retrieved.EmailNotifications)
	assert.Len(t, retrieved.ExtensionAttributes, 1)
	assert.Equal(t, 200, resp.StatusCode)

	// GetByName
	acc.LogTestStage(t, "Read", "Getting patch software title configuration by name: %s", name)
	retrievedByName, resp, err := svc.GetByNameV2(ctx, name)
	require.NoError(t, err, "Failed to get patch software title configuration by name")
	require.NotNil(t, retrievedByName)
	assert.Equal(t, created.ID, retrievedByName.ID)
	assert.Equal(t, name, retrievedByName.DisplayName)
	assert.Equal(t, 200, resp.StatusCode)

	// Update
	acc.LogTestStage(t, "Update", "Updating patch software title configuration ID: %s", created.ID)
	updated, resp, err := svc.UpdateByIDV2(ctx, created.ID, &patch_software_title_configurations.ResourcePatchSoftwareTitleConfiguration{
		DisplayName:        name,
		SoftwareTitleID:    softwareTitleID,
		UINotifications:    false,
		EmailNotifications: true,
		ExtensionAttributes: []patch_software_title_configurations.SubsetExtensionAttribute{
			{
				Accepted: false,
				EAID:     "1",
			},
		},
	})
	require.NoError(t, err, "Failed to update patch software title configuration")
	require.NotNil(t, updated)
	assert.Equal(t, created.ID, updated.ID)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify update
	acc.LogTestStage(t, "Verify", "Verifying patch software title configuration update")
	updatedRetrieved, resp, err := svc.GetByIDV2(ctx, created.ID)
	require.NoError(t, err, "Failed to get updated patch software title configuration")
	require.NotNil(t, updatedRetrieved)
	assert.Equal(t, created.ID, updatedRetrieved.ID)
	assert.Equal(t, 200, resp.StatusCode)

	// Delete
	acc.LogTestStage(t, "Delete", "Deleting patch software title configuration ID: %s", created.ID)
	resp, err = svc.DeleteByIDV2(ctx, created.ID)
	require.NoError(t, err, "Failed to delete patch software title configuration")
	assert.Equal(t, 200, resp.StatusCode)

	// Verify deletion
	acc.LogTestStage(t, "Verify", "Verifying patch software title configuration deletion")
	_, resp, err = svc.GetByIDV2(ctx, created.ID)
	assert.Error(t, err, "Expected error when getting deleted patch software title configuration")
	if resp != nil {
		assert.Equal(t, 404, resp.StatusCode)
	}
}
