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

func TestUnit_CommandFlush_FlushByIDAndStatus_Computers(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushByIDAndStatusComputersMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "computers", "123", "Pending")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_ComputerGroups(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushByIDAndStatusComputerGroupsMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "computergroups", "456", "Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_MobileDevices(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushByIDAndStatusMobileDevicesMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "mobiledevices", "789", "Pending+Failed")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestUnit_CommandFlush_FlushByIDAndStatus_MobileDeviceGroups(t *testing.T) {
	mockClient := mocks.NewCommandFlushMock()
	mockClient.RegisterFlushByIDAndStatusMobileDeviceGroupsMock()
	svc := command_flush.NewService(mockClient)

	resp, err := svc.FlushByIDAndStatus(context.Background(), "mobiledevicegroups", "101112", "Pending")

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

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
