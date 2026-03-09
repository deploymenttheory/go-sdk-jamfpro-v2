package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: API Authorization
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Returns current authorization details for the API token
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service is read-only; returns details for the authenticated token
//     -- Tests: TestAcceptance_ApiAuthorization_get_v1
//     -- Flow: Get → Verify account username, authentication type non-empty
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Read operations (GetV1)
//   ✓ Field validation (account username, authentication type)
//
// Notes
// -----------------------------------------------------------------------------
//   • Returns details for the currently authenticated API token/credential
//   • AuthenticationType indicates the auth method (e.g. "OAUTH2", "BASIC")
//   • Account.Username will be non-empty for any authenticated session
//
// =============================================================================

// TestAcceptance_ApiAuthorization_get_v1 verifies the API authorization endpoint
// returns valid auth details for the authenticated session.
func TestAcceptance_ApiAuthorization_get_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.ApiAuthorization
	ctx := context.Background()

	acc.LogTestStage(t, "Get", "Getting API authorization details")

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	assert.NotEmpty(t, result.Account.Username, "account username should not be empty")
	assert.NotEmpty(t, result.AuthenticationType, "authentication type should not be empty")

	acc.LogTestSuccess(t, "GetV1: username=%q authenticationType=%q", result.Account.Username, result.AuthenticationType)
}
