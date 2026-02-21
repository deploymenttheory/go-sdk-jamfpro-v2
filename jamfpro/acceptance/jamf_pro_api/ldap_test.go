package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: LDAP
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetLdapGroupsV1(ctx, rsqlQuery) - Retrieves LDAP groups with optional RSQL filtering
//   • GetLdapServersV1(ctx) - Retrieves active LDAP/cloud identity provider servers
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service only provides read access to LDAP groups and servers
//     -- Tests: TestAcceptance_Ldap_GetLdapGroupsV1, TestAcceptance_Ldap_GetLdapServersV1
//     -- Flow: Get LDAP data → Verify response structure
//
//   ✗ Pattern 5: RSQL Filter Testing [MANDATORY - MISSING]
//     -- Reason: GetLdapGroupsV1 accepts rsqlQuery parameter for filtering
//     -- Tests: MISSING - Current test uses "q" parameter instead of RSQL "filter"
//     -- Flow: Should test with filter parameter for RSQL filtering
//     -- Status: MANDATORY test not implemented
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get LDAP groups (with query parameter "q")
//   ✓ Get LDAP servers
//   ✓ Verify response structure (200 status)
//   ✗ RSQL filtering for LDAP groups [MANDATORY - MISSING]
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a read-only endpoint - no mutations possible
//   • GetLdapGroupsV1 supports RSQL filtering via "filter" parameter (not "q")
//   • Current test uses "q" parameter which may be a search query, not RSQL
//   • GetLdapServersV1 returns all active LDAP/cloud identity provider servers
//   • LDAP groups are used for directory service integration
//   • No cleanup needed as this is a read-only operation
//   • TODO: Add RSQL filter test using map[string]string{"filter": "rsql-expression"}
//   • TODO: Clarify if "q" parameter is intentional or should be "filter" for RSQL
//
// =============================================================================

func TestAcceptance_Ldap_GetLdapGroupsV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Ldap
	ctx := context.Background()

	result, resp, err := svc.GetLdapGroupsV1(ctx, map[string]string{"q": "test"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result.Results)
}

func TestAcceptance_Ldap_GetLdapServersV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Ldap
	ctx := context.Background()

	result, resp, err := svc.GetLdapServersV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_Ldap_GetLdapGroupsV1_WithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Ldap
	ctx := context.Background()

	// First, get all LDAP groups to find one to filter for
	allGroups, allResp, err := svc.GetLdapGroupsV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, allGroups)
	assert.Equal(t, 200, allResp.StatusCode)

	if len(allGroups.Results) == 0 {
		t.Skip("No LDAP groups available in this environment; skipping RSQL filter test")
	}

	// Pick the first group to filter for
	targetGroup := allGroups.Results[0]
	acc.LogTestSuccess(t, "Found LDAP group to test RSQL filtering: %q", targetGroup.Name)

	// Test RSQL filtering by name
	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, targetGroup.Name),
	}

	filtered, filteredResp, err := svc.GetLdapGroupsV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, filtered)
	assert.Equal(t, 200, filteredResp.StatusCode)

	// Verify the target group appears in filtered results
	found := false
	for _, g := range filtered.Results {
		if g.ID == targetGroup.ID {
			found = true
			assert.Equal(t, targetGroup.Name, g.Name)
			break
		}
	}
	assert.True(t, found, "target LDAP group should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target group found=%v", len(filtered.Results), found)
}
