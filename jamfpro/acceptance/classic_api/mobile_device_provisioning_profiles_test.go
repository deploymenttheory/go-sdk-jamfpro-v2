package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_provisioning_profiles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_MobileDeviceProvisioningProfiles_Lifecycle exercises the full
// write/read/delete lifecycle: CreateByID → List → GetByID → GetByName →
// GetByUUID → UpdateByID → UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_MobileDeviceProvisioningProfiles_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMobileDeviceProvisioningProfiles
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test mobile device provisioning profile")

	profileName := acc.UniqueName("sdkv2_acc_acc-test-md-provisioning-profile")
	profileUUID := "550e8400-e29b-41d4-a716-" + fmt.Sprintf("%012d", time.Now().Unix()%1000000000000)
	createReq := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name:        profileName,
			DisplayName: profileName,
			UUID:        profileUUID,
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateByID(ctx1, 0, createReq)
	require.NoError(t, err, "CreateByID should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.General.ID, "created profile ID should be a positive integer")

	profileID := created.General.ID
	acc.LogTestSuccess(t, "mobile device provisioning profile created with ID=%d name=%q", profileID, profileName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, profileID)
		acc.LogCleanupDeleteError(t, "mobile device provisioning profile", fmt.Sprintf("%d", profileID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new profile appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing mobile device provisioning profiles to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, p := range list.Profiles {
		if p.ID == profileID {
			found = true
			assert.Equal(t, profileName, p.Name)
			break
		}
	}
	assert.True(t, found, "newly created profile should appear in list")
	acc.LogTestSuccess(t, "Profile ID=%d found in list (%d total)", profileID, len(list.Profiles))

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching profile by ID=%d", profileID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, profileID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, profileID, fetched.General.ID)
	assert.Equal(t, profileName, fetched.General.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching profile by name=%q", profileName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, profileName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, profileID, fetchedByName.General.ID)
	assert.Equal(t, profileName, fetchedByName.General.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.General.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 5. GetByUUID (if UUID is set)
	// ------------------------------------------------------------------
	if fetched.General.UUID != "" {
		acc.LogTestStage(t, "GetByUUID", "Fetching profile by UUID=%q", fetched.General.UUID)

		ctx5a, cancel5a := context.WithTimeout(ctx, acc.Config.RequestTimeout)
		defer cancel5a()

		fetchedByUUID, uuidResp, err := svc.GetByUUID(ctx5a, fetched.General.UUID)
		require.NoError(t, err, "GetByUUID should not return an error")
		require.NotNil(t, fetchedByUUID)
		assert.Equal(t, 200, uuidResp.StatusCode)
		assert.Equal(t, profileID, fetchedByUUID.General.ID)
		acc.LogTestSuccess(t, "GetByUUID: ID=%d", fetchedByUUID.General.ID)
	}

	// ------------------------------------------------------------------
	// 6. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("sdkv2_acc_acc-test-md-provisioning-profile-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating profile ID=%d to name=%q", profileID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name:        updatedName,
			DisplayName: updatedName,
			UUID:        fetched.General.UUID,
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx5, profileID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating profile name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name:        profileName,
			DisplayName: profileName,
			UUID:        fetched.General.UUID,
		},
	}
	reverted, revertResp, err := svc.UpdateByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 8. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetByID(ctx7, profileID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, profileName, verified.General.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.General.Name)

	// ------------------------------------------------------------------
	// 9. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting profile ID=%d", profileID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByID(ctx8, profileID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Profile ID=%d deleted", profileID)
}

// =============================================================================
// TestAcceptance_MobileDeviceProvisioningProfiles_DeleteByName creates a profile
// then deletes by name.
// =============================================================================

func TestAcceptance_MobileDeviceProvisioningProfiles_delete_by_name(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMobileDeviceProvisioningProfiles
	ctx := context.Background()

	profileName := acc.UniqueName("sdkv2_acc_acc-test-md-provisioning-profile-dbn")
	profileUUID := "550e8400-e29b-41d4-a716-" + fmt.Sprintf("%012d", time.Now().Unix()%1000000000000)
	createReq := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name:        profileName,
			DisplayName: profileName,
			UUID:        profileUUID,
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateByID(ctx1, 0, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	profileID := created.General.ID
	acc.LogTestSuccess(t, "Created profile ID=%d name=%q for delete-by-name test", profileID, profileName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, profileID)
		acc.LogCleanupDeleteError(t, "mobile device provisioning profile", fmt.Sprintf("%d", profileID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, profileName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Profile %q deleted by name", profileName)
}

// =============================================================================
// TestAcceptance_MobileDeviceProvisioningProfiles_ValidationErrors validates error handling.
// =============================================================================

func TestAcceptance_MobileDeviceProvisioningProfiles_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMobileDeviceProvisioningProfiles

	t.Run("GetByID_NegativeID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), -1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device provisioning profile ID must be a non-negative integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device provisioning profile name cannot be empty")
	})

	t.Run("CreateByID_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateByID(context.Background(), 0, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 0, &mobile_device_provisioning_profiles.RequestResource{
			General: mobile_device_provisioning_profiles.SubsetGeneral{Name: "test"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device provisioning profile ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &mobile_device_provisioning_profiles.RequestResource{
			General: mobile_device_provisioning_profiles.SubsetGeneral{Name: "x"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device provisioning profile name cannot be empty")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device provisioning profile ID must be a positive integer")
	})
}
