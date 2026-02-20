package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// uniqueName returns a name unique to the test run to avoid conflicts with
// existing data and between parallel test runs.
func uniqueAccountName(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixMilli())
}

// =============================================================================
// TestAcceptance_Accounts_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_Accounts_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Accounts
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test account")

	accountName := uniqueName("test-account")
	createReq := &accounts.RequestAccount{
		Name:         accountName,
		FullName:     "Test Account User",
		Email:        fmt.Sprintf("%s@example.com", accountName),
		EmailAddress: fmt.Sprintf("%s@example.com", accountName),
		Password:     "TestPassword123!",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Enabled:      "Enabled",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateAccount(ctx1, createReq)
	require.NoError(t, err, "CreateAccount should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created account ID should be a positive integer")

	accountID := created.ID
	acc.LogTestSuccess(t, "Account created with ID=%d name=%q", accountID, accountName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteAccountByID(cleanupCtx, accountID)
		acc.LogCleanupDeleteError(t, "account", fmt.Sprintf("%d", accountID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new account appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing accounts to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListAccounts(ctx2)
	require.NoError(t, err, "ListAccounts should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, u := range list.Users {
		if u.ID == accountID {
			found = true
			assert.Equal(t, accountName, u.Name)
			break
		}
	}
	if !found {
		for _, g := range list.Groups {
			if g.ID == accountID {
				found = true
				assert.Equal(t, accountName, g.Name)
				break
			}
		}
	}
	assert.True(t, found, "newly created account should appear in list")
	acc.LogTestSuccess(t, "Account ID=%d found in list", accountID)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching account by ID=%d", accountID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetAccountByID(ctx3, accountID)
	require.NoError(t, err, "GetAccountByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, accountID, fetched.ID)
	assert.Equal(t, accountName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching account by name=%q", accountName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetAccountByName(ctx4, accountName)
	require.NoError(t, err, "GetAccountByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, accountID, fetchedByName.ID)
	assert.Equal(t, accountName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("updated-account")
	acc.LogTestStage(t, "UpdateByID", "Updating account ID=%d to name=%q", accountID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &accounts.RequestAccount{
		Name:         updatedName,
		FullName:     "Updated Account User",
		Email:        fmt.Sprintf("%s@example.com", updatedName),
		EmailAddress: fmt.Sprintf("%s@example.com", updatedName),
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Enabled:      "Enabled",
	}
	updated, updateResp, err := svc.UpdateAccountByID(ctx5, accountID, updateReq)
	require.NoError(t, err, "UpdateAccountByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating account name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &accounts.RequestAccount{
		Name:         accountName,
		FullName:     "Test Account User",
		Email:        fmt.Sprintf("%s@example.com", accountName),
		EmailAddress: fmt.Sprintf("%s@example.com", accountName),
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Enabled:      "Enabled",
	}
	reverted, revertResp, err := svc.UpdateAccountByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateAccountByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetAccountByID(ctx7, accountID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, accountName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting account ID=%d", accountID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteAccountByID(ctx8, accountID)
	require.NoError(t, err, "DeleteAccountByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Account ID=%d deleted", accountID)
}

// =============================================================================
// TestAcceptance_Accounts_DeleteByName creates an account then deletes by name.
// =============================================================================

func TestAcceptance_Accounts_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Accounts
	ctx := context.Background()

	accountName := uniqueName("test-delete-account")
	createReq := &accounts.RequestAccount{
		Name:         accountName,
		FullName:     "Test Delete Account User",
		Email:        fmt.Sprintf("%s@example.com", accountName),
		EmailAddress: fmt.Sprintf("%s@example.com", accountName),
		Password:     "TestPassword123!",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Enabled:      "Enabled",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateAccount(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	accountID := created.ID
	acc.LogTestSuccess(t, "Created account ID=%d name=%q for delete-by-name test", accountID, accountName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteAccountByID(cleanupCtx, accountID)
		acc.LogCleanupDeleteError(t, "account", fmt.Sprintf("%d", accountID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteAccountByName(ctx2, accountName)
	require.NoError(t, err, "DeleteAccountByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Account %q deleted by name", accountName)
}

// =============================================================================
// TestAcceptance_Accounts_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_Accounts_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Accounts

	t.Run("GetAccountByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetAccountByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account ID must be a positive integer")
	})

	t.Run("GetAccountByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetAccountByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account name is required")
	})

	t.Run("CreateAccount_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateAccount(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateAccountByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateAccountByID(context.Background(), 0, &accounts.RequestAccount{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account ID must be a positive integer")
	})

	t.Run("UpdateAccountByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateAccountByName(context.Background(), "", &accounts.RequestAccount{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account name is required")
	})

	t.Run("DeleteAccountByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteAccountByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account ID must be a positive integer")
	})

	t.Run("DeleteAccountByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteAccountByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account name is required")
	})
}
