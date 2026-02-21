package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Pro Information
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV2(ctx) - Retrieves Jamf Pro information and feature flags
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service only provides read access to Jamf Pro system information
//     -- Tests: TestAcceptance_JamfProInformation_Get
//     -- Flow: Get information → Verify response structure and required fields
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get Jamf Pro information
//   ✓ Verify response structure (200 status)
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a read-only endpoint - no mutations possible
//   • Returns system information and feature flags for the Jamf Pro instance
//   • No cleanup needed as this is a read-only operation
//   • Feature flags indicate which features are enabled for the tenant
//
// =============================================================================

func TestAcceptance_JamfProInformation_Get(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProInformation
	ctx := context.Background()

	result, resp, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
