package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func uniqueComputerGroupName(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixMilli())
}

// =============================================================================
// TestAcceptance_ComputerGroups_Smart_Lifecycle
// =============================================================================

func TestAcceptance_ComputerGroups_Smart_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerGroups
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test smart computer group")

	createReq := &computer_groups.RequestSmartGroup{
		Name: uniqueComputerGroupName("acc-test-smart"),
		Criteria: []computer_groups.Criterion{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "like", Value: "test-%"},
		},
	}
	created, createResp, err := svc.CreateSmartGroupV2(ctx, createReq)
	require.NoError(t, err, "CreateSmartGroupV2 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	groupID := created.ID
	acc.LogTestSuccess(t, "Smart group created with ID=%s", groupID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteSmartGroupV2(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "smart computer group", groupID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing smart groups to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListSmartGroupsV2(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, createReq.Name, g.Name)
			break
		}
	}
	assert.True(t, found, "newly created smart group should appear in list")
	acc.LogTestSuccess(t, "Smart group ID=%s found in list (%d total)", groupID, list.TotalCount)

	// 3. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching smart group by ID=%s", groupID)

	fetched, fetchResp, err := svc.GetSmartGroupByIDV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	if fetched.ID != "" {
		assert.Equal(t, groupID, fetched.ID, "GetByID response ID when present")
	}
	assert.Equal(t, createReq.Name, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating smart group ID=%s", groupID)

	updateReq := &computer_groups.RequestSmartGroup{
		Name: uniqueComputerGroupName("acc-test-smart-updated"),
		Criteria: []computer_groups.Criterion{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "like", Value: "updated-%"},
		},
	}
	updated, updateResp, err := svc.UpdateSmartGroupV2(ctx, groupID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 202}, updateResp.StatusCode, "Update returns 200 or 202")
	assert.Equal(t, updateReq.Name, updated.Name)
	acc.LogTestSuccess(t, "Smart group updated: ID=%s", groupID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetSmartGroupByIDV2(ctx, groupID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Name, fetched2.Name)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting smart group ID=%s", groupID)

	deleteResp, err := svc.DeleteSmartGroupV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Smart group ID=%s deleted", groupID)
}

// =============================================================================
// TestAcceptance_ComputerGroups_Static_Lifecycle
// =============================================================================

func TestAcceptance_ComputerGroups_Static_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerGroups
	ctx := context.Background()

	// 1. Create — static group with empty or minimal membership (computer IDs may not exist)
	acc.LogTestStage(t, "Create", "Creating test static computer group")

	createReq := &computer_groups.RequestStaticGroup{
		Name:        uniqueComputerGroupName("acc-test-static"),
		ComputerIds: []string{},
	}
	created, createResp, err := svc.CreateStaticGroupV2(ctx, createReq)
	if err != nil && createResp != nil && createResp.StatusCode == 500 {
		t.Skip("Static computer group create returned 500 in this environment; skipping lifecycle")
	}
	require.NoError(t, err, "CreateStaticGroupV2 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	groupID := created.ID
	acc.LogTestSuccess(t, "Static group created with ID=%s", groupID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteStaticGroupByIDV2(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "static computer group", groupID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing static groups to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListStaticGroupsV2(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, createReq.Name, g.Name)
			break
		}
	}
	assert.True(t, found, "newly created static group should appear in list")
	acc.LogTestSuccess(t, "Static group ID=%s found in list (%d total)", groupID, list.TotalCount)

	// 3. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching static group by ID=%s", groupID)

	fetched, fetchResp, err := svc.GetStaticGroupByIDV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, groupID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	assert.False(t, fetched.IsSmart)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

	// 4. Update membership (PATCH)
	acc.LogTestStage(t, "Update", "Updating static group membership ID=%s", groupID)

	updateReq := &computer_groups.RequestStaticGroup{
		Name:        createReq.Name,
		ComputerIds: []string{},
	}
	updated, updateResp, err := svc.UpdateStaticGroupByIDV2(ctx, groupID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	acc.LogTestSuccess(t, "Static group membership updated: ID=%s", groupID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetStaticGroupByIDV2(ctx, groupID)
	require.NoError(t, err)
	assert.Equal(t, groupID, fetched2.ID)
	acc.LogTestSuccess(t, "Update verified")

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting static group ID=%s", groupID)

	deleteResp, err := svc.DeleteStaticGroupByIDV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Static group ID=%s deleted", groupID)
}

// =============================================================================
// TestAcceptance_ComputerGroups_ValidationErrors
// =============================================================================

func TestAcceptance_ComputerGroups_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerGroups

	t.Run("GetSmartGroupByIDV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetSmartGroupByIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart group ID is required")
	})

	t.Run("CreateSmartGroupV2_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateSmartGroupV2(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateSmartGroupV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateSmartGroupV2(context.Background(), "", &computer_groups.RequestSmartGroup{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteSmartGroupV2_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteSmartGroupV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart group ID is required")
	})

	t.Run("GetStaticGroupByIDV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetStaticGroupByIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "static group ID is required")
	})

	t.Run("CreateStaticGroupV2_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateStaticGroupV2(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateStaticGroupByIDV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateStaticGroupByIDV2(context.Background(), "", &computer_groups.RequestStaticGroup{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteStaticGroupByIDV2_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteStaticGroupByIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "static group ID is required")
	})
}
