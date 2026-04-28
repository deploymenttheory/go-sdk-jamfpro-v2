package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/users_inventory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Users Inventory (v1)
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   - ListV1(ctx, rsqlQuery) - Lists users with optional RSQL filtering
//   - GetByIDV1(ctx, id) - Retrieves a user by ID
//   - CreateV1(ctx, request) - Creates a new user
//   - UpdateByIDV1(ctx, id, request) - Updates an existing user
//   - DeleteByIDV1(ctx, id) - Deletes a user by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   - Pattern 1: Full CRUD Lifecycle
//   - Pattern 5: RSQL Filter Testing [MANDATORY]
//   - Pattern 7: Validation Errors
//
// =============================================================================

func TestAcceptance_UsersInventory_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.UsersInventory
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test user")

	username := acc.UniqueName("sdkv2_acc_user")
	createReq := &users_inventory.RequestUserInventory{
		Username: username,
		Realname: "SDK Test User",
		Email:    fmt.Sprintf("%s@example.com", username),
		Position: "Test Engineer",
	}
	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err, "CreateV1 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode())
	assert.NotEmpty(t, created.ID)

	userID := created.ID
	acc.LogTestSuccess(t, "User created with ID=%s", userID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, userID)
		acc.LogCleanupDeleteError(t, "user", userID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing users to verify creation")

	list, listResp, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode())

	found := false
	for _, u := range list.Results {
		if u.ID == userID {
			found = true
			assert.Equal(t, createReq.Username, u.Username)
			break
		}
	}
	assert.True(t, found, "newly created user should appear in list")
	acc.LogTestSuccess(t, "User ID=%s found in list (%d total)", userID, list.TotalCount)

	// 3. GetByID
	acc.LogTestStage(t, "GetByID", "Getting user by ID=%s", userID)

	fetched, fetchResp, err := svc.GetByIDV1(ctx, userID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode())
	assert.Equal(t, userID, fetched.ID)
	assert.Equal(t, createReq.Username, fetched.Username)
	acc.LogTestSuccess(t, "GetByID: username=%q", fetched.Username)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating user ID=%s", userID)

	updateReq := &users_inventory.RequestUserInventory{
		Username: username,
		Realname: "SDK Test User Updated",
		Email:    fmt.Sprintf("%s@example.com", username),
		Position: "Senior Test Engineer",
	}
	updateResp, err := svc.UpdateByIDV1(ctx, userID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updateResp)
	assert.Equal(t, 204, updateResp.StatusCode())
	acc.LogTestSuccess(t, "User updated: ID=%s", userID)

	// 5. Re-fetch to verify update
	fetched2, _, err := svc.GetByIDV1(ctx, userID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Realname, fetched2.Realname)
	assert.Equal(t, updateReq.Position, fetched2.Position)
	acc.LogTestSuccess(t, "Update verified: realname=%q position=%q", fetched2.Realname, fetched2.Position)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting user ID=%s", userID)

	deleteResp, err := svc.DeleteByIDV1(ctx, userID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode())
	acc.LogTestSuccess(t, "User ID=%s deleted", userID)
}

// =============================================================================
// TestAcceptance_UsersInventory_list_with_rsql_filter
// =============================================================================

func TestAcceptance_UsersInventory_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.UsersInventory
	ctx := context.Background()

	username := acc.UniqueName("sdkv2_acc_rsql_user")
	createReq := &users_inventory.RequestUserInventory{
		Username: username,
		Realname: "RSQL Test User",
		Email:    fmt.Sprintf("%s@example.com", username),
	}

	created, _, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	userID := created.ID
	acc.LogTestSuccess(t, "Created user ID=%s username=%q", userID, username)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, userID)
		acc.LogCleanupDeleteError(t, "user", userID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`username=="%s"`, username),
	}

	list, listResp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode())

	found := false
	for _, u := range list.Results {
		if u.ID == userID {
			found = true
			assert.Equal(t, username, u.Username)
			break
		}
	}
	assert.True(t, found, "user should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target user found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_UsersInventory_validation_errors
// =============================================================================

func TestAcceptance_UsersInventory_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.UsersInventory

	t.Run("GetByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user ID is required")
	})

	t.Run("CreateV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByIDV1_EmptyID", func(t *testing.T) {
		err2 := func() error {
			_, err := svc.UpdateByIDV1(context.Background(), "", &users_inventory.RequestUserInventory{Username: "x"})
			return err
		}()
		assert.Error(t, err2)
		assert.Contains(t, err2.Error(), "user ID is required")
	})

	t.Run("UpdateByIDV1_NilRequest", func(t *testing.T) {
		err2 := func() error {
			_, err := svc.UpdateByIDV1(context.Background(), "1", nil)
			return err
		}()
		assert.Error(t, err2)
		assert.Contains(t, err2.Error(), "request is required")
	})

	t.Run("DeleteByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user ID is required")
	})
}
