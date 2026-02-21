package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Time Zones
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx) - Retrieves all available time zones
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service only provides read access to available time zones
//     -- Tests: TestAcceptance_TimeZones_ListV1
//     -- Flow: List time zones → Verify response structure
//
//   Note: RSQL Filter Testing NOT applicable
//     -- ListV1 does not support filtering - returns all time zones
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ List all time zones
//   ✓ Verify response structure (200 status)
//   ✓ Verify non-empty result set
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a read-only endpoint - no mutations possible
//   • Returns all available time zones for the Jamf Pro instance
//   • Time zones used for scheduling and date/time operations
//   • No filtering or pagination - returns complete list
//   • No cleanup needed as this is a read-only operation
//
// =============================================================================

func TestAcceptance_TimeZones_ListV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.TimeZones
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result)
}
