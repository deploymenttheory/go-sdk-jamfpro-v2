package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/smart_computer_groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Smart Computer Groups
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • List(ctx, rsqlQuery) - Lists all smart computer groups
//   • GetByID(ctx, id) - Retrieves a smart computer group by ID
//   • GetByName(ctx, name) - Retrieves a smart computer group by name
//   • GetMembership(ctx, id) - Retrieves computer IDs in the group
//   • Create(ctx, request) - Creates a new smart computer group
//   • UpdateByID(ctx, id, request) - Updates an existing smart computer group
//   • DeleteByID(ctx, id) - Deletes a smart computer group by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Flow: Create → GetByID → Update → GetByID (verify) → Delete
//
//   ✓ Pattern 7: Validation Errors
//     -- Cases: Empty IDs, empty name, nil requests
//
// =============================================================================

func TestAcceptance_SmartComputerGroups_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SmartComputerGroups
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test smart computer group")

	createReq := &smart_computer_groups.RequestSmartGroup{
		Name:        acc.UniqueName("sdkv2_acc_acc-test-smart-cg"),
		Description: "Acceptance test smart computer group",
		Criteria: []smart_computer_groups.SubsetCriteria{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "is", Value: "*"},
		},
	}
	created, createResp, err := svc.Create(ctx, createReq)
	if err != nil && createResp != nil && createResp.StatusCode == 500 {
		t.Skip("Smart computer group create returned 500 in this environment; skipping lifecycle")
	}
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	groupID := created.ID
	acc.LogTestSuccess(t, "Smart computer group created with ID=%s", groupID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "smart computer group", groupID, delErr)
	})

	// 2. GetByID (with retry for eventual consistency)
	acc.LogTestStage(t, "GetByID", "Getting smart computer group by ID=%s", groupID)

	var fetched *smart_computer_groups.ResourceSmartGroup
	var fetchResp *interfaces.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		fetched, fetchResp, getErr = svc.GetByID(ctx, groupID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, groupID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

	// 3. GetByName
	acc.LogTestStage(t, "GetByName", "Getting smart computer group by name=%s", createReq.Name)

	byName, _, err := svc.GetByName(ctx, createReq.Name)
	require.NoError(t, err)
	require.NotNil(t, byName)
	assert.Equal(t, groupID, byName.ID)
	acc.LogTestSuccess(t, "GetByName: ID=%s", byName.ID)

	// 4. GetMembership
	acc.LogTestStage(t, "GetMembership", "Getting membership for smart computer group ID=%s", groupID)

	membership, memResp, err := svc.GetMembership(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, membership)
	assert.Equal(t, 200, memResp.StatusCode)
	assert.NotNil(t, membership.Members)
	acc.LogTestSuccess(t, "GetMembership: %d members", len(membership.Members))

	// 5. Update
	acc.LogTestStage(t, "Update", "Updating smart computer group ID=%s", groupID)

	updateReq := &smart_computer_groups.RequestSmartGroup{
		Name:        acc.UniqueName("sdkv2_acc_acc-test-smart-cg-updated"),
		Description: "Updated description",
		Criteria: []smart_computer_groups.SubsetCriteria{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "is", Value: "*"},
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx, groupID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	acc.LogTestSuccess(t, "Smart computer group updated: ID=%s", groupID)

	// 6. Re-fetch to verify
	fetched2, _, err := svc.GetByID(ctx, groupID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Name, fetched2.Name)
	assert.Equal(t, updateReq.Description, fetched2.Description)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

	// 7. Delete
	acc.LogTestStage(t, "Delete", "Deleting smart computer group ID=%s", groupID)

	deleteResp, err := svc.DeleteByID(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Smart computer group ID=%s deleted", groupID)
}

func TestAcceptance_SmartComputerGroups_list(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SmartComputerGroups

	result, resp, err := svc.List(context.Background(), map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
	assert.NotNil(t, result.Results)
}

// =============================================================================
// TestAcceptance_SmartComputerGroups_list_with_rsql_filter
// =============================================================================

func TestAcceptance_SmartComputerGroups_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SmartComputerGroups
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_rsql-smart-cg")
	createReq := &smart_computer_groups.RequestSmartGroup{
		Name:        name,
		Description: "Acceptance test RSQL filter smart computer group",
		Criteria: []smart_computer_groups.SubsetCriteria{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "is", Value: "*"},
		},
	}

	created, createResp, err := svc.Create(ctx, createReq)
	if err != nil && createResp != nil && createResp.StatusCode == 500 {
		t.Skip("Smart computer group create returned 500 in this environment; skipping RSQL filter test")
	}
	require.NoError(t, err)
	require.NotNil(t, created)

	groupID := created.ID
	acc.LogTestSuccess(t, "Created smart computer group ID=%s name=%q", groupID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "smart computer group", groupID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.List(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, name, g.Name)
			break
		}
	}
	assert.True(t, found, "smart computer group should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target group found=%v", list.TotalCount, found)
}

func TestAcceptance_SmartComputerGroups_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SmartComputerGroups

	t.Run("GetByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart computer group ID is required")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart computer group name is required")
	})

	t.Run("GetMembership_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetMembership(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart computer group ID is required")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), "", &smart_computer_groups.RequestSmartGroup{
			Name: "sdkv2_acc_x",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("UpdateByID_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteByID_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart computer group ID is required")
	})
}
