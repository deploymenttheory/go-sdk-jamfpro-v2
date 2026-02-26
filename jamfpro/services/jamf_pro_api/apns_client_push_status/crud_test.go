package apns_client_push_status

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/apns_client_push_status/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.APNSClientPushStatusMock) {
	t.Helper()
	mock := mocks.NewAPNSClientPushStatusMock()
	return NewService(mock), mock
}

// Test ListV1 with success response
func TestUnit_APNSClientPushStatus_ListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 3, result.TotalCount)
	assert.Len(t, result.Results, 3)
	assert.Equal(t, "MOBILE_DEVICE", result.Results[0].DeviceType)
	assert.Equal(t, "101", result.Results[0].ClientID)
	assert.Equal(t, "2024-01-15T10:30:00Z", result.Results[0].DisabledAt)
	assert.Equal(t, "101", result.Results[0].ManagementID)
}

// Test ListV1 with RSQL filter query for device type
func TestUnit_APNSClientPushStatus_ListV1_WithRSQLFilterDeviceType(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	rsqlQuery := map[string]string{
		"filter": `deviceType=="MOBILE_DEVICE"`,
		"sort":   "pushDisabledTime:asc",
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify RSQL query was passed to the client
	assert.NotNil(t, mock.LastRSQLQuery)
	assert.Equal(t, `deviceType=="MOBILE_DEVICE"`, mock.LastRSQLQuery["filter"])
	assert.Equal(t, "pushDisabledTime:asc", mock.LastRSQLQuery["sort"])
}

// Test ListV1 with RSQL filter query for date
func TestUnit_APNSClientPushStatus_ListV1_WithRSQLFilterDate(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	rsqlQuery := map[string]string{
		"filter": `disabledAt>2024-01-01T00:00:00Z`,
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify RSQL query was passed
	assert.Equal(t, `disabledAt>2024-01-01T00:00:00Z`, mock.LastRSQLQuery["filter"])
}

// Test ListV1 with pagination parameters
func TestUnit_APNSClientPushStatus_ListV1_WithPagination(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "50",
		"sort":      "pushDisabledTime:desc",
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify pagination parameters were passed
	assert.Equal(t, "0", mock.LastRSQLQuery["page"])
	assert.Equal(t, "50", mock.LastRSQLQuery["page-size"])
	assert.Equal(t, "pushDisabledTime:desc", mock.LastRSQLQuery["sort"])
}

// Test ListV1 with complex RSQL filter
func TestUnit_APNSClientPushStatus_ListV1_WithComplexRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	rsqlQuery := map[string]string{
		"filter": `deviceType=="COMPUTER";disabledAt>2024-01-01T00:00:00Z`,
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify complex RSQL query was passed
	assert.Equal(t, `deviceType=="COMPUTER";disabledAt>2024-01-01T00:00:00Z`, mock.LastRSQLQuery["filter"])
}

// Test ListV1 with sorting by management ID
func TestUnit_APNSClientPushStatus_ListV1_WithSorting(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	rsqlQuery := map[string]string{
		"sort": "managementId:asc",
	}

	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify sort parameter was passed
	assert.Equal(t, "managementId:asc", mock.LastRSQLQuery["sort"])
}

// Test EnableAllClientsV1 with success response
func TestUnit_APNSClientPushStatus_EnableAllClientsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterEnableAllClientsMock()

	resp, err := svc.EnableAllClientsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 202, resp.StatusCode)
}

// Test GetEnableAllClientsStatusV1 with success response
func TestUnit_APNSClientPushStatus_GetEnableAllClientsStatusV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetEnableAllClientsStatusMock()

	result, resp, err := svc.GetEnableAllClientsStatusV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "2024-11-06T14:30:00Z", result.RequestedTime)
	assert.Equal(t, "QUEUED", result.Status)
	require.NotNil(t, result.ProcessedTime)
	assert.Equal(t, "2024-11-06T15:30:00Z", *result.ProcessedTime)
}

// Test EnableClientV1 with success response
func TestUnit_APNSClientPushStatus_EnableClientV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterEnableClientMock()

	req := &EnableClientRequest{
		ManagementID: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
	}

	resp, err := svc.EnableClientV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

// Test EnableClientV1 with nil request returns error
func TestUnit_APNSClientPushStatus_EnableClientV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.EnableClientV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// Test EnableClientV1 with empty managementId returns error
func TestUnit_APNSClientPushStatus_EnableClientV1_EmptyManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &EnableClientRequest{ManagementID: ""}

	resp, err := svc.EnableClientV1(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "managementId is required")
}

func TestUnit_APNSClientPushStatus_EnableAllClientsV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.EnableAllClientsV1(context.Background())
	require.Error(t, err)
}

func TestUnit_APNSClientPushStatus_GetEnableAllClientsStatusV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetEnableAllClientsStatusV1(context.Background())
	require.Error(t, err)
}

func TestUnit_APNSClientPushStatus_ListV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.ListV1(context.Background(), nil)
	require.Error(t, err)
}
