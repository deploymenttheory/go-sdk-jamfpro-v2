package device_communication_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/device_communication_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.DeviceCommunicationSettingsMock) {
	t.Helper()
	mock := mocks.NewDeviceCommunicationSettingsMock()
	return NewService(mock), mock
}

func TestUnit_DeviceCommunicationSettings_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.AutoRenewMobileDeviceMdmProfileWhenCaRenewed)
	assert.Equal(t, 30, result.MdmProfileMobileDeviceExpirationLimitInDays)
}

func TestUnit_DeviceCommunicationSettings_UpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_DeviceCommunicationSettings_UpdateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterPutMock()

	request := &ResourceDeviceCommunicationSettings{
		MdmProfileMobileDeviceExpirationLimitInDays: 30,
		MdmProfileComputerExpirationLimitInDays:    30,
	}
	result, resp, err := svc.UpdateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_DeviceCommunicationSettings_GetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Settings updated", result.Results[0].Note)
}

func TestUnit_DeviceCommunicationSettings_GetHistoryV1_WithParams(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "date:desc"}
	result, resp, err := svc.GetHistoryV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_DeviceCommunicationSettings_GetHistoryV1_ApiError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryErrorMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
	assert.Contains(t, err.Error(), "failed to get device communication settings history")
}

func TestUnit_DeviceCommunicationSettings_GetHistoryV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryInvalidJSONMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to unmarshal page")
}

func TestUnit_DeviceCommunicationSettings_GetHistoryV1_InvalidHistoryItem(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryInvalidItemMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to decode history item")
}

func TestUnit_DeviceCommunicationSettings_GetV1_ApiError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetErrorMock()

	result, resp, err := svc.GetV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestUnit_DeviceCommunicationSettings_GetV1_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	result, resp, err := svc.GetV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response")
}

func TestUnit_DeviceCommunicationSettings_UpdateV1_ApiError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterPutErrorMock()

	request := &ResourceDeviceCommunicationSettings{MdmProfileMobileDeviceExpirationLimitInDays: 30}
	result, resp, err := svc.UpdateV1(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestUnit_DeviceCommunicationSettings_AddHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNotesMock()

	req := &RequestAddHistoryNotes{Note: "Added via SDK"}
	result, resp, err := svc.AddHistoryNotesV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Contains(t, result.Href, "/api/v1/device-communication-settings/history/1")
}

func TestUnit_DeviceCommunicationSettings_AddHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddHistoryNotesV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_DeviceCommunicationSettings_AddHistoryNotesV1_EmptyNote(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAddHistoryNotes{Note: ""}
	result, resp, err := svc.AddHistoryNotesV1(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "note is required")
}

func TestUnit_DeviceCommunicationSettings_AddHistoryNotesV1_ApiError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNotesErrorMock()

	req := &RequestAddHistoryNotes{Note: "Test note"}
	result, resp, err := svc.AddHistoryNotesV1(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestUnit_DeviceCommunicationSettings_NewService(t *testing.T) {
	mock := mocks.NewDeviceCommunicationSettingsMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
}
