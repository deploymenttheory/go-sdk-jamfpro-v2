package jamf_pro_api

import (
	"context"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/static_computer_groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Static Computer Groups
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV2(ctx, rsqlQuery) - Lists all static computer groups
//   • GetByIDV2(ctx, id) - Retrieves a static computer group by ID
//   • GetByNameV2(ctx, name) - Retrieves a static computer group by name
//   • CreateV2(ctx, request) - Creates a new static computer group
//   • UpdateByIDV2(ctx, id, request) - Updates an existing static computer group by ID
//   • DeleteByIDV2(ctx, id) - Deletes a static computer group by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Flow: Create → GetByID → GetByName → Update → Verify → Delete
//
//   ✓ Pattern 7: Validation Errors
//     -- Cases: Empty IDs, nil requests, empty name
//
// =============================================================================

func TestAcceptance_StaticComputerGroups_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.StaticComputerGroups
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test static computer group")

	createReq := &static_computer_groups.RequestStaticGroup{
		Name:        acc.UniqueName("acc-test-static-cg"),
		Description: "Acceptance test static computer group",
		Assignments: []string{},
	}
	created, createResp, err := svc.CreateV2(ctx, createReq)
	if err != nil && createResp != nil && createResp.StatusCode == 500 {
		t.Skip("Static computer group create returned 500 in this environment; skipping lifecycle")
	}
	require.NoError(t, err, "CreateV2 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	groupID := created.ID
	acc.LogTestSuccess(t, "Static computer group created with ID=%s", groupID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV2(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "static computer group", groupID, delErr)
	})

	// 2. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching static computer group by ID=%s", groupID)

	fetched, fetchResp, err := svc.GetByIDV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, groupID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	acc.LogTestSuccess(t, "GetByIDV2: name=%q", fetched.Name)

	// 3. GetByName
	acc.LogTestStage(t, "GetByName", "Fetching static computer group by name=%s", createReq.Name)

	byName, _, err := svc.GetByNameV2(ctx, createReq.Name)
	require.NoError(t, err)
	require.NotNil(t, byName)
	assert.Equal(t, groupID, byName.ID)
	acc.LogTestSuccess(t, "GetByNameV2: ID=%s", byName.ID)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating static computer group ID=%s", groupID)

	updateReq := &static_computer_groups.RequestStaticGroup{
		Name:        acc.UniqueName("acc-test-static-cg-updated"),
		Description: "Updated description",
		Assignments: []string{},
	}
	updated, updateResp, err := svc.UpdateByIDV2(ctx, groupID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	acc.LogTestSuccess(t, "Static computer group updated: ID=%s", groupID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetByIDV2(ctx, groupID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Name, fetched2.Name)
	assert.Equal(t, updateReq.Description, fetched2.Description)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting static computer group ID=%s", groupID)

	deleteResp, err := svc.DeleteByIDV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Static computer group ID=%s deleted", groupID)
}

func TestAcceptance_StaticComputerGroups_List(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.StaticComputerGroups

	result, resp, err := svc.ListV2(context.Background(), map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
	assert.NotNil(t, result.Results)
}

func TestAcceptance_StaticComputerGroups_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.StaticComputerGroups

	t.Run("GetByIDV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "static computer group ID is required")
	})

	t.Run("GetByNameV2_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByNameV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "static computer group name is required")
	})

	t.Run("CreateV2_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV2(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByIDV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV2(context.Background(), "", &static_computer_groups.RequestStaticGroup{
			Name:        "x",
			Assignments: []string{},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteByIDV2_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "static computer group ID is required")
	})
}
