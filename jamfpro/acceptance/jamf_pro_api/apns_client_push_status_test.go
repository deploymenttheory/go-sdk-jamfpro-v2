package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/apns_client_push_status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_APNSClientPushStatus_list_v1 tests retrieving MDM clients with push notifications disabled.
func TestAcceptance_APNSClientPushStatus_list_v1(t *testing.T) {
	acc.RequireClient(t)
	client := acc.Client

	// Test 1: List all clients with push disabled (no filter)
	t.Run("ListAll", func(t *testing.T) {
		result, resp, err := client.APNSClientPushStatus.ListV1(context.Background(), nil)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)
		assert.GreaterOrEqual(t, result.TotalCount, 0)
		t.Logf("Found %d total clients with push notifications disabled", result.TotalCount)
	})

	// Test 2: List with RSQL filter for mobile devices only
	t.Run("FilterByMobileDevices", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `deviceType=="MOBILE_DEVICE"`,
		}

		result, resp, err := client.APNSClientPushStatus.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify all returned entries are mobile devices
		for _, entry := range result.Results {
			assert.Equal(t, "MOBILE_DEVICE", entry.DeviceType, "Expected all entries to be mobile devices")
		}
		t.Logf("Found %d mobile devices with push disabled", len(result.Results))
	})

	// Test 3: List with RSQL filter for computers only
	t.Run("FilterByComputers", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `deviceType=="COMPUTER"`,
		}

		result, resp, err := client.APNSClientPushStatus.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify all returned entries are computers
		for _, entry := range result.Results {
			assert.Equal(t, "COMPUTER", entry.DeviceType, "Expected all entries to be computers")
		}
		t.Logf("Found %d computers with push disabled", len(result.Results))
	})

	// Test 4: List with RSQL filter for date range
	t.Run("FilterByDateRange", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `disabledAt>2024-01-01T00:00:00Z`,
			"sort":   "pushDisabledTime:desc",
		}

		result, resp, err := client.APNSClientPushStatus.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
		t.Logf("Found %d devices with push disabled after 2024-01-01", len(result.Results))
	})

	// Test 5: List with complex RSQL filter
	t.Run("ComplexRSQLFilter", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `deviceType=="COMPUTER";disabledAt>2024-01-01T00:00:00Z`,
			"sort":   "pushDisabledTime:asc",
		}

		result, resp, err := client.APNSClientPushStatus.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify all returned entries match the filter criteria
		for _, entry := range result.Results {
			assert.Equal(t, "COMPUTER", entry.DeviceType)
		}
		t.Logf("Found %d computers with push disabled after 2024-01-01", len(result.Results))
	})

	// Test 6: List with pagination parameters
	t.Run("WithPagination", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"page":      "0",
			"page-size": "5",
			"sort":      "pushDisabledTime:desc",
		}

		result, resp, err := client.APNSClientPushStatus.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
		// Note: GetPaginated fetches all pages, so we should get all results
		t.Logf("Retrieved %d entries (pagination handled automatically)", len(result.Results))
	})

	// Test 7: List with sorting by different fields
	t.Run("WithSortingByDeviceType", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"sort": "deviceType:asc,pushDisabledTime:desc",
		}

		result, resp, err := client.APNSClientPushStatus.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
		t.Logf("Retrieved %d entries sorted by device type then push disabled time", len(result.Results))
	})

	// Test 8: List with sorting by management ID
	t.Run("WithSortingByManagementId", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"sort": "managementId:asc",
		}

		result, resp, err := client.APNSClientPushStatus.ListV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
		t.Logf("Retrieved %d entries sorted by management ID ascending", len(result.Results))
	})
}

// TestAcceptance_APNSClientPushStatus_enable_all_clients_v1 tests enabling push for all clients.
func TestAcceptance_APNSClientPushStatus_enable_all_clients_v1(t *testing.T) {
	acc.RequireClient(t)
	client := acc.Client

	resp, err := client.APNSClientPushStatus.EnableAllClientsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	// API may return 200 or 202 depending on implementation
	assert.True(t, resp.StatusCode == 200 || resp.StatusCode == 202, "expected 200 or 202, got %d", resp.StatusCode)
	t.Logf("EnableAllClientsV1 completed with status %d", resp.StatusCode)
}

// TestAcceptance_APNSClientPushStatus_get_enable_all_clients_status_v1 tests retrieving enable-all-clients status.
func TestAcceptance_APNSClientPushStatus_get_enable_all_clients_status_v1(t *testing.T) {
	acc.RequireClient(t)
	client := acc.Client

	result, resp, err := client.APNSClientPushStatus.GetEnableAllClientsStatusV1(context.Background())
	// 404 is valid if no recent enable-all-clients request exists
	if err != nil && resp != nil && resp.StatusCode == 404 {
		t.Log("No recent enable-all-clients request (404) - skipping status check")
		return
	}
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, []string{"QUEUED", "STARTED", "COMPLETED"}, result.Status)
	t.Logf("Enable-all-clients status: %s (requested: %s)", result.Status, result.RequestedTime)
}

// TestAcceptance_APNSClientPushStatus_enable_client_v1 tests enabling push for a single client.
// This test requires a valid managementId from a device with push disabled.
// If no such device exists, the test may fail with 404 - we log and skip in that case.
func TestAcceptance_APNSClientPushStatus_enable_client_v1(t *testing.T) {
	acc.RequireClient(t)
	client := acc.Client

	// First list clients with push disabled to get a managementId
	listResult, listResp, err := client.APNSClientPushStatus.ListV1(context.Background(), map[string]string{"page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, listResult)
	assert.Equal(t, 200, listResp.StatusCode)

	if listResult.TotalCount == 0 {
		t.Skip("No clients with push disabled - cannot test EnableClientV1")
	}

	managementID := listResult.Results[0].ManagementID
	req := &apns_client_push_status.EnableClientRequest{ManagementID: managementID}

	resp, err := client.APNSClientPushStatus.EnableClientV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
	t.Logf("Successfully enabled push for client %s", managementID)
}
