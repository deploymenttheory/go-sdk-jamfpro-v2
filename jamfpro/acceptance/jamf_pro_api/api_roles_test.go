package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: API Roles
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - Lists API roles with optional RSQL filtering
//   • GetByIDV1(ctx, id) - Retrieves an API role by ID
//   • CreateV1(ctx, request) - Creates a new API role
//   • UpdateByIDV1(ctx, id, request) - Updates an existing API role
//   • DeleteByIDV1(ctx, id) - Deletes an API role by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_APIRoles_Lifecycle
//     -- Flow: Create → GetByID → Update → Delete
//
//   ✗ Pattern 5: RSQL Filter Testing [MANDATORY - MISSING]
//     -- Reason: ListV1 accepts rsqlQuery parameter for filtering
//     -- Tests: MISSING - Should be added as TestAcceptance_APIRoles_ListWithRSQLFilter
//     -- Flow: Create unique role → Filter with RSQL → Verify filtered results
//     -- Status: MANDATORY test not implemented
//
//   ✗ Pattern 7: Validation Errors [RECOMMENDED - MISSING]
//     -- Tests: MISSING - Should validate empty IDs, nil requests
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (single role creation)
//   ✓ Read operations (GetByID, List with pagination)
//   ✗ List with RSQL filtering [MANDATORY - MISSING]
//   ✓ Update operations (change display name)
//   ✓ Delete operations (single delete)
//   ✗ Input validation and error handling [RECOMMENDED - MISSING]
//   ✓ Cleanup and resource management
//
// Notes
// -----------------------------------------------------------------------------
//   • RSQL testing is MANDATORY because ListV1 supports filtering
//   • API roles define privilege sets for API integrations and OAuth clients
//   • Privileges must be valid Jamf Pro privilege names (e.g., "Read Computers")
//   • API roles are dependency for API integrations (created first in api_integrations tests)
//   • All tests register cleanup handlers to remove test roles
//   • Tests use timestamp-based unique names to avoid conflicts
//   • TODO: Add RSQL filter test (MANDATORY)
//   • TODO: Add validation error tests
//
// =============================================================================

func TestAcceptance_APIRoles_ListV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.APIRoles
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_APIRoles_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.APIRoles
	ctx := context.Background()

	name := fmt.Sprintf("acc-rsql-role-%d", time.Now().UnixMilli())
	createReq := &api_roles.RequestAPIRole{
		DisplayName: name,
		Privileges:  []string{"Read Computers"},
	}

	created, _, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	roleID := created.ID
	acc.LogTestSuccess(t, "Created API role ID=%s name=%q", roleID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, roleID)
		acc.LogCleanupDeleteError(t, "API role", roleID, delErr)
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
	for _, role := range list.Results {
		if role.ID == roleID {
			found = true
			assert.Equal(t, name, role.DisplayName)
			break
		}
	}
	assert.True(t, found, "API role should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target role found=%v", list.TotalCount, found)
}

func TestAcceptance_APIRoles_Lifecycle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.APIRoles
	ctx := context.Background()

	name := fmt.Sprintf("acc-api-role-%d", time.Now().UnixMilli())
	createReq := &api_roles.RequestAPIRole{
		DisplayName: name,
		Privileges:  []string{"Read Computers"},
	}
	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	roleID := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV1(cleanupCtx, roleID)
	})

	fetched, _, err := svc.GetByIDV1(ctx, roleID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, name, fetched.DisplayName)

	updatedName := name + "-updated"
	updateReq := &api_roles.RequestAPIRole{DisplayName: updatedName, Privileges: []string{"Read Computers"}}
	_, updateResp, err := svc.UpdateByIDV1(ctx, roleID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updateResp)
	assert.Contains(t, []int{200, 202}, updateResp.StatusCode)

	delResp, err := svc.DeleteByIDV1(ctx, roleID)
	require.NoError(t, err)
	require.NotNil(t, delResp)
	assert.Equal(t, 204, delResp.StatusCode)
}
