package computer_history_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_history"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_history/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_ComputerHistory_GetByID(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetByIDMock()
	svc := computer_history.NewComputerHistory(mockClient)

	resp, _, err := svc.GetByID(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test-MacBook-Pro", resp.General.Name)
	assert.Equal(t, "C02XYZ123456", resp.General.SerialNumber)
	assert.Len(t, resp.ComputerUsageLogs, 1)
	assert.Equal(t, "Login", resp.ComputerUsageLogs[0].Event)
	assert.Len(t, resp.Audits, 1)
	assert.Len(t, resp.PolicyLogs, 1)
	require.NotNil(t, resp.Commands)
	assert.Len(t, resp.Commands.Completed, 1)
	assert.Equal(t, "Update Inventory", resp.Commands.Completed[0].Name)
}

func TestUnit_ComputerHistory_GetByID_EmptyID(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ID cannot be empty")
}

func TestUnit_ComputerHistory_GetByIDAndSubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetByIDAndSubsetMock()
	svc := computer_history.NewComputerHistory(mockClient)

	resp, _, err := svc.GetByIDAndSubset(context.Background(), "1", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Test-MacBook-Pro", resp.General.Name)
}

func TestUnit_ComputerHistory_GetByIDAndSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByIDAndSubset(context.Background(), "1", "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

func TestUnit_ComputerHistory_GetByName(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetByNameMock()
	svc := computer_history.NewComputerHistory(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Test-MacBook-Pro")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "Test-MacBook-Pro", resp.General.Name)
}

func TestUnit_ComputerHistory_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name cannot be empty")
}

func TestUnit_ComputerHistory_GetByNameAndSubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetByNameAndSubsetMock()
	svc := computer_history.NewComputerHistory(mockClient)

	resp, _, err := svc.GetByNameAndSubset(context.Background(), "Test-MacBook-Pro", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "Test-MacBook-Pro", resp.General.Name)
}

func TestUnit_ComputerHistory_GetByUDID(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetByUDIDMock()
	svc := computer_history.NewComputerHistory(mockClient)

	udid := "00000000-0000-0000-0000-000000000001"
	resp, _, err := svc.GetByUDID(context.Background(), udid)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, udid, resp.General.UDID)
}

func TestUnit_ComputerHistory_GetByUDIDAndSubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetByUDIDAndSubsetMock()
	svc := computer_history.NewComputerHistory(mockClient)

	udid := "00000000-0000-0000-0000-000000000001"
	resp, _, err := svc.GetByUDIDAndSubset(context.Background(), udid, "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, udid, resp.General.UDID)
}

func TestUnit_ComputerHistory_GetBySerialNumber(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetBySerialNumberMock()
	svc := computer_history.NewComputerHistory(mockClient)

	serial := "C02XYZ123456"
	resp, _, err := svc.GetBySerialNumber(context.Background(), serial)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, serial, resp.General.SerialNumber)
}

func TestUnit_ComputerHistory_GetBySerialNumberAndSubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetBySerialNumberAndSubsetMock()
	svc := computer_history.NewComputerHistory(mockClient)

	serial := "C02XYZ123456"
	resp, _, err := svc.GetBySerialNumberAndSubset(context.Background(), serial, "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, serial, resp.General.SerialNumber)
}

func TestUnit_ComputerHistory_GetByMACAddress(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetByMACAddressMock()
	svc := computer_history.NewComputerHistory(mockClient)

	macAddr := "00:11:22:33:44:55"
	resp, _, err := svc.GetByMACAddress(context.Background(), macAddr)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, macAddr, resp.General.MacAddress)
}

func TestUnit_ComputerHistory_GetByMACAddressAndSubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterGetByMACAddressAndSubsetMock()
	svc := computer_history.NewComputerHistory(mockClient)

	macAddr := "00:11:22:33:44:55"
	resp, _, err := svc.GetByMACAddressAndSubset(context.Background(), macAddr, "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, macAddr, resp.General.MacAddress)
}

func TestUnit_ComputerHistory_GetByMACAddress_Empty(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByMACAddress(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "MAC address cannot be empty")
}

func TestUnit_ComputerHistory_NotFound(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByID(context.Background(), "999")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_ComputerHistory_GetByIDAndSubset_EmptyID(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByIDAndSubset(context.Background(), "", "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ID cannot be empty")
}

func TestUnit_ComputerHistory_GetByNameAndSubset_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByNameAndSubset(context.Background(), "", "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "name cannot be empty")
}

func TestUnit_ComputerHistory_GetByNameAndSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByNameAndSubset(context.Background(), "Test-MacBook-Pro", "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

func TestUnit_ComputerHistory_GetByUDID_Empty(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByUDID(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "UDID cannot be empty")
}

func TestUnit_ComputerHistory_GetByUDIDAndSubset_EmptyUDID(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByUDIDAndSubset(context.Background(), "", "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "UDID cannot be empty")
}

func TestUnit_ComputerHistory_GetByUDIDAndSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByUDIDAndSubset(context.Background(), "00000000-0000-0000-0000-000000000001", "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

func TestUnit_ComputerHistory_GetBySerialNumber_Empty(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetBySerialNumber(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "serial number cannot be empty")
}

func TestUnit_ComputerHistory_GetBySerialNumberAndSubset_EmptySerial(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetBySerialNumberAndSubset(context.Background(), "", "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "serial number cannot be empty")
}

func TestUnit_ComputerHistory_GetBySerialNumberAndSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetBySerialNumberAndSubset(context.Background(), "C02XYZ123456", "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

func TestUnit_ComputerHistory_GetByMACAddressAndSubset_EmptyMAC(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByMACAddressAndSubset(context.Background(), "", "General")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "MAC address cannot be empty")
}

func TestUnit_ComputerHistory_GetByMACAddressAndSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewComputerHistoryMock()
	svc := computer_history.NewComputerHistory(mockClient)

	_, _, err := svc.GetByMACAddressAndSubset(context.Background(), "00:11:22:33:44:55", "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

// Verify url.PathEscape behavior for MAC address.
// Note: colons are valid in URL paths and are not escaped by PathEscape.
func TestUnit_ComputerHistory_MACAddressPathEscape(t *testing.T) {
	macAddr := "00:11:22:33:44:55"
	escaped := url.PathEscape(macAddr)
	// PathEscape does not escape colons in paths - they are valid path characters
	assert.Equal(t, "00:11:22:33:44:55", escaped)
}
