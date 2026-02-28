package devices

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/devices/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.DevicesMock) {
	t.Helper()
	mock := mocks.NewDevicesMock()
	return NewService(mock), mock
}

func TestUnit_Devices_NewService(t *testing.T) {
	mock := mocks.NewDevicesMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
	assert.NotNil(t, svc.client)
}

func TestUnit_Devices_GetGroupsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetGroupsMock()

	result, resp, err := svc.GetGroupsV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 2)
	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, "All Devices", result[0].Name)
	assert.Equal(t, "2", result[1].ID)
	assert.Equal(t, "Marketing Laptops", result[1].Name)
}

func TestUnit_Devices_GetGroupsV1_EmptyArray(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetGroupsEmptyMock()

	result, resp, err := svc.GetGroupsV1(context.Background(), "2")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Empty(t, result)
}

func TestUnit_Devices_GetGroupsV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetGroupsV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Devices_GetGroupsV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetGroupsV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Contains(t, err.Error(), "NOT-FOUND")
}

func TestUnit_Devices_GetGroupsV1_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetGroupsV1(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get groups for device ID 1")
}
