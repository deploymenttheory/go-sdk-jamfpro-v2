package jamf_pro_api

import (
	"context"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dock_items"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Dock Items
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetByIDV1(ctx, id) - Retrieves a dock item by ID
//   • CreateV1(ctx, request) - Creates a new dock item
//   • UpdateByIDV1(ctx, id, request) - Updates an existing dock item
//   • DeleteByIDV1(ctx, id) - Deletes a dock item by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle (without List)
//     -- Reason: Service supports Create, Read (GetByID only), Update, Delete operations
//     -- Tests: TestAcceptance_DockItems_Lifecycle
//     -- Flow: Create → GetByID → Update → Verify → Delete
//     -- Note: No List operation available for this service
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_DockItems_ValidationErrors
//     -- Cases: Empty IDs, nil requests, missing required fields
//
//   Note: RSQL Filter Testing NOT applicable
//     -- No List operation available - service only supports individual CRUD operations
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (single dock item creation)
//   ✓ Read operations (GetByID only - no list available)
//   ✓ Update operations (full resource update)
//   ✓ Delete operations (single delete)
//   ✓ Input validation and error handling
//   ✓ Cleanup and resource management
//
// Notes
// -----------------------------------------------------------------------------
//   • No List operation exists for dock items - only individual CRUD
//   • Dock items define applications/folders displayed in macOS dock
//   • Type options: "App" (application) or "Folder"
//   • Path must be valid macOS application or folder path
//   • All tests register cleanup handlers to remove test dock items
//   • Tests use acc.UniqueName() to avoid conflicts in shared test environments
//   • Update response may not include body - verification done via re-fetch
//
// =============================================================================
// TestAcceptance_DockItems_Lifecycle exercises the full write/read/delete
// lifecycle: Create → GetByID → Update → GetByID (verify) → Delete.
// =============================================================================

func TestAcceptance_DockItems_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DockItems
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test dock item")

	createReq := &dock_items.RequestDockItem{
		Name: acc.UniqueName("acc-test-dock"),
		Path: "/Applications/Safari.app",
		Type: dock_items.TypeApp,
	}
	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err, "CreateDockItemV1 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	dockItemID := created.ID
	acc.LogTestSuccess(t, "Dock item created with ID=%s", dockItemID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, dockItemID)
		acc.LogCleanupDeleteError(t, "dock item", dockItemID, delErr)
	})

	// 2. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching dock item by ID=%s", dockItemID)

	fetched, fetchResp, err := svc.GetByIDV1(ctx, dockItemID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, dockItemID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	assert.Equal(t, createReq.Path, fetched.Path)
	assert.Equal(t, createReq.Type, fetched.Type)
	acc.LogTestSuccess(t, "GetByID: name=%q path=%q type=%q", fetched.Name, fetched.Path, fetched.Type)

	// 3. Update
	acc.LogTestStage(t, "Update", "Updating dock item ID=%s", dockItemID)

	updateReq := &dock_items.RequestDockItem{
		Name: acc.UniqueName("acc-test-dock-updated"),
		Path: "/Applications/Google Chrome.app",
		Type: dock_items.TypeApp,
	}
	updated, updateResp, err := svc.UpdateByIDV1(ctx, dockItemID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	// API may not return body on update; verify via re-fetch below
	acc.LogTestSuccess(t, "Dock item updated: ID=%s", dockItemID)

	// 4. Re-fetch to verify
	fetched2, _, err := svc.GetByIDV1(ctx, dockItemID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Name, fetched2.Name)
	assert.Equal(t, updateReq.Path, fetched2.Path)
	acc.LogTestSuccess(t, "Update verified: name=%q path=%q", fetched2.Name, fetched2.Path)

	// 5. Delete
	acc.LogTestStage(t, "Delete", "Deleting dock item ID=%s", dockItemID)

	deleteResp, err := svc.DeleteByIDV1(ctx, dockItemID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Dock item ID=%s deleted", dockItemID)
}

// =============================================================================
// TestAcceptance_DockItems_ValidationErrors
// =============================================================================

func TestAcceptance_DockItems_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DockItems

	t.Run("GetDockItemByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "dock item ID is required")
	})

	t.Run("CreateDockItemV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateDockItemByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "", &dock_items.RequestDockItem{
			Name: "x", Path: "/path", Type: dock_items.TypeApp,
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteDockItemByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "dock item ID is required")
	})
}
