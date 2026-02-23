package command_flush_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/command_flush"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/command_flush/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---- Computers - all status combinations ----

func TestUnit_CommandFlush_FlushByIDAndStatus_Computers_Pending(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushComputersPendingMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "computers", "123", "Pending")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_Computers_Failed(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushComputersFailedMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "computers", "123", "Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_Computers_PendingAndFailed(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushComputersPendingAndFailedMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "computers", "123", "Pending+Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

// ---- Computer Groups - all status combinations ----

func TestUnit_CommandFlush_FlushByIDAndStatus_ComputerGroups_Pending(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushComputerGroupsPendingMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "computergroups", "456", "Pending")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_ComputerGroups_Failed(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushComputerGroupsFailedMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "computergroups", "456", "Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_ComputerGroups_PendingAndFailed(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushComputerGroupsPendingAndFailedMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "computergroups", "456", "Pending+Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

// ---- Mobile Devices - all status combinations ----

func TestUnit_CommandFlush_FlushByIDAndStatus_MobileDevices_Pending(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushMobileDevicesPendingMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "mobiledevices", "789", "Pending")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_MobileDevices_Failed(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushMobileDevicesFailedMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "mobiledevices", "789", "Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_MobileDevices_PendingAndFailed(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushMobileDevicesPendingAndFailedMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "mobiledevices", "789", "Pending+Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

// ---- Mobile Device Groups - all status combinations ----

func TestUnit_CommandFlush_FlushByIDAndStatus_MobileDeviceGroups_Pending(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushMobileDeviceGroupsPendingMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "mobiledevicegroups", "101112", "Pending")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_MobileDeviceGroups_Failed(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushMobileDeviceGroupsFailedMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "mobiledevicegroups", "101112", "Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_MobileDeviceGroups_PendingAndFailed(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushMobileDeviceGroupsPendingAndFailedMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "mobiledevicegroups", "101112", "Pending+Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

// ---- FlushWithXML ----

func TestUnit_CommandFlush_FlushWithXML(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushWithXMLMock()
	svc := command_flush.NewService(mockClient)

	req := &command_flush.RequestCommandFlush{
		Status: "Pending",
		MobileDevices: &command_flush.MobileDevices{
			MobileDevice: []command_flush.DeviceID{
				{ID: 1},
				{ID: 2},
			},
		},
	}

	resp, err := svc.FlushWithXML(context.Background(), req)

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

// ---- Validation errors ----

func TestUnit_CommandFlush_FlushByIDAndStatus_InvalidIDType(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	svc := command_flush.NewService(mockClient)

	_, err := svc.FlushByIDAndStatus(context.Background(), "invalid", "123", "Pending")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid idType")
}

func TestUnit_CommandFlush_FlushByIDAndStatus_InvalidStatus(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	svc := command_flush.NewService(mockClient)

	_, err := svc.FlushByIDAndStatus(context.Background(), "computers", "123", "Invalid")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid status")
}

func TestUnit_CommandFlush_FlushWithXML_InvalidStatus(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	svc := command_flush.NewService(mockClient)

	req := &command_flush.RequestCommandFlush{
		Status: "Invalid",
		MobileDevices: &command_flush.MobileDevices{
			MobileDevice: []command_flush.DeviceID{{ID: 1}},
		},
	}

	_, err := svc.FlushWithXML(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid status")
}

func TestUnit_CommandFlush_FlushWithXML_NoDevices(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	svc := command_flush.NewService(mockClient)

	req := &command_flush.RequestCommandFlush{
		Status: "Pending",
	}

	_, err := svc.FlushWithXML(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "at least one device list")
}
