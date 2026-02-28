package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mdm_renewal"
	"github.com/stretchr/testify/assert"
)

// =============================================================================
// Acceptance Tests: MDM Renewal
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetDeviceCommonDetailsV1(ctx, clientManagementID) - Gets device MDM details
//   • GetRenewalStrategiesV1(ctx, clientManagementID) - Gets renewal strategies
//   • DeleteRenewalStrategiesV1(ctx, clientManagementID) - Deletes renewal strategies
//   • UpdateDeviceCommonDetailsV1(ctx, request) - Updates device common details (PATCH)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_MDMRenewal_validation_errors
//     -- Cases: Empty/nil arguments for all endpoints
//
// Notes
// -----------------------------------------------------------------------------
//   • Full lifecycle requires a real device with MDM renewal issues
//   • clientManagementId corresponds to a device's management ID in Jamf Pro
//   • Validation errors are the only tests that reliably run in all environments
//
// =============================================================================

// TestAcceptance_MDMRenewal_validation_errors verifies input validation for all operations.
func TestAcceptance_MDMRenewal_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MDMRenewal

	t.Run("GetDeviceCommonDetailsV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetDeviceCommonDetailsV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementId is required")
	})

	t.Run("GetRenewalStrategiesV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetRenewalStrategiesV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementId is required")
	})

	t.Run("DeleteRenewalStrategiesV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteRenewalStrategiesV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementId is required")
	})

	t.Run("UpdateDeviceCommonDetailsV1_NilRequest", func(t *testing.T) {
		_, err := svc.UpdateDeviceCommonDetailsV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateDeviceCommonDetailsV1_EmptyClientManagementID", func(t *testing.T) {
		_, err := svc.UpdateDeviceCommonDetailsV1(context.Background(), &mdm_renewal.RequestDeviceCommonDetailsUpdate{
			ClientManagementID: "",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "clientManagementId is required")
	})
}
