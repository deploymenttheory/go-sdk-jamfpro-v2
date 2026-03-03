package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/onboarding"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
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
//     -- Tests: TestAcceptance_Onboarding_get_v1, TestAcceptance_Onboarding_update_v1
//     -- Flow: Get → Update → Verify → Restore
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get current settings
//   ✓ Update settings
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

func TestAcceptance_Onboarding_get_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_Onboarding_update_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	current, _, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	updateReq := &onboarding.ResourceUpdateOnboardingSettings{
		Enabled:         !current.Enabled,
		OnboardingItems: convertToUpdateItems(current.OnboardingItems),
	}
	updated, resp, err := svc.UpdateV1(ctx, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	restoreReq := &onboarding.ResourceUpdateOnboardingSettings{
		Enabled:         current.Enabled,
		OnboardingItems: convertToUpdateItems(current.OnboardingItems),
	}
	_, _, _ = svc.UpdateV1(ctx, restoreReq)
}

func convertToUpdateItems(items []onboarding.OnboardingItemResponse) []onboarding.SubsetOnboardingItemRequest {
	result := make([]onboarding.SubsetOnboardingItemRequest, len(items))
	for i, item := range items {
		result[i] = onboarding.SubsetOnboardingItemRequest{
			ID:                    item.ID,
			EntityID:              item.EntityID,
			SelfServiceEntityType: item.SelfServiceEntityType,
			Priority:              item.Priority,
		}
	}
	return result
}

func TestAcceptance_Onboarding_get_eligible_apps_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	result, resp, err := svc.GetEligibleAppsV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_Onboarding_get_eligible_configuration_profiles_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	result, resp, err := svc.GetEligibleConfigurationProfilesV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_Onboarding_get_eligible_policies_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	result, resp, err := svc.GetEligiblePoliciesV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_Onboarding_get_history_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	// Add history note first
	noteReq := &shared.SharedHistoryNoteRequest{
		Note: "Acceptance test history note for onboarding",
	}
	addResult, addResp, err := svc.AddHistoryNotesV1(ctx, noteReq)
	require.NoError(t, err)
	require.NotNil(t, addResult)
	assert.Equal(t, 201, addResp.StatusCode)
	t.Logf("Added history note with ID: %s", addResult.ID)

	result, resp, err := svc.GetHistoryV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 1, "Should have at least the note we just added")
}

func TestAcceptance_Onboarding_add_history_notes_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Onboarding
	ctx := context.Background()

	noteReq := &onboarding.RequestAddHistoryNotes{
		Note: "Test history note from acceptance test",
	}

	result, resp, err := svc.AddHistoryNotesV1(ctx, noteReq)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotEmpty(t, result.ID)
}
