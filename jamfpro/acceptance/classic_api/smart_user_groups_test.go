package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/smart_user_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_SmartUserGroups_lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_SmartUserGroups_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicSmartUserGroups
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test smart user group")

	groupName := acc.UniqueName("sdkv2_acc_smart-usergrp")
	createReq := &smart_user_groups.RequestSmartUserGroup{
		Name:             groupName,
		IsSmart:          true,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created smart user group ID should be a positive integer")

	groupID := created.ID
	acc.LogTestSuccess(t, "Smart user group created with ID=%d name=%q", groupID, groupName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "smart user group", fmt.Sprintf("%d", groupID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new smart user group appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing user groups to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, groupName, g.Name)
			assert.True(t, g.IsSmart)
			break
		}
	}
	assert.True(t, found, "newly created smart user group should appear in list")
	acc.LogTestSuccess(t, "Smart user group ID=%d found in list (%d total)", groupID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Getting smart user group by ID=%d", groupID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, groupID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, groupID, fetched.ID)
	assert.Equal(t, groupName, fetched.Name)
	assert.True(t, fetched.IsSmart)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Getting smart user group by name=%q", groupName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, groupName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, groupID, fetchedByName.ID)
	assert.Equal(t, groupName, fetchedByName.Name)
	assert.True(t, fetchedByName.IsSmart)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("sdkv2_acc_smart-usergrp-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating smart user group ID=%d to name=%q", groupID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &smart_user_groups.RequestSmartUserGroup{
		Name:             updatedName,
		IsSmart:          true,
		IsNotifyOnChange: true,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 2,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "@example.com",
				},
				{
					Name:       "Username",
					Priority:   1,
					AndOr:      "and",
					SearchType: "like",
					Value:      "admin",
				},
			},
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx5, groupID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating smart user group name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &smart_user_groups.RequestSmartUserGroup{
		Name:             groupName,
		IsSmart:          true,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
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

	verified, verifyResp, err := svc.GetByID(ctx7, groupID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, groupName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting smart user group ID=%d", groupID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByID(ctx8, groupID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Smart user group ID=%d deleted", groupID)
}

// =============================================================================
// TestAcceptance_SmartUserGroups_delete_by_name creates a smart user group
// then deletes it by name.
// =============================================================================

func TestAcceptance_SmartUserGroups_delete_by_name(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicSmartUserGroups
	ctx := context.Background()

	groupName := acc.UniqueName("sdkv2_acc_smart-usergrp-dbn")
	createReq := &smart_user_groups.RequestSmartUserGroup{
		Name:             groupName,
		IsSmart:          true,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	groupID := created.ID
	acc.LogTestSuccess(t, "Created smart user group ID=%d name=%q for delete-by-name test", groupID, groupName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "smart user group", fmt.Sprintf("%d", groupID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, groupName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Smart user group %q deleted by name", groupName)
}

// =============================================================================
// TestAcceptance_SmartUserGroups_validation_errors validates error handling.
// =============================================================================

func TestAcceptance_SmartUserGroups_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicSmartUserGroups

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err, "GetByID with zero ID should return an error")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err, "GetByName with empty name should return an error")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err, "Create with nil request should return an error")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		req := &smart_user_groups.RequestSmartUserGroup{Name: "test"}
		_, _, err := svc.UpdateByID(context.Background(), 0, req)
		assert.Error(t, err, "UpdateByID with zero ID should return an error")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		req := &smart_user_groups.RequestSmartUserGroup{Name: "test"}
		_, _, err := svc.UpdateByName(context.Background(), "", req)
		assert.Error(t, err, "UpdateByName with empty name should return an error")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err, "DeleteByID with zero ID should return an error")
	})
}
