package classic_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_Policies_minimum_config tests creating a policy with the
// minimum required configuration.
// =============================================================================

func TestAcceptance_Policies_minimum_config(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create policy with minimum config
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating policy with minimum configuration")

	policyName := acc.UniqueName("sdkv2_acc_policy_min")
	createReq := createMinimalPolicy(t, policyName)

	_, policyID := createPolicyWithCleanup(t, ctx, svc, createReq)

	// ------------------------------------------------------------------
	// 2. Get by ID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching policy by ID=%d", policyID)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, policyID, fetched.General.ID)
	assert.Equal(t, policyName, fetched.General.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 3. Get by Name
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching policy by name=%q", policyName)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx3, policyName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, policyID, fetchedByName.General.ID)
	assert.Equal(t, policyName, fetchedByName.General.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.General.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 4. Update by ID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByID", "Updating policy by ID=%d", policyID)

	updateReq := fetched
	updateReq.General.Enabled = true

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	updated, updateResp, err := svc.UpdateByID(ctx4, policyID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, policyID, updated.ID, "updated policy ID should match")
	acc.LogTestSuccess(t, "UpdateByID: ID=%d", updated.ID)

	// ------------------------------------------------------------------
	// 5. List
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing all policies")

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	list, listResp, err := svc.List(ctx5)
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
	assert.True(t, found, "created policy should appear in list")
	acc.LogTestSuccess(t, "List: found policy ID=%d in list of %d policies", policyID, list.Size)
}

// =============================================================================
// TestAcceptance_Policies_validation_errors tests validation error handling.
// =============================================================================

func TestAcceptance_Policies_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(ctx, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy ID must be greater than 0")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy name cannot be empty")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(ctx, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		policy := createMinimalPolicy(t, "test")
		_, _, err := svc.UpdateByID(ctx, 0, policy)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy ID must be greater than 0")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		policy := createMinimalPolicy(t, "test")
		_, _, err := svc.UpdateByName(ctx, "", policy)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy name cannot be empty")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(ctx, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy ID must be greater than 0")
	})

	t.Run("DeleteByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteByName(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy name cannot be empty")
	})
}
