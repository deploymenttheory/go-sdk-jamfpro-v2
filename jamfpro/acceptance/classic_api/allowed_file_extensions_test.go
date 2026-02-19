package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/allowed_file_extensions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_AllowedFileExtensions_Lifecycle exercises the full
// write/read/delete lifecycle: Create → List → GetByID → GetByExtension →
// DeleteByID.
// Note: the Classic API does not support updating allowed file extensions.
// =============================================================================

func TestAcceptance_AllowedFileExtensions_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AllowedFileExtensions
	ctx := context.Background()

	// Use a unique extension suffix so we don't conflict with existing entries.
	// Extensions must be simple strings (no spaces or special chars).
	extension := fmt.Sprintf("acc%d", time.Now().UnixMilli()%100000)

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating allowed file extension %q", extension)

	createReq := &allowed_file_extensions.RequestAllowedFileExtension{
		Extension: extension,
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateAllowedFileExtension(ctx1, createReq)
	require.NoError(t, err, "CreateAllowedFileExtension should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created ID should be a positive integer")
	assert.Equal(t, extension, created.Extension)

	extID := created.ID
	acc.LogTestSuccess(t, "Allowed file extension created with ID=%d extension=%q", extID, extension)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteAllowedFileExtensionByID(cleanupCtx, extID)
		acc.LogCleanupDeleteError(t, "allowed file extension", fmt.Sprintf("%d", extID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new extension appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing allowed file extensions to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListAllowedFileExtensions(ctx2)
	require.NoError(t, err, "ListAllowedFileExtensions should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, e := range list.Results {
		if e.ID == extID {
			found = true
			assert.Equal(t, extension, e.Extension)
			break
		}
	}
	assert.True(t, found, "newly created allowed file extension should appear in list")
	acc.LogTestSuccess(t, "Extension ID=%d found in list (%d total)", extID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching allowed file extension by ID=%d", extID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetAllowedFileExtensionByID(ctx3, extID)
	require.NoError(t, err, "GetAllowedFileExtensionByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, extID, fetched.ID)
	assert.Equal(t, extension, fetched.Extension)
	acc.LogTestSuccess(t, "GetByID: ID=%d extension=%q", fetched.ID, fetched.Extension)

	// ------------------------------------------------------------------
	// 4. GetByExtension
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByExtension", "Fetching allowed file extension by extension=%q", extension)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByExt, fetchByExtResp, err := svc.GetAllowedFileExtensionByExtension(ctx4, extension)
	require.NoError(t, err, "GetAllowedFileExtensionByExtension should not return an error")
	require.NotNil(t, fetchedByExt)
	assert.Equal(t, 200, fetchByExtResp.StatusCode)
	assert.Equal(t, extID, fetchedByExt.ID)
	assert.Equal(t, extension, fetchedByExt.Extension)
	acc.LogTestSuccess(t, "GetByExtension: ID=%d extension=%q", fetchedByExt.ID, fetchedByExt.Extension)

	// ------------------------------------------------------------------
	// 5. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting allowed file extension ID=%d", extID)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	deleteResp, err := svc.DeleteAllowedFileExtensionByID(ctx5, extID)
	require.NoError(t, err, "DeleteAllowedFileExtensionByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Allowed file extension ID=%d deleted", extID)
}

// =============================================================================
// TestAcceptance_AllowedFileExtensions_ValidationErrors tests client-side
// validation without making any network calls.
// =============================================================================

func TestAcceptance_AllowedFileExtensions_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AllowedFileExtensions

	t.Run("GetAllowedFileExtensionByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetAllowedFileExtensionByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "allowed file extension ID must be a positive integer")
	})

	t.Run("GetAllowedFileExtensionByExtension_Empty", func(t *testing.T) {
		_, _, err := svc.GetAllowedFileExtensionByExtension(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "extension is required")
	})

	t.Run("CreateAllowedFileExtension_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateAllowedFileExtension(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteAllowedFileExtensionByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteAllowedFileExtensionByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "allowed file extension ID must be a positive integer")
	})
}
