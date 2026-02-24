package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_Accounts_ListV1 tests listing user accounts with various RSQL queries.
func TestAcceptance_Accounts_ListV1(t *testing.T) {
	acc.RequireClient(t)
	client := acc.Client

	// Test 1: List all accounts (no filter)
	t.Run("ListAll", func(t *testing.T) {
		result, resp, err := client.Accounts.ListV1(context.Background(), nil)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)
		assert.GreaterOrEqual(t, result.TotalCount, 0)
		t.Logf("Found %d total accounts", result.TotalCount)
	})

	// Test 2: List with RSQL filter for enabled accounts
	t.Run("FilterEnabledAccounts", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `accountStatus==Enabled`,
		}

		result, resp, err := client.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

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

		result, resp, err := client.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)

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

		result, resp, err := client.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)

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

		result, resp, err := client.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
		// Note: GetPaginated fetches all pages, so we should get all results
		t.Logf("Retrieved %d accounts (pagination handled automatically)", len(result.Results))
	})

	// Test 6: List with sorting
	t.Run("WithSorting", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"sort": "realname:desc",
		}

		result, resp, err := client.Accounts.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
		t.Logf("Retrieved %d accounts sorted by realname descending", len(result.Results))
	})
}

// TestAcceptance_Accounts_CRUD tests the full CRUD lifecycle for accounts.
// Note: This test is commented out by default as it modifies data.
// Uncomment and run carefully in a test environment.
/*
func TestAcceptance_Accounts_CRUD(t *testing.T) {
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

	created, resp, err := client.Accounts.CreateV1(context.Background(), createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotEmpty(t, created.ID)
	t.Logf("Created account with ID: %s", created.ID)

	// Clean up
	defer func() {
		_, err := client.Accounts.DeleteByIDV1(context.Background(), created.ID)
		if err != nil {
			t.Logf("Warning: Failed to delete test account: %v", err)
		}
	}()

	// Get by ID
	retrieved, resp, err := client.Accounts.GetByIDV1(context.Background(), created.ID)
	require.NoError(t, err)
	require.NotNil(t, retrieved)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "testuser_acceptance", retrieved.Username)
	t.Logf("Retrieved account: %s", retrieved.Username)

	// Verify account appears in list with RSQL filter
	rsqlQuery := map[string]string{
		"filter": `username=="testuser_acceptance"`,
	}
	listResult, resp, err := client.Accounts.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, len(listResult.Results), 1, "Created account should appear in filtered list")
	t.Logf("Found created account in list")
}
*/
