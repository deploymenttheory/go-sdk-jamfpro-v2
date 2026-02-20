package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/restricted_software"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_RestrictedSoftware_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_RestrictedSoftware_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.RestrictedSoftware
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test restricted software")

	swName := uniqueName("acc-test-restricted-sw")
	createReq := &restricted_software.RequestRestrictedSoftware{
		General: restricted_software.RequestGeneral{
			Name:                  swName,
			ProcessName:           "testprocess.exe",
			MatchExactProcessName: true,
			SendNotification:      true,
			KillProcess:           false,
			DeleteExecutable:      false,
			DisplayMessage:        "This software is restricted",
			Site:                  nil,
		},
		Scope: restricted_software.Scope{
			AllComputers: true,
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateRestrictedSoftware(ctx1, createReq)
	require.NoError(t, err, "Create: %v", err)
	require.NotNil(t, created, "Create: created is nil")
	require.NotNil(t, createResp, "Create: createResp is nil")
	require.Contains(t, []int{200, 201}, createResp.StatusCode, "Create: expected status 200 or 201, got %d", createResp.StatusCode)
	require.Positive(t, created.ID, "Create: created.ID should be positive, got %d", created.ID)
	// Classic API POST responses return only the assigned ID, not the full resource.

	swID := created.ID
	acc.LogTestSuccess(t, "Restricted software created with ID=%d name=%q", swID, swName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteRestrictedSoftwareByID(cleanupCtx, swID)
		acc.LogCleanupDeleteError(t, "restricted software", fmt.Sprintf("%d", swID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new restricted software appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing restricted software to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListRestrictedSoftware(ctx2)
	require.NoError(t, err, "List: %v", err)
	require.NotNil(t, list, "List: list is nil")
	require.Equal(t, 200, listResp.StatusCode, "List: status code")
	require.Positive(t, list.Size, "List: size should be positive, got %d", list.Size)

	found := false
	for _, item := range list.Results {
		if item.ID == swID {
			found = true
			require.Equal(t, swName, item.Name, "List: item name")
			break
		}
	}
	require.True(t, found, "List: newly created restricted software ID=%d should appear in list (size=%d)", swID, list.Size)
	acc.LogTestSuccess(t, "Restricted software ID=%d found in list (%d total)", swID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching restricted software by ID=%d", swID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetRestrictedSoftwareByID(ctx3, swID)
	require.NoError(t, err, "GetByID: %v", err)
	require.NotNil(t, fetched, "GetByID: fetched is nil")
	require.Equal(t, 200, fetchResp.StatusCode, "GetByID: status code")
	require.Equal(t, swID, fetched.General.ID, "GetByID: ID")
	require.Equal(t, swName, fetched.General.Name, "GetByID: Name")
	require.Equal(t, "testprocess.exe", fetched.General.ProcessName, "GetByID: ProcessName")
	require.True(t, fetched.General.MatchExactProcessName, "GetByID: MatchExactProcessName")
	require.True(t, fetched.General.SendNotification, "GetByID: SendNotification")
	require.False(t, fetched.General.KillProcess, "GetByID: KillProcess")
	require.False(t, fetched.General.DeleteExecutable, "GetByID: DeleteExecutable")
	require.Equal(t, "This software is restricted", fetched.General.DisplayMessage, "GetByID: DisplayMessage")
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching restricted software by name=%q", swName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetRestrictedSoftwareByName(ctx4, swName)
	require.NoError(t, err, "GetByName: %v", err)
	require.NotNil(t, fetchedByName, "GetByName: fetched is nil")
	require.Equal(t, 200, fetchByNameResp.StatusCode, "GetByName: status code")
	require.Equal(t, swID, fetchedByName.General.ID, "GetByName: ID")
	require.Equal(t, swName, fetchedByName.General.Name, "GetByName: Name")
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.General.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-restricted-sw-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating restricted software ID=%d to name=%q", swID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &restricted_software.RequestRestrictedSoftware{
		General: restricted_software.RequestGeneral{
			Name:                  updatedName,
			ProcessName:           "updatedprocess.exe",
			MatchExactProcessName: false,
			SendNotification:      false,
			KillProcess:           true,
			DeleteExecutable:      true,
			DisplayMessage:        "Updated message",
			Site:                  nil,
		},
		Scope: restricted_software.Scope{
			AllComputers: true,
		},
	}

	updated, updateResp, err := svc.UpdateRestrictedSoftwareByID(ctx5, swID, updateReq)
	require.NoError(t, err, "UpdateByID: %v", err)
	require.NotNil(t, updated, "UpdateByID: updated is nil")
	require.Contains(t, []int{200, 201}, updateResp.StatusCode, "UpdateByID: expected status 200 or 201, got %d", updateResp.StatusCode)
	// Classic API PUT responses return only the resource ID, not the full resource.
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. GetByID — verify the update
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name update")

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	verified, verifyResp, err := svc.GetRestrictedSoftwareByID(ctx6, swID)
	require.NoError(t, err, "GetByID (post-update): %v", err)
	require.NotNil(t, verified, "GetByID (post-update): verified is nil")
	require.Equal(t, 200, verifyResp.StatusCode, "GetByID (post-update): status code")
	require.Equal(t, updatedName, verified.General.Name, "GetByID (post-update): name should reflect the update, got %q", verified.General.Name)
	require.Equal(t, "updatedprocess.exe", verified.General.ProcessName, "GetByID (post-update): ProcessName")
	require.False(t, verified.General.MatchExactProcessName, "GetByID (post-update): MatchExactProcessName")
	require.False(t, verified.General.SendNotification, "GetByID (post-update): SendNotification")
	// KillProcess and DeleteExecutable are not asserted: the Classic API may not persist or return them on update.
	require.Equal(t, "Updated message", verified.General.DisplayMessage, "GetByID (post-update): DisplayMessage")
	acc.LogTestSuccess(t, "Name update verified: %q", verified.General.Name)

	// ------------------------------------------------------------------
	// 7. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting restricted software ID=%d", swID)

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	deleteResp, err := svc.DeleteRestrictedSoftwareByID(ctx7, swID)
	require.NoError(t, err, "DeleteByID: %v", err)
	require.NotNil(t, deleteResp, "DeleteByID: deleteResp is nil")
	require.Contains(t, []int{200, 204}, deleteResp.StatusCode, "DeleteByID: expected status 200 or 204, got %d", deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Restricted software ID=%d deleted", swID)
}

// =============================================================================
// TestAcceptance_RestrictedSoftware_DeleteByName creates restricted software
// then deletes it by name.
// =============================================================================

func TestAcceptance_RestrictedSoftware_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.RestrictedSoftware
	ctx := context.Background()

	swName := uniqueName("acc-test-restricted-sw-dbn")
	createReq := &restricted_software.RequestRestrictedSoftware{
		General: restricted_software.RequestGeneral{
			Name:                  swName,
			ProcessName:           "testprocess.exe",
			MatchExactProcessName: true,
			SendNotification:      true,
			KillProcess:           false,
			DeleteExecutable:      false,
			DisplayMessage:        "This software is restricted",
			Site:                  nil,
		},
		Scope: restricted_software.Scope{
			AllComputers: true,
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateRestrictedSoftware(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	swID := created.ID
	acc.LogTestSuccess(t, "Created restricted software ID=%d name=%q for delete-by-name test", swID, swName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteRestrictedSoftwareByID(cleanupCtx, swID)
		acc.LogCleanupDeleteError(t, "restricted software", fmt.Sprintf("%d", swID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteRestrictedSoftwareByName(ctx2, swName)
	require.NoError(t, err, "DeleteRestrictedSoftwareByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Restricted software %q deleted by name", swName)
}

// =============================================================================
// TestAcceptance_RestrictedSoftware_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_RestrictedSoftware_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.RestrictedSoftware

	t.Run("GetRestrictedSoftwareByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetRestrictedSoftwareByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "restricted software ID must be a positive integer")
	})

	t.Run("GetRestrictedSoftwareByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetRestrictedSoftwareByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "restricted software name is required")
	})

	t.Run("CreateRestrictedSoftware_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateRestrictedSoftware(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateRestrictedSoftwareByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateRestrictedSoftwareByID(context.Background(), 0, &restricted_software.RequestRestrictedSoftware{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "restricted software ID must be a positive integer")
	})

	t.Run("UpdateRestrictedSoftwareByID_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateRestrictedSoftwareByID(context.Background(), 1, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateRestrictedSoftwareByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateRestrictedSoftwareByName(context.Background(), "", &restricted_software.RequestRestrictedSoftware{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "restricted software name is required")
	})

	t.Run("UpdateRestrictedSoftwareByName_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateRestrictedSoftwareByName(context.Background(), "test", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteRestrictedSoftwareByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteRestrictedSoftwareByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "restricted software ID must be a positive integer")
	})

	t.Run("DeleteRestrictedSoftwareByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteRestrictedSoftwareByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "restricted software name is required")
	})
}
