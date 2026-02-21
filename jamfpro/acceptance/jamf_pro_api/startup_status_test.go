package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Startup Status
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Retrieves Jamf Pro server startup status
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service only provides read access to server startup status
//     -- Tests: TestAcceptance_StartupStatus_GetV1
//     -- Flow: Get status → Verify response structure
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get startup status
//   ✓ Verify response structure (200 status)
//   ✓ Verify Step field is present
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a read-only endpoint - no mutations possible
//   • Returns server startup status information
//   • Useful for monitoring server initialization state
//   • No cleanup needed as this is a read-only operation
//
// =============================================================================

func TestAcceptance_StartupStatus_GetV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.StartupStatus
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.Step)
}
