package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
)

// =============================================================================
// Acceptance Tests: MDM Commands
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • BlankPush(ctx, clientManagementIDs) - Sends blank push to devices
//   • SendCommand(ctx, req) - Sends MDM command for creation and queuing
//   • DeployPackage(ctx, req) - Deploys package via MDM
//   • RenewProfile(ctx, req) - Renews MDM profiles for device UDIDs
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_MDM_ValidationErrors
//     -- Cases: Empty IDs, nil requests
//
// Notes
// -----------------------------------------------------------------------------
//   • MDM commands require enrolled devices; acceptance tests focus on validation
//   • Blank push, send command, deploy package, and renew profile are not exercised
//     against a real tenant to avoid affecting production devices
//
// =============================================================================
// TestAcceptance_MDM_ValidationErrors
// =============================================================================

func TestAcceptance_MDM_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MDM

	t.Run("BlankPush_EmptyClientManagementIDs", func(t *testing.T) {
		_, _, err := svc.BlankPush(context.Background(), []string{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementIDs is required")
	})

	t.Run("BlankPush_NilClientManagementIDs", func(t *testing.T) {
		_, _, err := svc.BlankPush(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementIDs is required")
	})

	t.Run("SendCommand_NilRequest", func(t *testing.T) {
		_, _, err := svc.SendCommand(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeployPackage_NilRequest", func(t *testing.T) {
		_, _, err := svc.DeployPackage(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("RenewProfile_NilRequest", func(t *testing.T) {
		_, _, err := svc.RenewProfile(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})
}

// =============================================================================
// TestAcceptance_MDM_BlankPush (optional - requires enrolled device)
// =============================================================================
// Uncomment and set JAMF_MDM_TEST_DEVICE_ID to test blank push against a real device.
//
// func TestAcceptance_MDM_blank_push(t *testing.T) {
// 	acc.RequireClient(t)
// 	deviceID := os.Getenv("JAMF_MDM_TEST_DEVICE_ID")
// 	if deviceID == "" {
// 		t.Skip("JAMF_MDM_TEST_DEVICE_ID not set, skipping blank push test")
// 	}
// 	svc := acc.Client.MDM
// 	result, resp, err := svc.BlankPush(context.Background(), []string{deviceID})
// 	require.NoError(t, err)
// 	require.NotNil(t, result)
// 	assert.Equal(t, 200, resp.StatusCode)
// }
