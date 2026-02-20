package classic_api

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ibeacons"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_IBeacons_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_IBeacons_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.IBeacons
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test iBeacon")

	beaconName := uniqueName("acc-test-ibeacon")
	createReq := &ibeacons.RequestIBeacon{
		Name:  beaconName,
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 1,
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateIBeacon(ctx1, createReq)
	require.NoError(t, err, "CreateIBeacon should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created iBeacon ID should be a positive integer")

	beaconID := created.ID
	acc.LogTestSuccess(t, "iBeacon created with ID=%d name=%q", beaconID, beaconName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteIBeaconByID(cleanupCtx, beaconID)
		acc.LogCleanupDeleteError(t, "iBeacon", fmt.Sprintf("%d", beaconID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new iBeacon appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing iBeacons to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListIBeacons(ctx2)
	require.NoError(t, err, "ListIBeacons should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, b := range list.Results {
		if b.ID == beaconID {
			found = true
			assert.Equal(t, beaconName, b.Name)
			break
		}
	}
	assert.True(t, found, "newly created iBeacon should appear in list")
	acc.LogTestSuccess(t, "iBeacon ID=%d found in list (%d total)", beaconID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching iBeacon by ID=%d", beaconID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetIBeaconByID(ctx3, beaconID)
	require.NoError(t, err, "GetIBeaconByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, beaconID, fetched.ID)
	assert.Equal(t, beaconName, fetched.Name)
	assert.Equal(t, "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0", strings.ToUpper(fetched.UUID))
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q uuid=%s", fetched.ID, fetched.Name, fetched.UUID)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching iBeacon by name=%q", beaconName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetIBeaconByName(ctx4, beaconName)
	require.NoError(t, err, "GetIBeaconByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, beaconID, fetchedByName.ID)
	assert.Equal(t, beaconName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-ibeacon-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating iBeacon ID=%d to name=%q", beaconID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &ibeacons.RequestIBeacon{
		Name:  updatedName,
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 2,
	}
	updated, updateResp, err := svc.UpdateIBeaconByID(ctx5, beaconID, updateReq)
	require.NoError(t, err, "UpdateIBeaconByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating iBeacon name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &ibeacons.RequestIBeacon{
		Name:  beaconName,
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 1,
	}
	reverted, revertResp, err := svc.UpdateIBeaconByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateIBeaconByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetIBeaconByID(ctx7, beaconID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, beaconName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting iBeacon ID=%d", beaconID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteIBeaconByID(ctx8, beaconID)
	require.NoError(t, err, "DeleteIBeaconByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "iBeacon ID=%d deleted", beaconID)
}

// =============================================================================
// TestAcceptance_IBeacons_DeleteByName creates an iBeacon then deletes by name.
// =============================================================================

func TestAcceptance_IBeacons_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.IBeacons
	ctx := context.Background()

	beaconName := uniqueName("acc-test-ibeacon-dbn")
	createReq := &ibeacons.RequestIBeacon{
		Name:  beaconName,
		UUID:  "F7826DA6-4FA2-4E98-8024-BC5B71E0893E",
		Major: 2,
		Minor: 1,
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateIBeacon(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	beaconID := created.ID
	acc.LogTestSuccess(t, "Created iBeacon ID=%d name=%q for delete-by-name test", beaconID, beaconName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteIBeaconByID(cleanupCtx, beaconID)
		acc.LogCleanupDeleteError(t, "iBeacon", fmt.Sprintf("%d", beaconID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteIBeaconByName(ctx2, beaconName)
	require.NoError(t, err, "DeleteIBeaconByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "iBeacon %q deleted by name", beaconName)
}

// =============================================================================
// TestAcceptance_IBeacons_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_IBeacons_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.IBeacons

	t.Run("GetIBeaconByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetIBeaconByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "iBeacon ID must be a positive integer")
	})

	t.Run("GetIBeaconByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetIBeaconByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "iBeacon name is required")
	})

	t.Run("CreateIBeacon_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateIBeacon(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateIBeaconByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateIBeaconByID(context.Background(), 0, &ibeacons.RequestIBeacon{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "iBeacon ID must be a positive integer")
	})

	t.Run("UpdateIBeaconByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateIBeaconByName(context.Background(), "", &ibeacons.RequestIBeacon{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "iBeacon name is required")
	})

	t.Run("DeleteIBeaconByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteIBeaconByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "iBeacon ID must be a positive integer")
	})

	t.Run("DeleteIBeaconByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteIBeaconByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "iBeacon name is required")
	})
}
