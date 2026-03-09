package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Remote Assist
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListSessionsV1(ctx) - Lists up to 100 recent sessions (no pagination)
//   • GetSessionByIDV1(ctx, id) - Gets a session by ID (v1)
//   • ListSessionsV2(ctx, rsqlQuery) - Lists sessions with pagination and RSQL
//   • GetSessionByIDV2(ctx, id) - Gets a session by ID with details (v2)
//   • ExportSessionsV2(ctx, request, acceptType) - Exports sessions as CSV/JSON
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Tests: TestAcceptance_JamfRemoteAssist_list_sessions
//     -- Flow: ListSessionsV1 → ListSessionsV2 → if sessions exist: GetByID
//     -- Note: Skips per-session tests gracefully if no sessions exist
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_JamfRemoteAssist_validation_errors
//     -- Cases: GetSessionByIDV1/V2 with empty ID
//
// Notes
// -----------------------------------------------------------------------------
//   • Session history may be empty in environments with no remote assist usage
//   • RSQL filter keys: sessionId, deviceId, sessionAdminId
//
// =============================================================================

// TestAcceptance_JamfRemoteAssist_list_sessions verifies session listing and retrieval.
func TestAcceptance_JamfRemoteAssist_list_sessions(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.JamfRemoteAssist
	ctx := context.Background()

	// 1. ListSessionsV1
	acc.LogTestStage(t, "ListSessionsV1", "Listing recent remote assist sessions (v1)")

	sessionsV1, respV1, err := svc.ListSessionsV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, respV1)
	assert.Equal(t, 200, respV1.StatusCode())
	assert.GreaterOrEqual(t, len(sessionsV1), 0)
	acc.LogTestSuccess(t, "ListSessionsV1: %d session(s)", len(sessionsV1))

	// 2. ListSessionsV2 (with pagination)
	acc.LogTestStage(t, "ListSessionsV2", "Listing remote assist sessions with pagination (v2)")

	listV2, respV2, err := svc.ListSessionsV2(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, listV2)
	require.NotNil(t, respV2)
	assert.Equal(t, 200, respV2.StatusCode())
	assert.GreaterOrEqual(t, listV2.TotalCount, 0)
	acc.LogTestSuccess(t, "ListSessionsV2: totalCount=%d", listV2.TotalCount)

	// 3. GetSessionByIDV1 and V2 — only if sessions exist
	if len(sessionsV1) == 0 {
		acc.LogTestWarning(t, "No remote assist sessions found; skipping per-session fetch tests")
		return
	}

	sessionID := sessionsV1[0].SessionID
	if sessionID == "" {
		acc.LogTestWarning(t, "First session has no sessionId; skipping per-session fetch tests")
		return
	}

	// GetSessionByIDV1
	acc.LogTestStage(t, "GetSessionByIDV1", "Getting session ID=%s (v1)", sessionID)

	sessionV1, getV1Resp, err := svc.GetSessionByIDV1(ctx, sessionID)
	require.NoError(t, err)
	require.NotNil(t, sessionV1)
	assert.Equal(t, 200, getV1Resp.StatusCode())
	assert.Equal(t, sessionID, sessionV1.SessionID)
	acc.LogTestSuccess(t, "GetSessionByIDV1: sessionID=%s statusType=%s", sessionV1.SessionID, sessionV1.StatusType)

	// GetSessionByIDV2
	acc.LogTestStage(t, "GetSessionByIDV2", "Getting session ID=%s (v2)", sessionID)

	sessionV2, getV2Resp, err := svc.GetSessionByIDV2(ctx, sessionID)
	require.NoError(t, err)
	require.NotNil(t, sessionV2)
	assert.Equal(t, 200, getV2Resp.StatusCode())
	assert.Equal(t, sessionID, sessionV2.SessionID)
	acc.LogTestSuccess(t, "GetSessionByIDV2: sessionID=%s statusType=%s", sessionV2.SessionID, sessionV2.StatusType)
}

// TestAcceptance_JamfRemoteAssist_validation_errors verifies input validation.
func TestAcceptance_JamfRemoteAssist_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.JamfRemoteAssist

	t.Run("GetSessionByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetSessionByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "session ID is required")
	})

	t.Run("GetSessionByIDV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetSessionByIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "session ID is required")
	})
}
