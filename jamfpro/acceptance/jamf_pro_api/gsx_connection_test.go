package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: GSX Connection
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Retrieves GSX connection settings
//   • UpdateV1(ctx, request) - Updates GSX connection settings via PATCH
//   • GetHistoryV1(ctx, rsqlQuery) - Retrieves GSX connection history
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration Testing
//     -- Reason: GSX Connection is a singleton settings object (no Create/Delete)
//     -- Tests: TestAcceptance_GSXConnection_GetAndUpdate
//     -- Flow: Get current → Update → Get (verify) → Restore original
//
//   ✓ Pattern 4: History/Audit Trail Testing
//     -- Reason: Service provides GetHistoryV1 for audit trail
//     -- Tests: TestAcceptance_GSXConnection_History
//     -- Flow: Get history → Verify structure
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_GSXConnection_ValidationErrors
//     -- Cases: Nil request validation
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get operations (retrieve current settings)
//   ✓ Update operations (modify settings via PATCH)
//   ✓ History operations (retrieve audit trail with optional sorting)
//   ✓ Input validation and error handling
//   ✓ Settings restoration (restore original config after tests)
//
// Notes
// -----------------------------------------------------------------------------
//   • GSX Connection is a settings object, not a CRUD resource
//   • Tests preserve original settings and restore them after modifications
//   • History endpoint supports optional RSQL query for sorting
//   • All tests use proper cleanup to restore state
//
// =============================================================================

func TestAcceptance_GSXConnection_get_and_update(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.GSXConnection
	ctx := context.Background()

	// Get current settings to preserve them
	original, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, original)
	assert.Equal(t, 200, resp.StatusCode)

	// Register cleanup to restore original settings
	t.Cleanup(func() {
		_, _, err := svc.UpdateV1(ctx, original)
		if err != nil {
			t.Logf("Warning: Failed to restore original GSX connection settings: %v", err)
		}
	})

	// Verify structure of retrieved settings
	assert.NotNil(t, original.GsxKeystore)

	// Update with modified settings (toggle enabled state as a test)
	modified := *original
	modified.Enabled = !original.Enabled

	updated, resp, err := svc.UpdateV1(ctx, &modified)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify update was applied
	assert.Equal(t, modified.Enabled, updated.Enabled)

	// Get again to verify persistence
	current, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, modified.Enabled, current.Enabled)
}

func TestAcceptance_GSXConnection_history(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.GSXConnection
	ctx := context.Background()

	// Get history without filtering
	history, resp, err := svc.GetHistoryV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, history)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify response structure
	assert.GreaterOrEqual(t, history.TotalCount, 0)

	// If there are history entries, verify structure
	if history.TotalCount > 0 && len(history.Results) > 0 {
		entry := history.Results[0]
		assert.NotEmpty(t, entry.ID)
		assert.NotEmpty(t, entry.Username)
		assert.NotEmpty(t, entry.Date)
	}

	// Test with sorting parameter
	rsqlQuery := map[string]string{"sort": "date:desc"}
	historyFiltered, resp, err := svc.GetHistoryV1(ctx, rsqlQuery)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, historyFiltered)
}

func TestAcceptance_GSXConnection_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.GSXConnection
	ctx := context.Background()

	// Test nil request validation
	result, resp, err := svc.UpdateV1(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}
