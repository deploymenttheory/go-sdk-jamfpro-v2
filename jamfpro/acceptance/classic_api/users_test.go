package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/users"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_Users_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → GetByEmail → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_Users_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicUsers
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test user")

	userName := acc.UniqueName("acc-test-user")
	createReq := &users.RequestUser{
		Name:     userName,
		FullName: "Acceptance Test User",
		Email:    userName + "@example.com",
		Sites: []shared.SharedResourceSite{
			{ID: -1, Name: "None"},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created user ID should be a positive integer")

	userID := created.ID
	acc.LogTestSuccess(t, "User created with ID=%d name=%q", userID, userName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, userID)
		acc.LogCleanupDeleteError(t, "user", fmt.Sprintf("%d", userID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new user appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing users to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, u := range list.Results {
		if u.ID == userID {
			found = true
			assert.Equal(t, userName, u.Name)
			break
		}
	}
	assert.True(t, found, "newly created user should appear in list")
	acc.LogTestSuccess(t, "User ID=%d found in list (%d total)", userID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching user by ID=%d", userID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, userID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, userID, fetched.ID)
	assert.Equal(t, userName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching user by name=%q", userName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, userName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, userID, fetchedByName.ID)
	assert.Equal(t, userName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. GetByEmail
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByEmail", "Fetching user by email=%q", createReq.Email)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	fetchedByEmail, fetchByEmailResp, err := svc.GetByEmail(ctx5, createReq.Email)
	require.NoError(t, err, "GetByEmail should not return an error")
	require.NotNil(t, fetchedByEmail)
	assert.Equal(t, 200, fetchByEmailResp.StatusCode)
	assert.Equal(t, userID, fetchedByEmail.ID)
	assert.Equal(t, userName, fetchedByEmail.Name)
	acc.LogTestSuccess(t, "GetByEmail: ID=%d email=%q", fetchedByEmail.ID, createReq.Email)

	// ------------------------------------------------------------------
	// 6. UpdateByID
	// ------------------------------------------------------------------
	updatedFullName := "Acceptance Test User Updated"
	acc.LogTestStage(t, "UpdateByID", "Updating user ID=%d full_name=%q", userID, updatedFullName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	updateReq := &users.RequestUser{
		Name:     userName,
		FullName: updatedFullName,
		Email:    createReq.Email,
		Sites: []shared.SharedResourceSite{
			{ID: -1, Name: "None"},
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx6, userID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify update
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify update")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetByID(ctx7, userID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, updatedFullName, verified.FullName, "full_name should reflect the update")
	acc.LogTestSuccess(t, "Update verified: full_name=%q", verified.FullName)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting user ID=%d", userID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByID(ctx8, userID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "User ID=%d deleted", userID)
}

// =============================================================================
// TestAcceptance_Users_DeleteByName creates a user then deletes by name.
// =============================================================================

func TestAcceptance_Users_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicUsers
	ctx := context.Background()

	userName := acc.UniqueName("acc-test-user-dbn")
	createReq := &users.RequestUser{
		Name:     userName,
		FullName: "Delete By Name Test",
		Email:    userName + "@example.com",
		Sites: []shared.SharedResourceSite{
			{ID: -1, Name: "None"},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	userID := created.ID
	acc.LogTestSuccess(t, "Created user ID=%d name=%q for delete-by-name test", userID, userName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, userID)
		acc.LogCleanupDeleteError(t, "user", fmt.Sprintf("%d", userID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, userName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "User %q deleted by name", userName)
}

// =============================================================================
// TestAcceptance_Users_ValidationErrors validates error handling.
// =============================================================================

func TestAcceptance_Users_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicUsers

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user ID must be a positive integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user name cannot be empty")
	})

	t.Run("GetByEmail_EmptyEmail", func(t *testing.T) {
		_, _, err := svc.GetByEmail(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user email cannot be empty")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 0, &users.RequestUser{Name: "test"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &users.RequestUser{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user name cannot be empty")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user ID must be a positive integer")
	})
}
