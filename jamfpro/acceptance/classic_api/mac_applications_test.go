package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mac_applications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_MacApplications_lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
//
// Note: Mac applications require VPP configuration. This test may fail if
// no VPP account is configured or if the bundle ID is not available in VPP.
// =============================================================================

func TestAcceptance_MacApplications_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMacApplications
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test Mac application")

	appName := acc.UniqueName("sdkv2_acc_acc-test-macapp")
	createReq := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           appName,
			Version:        "1.0",
			BundleID:       "com.apple.Safari",
			URL:            "https://www.apple.com/safari/",
			DeploymentType: "Install Automatically/Prompt Users to Install",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: mac_applications.SubsetScope{
			AllComputers: boolPtr(true),
			AllJSSUsers:  boolPtr(false),
		},
		SelfService: mac_applications.SubsetSelfService{
			InstallButtonText:      "Install",
			SelfServiceDescription: "Safari web browser",
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created Mac application ID should be a positive integer")

	appID := created.ID
	acc.LogTestSuccess(t, "Mac application created with ID=%d name=%q", appID, appName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, appID)
		acc.LogCleanupDeleteError(t, "mac application", fmt.Sprintf("%d", appID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new Mac application appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing Mac applications to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, a := range list.Results {
		if a.ID == appID {
			found = true
			assert.Equal(t, appName, a.Name)
			break
		}
	}
	assert.True(t, found, "newly created Mac application should appear in list")
	acc.LogTestSuccess(t, "Mac application ID=%d found in list (%d total)", appID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Getting Mac application by ID=%d", appID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, appID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, appID, fetched.General.ID)
	assert.Equal(t, appName, fetched.General.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Getting Mac application by name=%q", appName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, appName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, appID, fetchedByName.General.ID)
	assert.Equal(t, appName, fetchedByName.General.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.General.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 5. GetByIDAndSubset
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByIDAndSubset", "Getting Mac application ID=%d subset=General", appID)

	ctx5a, cancel5a := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5a()

	fetchedSubset, subsetResp, err := svc.GetByIDAndSubset(ctx5a, appID, "General")
	require.NoError(t, err, "GetByIDAndSubset should not return an error")
	require.NotNil(t, fetchedSubset)
	assert.Equal(t, 200, subsetResp.StatusCode)
	assert.Equal(t, appName, fetchedSubset.General.Name)
	acc.LogTestSuccess(t, "GetByIDAndSubset: name=%q", fetchedSubset.General.Name)

	// ------------------------------------------------------------------
	// 6. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("sdkv2_acc_acc-test-macapp-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating Mac application ID=%d to name=%q", appID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           updatedName,
			Version:        "1.1",
			BundleID:       "com.apple.Safari",
			URL:            "https://www.apple.com/safari/",
			DeploymentType: "Install Automatically/Prompt Users to Install",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: mac_applications.SubsetScope{
			AllComputers: boolPtr(true),
			AllJSSUsers:  boolPtr(false),
		},
		SelfService: mac_applications.SubsetSelfService{
			InstallButtonText:      "Install",
			SelfServiceDescription: "Safari web browser (updated)",
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx5, appID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode())

	// ------------------------------------------------------------------
	// 7. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating Mac application name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           appName,
			Version:        "1.0",
			BundleID:       "com.apple.Safari",
			URL:            "https://www.apple.com/safari/",
			DeploymentType: "Install Automatically/Prompt Users to Install",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: mac_applications.SubsetScope{
			AllComputers: boolPtr(true),
			AllJSSUsers:  boolPtr(false),
		},
		SelfService: mac_applications.SubsetSelfService{
			InstallButtonText:      "Install",
			SelfServiceDescription: "Safari web browser",
		},
	}
	reverted, revertResp, err := svc.UpdateByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode())

	// ------------------------------------------------------------------
	// 8. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetByID(ctx7, appID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, appName, verified.General.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.General.Name)

	// ------------------------------------------------------------------
	// 9. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting Mac application ID=%d", appID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByID(ctx8, appID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Mac application ID=%d deleted", appID)
}

// =============================================================================
// TestAcceptance_MacApplications_delete_by_name creates a Mac application then deletes by name.
// =============================================================================

func TestAcceptance_MacApplications_delete_by_name(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMacApplications
	ctx := context.Background()

	appName := acc.UniqueName("sdkv2_acc_acc-test-macapp-dbn")
	createReq := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           appName,
			Version:        "1.0",
			BundleID:       "com.apple.Safari",
			URL:            "https://www.apple.com/safari/",
			DeploymentType: "Install Automatically/Prompt Users to Install",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: mac_applications.SubsetScope{
			AllComputers: boolPtr(true),
			AllJSSUsers:  boolPtr(false),
		},
		SelfService: mac_applications.SubsetSelfService{
			InstallButtonText: "Install",
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	appID := created.ID
	acc.LogTestSuccess(t, "Created Mac application ID=%d name=%q for delete-by-name test", appID, appName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, appID)
		acc.LogCleanupDeleteError(t, "mac application", fmt.Sprintf("%d", appID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, appName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Mac application %q deleted by name", appName)
}

// =============================================================================
// TestAcceptance_MacApplications_validation_errors validates error handling.
// =============================================================================

func TestAcceptance_MacApplications_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMacApplications

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mac application ID must be a positive integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mac application name cannot be empty")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 0, &mac_applications.Resource{
			General: mac_applications.SubsetGeneral{Name: "sdkv2_acc_test", BundleID: "com.test.app"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mac application ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &mac_applications.Resource{
			General: mac_applications.SubsetGeneral{Name: "sdkv2_acc_x", BundleID: "com.test.app"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mac application name cannot be empty")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mac application ID must be a positive integer")
	})
}
