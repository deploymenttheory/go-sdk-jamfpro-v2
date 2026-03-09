package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_configuration_profiles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_MobileDeviceConfigurationProfiles_lifecycle exercises the full
// write/read/delete lifecycle: Create → List → GetByID → GetByName →
// UpdateByID → UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_MobileDeviceConfigurationProfiles_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicAPI.MobileDeviceConfigurationProfiles
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test mobile device configuration profile")

	profileName := acc.UniqueName("sdkv2_acc_acc-test-md-profile")
	createReq := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:             profileName,
			Description:      "Acceptance test profile",
			DeploymentMethod: "Install Automatically",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: &mobile_device_configuration_profiles.SubsetScope{
			AllMobileDevices: true,
			AllJSSUsers:      false,
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode(), "expected 200 or 201")
	assert.Positive(t, created.ID, "created profile ID should be a positive integer")

	profileID := created.ID
	acc.LogTestSuccess(t, "mobile device configuration profile created with ID=%d name=%q", profileID, profileName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, profileID)
		acc.LogCleanupDeleteError(t, "mobile device configuration profile", fmt.Sprintf("%d", profileID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new profile appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing mobile device configuration profiles to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode())

	found := false
	for _, p := range list.ConfigurationProfiles {
		if p.ID == profileID {
			found = true
			assert.Equal(t, profileName, p.Name)
			break
		}
	}
	assert.True(t, found, "newly created profile should appear in list")
	acc.LogTestSuccess(t, "Profile ID=%d found in list (%d total)", profileID, len(list.ConfigurationProfiles))

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Getting profile by ID=%d", profileID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, profileID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode())
	assert.Equal(t, profileID, fetched.General.ID)
	assert.Equal(t, profileName, fetched.General.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Getting profile by name=%q", profileName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, profileName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode())
	assert.Equal(t, profileID, fetchedByName.General.ID)
	assert.Equal(t, profileName, fetchedByName.General.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.General.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 5. GetByIDWithSubset
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByIDWithSubset", "Getting profile by ID=%d with subset General", profileID)

	ctx5a, cancel5a := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5a()

	fetchedSubset, subsetResp, err := svc.GetByIDWithSubset(ctx5a, profileID, "General")
	require.NoError(t, err, "GetByIDWithSubset should not return an error")
	require.NotNil(t, fetchedSubset)
	assert.Equal(t, 200, subsetResp.StatusCode())
	assert.Equal(t, profileID, fetchedSubset.General.ID)
	acc.LogTestSuccess(t, "GetByIDWithSubset: ID=%d", fetchedSubset.General.ID)

	// ------------------------------------------------------------------
	// 6. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("sdkv2_acc_acc-test-md-profile-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating profile ID=%d to name=%q", profileID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:             updatedName,
			Description:      "Updated acceptance test profile",
			DeploymentMethod: "Install Automatically",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: &mobile_device_configuration_profiles.SubsetScope{
			AllMobileDevices: true,
			AllJSSUsers:      false,
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx5, profileID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode(), "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode())

	// ------------------------------------------------------------------
	// 7. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating profile name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:             profileName,
			Description:      "Acceptance test profile",
			DeploymentMethod: "Install Automatically",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: &mobile_device_configuration_profiles.SubsetScope{
			AllMobileDevices: true,
			AllJSSUsers:      false,
		},
	}
	reverted, revertResp, err := svc.UpdateByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode(), "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode())

	// ------------------------------------------------------------------
	// 8. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetByID(ctx7, profileID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode())
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
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode())
	acc.LogTestSuccess(t, "Profile ID=%d deleted", profileID)
}

// =============================================================================
// TestAcceptance_MobileDeviceConfigurationProfiles_delete_by_name creates a profile
// then deletes by name.
// =============================================================================

func TestAcceptance_MobileDeviceConfigurationProfiles_delete_by_name(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicAPI.MobileDeviceConfigurationProfiles
	ctx := context.Background()

	profileName := acc.UniqueName("sdkv2_acc_acc-test-md-profile-dbn")
	createReq := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:             profileName,
			DeploymentMethod: "Install Automatically",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: &mobile_device_configuration_profiles.SubsetScope{
			AllMobileDevices: true,
			AllJSSUsers:      false,
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	profileID := created.ID
	acc.LogTestSuccess(t, "Created profile ID=%d name=%q for delete-by-name test", profileID, profileName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, profileID)
		acc.LogCleanupDeleteError(t, "mobile device configuration profile", fmt.Sprintf("%d", profileID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, profileName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode())
	acc.LogTestSuccess(t, "Profile %q deleted by name", profileName)
}

// =============================================================================
// TestAcceptance_MobileDeviceConfigurationProfiles_validation_errors validates error handling.
// =============================================================================

func TestAcceptance_MobileDeviceConfigurationProfiles_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicAPI.MobileDeviceConfigurationProfiles

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device configuration profile ID must be a positive integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device configuration profile name cannot be empty")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 0, &mobile_device_configuration_profiles.RequestResource{
			General: mobile_device_configuration_profiles.SubsetGeneral{Name: "sdkv2_acc_test"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device configuration profile ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &mobile_device_configuration_profiles.RequestResource{
			General: mobile_device_configuration_profiles.SubsetGeneral{Name: "sdkv2_acc_x"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device configuration profile name cannot be empty")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device configuration profile ID must be a positive integer")
	})
}
