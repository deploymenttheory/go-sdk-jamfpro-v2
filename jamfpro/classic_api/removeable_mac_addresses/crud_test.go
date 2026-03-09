package removeable_mac_addresses

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/removeable_mac_addresses/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh RemoveableMacAddressesMock.
func setupMockService(t *testing.T) (*RemoveableMacAddresses, *mocks.RemoveableMacAddressesMock) {
	t.Helper()
	mock := mocks.NewRemoveableMacAddressesMock()
	return NewRemoveableMacAddresses(mock), mock
}

// =============================================================================
// ListRemoveableMacAddresses
// =============================================================================

func TestUnit_RemoveableMacAddresses_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "AA:BB:CC:DD:EE:FF", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "11:22:33:44:55:66", result.Results[1].Name)
}

// =============================================================================
// GetRemoveableMacAddressByID
// =============================================================================

func TestUnit_RemoveableMacAddresses_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AA:BB:CC:DD:EE:FF", result.Name)
}

func TestUnit_RemoveableMacAddresses_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

func TestUnit_RemoveableMacAddresses_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

func TestUnit_RemoveableMacAddresses_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// =============================================================================
// GetRemoveableMacAddressByName
// =============================================================================

func TestUnit_RemoveableMacAddresses_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "AA:BB:CC:DD:EE:FF")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AA:BB:CC:DD:EE:FF", result.Name)
}

func TestUnit_RemoveableMacAddresses_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address name is required")
}

// =============================================================================
// CreateRemoveableMacAddress
// =============================================================================

func TestUnit_RemoveableMacAddresses_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestRemoveableMacAddress{Name: "AA:BB:CC:DD:EE:FF"}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 100, result.ID)
	assert.Equal(t, "AA:BB:CC:DD:EE:FF", result.Name)
}

func TestUnit_RemoveableMacAddresses_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_RemoveableMacAddresses_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestRemoveableMacAddress{Name: "AA:BB:CC:DD:EE:FF"}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

// =============================================================================
// UpdateRemoveableMacAddressByID
// =============================================================================

func TestUnit_RemoveableMacAddresses_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDMock()

	req := &RequestRemoveableMacAddress{Name: "Updated MAC Address"}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Updated MAC Address", result.Name)
}

func TestUnit_RemoveableMacAddresses_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestRemoveableMacAddress{Name: "Updated MAC Address"}
	result, resp, err := svc.UpdateByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

func TestUnit_RemoveableMacAddresses_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateRemoveableMacAddressByName
// =============================================================================

func TestUnit_RemoveableMacAddresses_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByNameMock()

	req := &RequestRemoveableMacAddress{Name: "Updated MAC Address"}
	result, resp, err := svc.UpdateByName(context.Background(), "AA:BB:CC:DD:EE:FF", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Updated MAC Address", result.Name)
}

func TestUnit_RemoveableMacAddresses_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestRemoveableMacAddress{Name: "Updated MAC Address"}
	result, resp, err := svc.UpdateByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address name is required")
}

func TestUnit_RemoveableMacAddresses_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "AA:BB:CC:DD:EE:FF", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteRemoveableMacAddressByID
// =============================================================================

func TestUnit_RemoveableMacAddresses_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_RemoveableMacAddresses_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

func TestUnit_RemoveableMacAddresses_DeleteByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

// =============================================================================
// DeleteRemoveableMacAddressByName
// =============================================================================

func TestUnit_RemoveableMacAddresses_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "AA:BB:CC:DD:EE:FF")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_RemoveableMacAddresses_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address name is required")
}
