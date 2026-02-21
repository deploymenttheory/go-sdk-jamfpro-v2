package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Device Communication Settings
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Retrieves current device communication settings
//   • UpdateV1(ctx, request) - Updates device communication settings (PUT)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✗ Pattern 2: Settings/Configuration [INCOMPLETE]
//     -- Reason: Singleton settings that cannot be created or deleted, only updated
//     -- Tests: TestAcceptance_DeviceCommunicationSettings_GetV1, TestAcceptance_DeviceCommunicationSettings_UpdateV1
//     -- Flow: Get → Update (same values) - MISSING proper update/restore cycle
//     -- Status: Tests exist but don't implement full Pattern 2 workflow
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get current settings
//   ✓ Update settings (but using same values - no actual change tested)
//   ✗ Update with changed settings (not yet tested)
//   ✗ Verify updated settings (not yet tested)
//   ✗ Restore original settings (not yet tested - should be added)
//   ✗ Verify restoration (not yet tested - should be added)
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a singleton configuration - no create/delete operations
//   • Settings affect device communication behavior across the tenant
//   • TODO: Implement full Pattern 2 test with Get → Update → Verify → Restore → Verify cycle
//   • Current UpdateV1 test just updates with same values - should test actual changes
//   • Suggested safe change for testing: toggle a boolean field temporarily then restore
//   • UpdateV1 uses PUT (full resource replacement)
//
// =============================================================================

func TestAcceptance_DeviceCommunicationSettings_GetV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.DeviceCommunicationSettings
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_DeviceCommunicationSettings_UpdateV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.DeviceCommunicationSettings
	ctx := context.Background()

	current, _, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	request := *current
	updated, resp, err := svc.UpdateV1(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, updated)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
