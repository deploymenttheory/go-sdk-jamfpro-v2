package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Pro Version
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Retrieves the current Jamf Pro server version
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service only provides read access to system version information
//     -- Tests: TestAcceptance_JamfProVersion_GetV1
//     -- Flow: Get version → Verify response structure and required fields
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get version information
//   ✓ Verify response structure (200 status)
//   ✓ Validate required fields (Version is present and non-empty)
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a read-only endpoint - no mutations possible
//   • Version format is typically "X.Y.Z" (e.g., "11.2.0")
//   • Version information is used by acc.GreaterThanJamfProVersion() helper
//   • No cleanup needed as this is a read-only operation
//
// =============================================================================

func TestAcceptance_JamfProVersion_GetV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProVersion
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result.Version)
	assert.NotEmpty(t, *result.Version)
}
