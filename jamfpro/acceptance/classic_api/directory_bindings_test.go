package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/directory_bindings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_DirectoryBindings_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_DirectoryBindings_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DirectoryBindings
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test directory binding")

	bindingName := uniqueName("acc-test-dirbinding")
	createReq := &directory_bindings.RequestDirectoryBinding{
		Name:     bindingName,
		Priority: 1,
		Domain:   "test.example.com",
		Type:     "Active Directory",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateDirectoryBinding(ctx1, createReq)
	require.NoError(t, err, "CreateDirectoryBinding should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created directory binding ID should be a positive integer")

	bindingID := created.ID
	acc.LogTestSuccess(t, "Directory binding created with ID=%d name=%q", bindingID, bindingName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteDirectoryBindingByID(cleanupCtx, bindingID)
		acc.LogCleanupDeleteError(t, "directory binding", fmt.Sprintf("%d", bindingID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new binding appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing directory bindings to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListDirectoryBindings(ctx2)
	require.NoError(t, err, "ListDirectoryBindings should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, b := range list.Results {
		if b.ID == bindingID {
			found = true
			assert.Equal(t, bindingName, b.Name)
			break
		}
	}
	assert.True(t, found, "newly created directory binding should appear in list")
	acc.LogTestSuccess(t, "Directory binding ID=%d found in list (%d total)", bindingID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching directory binding by ID=%d", bindingID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetDirectoryBindingByID(ctx3, bindingID)
	require.NoError(t, err, "GetDirectoryBindingByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, bindingID, fetched.ID)
	assert.Equal(t, bindingName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching directory binding by name=%q", bindingName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetDirectoryBindingByName(ctx4, bindingName)
	require.NoError(t, err, "GetDirectoryBindingByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, bindingID, fetchedByName.ID)
	assert.Equal(t, bindingName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-dirbinding-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating directory binding ID=%d to name=%q", bindingID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &directory_bindings.RequestDirectoryBinding{
		Name:     updatedName,
		Priority: 1,
		Domain:   "test.example.com",
		Type:     "Active Directory",
	}
	updated, updateResp, err := svc.UpdateDirectoryBindingByID(ctx5, bindingID, updateReq)
	require.NoError(t, err, "UpdateDirectoryBindingByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating directory binding name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &directory_bindings.RequestDirectoryBinding{
		Name:     bindingName,
		Priority: 1,
		Domain:   "test.example.com",
		Type:     "Active Directory",
	}
	reverted, revertResp, err := svc.UpdateDirectoryBindingByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateDirectoryBindingByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetDirectoryBindingByID(ctx7, bindingID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, bindingName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting directory binding ID=%d", bindingID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteDirectoryBindingByID(ctx8, bindingID)
	require.NoError(t, err, "DeleteDirectoryBindingByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Directory binding ID=%d deleted", bindingID)
}

// =============================================================================
// TestAcceptance_DirectoryBindings_DeleteByName creates a binding then deletes by name.
// =============================================================================

func TestAcceptance_DirectoryBindings_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DirectoryBindings
	ctx := context.Background()

	bindingName := uniqueName("acc-test-dirbinding-dbn")
	createReq := &directory_bindings.RequestDirectoryBinding{
		Name:     bindingName,
		Priority: 1,
		Domain:   "test.example.com",
		Type:     "Active Directory",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateDirectoryBinding(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	bindingID := created.ID
	acc.LogTestSuccess(t, "Created directory binding ID=%d name=%q for delete-by-name test", bindingID, bindingName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteDirectoryBindingByID(cleanupCtx, bindingID)
		acc.LogCleanupDeleteError(t, "directory binding", fmt.Sprintf("%d", bindingID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteDirectoryBindingByName(ctx2, bindingName)
	require.NoError(t, err, "DeleteDirectoryBindingByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Directory binding %q deleted by name", bindingName)
}

// =============================================================================
// TestAcceptance_DirectoryBindings_ValidationErrors tests client-side validation.
// =============================================================================

func TestAcceptance_DirectoryBindings_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DirectoryBindings

	t.Run("GetDirectoryBindingByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetDirectoryBindingByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
	})

	t.Run("GetDirectoryBindingByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetDirectoryBindingByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "directory binding name is required")
	})

	t.Run("CreateDirectoryBinding_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateDirectoryBinding(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateDirectoryBindingByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateDirectoryBindingByID(context.Background(), 0, &directory_bindings.RequestDirectoryBinding{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
	})

	t.Run("UpdateDirectoryBindingByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateDirectoryBindingByName(context.Background(), "", &directory_bindings.RequestDirectoryBinding{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "directory binding name is required")
	})

	t.Run("DeleteDirectoryBindingByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteDirectoryBindingByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
	})

	t.Run("DeleteDirectoryBindingByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteDirectoryBindingByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "directory binding name is required")
	})
}
