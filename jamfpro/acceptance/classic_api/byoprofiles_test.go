package classic_api

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/byoprofiles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_BYOProfiles_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_BYOProfiles_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.BYOProfiles
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test BYO profile")

	profileName := uniqueName("acc-test-byoprofile")
	createReq := &byoprofiles.RequestBYOProfile{
		General: byoprofiles.GeneralSettings{
			Name:        profileName,
			Enabled:     true,
			Description: "Acceptance test BYO profile",
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateBYOProfile(ctx1, createReq)
	if err != nil {
		var apiErr *client.APIError
		if errors.As(err, &apiErr) && apiErr.StatusCode == 409 && strings.Contains(apiErr.Message, "Unable to update the database") {
			t.Skip("BYO profile create returned 409 in this environment; skipping lifecycle")
		}
		require.NoError(t, err, "CreateBYOProfile should not return an error")
	}
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created BYO profile ID should be a positive integer")

	profileID := created.ID
	acc.LogTestSuccess(t, "BYO profile created with ID=%d name=%q", profileID, profileName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteBYOProfileByID(cleanupCtx, profileID)
		acc.LogCleanupDeleteError(t, "BYO profile", fmt.Sprintf("%d", profileID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new BYO profile appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing BYO profiles to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListBYOProfiles(ctx2)
	require.NoError(t, err, "ListBYOProfiles should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, p := range list.Results {
		if p.ID == profileID {
			found = true
			assert.Equal(t, profileName, p.Name)
			break
		}
	}
	assert.True(t, found, "newly created BYO profile should appear in list")
	acc.LogTestSuccess(t, "BYO profile ID=%d found in list (%d total)", profileID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching BYO profile by ID=%d", profileID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetBYOProfileByID(ctx3, profileID)
	require.NoError(t, err, "GetBYOProfileByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, profileID, fetched.ID)
	assert.Equal(t, profileName, fetched.General.Name)
	assert.Equal(t, "Acceptance test BYO profile", fetched.General.Description)
	assert.True(t, fetched.General.Enabled)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching BYO profile by name=%q", profileName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetBYOProfileByName(ctx4, profileName)
	require.NoError(t, err, "GetBYOProfileByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, profileID, fetchedByName.ID)
	assert.Equal(t, profileName, fetchedByName.General.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-byoprofile-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating BYO profile ID=%d to name=%q", profileID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &byoprofiles.RequestBYOProfile{
		General: byoprofiles.GeneralSettings{
			Name:        updatedName,
			Enabled:     false,
			Description: "Updated description",
		},
	}
	updated, updateResp, err := svc.UpdateBYOProfileByID(ctx5, profileID, updateReq)
	require.NoError(t, err, "UpdateBYOProfileByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating BYO profile name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &byoprofiles.RequestBYOProfile{
		General: byoprofiles.GeneralSettings{
			Name:        profileName,
			Enabled:     true,
			Description: "Reverted description",
		},
	}
	reverted, revertResp, err := svc.UpdateBYOProfileByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateBYOProfileByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetBYOProfileByID(ctx7, profileID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, profileName, verified.General.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.General.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting BYO profile ID=%d", profileID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteBYOProfileByID(ctx8, profileID)
	require.NoError(t, err, "DeleteBYOProfileByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "BYO profile ID=%d deleted", profileID)
}

// =============================================================================
// TestAcceptance_BYOProfiles_DeleteByName creates a BYO profile then deletes by name.
// =============================================================================

func TestAcceptance_BYOProfiles_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.BYOProfiles
	ctx := context.Background()

	profileName := uniqueName("acc-test-byoprofile-del")
	createReq := &byoprofiles.RequestBYOProfile{
		General: byoprofiles.GeneralSettings{
			Name:        profileName,
			Enabled:     true,
			Description: "Test BYO profile for delete by name",
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateBYOProfile(ctx1, createReq)
	if err != nil {
		var apiErr *client.APIError
		if errors.As(err, &apiErr) && apiErr.StatusCode == 409 && strings.Contains(apiErr.Message, "Unable to update the database") {
			t.Skip("BYO profile create returned 409 in this environment; skipping delete-by-name")
		}
		require.NoError(t, err)
	}
	require.NotNil(t, created)

	profileID := created.ID
	acc.LogTestSuccess(t, "Created BYO profile ID=%d name=%q for delete-by-name test", profileID, profileName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteBYOProfileByID(cleanupCtx, profileID)
		acc.LogCleanupDeleteError(t, "BYO profile", fmt.Sprintf("%d", profileID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteBYOProfileByName(ctx2, profileName)
	require.NoError(t, err, "DeleteBYOProfileByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "BYO profile %q deleted by name", profileName)
}

// =============================================================================
// TestAcceptance_BYOProfiles_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_BYOProfiles_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.BYOProfiles

	t.Run("GetBYOProfileByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetBYOProfileByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
	})

	t.Run("GetBYOProfileByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetBYOProfileByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "BYO profile name is required")
	})

	t.Run("CreateBYOProfile_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateBYOProfile(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateBYOProfileByID_ZeroID", func(t *testing.T) {
		req := &byoprofiles.RequestBYOProfile{
			General: byoprofiles.GeneralSettings{Name: "x"},
		}
		_, _, err := svc.UpdateBYOProfileByID(context.Background(), 0, req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
	})

	t.Run("UpdateBYOProfileByName_EmptyName", func(t *testing.T) {
		req := &byoprofiles.RequestBYOProfile{
			General: byoprofiles.GeneralSettings{Name: "x"},
		}
		_, _, err := svc.UpdateBYOProfileByName(context.Background(), "", req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "BYO profile name is required")
	})

	t.Run("DeleteBYOProfileByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteBYOProfileByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
	})

	t.Run("DeleteBYOProfileByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteBYOProfileByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "BYO profile name is required")
	})
}
