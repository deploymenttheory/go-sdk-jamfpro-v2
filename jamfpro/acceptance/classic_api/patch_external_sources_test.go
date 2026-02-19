package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/patch_external_sources"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_PatchExternalSources_Lifecycle exercises the full
// write/read/delete lifecycle: Create → List → GetByID → GetByName →
// UpdateByID → UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_PatchExternalSources_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.PatchExternalSources
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test patch external source")

	sourceName := uniqueName("acc-test-patchsrc")
	createReq := &patch_external_sources.RequestPatchExternalSource{
		Name:       sourceName,
		HostName:   "patches.example.com",
		SSLEnabled: true,
		Port:       443,
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreatePatchExternalSource(ctx1, createReq)
	require.NoError(t, err, "CreatePatchExternalSource should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created patch external source ID should be a positive integer")

	sourceID := created.ID
	acc.LogTestSuccess(t, "Patch external source created with ID=%d name=%q", sourceID, sourceName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeletePatchExternalSourceByID(cleanupCtx, sourceID)
		acc.LogCleanupDeleteError(t, "patch external source", fmt.Sprintf("%d", sourceID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new source appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing patch external sources to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListPatchExternalSources(ctx2)
	require.NoError(t, err, "ListPatchExternalSources should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, s := range list.Results {
		if s.ID == sourceID {
			found = true
			assert.Equal(t, sourceName, s.Name)
			break
		}
	}
	assert.True(t, found, "newly created patch external source should appear in list")
	acc.LogTestSuccess(t, "Patch external source ID=%d found in list (%d total)", sourceID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching patch external source by ID=%d", sourceID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetPatchExternalSourceByID(ctx3, sourceID)
	require.NoError(t, err, "GetPatchExternalSourceByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, sourceID, fetched.ID)
	assert.Equal(t, sourceName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching patch external source by name=%q", sourceName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetPatchExternalSourceByName(ctx4, sourceName)
	require.NoError(t, err, "GetPatchExternalSourceByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, sourceID, fetchedByName.ID)
	assert.Equal(t, sourceName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-patchsrc-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating patch external source ID=%d to name=%q", sourceID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &patch_external_sources.RequestPatchExternalSource{
		Name:       updatedName,
		HostName:   "patches.example.com",
		SSLEnabled: true,
		Port:       443,
	}
	updated, updateResp, err := svc.UpdatePatchExternalSourceByID(ctx5, sourceID, updateReq)
	require.NoError(t, err, "UpdatePatchExternalSourceByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating patch external source name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &patch_external_sources.RequestPatchExternalSource{
		Name:       sourceName,
		HostName:   "patches.example.com",
		SSLEnabled: true,
		Port:       443,
	}
	reverted, revertResp, err := svc.UpdatePatchExternalSourceByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdatePatchExternalSourceByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetPatchExternalSourceByID(ctx7, sourceID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, sourceName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting patch external source ID=%d", sourceID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeletePatchExternalSourceByID(ctx8, sourceID)
	require.NoError(t, err, "DeletePatchExternalSourceByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Patch external source ID=%d deleted", sourceID)
}

// =============================================================================
// TestAcceptance_PatchExternalSources_ValidationErrors tests client-side validation.
// =============================================================================

func TestAcceptance_PatchExternalSources_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.PatchExternalSources

	t.Run("GetPatchExternalSourceByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetPatchExternalSourceByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
	})

	t.Run("GetPatchExternalSourceByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetPatchExternalSourceByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "patch external source name is required")
	})

	t.Run("CreatePatchExternalSource_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreatePatchExternalSource(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdatePatchExternalSourceByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdatePatchExternalSourceByID(context.Background(), 0, &patch_external_sources.RequestPatchExternalSource{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
	})

	t.Run("UpdatePatchExternalSourceByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdatePatchExternalSourceByName(context.Background(), "", &patch_external_sources.RequestPatchExternalSource{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "patch external source name is required")
	})

	t.Run("DeletePatchExternalSourceByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeletePatchExternalSourceByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
	})
}
