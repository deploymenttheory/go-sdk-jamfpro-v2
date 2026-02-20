package removeable_mac_addresses

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/removeable_mac_addresses/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh RemoveableMacAddressesMock.
func setupMockService(t *testing.T) (*Service, *mocks.RemoveableMacAddressesMock) {
	t.Helper()
	mock := mocks.NewRemoveableMacAddressesMock()
	return NewService(mock), mock
}

// =============================================================================
// ListRemoveableMacAddresses
// =============================================================================

func TestUnitListRemoveableMacAddresses_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListRemoveableMacAddressesMock()

	result, resp, err := svc.ListRemoveableMacAddresses(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
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

func TestUnitGetRemoveableMacAddressByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetRemoveableMacAddressByIDMock()

	result, resp, err := svc.GetRemoveableMacAddressByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AA:BB:CC:DD:EE:FF", result.Name)
}

func TestUnitGetRemoveableMacAddressByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetRemoveableMacAddressByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

func TestUnitGetRemoveableMacAddressByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetRemoveableMacAddressByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

func TestUnitGetRemoveableMacAddressByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetRemoveableMacAddressByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetRemoveableMacAddressByName
// =============================================================================

func TestUnitGetRemoveableMacAddressByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetRemoveableMacAddressByNameMock()

	result, resp, err := svc.GetRemoveableMacAddressByName(context.Background(), "AA:BB:CC:DD:EE:FF")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AA:BB:CC:DD:EE:FF", result.Name)
}

func TestUnitGetRemoveableMacAddressByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetRemoveableMacAddressByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address name is required")
}

// =============================================================================
// CreateRemoveableMacAddress
// =============================================================================

func TestUnitCreateRemoveableMacAddress_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateRemoveableMacAddressMock()

	req := &RequestRemoveableMacAddress{Name: "AA:BB:CC:DD:EE:FF"}
	result, resp, err := svc.CreateRemoveableMacAddress(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
	assert.Equal(t, "AA:BB:CC:DD:EE:FF", result.Name)
}

func TestUnitCreateRemoveableMacAddress_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateRemoveableMacAddress(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateRemoveableMacAddress_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestRemoveableMacAddress{Name: "AA:BB:CC:DD:EE:FF"}
	result, resp, err := svc.CreateRemoveableMacAddress(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateRemoveableMacAddressByID
// =============================================================================

func TestUnitUpdateRemoveableMacAddressByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateRemoveableMacAddressByIDMock()

	req := &RequestRemoveableMacAddress{Name: "Updated MAC Address"}
	result, resp, err := svc.UpdateRemoveableMacAddressByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Updated MAC Address", result.Name)
}

func TestUnitUpdateRemoveableMacAddressByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestRemoveableMacAddress{Name: "Updated MAC Address"}
	result, resp, err := svc.UpdateRemoveableMacAddressByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

func TestUnitUpdateRemoveableMacAddressByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateRemoveableMacAddressByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateRemoveableMacAddressByName
// =============================================================================

func TestUnitUpdateRemoveableMacAddressByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateRemoveableMacAddressByNameMock()

	req := &RequestRemoveableMacAddress{Name: "Updated MAC Address"}
	result, resp, err := svc.UpdateRemoveableMacAddressByName(context.Background(), "AA:BB:CC:DD:EE:FF", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Updated MAC Address", result.Name)
}

func TestUnitUpdateRemoveableMacAddressByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestRemoveableMacAddress{Name: "Updated MAC Address"}
	result, resp, err := svc.UpdateRemoveableMacAddressByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address name is required")
}

func TestUnitUpdateRemoveableMacAddressByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateRemoveableMacAddressByName(context.Background(), "AA:BB:CC:DD:EE:FF", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteRemoveableMacAddressByID
// =============================================================================

func TestUnitDeleteRemoveableMacAddressByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteRemoveableMacAddressByIDMock()

	resp, err := svc.DeleteRemoveableMacAddressByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteRemoveableMacAddressByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteRemoveableMacAddressByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

func TestUnitDeleteRemoveableMacAddressByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteRemoveableMacAddressByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
}

// =============================================================================
// DeleteRemoveableMacAddressByName
// =============================================================================

func TestUnitDeleteRemoveableMacAddressByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteRemoveableMacAddressByNameMock()

	resp, err := svc.DeleteRemoveableMacAddressByName(context.Background(), "AA:BB:CC:DD:EE:FF")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteRemoveableMacAddressByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteRemoveableMacAddressByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "removeable MAC address name is required")
}
