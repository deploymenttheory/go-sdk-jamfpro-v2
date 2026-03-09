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
// Acceptance Tests: Classic LDAP
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetMappingsByIDV1(ctx, id) - Returns LDAP attribute mappings for an on-prem
//                                   LDAP configuration by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Reason: Service is read-only; requires an existing on-prem LDAP server
//     -- Tests: TestAcceptance_ClassicLdap_get_mappings_v1
//     -- Flow: List LDAP servers → skip if none → GetMappings on first result
//
// Notes
// -----------------------------------------------------------------------------
//   • Endpoint: GET /api/v1/classic-ldap/{id}
//   • The {id} is the Classic LDAP server integer ID (ldapServerId)
//   • Test skips gracefully if no on-prem LDAP servers are configured
//
// =============================================================================

func TestAcceptance_ClassicLdap_get_mappings_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.ClassicLdap
	ctx := context.Background()

	// Get on-prem LDAP servers to find a valid ID
	servers, _, err := acc.Client.JamfProAPI.Ldap.GetLdapServersOnlyV1(ctx)
	require.NoError(t, err)

	if len(servers) == 0 {
		t.Skip("No on-prem LDAP servers configured; skipping ClassicLdap GetMappings")
	}

	// Use the first LDAP server's ID
	ldapID := fmt.Sprintf("%d", servers[0].ID)
	acc.LogTestStage(t, "GetMappings", "Getting LDAP mappings for server ID=%s", ldapID)

	result, resp, err := svc.GetMappingsByIDV1(ctx, ldapID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	acc.LogTestSuccess(t, "GetMappingsByIDV1: ldapServerID=%s userObjectMapUsernameTo=%q",
		ldapID, result.UserObjectMapUsernameTo)
}
