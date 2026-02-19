package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/vpp_accounts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_VPPAccounts_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → UpdateByID → DeleteByID.
// Note: the Classic API VPP accounts resource does not support ByName operations.
// =============================================================================

func TestAcceptance_VPPAccounts_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VPPAccounts
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test VPP account")

	accountName := uniqueName("acc-test-vpp")
	createReq := &vpp_accounts.RequestVPPAccount{
		Name:    accountName,
		Country: "US",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateVPPAccount(ctx1, createReq)
	require.NoError(t, err, "CreateVPPAccount should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created VPP account ID should be a positive integer")

	accountID := created.ID
	acc.LogTestSuccess(t, "VPP account created with ID=%d name=%q", accountID, accountName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteVPPAccountByID(cleanupCtx, accountID)
		acc.LogCleanupDeleteError(t, "VPP account", fmt.Sprintf("%d", accountID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new account appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing VPP accounts to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListVPPAccounts(ctx2)
	require.NoError(t, err, "ListVPPAccounts should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, a := range list.Results {
		if a.ID == accountID {
			found = true
			assert.Equal(t, accountName, a.Name)
			break
		}
	}
	assert.True(t, found, "newly created VPP account should appear in list")
	acc.LogTestSuccess(t, "VPP account ID=%d found in list (%d total)", accountID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching VPP account by ID=%d", accountID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetVPPAccountByID(ctx3, accountID)
	require.NoError(t, err, "GetVPPAccountByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, accountID, fetched.ID)
	assert.Equal(t, accountName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-vpp-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating VPP account ID=%d to name=%q", accountID, updatedName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	updateReq := &vpp_accounts.RequestVPPAccount{
		Name:    updatedName,
		Country: "US",
	}
	updated, updateResp, err := svc.UpdateVPPAccountByID(ctx4, accountID, updateReq)
	require.NoError(t, err, "UpdateVPPAccountByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 5. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting VPP account ID=%d", accountID)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	deleteResp, err := svc.DeleteVPPAccountByID(ctx5, accountID)
	require.NoError(t, err, "DeleteVPPAccountByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "VPP account ID=%d deleted", accountID)
}

// =============================================================================
// TestAcceptance_VPPAccounts_ValidationErrors tests client-side validation.
// =============================================================================

func TestAcceptance_VPPAccounts_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VPPAccounts

	t.Run("GetVPPAccountByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetVPPAccountByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
	})

	t.Run("CreateVPPAccount_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateVPPAccount(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateVPPAccountByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateVPPAccountByID(context.Background(), 0, &vpp_accounts.RequestVPPAccount{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
	})

	t.Run("DeleteVPPAccountByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteVPPAccountByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
	})
}
