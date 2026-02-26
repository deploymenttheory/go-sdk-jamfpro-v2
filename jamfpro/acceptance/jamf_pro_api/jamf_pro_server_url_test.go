package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Pro Server URL
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Retrieves Jamf Pro server URL settings
//   • UpdateV1(ctx, request) - Updates Jamf Pro server URL settings
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration Testing
//     -- Reason: Jamf Pro Server URL is a singleton settings object (no Create/Delete)
//     -- Tests: TestAcceptance_JamfProServerURL_GetAndUpdate
//     -- Flow: Get current → Update → Get (verify) → Restore original
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_JamfProServerURL_ValidationErrors
//     -- Cases: Nil request validation
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get operations (retrieve current settings)
//   ✓ Update operations (modify settings via PUT)
//   ✓ Input validation and error handling
//   ✓ Settings restoration (restore original config after tests)
//
// Notes
// -----------------------------------------------------------------------------
//   • Jamf Pro Server URL is a settings object, not a CRUD resource
//   • Tests preserve original settings and restore them after modifications
//   • All tests use proper cleanup to restore state
//
// =============================================================================

func TestAcceptance_JamfProServerURL_get_and_update(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProServerURL
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
			t.Logf("Warning: Failed to restore original Jamf Pro server URL settings: %v", err)
		}
	})

	// Verify structure of retrieved settings
	assert.NotEmpty(t, original.URL)

	// Update with modified settings
	modified := *original
	modified.URL = "https://jamf-test.example.com"
	modified.UnsecuredEnrollmentUrl = "http://jamf-test.example.com:8080"

	updated, resp, err := svc.UpdateV1(ctx, &modified)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify update was applied
	assert.Equal(t, modified.URL, updated.URL)
	assert.Equal(t, modified.UnsecuredEnrollmentUrl, updated.UnsecuredEnrollmentUrl)

	// Get again to verify persistence
	current, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, modified.URL, current.URL)
	assert.Equal(t, modified.UnsecuredEnrollmentUrl, current.UnsecuredEnrollmentUrl)
}

func TestAcceptance_JamfProServerURL_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProServerURL
	ctx := context.Background()

	// Test nil request validation
	result, resp, err := svc.UpdateV1(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}
