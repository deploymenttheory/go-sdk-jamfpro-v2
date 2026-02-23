package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/file_share_distribution_points"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_FileShareDistributionPoints_Lifecycle exercises the full
// write/read/delete lifecycle: Create → List → GetByID → GetByName →
// UpdateByID → UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_FileShareDistributionPoints_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicFileShareDistributionPoints
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test file share distribution point")

	dpName := acc.UniqueName("acc-test-fsdp")
	createReq := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name:                  dpName,
		IsMaster:              false,
		LocalPath:             "/path/to/share",
		ConnectionType:        "SMB",
		ShareName:             "JamfShare",
		SharePort:             445,
		HTTPDownloadsEnabled:  true,
		HTTPURL:               "http://192.168.1.100:8080",
		NoAuthenticationRequired: false,
		UsernamePasswordRequired:  true,
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created distribution point ID should be a positive integer")

	dpID := created.ID
	acc.LogTestSuccess(t, "File share distribution point created with ID=%d name=%q", dpID, dpName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, dpID)
		acc.LogCleanupDeleteError(t, "file share distribution point", fmt.Sprintf("%d", dpID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new distribution point appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing file share distribution points to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, d := range list.Results {
		if d.ID == dpID {
			found = true
			assert.Equal(t, dpName, d.Name)
			break
		}
	}
	assert.True(t, found, "newly created distribution point should appear in list")
	acc.LogTestSuccess(t, "Distribution point ID=%d found in list (%d total)", dpID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching distribution point by ID=%d", dpID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, dpID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, dpID, fetched.ID)
	assert.Equal(t, dpName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching distribution point by name=%q", dpName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, dpName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, dpID, fetchedByName.ID)
	assert.Equal(t, dpName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("acc-test-fsdp-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating distribution point ID=%d to name=%q", dpID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name:                     updatedName,
		IsMaster:                 false,
		LocalPath:                "/path/to/share",
		ConnectionType:           "SMB",
		ShareName:                "JamfShare",
		SharePort:                445,
		HTTPDownloadsEnabled:     true,
		HTTPURL:                  "http://192.168.1.100:8080",
		NoAuthenticationRequired: false,
		UsernamePasswordRequired: true,
	}
	updated, updateResp, err := svc.UpdateByID(ctx5, dpID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating distribution point name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name:                     dpName,
		IsMaster:                 false,
		LocalPath:                "/path/to/share",
		ConnectionType:           "SMB",
		ShareName:                "JamfShare",
		SharePort:                445,
		HTTPDownloadsEnabled:     true,
		HTTPURL:                  "http://192.168.1.100:8080",
		NoAuthenticationRequired: false,
		UsernamePasswordRequired: true,
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

	verified, verifyResp, err := svc.GetByID(ctx7, dpID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, dpName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting distribution point ID=%d", dpID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByID(ctx8, dpID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Distribution point ID=%d deleted", dpID)
}

// =============================================================================
// TestAcceptance_FileShareDistributionPoints_DeleteByName creates a
// distribution point then deletes by name.
// =============================================================================

func TestAcceptance_FileShareDistributionPoints_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicFileShareDistributionPoints
	ctx := context.Background()

	dpName := acc.UniqueName("acc-test-fsdp-dbn")
	createReq := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name:     dpName,
		IsMaster: false,
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	dpID := created.ID
	acc.LogTestSuccess(t, "Created distribution point ID=%d name=%q for delete-by-name test", dpID, dpName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, dpID)
		acc.LogCleanupDeleteError(t, "distribution point", fmt.Sprintf("%d", dpID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, dpName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Distribution point %q deleted by name", dpName)
}

// =============================================================================
// TestAcceptance_FileShareDistributionPoints_ValidationErrors validates error handling.
// =============================================================================

func TestAcceptance_FileShareDistributionPoints_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicFileShareDistributionPoints

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "distribution point ID must be a positive integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "distribution point name cannot be empty")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 0, &file_share_distribution_points.RequestFileShareDistributionPoint{Name: "test"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "distribution point ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &file_share_distribution_points.RequestFileShareDistributionPoint{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "distribution point name cannot be empty")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "distribution point ID must be a positive integer")
	})
}
