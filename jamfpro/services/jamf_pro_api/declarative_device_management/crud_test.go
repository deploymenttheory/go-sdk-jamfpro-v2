package declarative_device_management

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/declarative_device_management/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForceSyncV1(t *testing.T) {
	mock := mocks.NewDeclarativeDeviceManagementMock()
	mock.RegisterForceSyncMock("test-client-id")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.ForceSyncV1(ctx, "test-client-id")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestForceSyncV1_EmptyClientManagementID(t *testing.T) {
	mock := mocks.NewDeclarativeDeviceManagementMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.ForceSyncV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestGetStatusItemsV1(t *testing.T) {
	mock := mocks.NewDeclarativeDeviceManagementMock()
	mock.RegisterGetStatusItemsMock("test-client-id")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetStatusItemsV1(ctx, "test-client-id")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result.StatusItems, 2)
	assert.Equal(t, "device.model.family", result.StatusItems[0].Key)
	assert.Equal(t, "Mac", result.StatusItems[0].Value)
}

func TestGetStatusItemsV1_EmptyClientManagementID(t *testing.T) {
	mock := mocks.NewDeclarativeDeviceManagementMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetStatusItemsV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestGetStatusItemByKeyV1(t *testing.T) {
	mock := mocks.NewDeclarativeDeviceManagementMock()
	mock.RegisterGetStatusItemByKeyMock("test-client-id", "device.model.family")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetStatusItemByKeyV1(ctx, "test-client-id", "device.model.family")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "device.model.family", result.Key)
	assert.Equal(t, "Mac", result.Value)
}

func TestGetStatusItemByKeyV1_EmptyClientManagementID(t *testing.T) {
	mock := mocks.NewDeclarativeDeviceManagementMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetStatusItemByKeyV1(ctx, "", "device.model.family")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestGetStatusItemByKeyV1_EmptyKey(t *testing.T) {
	mock := mocks.NewDeclarativeDeviceManagementMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetStatusItemByKeyV1(ctx, "test-client-id", "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "key is required")
}
