package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Patch Management
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • AcceptDisclaimerV2(ctx) - Accepts the Patch Management disclaimer
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern: One-time mutation
//     -- Reason: AcceptDisclaimerV2 is a one-time operation
//     -- Tests: TestAcceptance_PatchManagement_AcceptDisclaimerV2
//     -- Flow: Accept disclaimer → Verify 200 status
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Accept Patch Management disclaimer
//   ✓ Verify response structure (200 status)
//
// Notes
// -----------------------------------------------------------------------------
//   • AcceptDisclaimerV2 is a one-time operation - accepting the disclaimer
//     may affect subsequent tests. Use with caution in shared environments.
//
// =============================================================================

func TestAcceptance_PatchManagement_AcceptDisclaimerV2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.PatchManagement
	ctx := context.Background()

	resp, err := svc.AcceptDisclaimerV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
