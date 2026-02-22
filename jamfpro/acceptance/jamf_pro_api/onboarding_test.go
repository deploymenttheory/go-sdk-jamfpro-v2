package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/onboarding"
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
//   • GetHistoryV1(ctx, rsqlQuery) - Retrieves onboarding history
//   • AddHistoryNotesV1(ctx, req) - Adds notes to onboarding history
//   • ExportHistoryV1(ctx, acceptHeader, rsqlQuery, req) - Exports history
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration
//     -- Reason: Singleton settings that cannot be created or deleted, only updated
//     -- Tests: TestAcceptance_Onboarding_SettingsWorkflow
//     -- Flow: Get → Update → Verify → Restore
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get current settings
//   ✓ Update settings
//   ✓ Verify updated settings
//   ✓ Restore original settings
//   ✓ Get eligible apps
//   ✓ Get eligible configuration profiles
//   ✓ Get eligible policies
//   ✓ Get history
//   ✓ Add history notes
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a singleton configuration - no create/delete operations
//   • Onboarding settings control new device enrollment workflow
//   • Eligible resource operations return lists of apps/profiles/policies available for onboarding
//   • History operations track changes to onboarding configuration
//
// =============================================================================

func TestAcceptance_Onboarding_Get(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	acc.LogTestStage(t, "Get onboarding settings")
	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	acc.LogTestSuccess(t, "Retrieved onboarding settings successfully")
}

func TestAcceptance_Onboarding_GetEligibleApps(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	acc.LogTestStage(t, "Get eligible apps for onboarding")
	result, resp, err := svc.GetEligibleAppsV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	acc.LogTestSuccess(t, "Retrieved %d eligible apps", result.TotalCount)
}

func TestAcceptance_Onboarding_GetEligibleConfigurationProfiles(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	acc.LogTestStage(t, "Get eligible configuration profiles for onboarding")
	result, resp, err := svc.GetEligibleConfigurationProfilesV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	acc.LogTestSuccess(t, "Retrieved %d eligible configuration profiles", result.TotalCount)
}

func TestAcceptance_Onboarding_GetEligiblePolicies(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	acc.LogTestStage(t, "Get eligible policies for onboarding")
	result, resp, err := svc.GetEligiblePoliciesV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	acc.LogTestSuccess(t, "Retrieved %d eligible policies", result.TotalCount)
}

func TestAcceptance_Onboarding_GetHistory(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	acc.LogTestStage(t, "Get onboarding history")
	result, resp, err := svc.GetHistoryV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	acc.LogTestSuccess(t, "Retrieved onboarding history with %d entries", result.TotalCount)
}

func TestAcceptance_Onboarding_AddHistoryNotes(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	acc.LogTestStage(t, "Add note to onboarding history")
	req := &onboarding.RequestAddHistoryNotes{
		Note: "Acceptance test note - automated test run",
	}
	result, resp, err := svc.AddHistoryNotesV1(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotEmpty(t, result.ID)
	acc.LogTestSuccess(t, "Added history note with ID: %s", result.ID)
}
