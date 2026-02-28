package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_apps"
	"github.com/stretchr/testify/assert"
)

// =============================================================================
// Acceptance Tests: Mobile Device Apps
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ReinstallAppConfigV1(ctx, request) - Redeploys managed app configuration
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_MobileDeviceApps_validation_errors
//     -- Cases: nil request, empty reinstall code
//
// Notes
// -----------------------------------------------------------------------------
//   • Full lifecycle requires the $APP_CONFIG_REINSTALL_CODE from a managed app
//   • This code is device-specific and cannot be sourced from the API alone
//   • The endpoint does not require authorization, only the reinstall code
//
// =============================================================================

// TestAcceptance_MobileDeviceApps_validation_errors verifies input validation.
func TestAcceptance_MobileDeviceApps_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceApps

	t.Run("ReinstallAppConfigV1_NilRequest", func(t *testing.T) {
		_, err := svc.ReinstallAppConfigV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("ReinstallAppConfigV1_EmptyReinstallCode", func(t *testing.T) {
		_, err := svc.ReinstallAppConfigV1(context.Background(), &mobile_device_apps.RequestReinstallAppConfig{
			ReinstallCode: "",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "reinstallCode is required")
	})
}
