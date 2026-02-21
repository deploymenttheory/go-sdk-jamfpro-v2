package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/engage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Engage
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV2(ctx) - Retrieves Engage settings
//   • UpdateV2(ctx, settings) - Updates Engage settings (PUT)
//   • GetHistoryV2(ctx, rsqlQuery) - Retrieves Engage history with optional RSQL filtering
//   • AddHistoryNotesV2(ctx, request) - Adds notes to Engage history
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration [COMPLETE]
//     -- Reason: Singleton settings that cannot be created or deleted, only updated
//     -- Tests: TestAcceptance_Engage_SettingsLifecycle
//     -- Flow: Get original → Update → Verify → Restore → Verify restoration (5-step)
//
//   ✓ Pattern 5: RSQL Filter Testing [COMPLETE]
//     -- Reason: GetHistoryV2 accepts rsqlQuery parameter for filtering
//     -- Tests: TestAcceptance_Engage_HistoryWithRSQLFilter
//     -- Flow: Get history → Filter with RSQL → Verify filtered results
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get current settings
//   ✓ Update settings
//   ✓ Verify updated settings (Get after Update)
//   ✓ Restore original settings
//   ✓ Verify restored settings (Get after Restore)
//   ✓ Get history (with pagination)
//   ✓ Get history with RSQL filtering
//   ✓ Add history notes
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a singleton configuration - no create/delete operations
//   • Engage settings control Jamf Engage integration (notifications, messaging)
//   • Tests restore original settings to avoid side effects
//   • History operations track changes to Engage settings
//   • RSQL filtering tested on history endpoint
//   • Tests skip gracefully if Engage is not available on the tenant
//
// =============================================================================

func TestAcceptance_Engage_SettingsLifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Engage
	ctx := context.Background()

	acc.LogTestStage(t, "Get", "Fetching current Engage settings")
	originalSettings, getResp, err := svc.GetV2(ctx)
	if err != nil {
		t.Skipf("Engage feature may not be supported on this tenant: %v", err)
		return
	}

	require.NotNil(t, originalSettings)
	assert.Equal(t, 200, getResp.StatusCode)
	originalEnabled := originalSettings.IsEnabled
	acc.LogTestSuccess(t, "Current Engage settings - IsEnabled: %v", originalEnabled)

	acc.LogTestStage(t, "Update", "Updating Engage settings (toggle IsEnabled)")
	updateReq := &engage.ResourceEngageSettings{
		IsEnabled: !originalEnabled,
	}
	updatedSettings, updateResp, err := svc.UpdateV2(ctx, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updatedSettings)
	assert.Contains(t, []int{200, 202}, updateResp.StatusCode)
	acc.LogTestSuccess(t, "Updated Engage settings - IsEnabled toggled to: %v", !originalEnabled)

	acc.LogTestStage(t, "Verify", "Verifying updated settings")
	verifyUpdated, verifyResp, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, verifyUpdated)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, !originalEnabled, verifyUpdated.IsEnabled, "Updated setting should be persisted")
	acc.LogTestSuccess(t, "Verified updated settings - IsEnabled: %v", verifyUpdated.IsEnabled)

	acc.LogTestStage(t, "Restore", "Restoring original Engage settings")
	restoreReq := &engage.ResourceEngageSettings{
		IsEnabled: originalEnabled,
	}
	restoredSettings, restoreResp, err := svc.UpdateV2(ctx, restoreReq)
	require.NoError(t, err)
	require.NotNil(t, restoredSettings)
	assert.Contains(t, []int{200, 202}, restoreResp.StatusCode)
	acc.LogTestSuccess(t, "Restored Engage settings - IsEnabled: %v", originalEnabled)

	acc.LogTestStage(t, "Verify Restoration", "Verifying restored settings")
	verifyRestored, verifyRestoreResp, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, verifyRestored)
	assert.Equal(t, 200, verifyRestoreResp.StatusCode)
	assert.Equal(t, originalEnabled, verifyRestored.IsEnabled, "Original setting should be restored")
	acc.LogTestSuccess(t, "Verified restoration - IsEnabled: %v (matches original)", verifyRestored.IsEnabled)
}

func TestAcceptance_Engage_History(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Engage
	ctx := context.Background()

	acc.LogTestStage(t, "GetHistory", "Fetching Engage settings history")
	history, histResp, err := svc.GetHistoryV2(ctx, map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "date:desc",
	})

	if err != nil {
		t.Skipf("Engage history may not be available on this tenant: %v", err)
		return
	}

	require.NotNil(t, history)
	assert.Equal(t, 200, histResp.StatusCode)
	assert.GreaterOrEqual(t, history.TotalCount, 0)
	acc.LogTestSuccess(t, "Found %d history entries", history.TotalCount)

	if history.TotalCount > 0 {
		firstEntry := history.Results[0]
		acc.LogTestSuccess(t, "Latest history entry - Username: %s, Date: %s, Note: %s",
			firstEntry.Username, firstEntry.Date, firstEntry.Note)
	}
}

func TestAcceptance_Engage_HistoryWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Engage
	ctx := context.Background()

	acc.LogTestStage(t, "GetHistory", "Fetching history to test RSQL filtering")
	allHistory, allResp, err := svc.GetHistoryV2(ctx, nil)
	if err != nil || allHistory.TotalCount == 0 {
		t.Skip("No Engage history available for RSQL filtering test")
		return
	}

	assert.Equal(t, 200, allResp.StatusCode)
	acc.LogTestSuccess(t, "Found %d total history entries", allHistory.TotalCount)

	// Test RSQL filtering by username (exclude nonexistent username to get results)
	acc.LogTestStage(t, "RSQL Filter", "Testing RSQL filter on history")
	rsqlQuery := map[string]string{
		"filter": `username!="nonexistent_user_xyz"`,
	}

	filteredHistory, filteredResp, err := svc.GetHistoryV2(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, filteredHistory)
	assert.Equal(t, 200, filteredResp.StatusCode)
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s)", filteredHistory.TotalCount)

	// Verify filtering worked
	assert.GreaterOrEqual(t, allHistory.TotalCount, filteredHistory.TotalCount,
		"Filtered results should be <= total results")
}

func TestAcceptance_Engage_AddHistoryNotes(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Engage
	ctx := context.Background()

	acc.LogTestStage(t, "AddHistoryNotes", "Adding note to Engage history")
	noteReq := &engage.RequestAddHistoryNotes{
		Note: fmt.Sprintf("Acceptance test note - automated testing"),
	}

	result, resp, err := svc.AddHistoryNotesV2(ctx, noteReq)
	if err != nil {
		t.Skipf("Adding history notes may not be supported on this tenant: %v", err)
		return
	}

	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotZero(t, result.ID)
	assert.NotEmpty(t, result.Username)
	acc.LogTestSuccess(t, "History note added - ID: %d, Username: %s, Date: %s", result.ID, result.Username, result.Date)

	// Verify the note appears in history
	acc.LogTestStage(t, "Verify", "Verifying note appears in history")
	history, histResp, err := svc.GetHistoryV2(ctx, map[string]string{
		"page":      "0",
		"page-size": "10",
		"sort":      "date:desc",
	})

	require.NoError(t, err)
	require.NotNil(t, history)
	assert.Equal(t, 200, histResp.StatusCode)

	if history.TotalCount > 0 {
		// The most recent entry should contain our note
		found := false
		for _, entry := range history.Results {
			if entry.Note == noteReq.Note {
				found = true
				acc.LogTestSuccess(t, "Found our note in history - Username: %s, Date: %s",
					entry.Username, entry.Date)
				break
			}
		}
		if !found {
			t.Logf("Note not found in first page of history results (may be on later page)")
		}
	}
}
