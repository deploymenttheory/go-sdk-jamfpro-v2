package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computers"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_Computers_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
//
// Note: Creating a computer via Classic API may require a valid serial number
// and MAC address. Adjust or skip if your Jamf instance restricts creation.
// =============================================================================

func TestAcceptance_Computers_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputers
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test computer")

	computerName := acc.UniqueName("sdkv2_acc_acc-test-computer")
	createReq := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			Name:         computerName,
			MacAddress:   "00:11:22:33:44:55",
			SerialNumber: fmt.Sprintf("ACC%d", time.Now().UnixMilli()),
			Site: shared.SharedResourceSite{
				ID:   -1,
				Name: "none",
			},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	if err != nil {
		t.Skipf("Create computer skipped (may require valid serial/MAC): %v", err)
		return
	}
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.General.ID, "created computer ID should be a positive integer")

	computerID := fmt.Sprintf("%d", created.General.ID)
	acc.LogTestSuccess(t, "Computer created with ID=%s name=%q", computerID, computerName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, computerID)
		acc.LogCleanupDeleteError(t, "computer", computerID, delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new computer appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing computers to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, c := range list.Results {
		if fmt.Sprintf("%d", c.ID) == computerID {
			found = true
			assert.Equal(t, computerName, c.Name)
			break
		}
	}
	assert.True(t, found, "newly created computer should appear in list")
	acc.LogTestSuccess(t, "Computer ID=%s found in list (%d total)", computerID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching computer by ID=%s", computerID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, computerID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, created.General.ID, fetched.General.ID)
	assert.Equal(t, computerName, fetched.General.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching computer by name=%q", computerName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, computerName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, created.General.ID, fetchedByName.General.ID)
	assert.Equal(t, computerName, fetchedByName.General.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.General.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("sdkv2_acc_acc-test-computer-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating computer ID=%s to name=%q", computerID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			ID:            created.General.ID,
			Name:          updatedName,
			MacAddress:    fetched.General.MacAddress,
			SerialNumber:  fetched.General.SerialNumber,
			Site:          shared.SharedResourceSite{ID: -1, Name: "none"},
		},
		Location:    fetched.Location,
		Purchasing:  fetched.Purchasing,
		Peripherals: fetched.Peripherals,
		Hardware:    fetched.Hardware,
		Security:    fetched.Security,
		Software:    fetched.Software,
		GroupsAccounts: fetched.GroupsAccounts,
	}
	updated, updateResp, err := svc.UpdateByID(ctx5, computerID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating computer name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			ID:            created.General.ID,
			Name:          computerName,
			MacAddress:    fetched.General.MacAddress,
			SerialNumber:  fetched.General.SerialNumber,
			Site:          shared.SharedResourceSite{ID: -1, Name: "none"},
		},
		Location:    fetched.Location,
		Purchasing:  fetched.Purchasing,
		Peripherals: fetched.Peripherals,
		Hardware:    fetched.Hardware,
		Security:    fetched.Security,
		Software:    fetched.Software,
		GroupsAccounts: fetched.GroupsAccounts,
	}
	reverted, revertResp, err := svc.UpdateByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetByID(ctx7, computerID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, computerName, verified.General.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.General.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting computer ID=%s", computerID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByID(ctx8, computerID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Computer ID=%s deleted", computerID)
}

// =============================================================================
// TestAcceptance_Computers_DeleteByName creates a computer then deletes by name.
// =============================================================================

func TestAcceptance_Computers_delete_by_name(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputers
	ctx := context.Background()

	computerName := acc.UniqueName("sdkv2_acc_acc-test-computer-dbn")
	createReq := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			Name:         computerName,
			MacAddress:   "00:11:22:33:44:66",
			SerialNumber: fmt.Sprintf("ACC%d", time.Now().UnixMilli()),
			Site:         shared.SharedResourceSite{ID: -1, Name: "none"},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	if err != nil {
		t.Skipf("Create computer skipped (may require valid serial/MAC): %v", err)
		return
	}
	require.NotNil(t, created)

	computerID := fmt.Sprintf("%d", created.General.ID)
	acc.LogTestSuccess(t, "Created computer ID=%s name=%q for delete-by-name test", computerID, computerName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, computerID)
		acc.LogCleanupDeleteError(t, "computer", computerID, delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, computerName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Computer %q deleted by name", computerName)
}

// =============================================================================
// TestAcceptance_Computers_ValidationErrors validates error handling.
// =============================================================================

func TestAcceptance_Computers_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputers

	t.Run("GetByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer ID cannot be empty")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer name cannot be empty")
	})

	t.Run("Create_NilComputer", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer is required")
	})

	t.Run("UpdateByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), "", &computers.ResponseComputer{General: computers.ComputerSubsetGeneral{Name: "test"}})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer ID cannot be empty")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &computers.ResponseComputer{General: computers.ComputerSubsetGeneral{Name: "x"}})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer name cannot be empty")
	})

	t.Run("DeleteByID_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer ID cannot be empty")
	})
}
