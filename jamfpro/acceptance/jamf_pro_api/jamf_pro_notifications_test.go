package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Pro Notifications
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetForUserAndSiteV1(ctx) - Gets all notifications for current user/site
//   • DeleteByTypeAndIDV1(ctx, type, id) - Deletes a notification by type and ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Tests: TestAcceptance_JamfProNotifications_get_v1
//     -- Flow: GetForUserAndSiteV1 → verify response (may be empty)
//     -- Note: Notifications are not deleted in tests to avoid disrupting real alerts
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_JamfProNotifications_validation_errors
//     -- Cases: DeleteByTypeAndIDV1 with empty type/id
//
// Notes
// -----------------------------------------------------------------------------
//   • Notifications list may be empty in clean environments — that is acceptable
//   • Notifications are not deleted during tests to preserve real system alerts
//
// =============================================================================

// TestAcceptance_JamfProNotifications_get_v1 fetches notifications for the current user/site.
func TestAcceptance_JamfProNotifications_get_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.JamfProNotifications
	ctx := context.Background()

	acc.LogTestStage(t, "GetForUserAndSiteV1", "Getting notifications for current user and site")

	notifications, resp, err := svc.GetForUserAndSiteV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.GreaterOrEqual(t, len(notifications), 0)

	acc.LogTestSuccess(t, "GetForUserAndSiteV1: %d notification(s) returned", len(notifications))

	if len(notifications) > 0 {
		first := notifications[0]
		acc.LogTestSuccess(t, "First notification: type=%s id=%s", first.Type, first.ID)
	}
}

// TestAcceptance_JamfProNotifications_validation_errors verifies input validation.
func TestAcceptance_JamfProNotifications_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.JamfProNotifications

	t.Run("DeleteByTypeAndIDV1_EmptyType", func(t *testing.T) {
		_, err := svc.DeleteByTypeAndIDV1(context.Background(), "", "1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "notification type is required")
	})

	t.Run("DeleteByTypeAndIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByTypeAndIDV1(context.Background(), "PATCH_UPDATE", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "notification id is required")
	})
}
