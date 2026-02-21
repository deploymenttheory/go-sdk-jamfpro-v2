package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Client Check-In Settings
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV3(ctx) - Retrieves current client check-in settings
//   • UpdateV3(ctx, request) - Updates client check-in settings
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration
//     -- Reason: Singleton settings that cannot be created or deleted
//     -- Tests: TestAcceptance_ClientCheckin_Get
//     -- Flow: Get current settings → Verify structure and values
//     -- Note: Full update/restore cycle not implemented yet (should be added)
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get current settings
//   ✓ Verify response structure
//   ✓ Validate required fields (CheckInFrequency)
//   ✗ Update settings (not yet tested - should be added)
//   ✗ Restore original settings (not yet tested - should be added)
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a singleton configuration - no create/delete operations
//   • Settings affect all managed devices in the tenant
//   • TODO: Add full Pattern 2 test with Get → Update → Verify → Restore cycle
//   • CheckInFrequency validation: must be >= 0
//
// =============================================================================

func TestAcceptance_ClientCheckin_Get(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClientCheckin
	ctx := context.Background()

	result, resp, err := svc.GetV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.CheckInFrequency, 0)
}
