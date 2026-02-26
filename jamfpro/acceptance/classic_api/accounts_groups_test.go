package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts_groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_AccountGroups_lifecycle exercises the full write/read/delete
// lifecycle: Create → GetByID → GetByName → UpdateByID → UpdateByName →
// GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_AccountGroups_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicAccountGroups
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test account group")

	groupName := acc.UniqueName("sdkv2_acc_test-account-group")
	createReq := &accounts_groups.RequestAccountGroup{
		Name:         groupName,
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created account group ID should be a positive integer")

	groupID := created.ID
	acc.LogTestSuccess(t, "Account group created with ID=%d name=%q", groupID, groupName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "account group", fmt.Sprintf("%d", groupID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching account group by ID=%d", groupID)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, groupID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, groupID, fetched.ID)
	assert.Equal(t, groupName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 3. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching account group by name=%q", groupName)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx3, groupName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, groupID, fetchedByName.ID)
	assert.Equal(t, groupName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 4. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("sdkv2_acc_updated-account-group")
	acc.LogTestStage(t, "UpdateByID", "Updating account group ID=%d to name=%q", groupID, updatedName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	updateReq := &accounts_groups.RequestAccountGroup{
		Name:         updatedName,
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
	}
	updated, updateResp, err := svc.UpdateByID(ctx4, groupID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 5. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating account group name=%q back to original", updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	revertReq := &accounts_groups.RequestAccountGroup{
		Name:         groupName,
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
	}
	reverted, revertResp, err := svc.UpdateByName(ctx5, updatedName, revertReq)
	require.NoError(t, err, "UpdateByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	verified, verifyResp, err := svc.GetByID(ctx6, groupID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, groupName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 7. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting account group ID=%d", groupID)

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	deleteResp, err := svc.DeleteByID(ctx7, groupID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Account group ID=%d deleted", groupID)
}
