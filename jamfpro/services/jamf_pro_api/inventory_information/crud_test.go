package inventory_information

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/inventory_information/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.InventoryInformationMock) {
	t.Helper()
	mock := mocks.NewInventoryInformationMock()
	return NewService(mock), mock
}

func TestUnit_InventoryInformation_NewService(t *testing.T) {
	mock := mocks.NewInventoryInformationMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
}

func TestUnit_InventoryInformation_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV1Mock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1200, result.ManagedComputers)
	assert.Equal(t, 110, result.UnmanagedComputers)
	assert.Equal(t, 850, result.ManagedDevices)
	assert.Equal(t, 45, result.UnmanagedDevices)
}

func TestUnit_InventoryInformation_GetV1_Error_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get inventory information")
}

func TestUnit_InventoryInformation_GetV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV1NotFoundErrorMock()

	result, resp, err := svc.GetV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Contains(t, err.Error(), "Jamf Pro API error")
}

func TestUnit_InventoryInformation_GetV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterInvalidJSONMock()

	result, resp, err := svc.GetV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, err.Error(), "failed to get inventory information")
}

