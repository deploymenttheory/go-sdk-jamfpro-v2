package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Devices
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetGroupsV1(ctx, id) - Returns groups that the specified device belongs to
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Reason: Service is read-only; requires an existing mobile device
//     -- Tests: TestAcceptance_Devices_get_groups_v1
//     -- Flow: Get first mobile device from groups → GetGroups
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_Devices_validation_errors
//     -- Cases: GetGroupsV1("") → "id is required"
//
// Notes
// -----------------------------------------------------------------------------
//   • Device IDs are sourced from mobile device group membership
//   • Test skips gracefully if no mobile devices exist
//
// =============================================================================

// TestAcceptance_Devices_get_groups_v1 fetches device groups for an existing mobile device.
func TestAcceptance_Devices_get_groups_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Devices
	ctx := context.Background()

	// Get a mobile device ID from smart mobile device groups membership
	groupList, _, err := acc.Client.SmartMobileDeviceGroups.List(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, groupList)

	if groupList.TotalCount == 0 {
		t.Skip("No smart mobile device groups exist; skipping Devices GetGroups")
	}

	// Get membership of the first group to find a device ID
	membership, _, err := acc.Client.SmartMobileDeviceGroups.GetMembership(ctx, groupList.Results[0].GroupID, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, membership)

	if membership.TotalCount == 0 {
		t.Skip("Smart mobile device group has no members; skipping Devices GetGroups")
	}

	deviceID := membership.Results[0].MobileDeviceId
	acc.LogTestStage(t, "GetGroups", "Fetching groups for device ID=%s", deviceID)

	groups, resp, err := svc.GetGroupsV1(ctx, deviceID)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, len(groups), 0)

	acc.LogTestSuccess(t, "GetGroupsV1: deviceID=%s groupCount=%d", deviceID, len(groups))
}

// TestAcceptance_Devices_validation_errors verifies input validation.
func TestAcceptance_Devices_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Devices

	t.Run("GetGroupsV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetGroupsV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})
}
