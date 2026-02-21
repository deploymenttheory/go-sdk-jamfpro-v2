package jamf_pro_api

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_integrations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: API Integrations
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - Lists API integrations with optional RSQL filtering
//   • GetByIDV1(ctx, id) - Retrieves an API integration by ID
//   • GetByNameV1(ctx, name) - Retrieves an API integration by display name (helper)
//   • CreateV1(ctx, request) - Creates a new API integration
//   • UpdateByIDV1(ctx, id, request) - Updates an existing API integration
//   • UpdateByNameV1(ctx, name, request) - Updates by display name (helper)
//   • DeleteByIDV1(ctx, id) - Deletes an API integration by ID
//   • DeleteByNameV1(ctx, name) - Deletes by display name (helper)
//   • RefreshClientCredentialsByIDV1(ctx, id) - Creates/refreshes OAuth client credentials
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_ApiIntegrations_Lifecycle
//     -- Flow: Create role (dependency) → Create → GetByID → GetByName → RefreshCredentials → Update → Delete
//
//   ✗ Pattern 5: RSQL Filter Testing [MANDATORY - MISSING]
//     -- Reason: ListV1 accepts rsqlQuery parameter for filtering
//     -- Tests: MISSING - Should be added as TestAcceptance_ApiIntegrations_ListWithRSQLFilter
//     -- Flow: Create unique integration → Filter with RSQL → Verify filtered results
//     -- Status: MANDATORY test not implemented
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (with API role dependency)
//   ✓ Read operations (GetByID, GetByName, List without filter)
//   ✗ List with RSQL filtering [MANDATORY - MISSING]
//   ✓ Update operations (toggle enabled state)
//   ✓ Delete operations (single delete)
//   ✓ OAuth client credentials operations (refresh)
//   ✗ Update by name operations (not yet tested)
//   ✗ Delete by name operations (not yet tested)
//   ✗ Input validation and error handling (not yet tested)
//
// Notes
// -----------------------------------------------------------------------------
//   • RSQL testing is MANDATORY because ListV1 supports filtering
//   • API integrations require valid API role names in AuthorizationScopes
//   • Test creates API role first as dependency (proper cleanup order maintained)
//   • RefreshClientCredentialsByIDV1 generates OAuth client ID and secret
//   • AccessTokenLifetimeSeconds defines OAuth token validity (e.g., 3600 = 1 hour)
//   • Allowed RSQL filter fields: id, displayName
//   • GetByName/UpdateByName/DeleteByName are helper methods using ListV1
//   • All tests register cleanup handlers to remove test resources
//   • Tests use acc.UniqueName() to avoid conflicts
//   • TODO: Add RSQL filter test (MANDATORY)
//   • TODO: Add validation error tests
//
// =============================================================================

func TestAcceptance_ApiIntegrations_ListV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ApiIntegrations
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_ApiIntegrations_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ApiIntegrations
	roleSvc := acc.Client.APIRoles
	ctx := context.Background()

	name := acc.UniqueName("acc-rsql-integration")

	// Create API role first (dependency)
	roleName := acc.UniqueName("acc-rsql-integration-role")
	roleReq := &api_roles.RequestAPIRole{
		DisplayName: roleName,
		Privileges:  []string{"Read Computers"},
	}
	createdRole, _, err := roleSvc.CreateV1(ctx, roleReq)
	require.NoError(t, err)
	require.NotNil(t, createdRole)

	roleID := createdRole.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := roleSvc.DeleteByIDV1(cleanupCtx, roleID)
		acc.LogCleanupDeleteError(t, "API role", roleID, delErr)
	})

	// Create API integration
	created, _, err := svc.CreateV1(ctx, &api_integrations.ResourceApiIntegration{
		DisplayName:                name,
		Enabled:                    true,
		AuthorizationScopes:        []string{roleName},
		AccessTokenLifetimeSeconds: 3600,
	})
	require.NoError(t, err)
	require.NotNil(t, created)

	integrationID := strconv.Itoa(created.ID)
	acc.LogTestSuccess(t, "Created API integration ID=%s name=%q", integrationID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, integrationID)
		acc.LogCleanupDeleteError(t, "API integration", integrationID, delErr)
	})

	// Test RSQL filtering by displayName
	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`displayName=="%s"`, name),
	}

	list, listResp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, integration := range list.Results {
		if integration.ID == created.ID {
			found = true
			assert.Equal(t, name, integration.DisplayName)
			break
		}
	}
	assert.True(t, found, "API integration should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target integration found=%v", list.TotalCount, found)
}

func TestAcceptance_ApiIntegrations_Lifecycle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ApiIntegrations
	roleSvc := acc.Client.APIRoles
	ctx := context.Background()
	name := acc.UniqueName("acc-api-integration")

	// Create a valid API role first (dependency chain: api_roles → api_integrations)
	roleName := acc.UniqueName("acc-api-integration-role")
	roleReq := &api_roles.RequestAPIRole{
		DisplayName: roleName,
		Privileges:  []string{"Read Computers"},
	}
	createdRole, roleResp, err := roleSvc.CreateV1(ctx, roleReq)
	require.NoError(t, err)
	require.NotNil(t, createdRole)
	require.NotNil(t, roleResp)
	assert.Contains(t, []int{200, 201}, roleResp.StatusCode)

	roleID := createdRole.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = roleSvc.DeleteByIDV1(cleanupCtx, roleID)
	})

	// Create API integration using the valid role name
	created, resp, err := svc.CreateV1(ctx, &api_integrations.ResourceApiIntegration{
		DisplayName:                name,
		Enabled:                    true,
		AuthorizationScopes:        []string{roleName},
		AccessTokenLifetimeSeconds: 3600,
	})
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 200, resp.StatusCode)

	idStr := strconv.Itoa(created.ID)
	acc.Cleanup(t, func() {
		_, _ = svc.DeleteByIDV1(ctx, idStr)
	})

	getByID, resp, err := svc.GetByIDV1(ctx, idStr)
	require.NoError(t, err)
	require.NotNil(t, getByID)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, name, getByID.DisplayName)

	byName, resp, err := svc.GetByNameV1(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, byName)
	assert.Equal(t, name, byName.DisplayName)

	creds, resp, err := svc.RefreshClientCredentialsByIDV1(ctx, idStr)
	require.NoError(t, err)
	require.NotNil(t, creds)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, creds.ClientID)
	assert.NotEmpty(t, creds.ClientSecret)

	updated, resp, err := svc.UpdateByIDV1(ctx, idStr, &api_integrations.ResourceApiIntegration{
		DisplayName:                name,
		Enabled:                    false,
		AuthorizationScopes:        created.AuthorizationScopes,
		AccessTokenLifetimeSeconds: 3600,
	})
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, resp.StatusCode)

	delResp, err := svc.DeleteByIDV1(ctx, idStr)
	require.NoError(t, err)
	assert.Equal(t, 204, delResp.StatusCode)
}
