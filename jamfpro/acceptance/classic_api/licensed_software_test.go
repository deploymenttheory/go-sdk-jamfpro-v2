package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/licensed_software"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_LicensedSoftware_lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_LicensedSoftware_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicLicensedSoftware
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test licensed software")

	lsName := acc.UniqueName("sdkv2_acc_acc-test-licensed-software")
	createReq := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name:      lsName,
			Publisher: "Acceptance Test Publisher",
			Platform:  "Mac",
			Site:      shared.SharedResourceSite{ID: -1, Name: "None"},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created licensed software ID should be a positive integer")

	lsID := created.ID
	acc.LogTestSuccess(t, "Licensed software created with ID=%d name=%q", lsID, lsName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, lsID)
		acc.LogCleanupDeleteError(t, "licensed software", fmt.Sprintf("%d", lsID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new licensed software appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing licensed software to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, len(list.Results), "results should not be empty")

	found := false
	for _, ls := range list.Results {
		if ls.ID == lsID {
			found = true
			assert.Equal(t, lsName, ls.Name)
			break
		}
	}
	assert.True(t, found, "newly created licensed software should appear in list")
	acc.LogTestSuccess(t, "Licensed software ID=%d found in list (%d total)", lsID, len(list.Results))

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching licensed software by ID=%d", lsID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, lsID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, lsID, fetched.General.ID)
	assert.Equal(t, lsName, fetched.General.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching licensed software by name=%q", lsName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, lsName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, lsID, fetchedByName.General.ID)
	assert.Equal(t, lsName, fetchedByName.General.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.General.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("sdkv2_acc_acc-test-licensed-software-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating licensed software ID=%d to name=%q", lsID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name:      updatedName,
			Publisher: "Acceptance Test Publisher Updated",
			Platform:  "Mac",
			Site:      shared.SharedResourceSite{ID: -1, Name: "None"},
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx5, lsID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating licensed software name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name:      lsName,
			Publisher: "Acceptance Test Publisher",
			Platform:  "Mac",
			Site:      shared.SharedResourceSite{ID: -1, Name: "None"},
		},
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

	verified, verifyResp, err := svc.GetByID(ctx7, lsID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, lsName, verified.General.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.General.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting licensed software ID=%d", lsID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByID(ctx8, lsID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Licensed software ID=%d deleted", lsID)
}

// =============================================================================
// TestAcceptance_LicensedSoftware_delete_by_name creates licensed software then deletes by name.
// =============================================================================

func TestAcceptance_LicensedSoftware_delete_by_name(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicLicensedSoftware
	ctx := context.Background()

	lsName := acc.UniqueName("sdkv2_acc_acc-test-licensed-software-dbn")
	createReq := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name:      lsName,
			Publisher: "Acceptance Test",
			Platform:  "Mac",
			Site:      shared.SharedResourceSite{ID: -1, Name: "None"},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	lsID := created.ID
	acc.LogTestSuccess(t, "Created licensed software ID=%d name=%q for delete-by-name test", lsID, lsName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, lsID)
		acc.LogCleanupDeleteError(t, "licensed software", fmt.Sprintf("%d", lsID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, lsName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Licensed software %q deleted by name", lsName)
}

// =============================================================================
// TestAcceptance_LicensedSoftware_validation_errors validates error handling.
// =============================================================================

func TestAcceptance_LicensedSoftware_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicLicensedSoftware

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "licensed software ID must be a positive integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "licensed software name cannot be empty")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 0, &licensed_software.Resource{
			General: licensed_software.SubsetGeneral{Name: "sdkv2_acc_test", Publisher: "Test", Platform: "Mac"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "licensed software ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &licensed_software.Resource{
			General: licensed_software.SubsetGeneral{Name: "sdkv2_acc_x", Publisher: "Test", Platform: "Mac"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "licensed software name cannot be empty")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "licensed software ID must be a positive integer")
	})
}
