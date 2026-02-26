package mobile_devices_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_devices"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_devices/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDevices_List(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterListMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "iPhone-01", resp.Results[0].Name)
	assert.Equal(t, 1, resp.Results[0].ID)
	assert.Equal(t, "John's iPhone", resp.Results[0].DeviceName)
	assert.Equal(t, "iPad-01", resp.Results[1].Name)
	assert.Equal(t, 2, resp.Results[1].ID)
}

func TestUnit_MobileDevices_GetByID(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterGetMobileDeviceByIDMock()
	svc := mobile_devices.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "iPhone-01", resp.General.Name)
	assert.Equal(t, "John's iPhone", resp.General.DeviceName)
	assert.Equal(t, "DNQPX1234567", resp.General.SerialNumber)
	assert.Equal(t, "00008030-001234567890001E", resp.General.UDID)
	assert.Equal(t, "iPhone", resp.General.Model)
	assert.Equal(t, "iPhone14,2", resp.General.ModelIdentifier)
}

func TestUnit_MobileDevices_GetByID_EmptyID(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device ID cannot be empty")
}

func TestUnit_MobileDevices_GetByName(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterGetMobileDeviceByNameMock()
	svc := mobile_devices.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "iPhone-01")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "iPhone-01", resp.General.Name)
}

func TestUnit_MobileDevices_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device name cannot be empty")
}

func TestUnit_MobileDevices_GetByIDAndDataSubset(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterGetMobileDeviceByIDAndDataSubsetMock()
	svc := mobile_devices.NewService(mockClient)

	resp, _, err := svc.GetByIDAndDataSubset(context.Background(), "1", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "iPhone-01", resp.General.Name)
}

func TestUnit_MobileDevices_GetByIDAndDataSubset_EmptyID(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	_, _, err := svc.GetByIDAndDataSubset(context.Background(), "", "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device ID cannot be empty")
}

func TestUnit_MobileDevices_GetByIDAndDataSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	_, _, err := svc.GetByIDAndDataSubset(context.Background(), "1", "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "data subset cannot be empty")
}

func TestUnit_MobileDevices_GetByNameAndDataSubset(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterGetMobileDeviceByNameAndDataSubsetMock()
	svc := mobile_devices.NewService(mockClient)

	resp, _, err := svc.GetByNameAndDataSubset(context.Background(), "iPhone-01", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "iPhone-01", resp.General.Name)
}

func TestUnit_MobileDevices_GetByNameAndDataSubset_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	_, _, err := svc.GetByNameAndDataSubset(context.Background(), "", "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device name cannot be empty")
}

func TestUnit_MobileDevices_Create(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterCreateMobileDeviceMock()
	svc := mobile_devices.NewService(mockClient)

	device := &mobile_devices.ResponseMobileDevice{
		General: mobile_devices.MobileDeviceSubsetGeneral{
			DisplayName: "test-device-01",
			DeviceName:  "Test iPhone",
			Name:        "test-device-01",
			SerialNumber: "TEST1234567",
			UDID:        "00008030-001234567890001E",
			Model:       "iPhone",
			ModelIdentifier: "iPhone14,2",
		},
	}

	resp, _, err := svc.Create(context.Background(), device)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.General.ID)
	assert.Equal(t, "test-device-01", resp.General.Name)
}

func TestUnit_MobileDevices_Create_NilDevice(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device is required")
}

func TestUnit_MobileDevices_UpdateByID(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterUpdateMobileDeviceByIDMock()
	svc := mobile_devices.NewService(mockClient)

	device := &mobile_devices.ResponseMobileDevice{
		General: mobile_devices.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "iPhone-01-Updated",
			DeviceName:  "John's iPhone Updated",
			Name:        "iPhone-01-Updated",
			SerialNumber: "DNQPX1234567",
			UDID:        "00008030-001234567890001E",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), "1", device)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "iPhone-01-Updated", resp.General.Name)
}

func TestUnit_MobileDevices_UpdateByID_EmptyID(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	device := &mobile_devices.ResponseMobileDevice{
		General: mobile_devices.MobileDeviceSubsetGeneral{Name: "test"},
	}

	_, _, err := svc.UpdateByID(context.Background(), "", device)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device ID cannot be empty")
}

func TestUnit_MobileDevices_UpdateByName(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterUpdateMobileDeviceByNameMock()
	svc := mobile_devices.NewService(mockClient)

	device := &mobile_devices.ResponseMobileDevice{
		General: mobile_devices.MobileDeviceSubsetGeneral{
			ID:          1,
			DisplayName: "iPhone-01-Updated",
			DeviceName:  "John's iPhone Updated",
			Name:        "iPhone-01-Updated",
			SerialNumber: "DNQPX1234567",
			UDID:        "00008030-001234567890001E",
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "iPhone-01", device)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "iPhone-01-Updated", resp.General.Name)
}

func TestUnit_MobileDevices_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	device := &mobile_devices.ResponseMobileDevice{
		General: mobile_devices.MobileDeviceSubsetGeneral{Name: "test"},
	}

	_, _, err := svc.UpdateByName(context.Background(), "", device)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device name cannot be empty")
}

func TestUnit_MobileDevices_DeleteByID(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterDeleteMobileDeviceByIDMock()
	svc := mobile_devices.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), "1")

	require.NoError(t, err)
}

func TestUnit_MobileDevices_DeleteByID_EmptyID(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device ID cannot be empty")
}

func TestUnit_MobileDevices_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterDeleteMobileDeviceByNameMock()
	svc := mobile_devices.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "iPhone-01")

	require.NoError(t, err)
}

func TestUnit_MobileDevices_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device name cannot be empty")
}

func TestUnit_MobileDevices_NotFound(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := mobile_devices.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), "999")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_MobileDevices_List_Error(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_MobileDevices_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.GetByName(context.Background(), "iPhone")
	require.Error(t, err)
}

func TestUnit_MobileDevices_GetByIDAndDataSubset_Error(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.GetByIDAndDataSubset(context.Background(), "1", "General")
	require.Error(t, err)
}

func TestUnit_MobileDevices_GetByNameAndDataSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.GetByNameAndDataSubset(context.Background(), "iPhone", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "data subset cannot be empty")
}

func TestUnit_MobileDevices_GetByNameAndDataSubset_Error(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.GetByNameAndDataSubset(context.Background(), "iPhone", "General")
	require.Error(t, err)
}

func TestUnit_MobileDevices_Create_Error(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.Create(context.Background(), &mobile_devices.ResponseMobileDevice{})
	require.Error(t, err)
}

func TestUnit_MobileDevices_UpdateByID_NilDevice(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device is required")
}

func TestUnit_MobileDevices_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), "1", &mobile_devices.ResponseMobileDevice{})
	require.Error(t, err)
}

func TestUnit_MobileDevices_UpdateByName_NilDevice(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "iPhone", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device is required")
}

func TestUnit_MobileDevices_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "iPhone", &mobile_devices.ResponseMobileDevice{})
	require.Error(t, err)
}

func TestUnit_MobileDevices_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, err := svc.DeleteByID(context.Background(), "1")
	require.Error(t, err)
}

func TestUnit_MobileDevices_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDevicesMock()
	svc := mobile_devices.NewService(mockClient)
	_, err := svc.DeleteByName(context.Background(), "iPhone")
	require.Error(t, err)
}
