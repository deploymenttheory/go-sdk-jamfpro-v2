package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: User Sessions (v1)
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   - GetActiveV1(ctx) - Returns detailed information about currently logged-in users
//   - GetCountV1(ctx)  - Returns the number of currently logged-in users
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   - Pattern 3: Read-Only Information
//
// =============================================================================

// TestAcceptance_UserSessions_get_active_v1 verifies the active user sessions endpoint.
func TestAcceptance_UserSessions_get_active_v1(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "get", "retrieving active user sessions")
	result, resp, err := acc.Client.JamfProAPI.UserSessions.GetActiveV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.GreaterOrEqual(t, len(result), 0)
	acc.LogTestSuccess(t, "active user sessions count=%d", len(result))
}

// TestAcceptance_UserSessions_get_count_v1 verifies the user session count endpoint.
func TestAcceptance_UserSessions_get_count_v1(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "get", "retrieving user session count")
	result, resp, err := acc.Client.JamfProAPI.UserSessions.GetCountV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.GreaterOrEqual(t, result.Count, 0)
	acc.LogTestSuccess(t, "user session count=%d", result.Count)
}
