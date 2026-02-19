package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dock_items"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func uniqueDockItemName(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixMilli())
}

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
		Name: uniqueDockItemName("acc-test-dock"),
		Path: "/Applications/Safari.app",
		Type: dock_items.TypeApp,
	}
	created, createResp, err := svc.CreateDockItemV1(ctx, createReq)
	require.NoError(t, err, "CreateDockItemV1 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	dockItemID := created.ID
	acc.LogTestSuccess(t, "Dock item created with ID=%s", dockItemID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteDockItemByIDV1(cleanupCtx, dockItemID)
		acc.LogCleanupDeleteError(t, "dock item", dockItemID, delErr)
	})

	// 2. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching dock item by ID=%s", dockItemID)

	fetched, fetchResp, err := svc.GetDockItemByIDV1(ctx, dockItemID)
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
		Name: uniqueDockItemName("acc-test-dock-updated"),
		Path: "/Applications/Google Chrome.app",
		Type: dock_items.TypeApp,
	}
	updated, updateResp, err := svc.UpdateDockItemByIDV1(ctx, dockItemID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, updateReq.Name, updated.Name)
	assert.Equal(t, updateReq.Path, updated.Path)
	acc.LogTestSuccess(t, "Dock item updated: ID=%s", dockItemID)

	// 4. Re-fetch to verify
	fetched2, _, err := svc.GetDockItemByIDV1(ctx, dockItemID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Name, fetched2.Name)
	assert.Equal(t, updateReq.Path, fetched2.Path)
	acc.LogTestSuccess(t, "Update verified: name=%q path=%q", fetched2.Name, fetched2.Path)

	// 5. Delete
	acc.LogTestStage(t, "Delete", "Deleting dock item ID=%s", dockItemID)

	deleteResp, err := svc.DeleteDockItemByIDV1(ctx, dockItemID)
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
		_, _, err := svc.GetDockItemByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "dock item ID is required")
	})

	t.Run("CreateDockItemV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateDockItemV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateDockItemByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateDockItemByIDV1(context.Background(), "", &dock_items.RequestDockItem{
			Name: "x", Path: "/path", Type: dock_items.TypeApp,
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteDockItemByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteDockItemByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "dock item ID is required")
	})
}
