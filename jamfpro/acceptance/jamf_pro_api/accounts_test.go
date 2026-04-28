package jamf_pro_api

import (
	"context"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/accounts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_Accounts_lifecycle tests the full CRUD lifecycle for accounts
// including the new UpdateByIDV1 method added in Jamf Pro 11.27.0.
func TestAcceptance_Accounts_lifecycle(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()
	svc := acc.Client.JamfProAPI.Accounts

	// 1. CREATE
	acc.LogTestStage(t, "create", "creating test account")
	createReq := &accounts.RequestAccount{
		Username:       acc.UniqueName("sdkv2_acc_account"),
		Realname:       "SDK v2 Acceptance Test Account",
		Email:          "sdkv2acc@example.com",
		PlainPassword:  "AcceptanceTest123!",
		AccessLevel:    "FullAccess",
		PrivilegeLevel: "AUDITOR",
		AccountStatus:  "Enabled",
		AccountType:    "DEFAULT",
	}

	created, resp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 201, resp.StatusCode())
	assert.NotEmpty(t, created.ID)
	acc.LogTestSuccess(t, "created account id=%s username=%s", created.ID, created.Username)

	acc.Cleanup(t, func() {
		ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if _, err := svc.DeleteByIDV1(ctx2, created.ID); err != nil {
			acc.LogCleanupDeleteError(t, "account", created.ID, err)
		}
	})

	// 2. GET by ID
	acc.LogTestStage(t, "get", "retrieving created account by ID")
	fetched, resp, err := svc.GetByIDV1(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, created.ID, fetched.ID)
	assert.Equal(t, createReq.Username, fetched.Username)

	// 3. LIST and find the created account
	acc.LogTestStage(t, "list", "listing accounts")
	list, resp, err := svc.ListV1(ctx, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	found := false
	for _, a := range list.Results {
		if a.ID == created.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "created account should appear in list")

	// 4. UPDATE
	acc.LogTestStage(t, "update", "updating account")
	updateReq := &accounts.RequestAccount{
		Realname:       "SDK v2 Acceptance Updated",
		Email:          "sdkv2acc_updated@example.com",
		AccessLevel:    "FullAccess",
		PrivilegeLevel: "AUDITOR",
		AccountStatus:  "Enabled",
		AccountType:    "DEFAULT",
	}
	updated, resp, err := svc.UpdateByIDV1(ctx, created.ID, updateReq)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "SDK v2 Acceptance Updated", updated.Realname)

	// 5. VERIFY update persisted
	acc.LogTestStage(t, "verify", "verifying update persisted")
	verified, resp, err := svc.GetByIDV1(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "SDK v2 Acceptance Updated", verified.Realname)

	// 6. DELETE
	acc.LogTestStage(t, "delete", "deleting account")
	resp, err = svc.DeleteByIDV1(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
	acc.LogTestSuccess(t, "deleted account id=%s", created.ID)
}

// TestAcceptance_Accounts_list_with_rsql_filter tests RSQL filtering on accounts.
func TestAcceptance_Accounts_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()
	svc := acc.Client.JamfProAPI.Accounts

	// Create unique account to filter on
	acc.LogTestStage(t, "setup", "creating account for RSQL filter test")
	uniqueUsername := acc.UniqueName("sdkv2_acc_rsql")
	createReq := &accounts.RequestAccount{
		Username:       uniqueUsername,
		Realname:       "SDK v2 RSQL Test",
		Email:          "sdkv2rsql@example.com",
		PlainPassword:  "AcceptanceTest123!",
		AccessLevel:    "FullAccess",
		PrivilegeLevel: "AUDITOR",
		AccountStatus:  "Enabled",
		AccountType:    "DEFAULT",
	}
	created, resp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode())

	acc.Cleanup(t, func() {
		ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if _, err := svc.DeleteByIDV1(ctx2, created.ID); err != nil {
			acc.LogCleanupDeleteError(t, "account", created.ID, err)
		}
	})

	// Filter by username
	acc.LogTestStage(t, "filter", "listing with RSQL filter")
	rsqlQuery := map[string]string{
		"filter": `username=="` + uniqueUsername + `"`,
	}
	result, resp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	require.GreaterOrEqual(t, len(result.Results), 1, "filtered list should contain the created account")
	assert.Equal(t, uniqueUsername, result.Results[0].Username)
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s)", len(result.Results))
}

// TestAcceptance_Accounts_validation_errors tests parameter validation.
func TestAcceptance_Accounts_validation_errors(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.Accounts
	ctx := context.Background()

	t.Run("GetByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account ID is required")
	})

	t.Run("UpdateByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(ctx, "", &accounts.RequestAccount{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account ID is required")
	})

	t.Run("UpdateByIDV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(ctx, "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("CreateV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV1(ctx, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV1(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account ID is required")
	})
}

// TestAcceptance_Accounts_list_v1 tests listing user accounts with various RSQL queries.
func TestAcceptance_Accounts_list_v1(t *testing.T) {
	acc.RequireClient(t)
	client := acc.Client

	// Test 1: List all accounts (no filter)
	t.Run("ListAll", func(t *testing.T) {
		result, resp, err := client.JamfProAPI.Accounts.ListV1(context.Background(), nil)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode())
		assert.GreaterOrEqual(t, result.TotalCount, 0)
		t.Logf("Found %d total accounts", result.TotalCount)
	})

	// Test 2: List with RSQL filter for enabled accounts
	t.Run("FilterEnabledAccounts", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `accountStatus==Enabled`,
		}

		result, resp, err := client.JamfProAPI.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode())

		// Verify all returned accounts are enabled
		for _, account := range result.Results {
			assert.Equal(t, "Enabled", account.AccountStatus, "Expected all accounts to have Enabled status")
		}
		t.Logf("Found %d enabled accounts", len(result.Results))
	})

	// Test 3: List with RSQL filter for administrators
	t.Run("FilterAdministrators", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `privilegeLevel==ADMINISTRATOR`,
		}

		result, resp, err := client.JamfProAPI.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode())

		// Verify all returned accounts are administrators
		for _, account := range result.Results {
			assert.Equal(t, "ADMINISTRATOR", account.PrivilegeLevel, "Expected all accounts to have ADMINISTRATOR privilege level")
		}
		t.Logf("Found %d administrator accounts", len(result.Results))
	})

	// Test 4: List with complex RSQL filter
	t.Run("ComplexRSQLFilter", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `accountStatus==Enabled and privilegeLevel==ADMINISTRATOR and failedLoginAttempts==0`,
			"sort":   "username:asc",
		}

		result, resp, err := client.JamfProAPI.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode())

		// Verify all returned accounts match the filter criteria
		for _, account := range result.Results {
			assert.Equal(t, "Enabled", account.AccountStatus)
			assert.Equal(t, "ADMINISTRATOR", account.PrivilegeLevel)
			assert.Equal(t, 0, account.FailedLoginAttempts)
		}
		t.Logf("Found %d enabled administrators with no failed login attempts", len(result.Results))
	})

	// Test 5: List with pagination parameters
	t.Run("WithPagination", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"page":      "0",
			"page-size": "5",
			"sort":      "username:asc",
		}

		result, resp, err := client.JamfProAPI.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode())
		// Note: GetPaginated fetches all pages, so we should get all results
		t.Logf("Retrieved %d accounts (pagination handled automatically)", len(result.Results))
	})

	// Test 6: List with sorting
	t.Run("WithSorting", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"sort": "realname:desc",
		}

		result, resp, err := client.JamfProAPI.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode())
		t.Logf("Retrieved %d accounts sorted by realname descending", len(result.Results))
	})
}

// TestAcceptance_Accounts_crud tests the full CRUD lifecycle for accounts.
// Note: This test is commented out by default as it modifies data.
// Uncomment and run carefully in a test environment.
/*
func TestAcceptance_Accounts_crud(t *testing.T) {
	acceptance.SkipIfNotEnabled(t)
	client := acceptance.NewClient(t)

	// Create test account
	createReq := &accounts.RequestAccount{
		Username:       "testuser_acceptance",
		Realname:       "Test User Acceptance",
		Email:          "test@example.com",
		PlainPassword:  "SecurePassword123!",
		AccessLevel:    "FullAccess",
		PrivilegeLevel: "AUDITOR",
		AccountStatus:  "Enabled",
		AccountType:    "DEFAULT",
	}

	created, resp, err := client.JamfProAPI.Accounts.CreateV1(context.Background(), createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 201, resp.StatusCode())
	assert.NotEmpty(t, created.ID)
	t.Logf("Created account with ID: %s", created.ID)

	// Clean up
	defer func() {
		_, err := client.JamfProAPI.Accounts.DeleteByIDV1(context.Background(), created.ID)
		if err != nil {
			t.Logf("Warning: Failed to delete test account: %v", err)
		}
	}()

	// Get by ID
	retrieved, resp, err := client.JamfProAPI.Accounts.GetByIDV1(context.Background(), created.ID)
	require.NoError(t, err)
	require.NotNil(t, retrieved)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "testuser_acceptance", retrieved.Username)
	t.Logf("Retrieved account: %s", retrieved.Username)

	// Verify account appears in list with RSQL filter
	rsqlQuery := map[string]string{
		"filter": `username=="testuser_acceptance"`,
	}
	listResult, resp, err := client.JamfProAPI.Accounts.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.GreaterOrEqual(t, len(listResult.Results), 1, "Created account should appear in filtered list")
	t.Logf("Found created account in list")
}
*/
