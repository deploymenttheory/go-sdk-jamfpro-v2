package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/policies"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


// =============================================================================
// TestAcceptance_Policies_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_Policies_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Policies
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test policy")

	policyName := acc.UniqueName("acc-test-policy")
	createReq := &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:      policyName,
			Enabled:   true,
			Frequency: "Once per computer",
		},
		Scope: policies.PolicySubsetScope{
			AllComputers: true,
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created policy ID should be a positive integer")

	policyID := created.ID
	acc.LogTestSuccess(t, "Policy created with ID=%d name=%q", policyID, policyName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, policyID)
		acc.LogCleanupDeleteError(t, "policy", fmt.Sprintf("%d", policyID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new policy appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing policies to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, p := range list.Results {
		if p.ID == policyID {
			found = true
			assert.Equal(t, policyName, p.Name)
			break
		}
	}
	assert.True(t, found, "newly created policy should appear in list")
	acc.LogTestSuccess(t, "Policy ID=%d found in list (%d total)", policyID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching policy by ID=%d", policyID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, policyID, fetched.General.ID)
	assert.Equal(t, policyName, fetched.General.Name)
	assert.True(t, fetched.General.Enabled)
	assert.True(t, fetched.Scope.AllComputers)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching policy by name=%q", policyName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, policyName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, policyID, fetchedByName.General.ID)
	assert.Equal(t, policyName, fetchedByName.General.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.General.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("acc-test-policy-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating policy ID=%d to name=%q", policyID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:      updatedName,
			Enabled:   false,
			Frequency: "Ongoing",
		},
		Scope: policies.PolicySubsetScope{
			AllComputers: true,
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx5, policyID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating policy name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:      policyName,
			Enabled:   true,
			Frequency: "Once per computer",
		},
		Scope: policies.PolicySubsetScope{
			AllComputers: true,
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

	verified, verifyResp, err := svc.GetByID(ctx7, policyID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, policyName, verified.General.Name, "name should reflect the revert")
	assert.True(t, verified.General.Enabled, "enabled should be true after revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.General.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting policy ID=%d", policyID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByID(ctx8, policyID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Policy ID=%d deleted", policyID)
}

// =============================================================================
// TestAcceptance_Policies_DeleteByName creates a policy then deletes by name.
// =============================================================================

func TestAcceptance_Policies_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Policies
	ctx := context.Background()

	policyName := acc.UniqueName("acc-test-policy-del")
	createReq := &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:      policyName,
			Enabled:   true,
			Frequency: "Once per computer",
		},
		Scope: policies.PolicySubsetScope{
			AllComputers: true,
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	policyID := created.ID
	acc.LogTestSuccess(t, "Created policy ID=%d name=%q for delete-by-name test", policyID, policyName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, policyID)
		acc.LogCleanupDeleteError(t, "policy", fmt.Sprintf("%d", policyID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, policyName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Policy %q deleted by name", policyName)
}

// =============================================================================
// TestAcceptance_Policies_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_Policies_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Policies

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy ID must be a positive integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy name is required")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 0, &policies.ResourcePolicy{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &policies.ResourcePolicy{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy name is required")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy ID must be a positive integer")
	})

	t.Run("DeleteByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy name is required")
	})
}
