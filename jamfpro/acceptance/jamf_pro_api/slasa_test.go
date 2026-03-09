package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: SLASA (Software License Agreement Service Acceptance)
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetStatusV1(ctx) - Retrieves SLASA acceptance status
//   • AcceptV1(ctx) - Accepts the SLASA (Software License Agreement)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: GetStatusV1 is read-only; AcceptV1 is a one-time mutation
//     -- Tests: TestAcceptance_SLASA_get_status_v1
//     -- Flow: Get status → Verify response structure
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get SLASA status
//   ✓ Verify response structure (200 status)
//   ✓ Verify SLASAAcceptanceStatus field is present
//
// Notes
// -----------------------------------------------------------------------------
//   • AcceptV1 is a one-time operation - accepting SLASA may affect subsequent
//     tests. Use with caution in shared environments.
//   • GetStatusV1 is safe to run repeatedly.
//
// =============================================================================

func TestAcceptance_SLASA_get_status_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.Slasa
	ctx := context.Background()

	result, resp, err := svc.GetStatusV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.SLASAAcceptanceStatus)
	assert.Contains(t, []string{"ACCEPTED", "NOT_ACCEPTED"}, result.SLASAAcceptanceStatus)
}
