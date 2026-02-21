package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Onboarding
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Retrieves current onboarding settings
//   • UpdateV1(ctx, request) - Updates onboarding settings (PUT)
//   • GetEligibleAppsV1(ctx, query) - Retrieves eligible apps for onboarding
//   • GetEligibleConfigurationProfilesV1(ctx, query) - Retrieves eligible config profiles
//   • GetEligiblePoliciesV1(ctx, query) - Retrieves eligible policies
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✗ Pattern 2: Settings/Configuration [INCOMPLETE]
//     -- Reason: Singleton settings that cannot be created or deleted, only updated
//     -- Tests: TestAcceptance_Onboarding_Get
//     -- Flow: Get only - MISSING update/restore cycle
//     -- Status: Tests exist but don't implement full Pattern 2 workflow
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get current settings
//   ✗ Update settings (not yet tested - should be added)
//   ✗ Verify updated settings (not yet tested)
//   ✗ Restore original settings (not yet tested - should be added)
//   ✗ Get eligible apps (not yet tested)
//   ✗ Get eligible configuration profiles (not yet tested)
//   ✗ Get eligible policies (not yet tested)
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a singleton configuration - no create/delete operations
//   • Onboarding settings control new device enrollment workflow
//   • Eligible resource operations return lists of apps/profiles/policies available for onboarding
//   • TODO: Implement full Pattern 2 test with Get → Update → Verify → Restore cycle
//   • TODO: Add tests for eligible resources endpoints
//
// =============================================================================

func TestAcceptance_Onboarding_Get(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()
	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
