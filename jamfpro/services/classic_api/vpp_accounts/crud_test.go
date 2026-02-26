package vpp_accounts

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/vpp_accounts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.VPPAccountsMock) {
	t.Helper()
	mock := mocks.NewVPPAccountsMock()
	return NewService(mock), mock
}

// =============================================================================
// ListVPPAccounts
// =============================================================================

func TestUnit_VppAccounts_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListVPPAccountsMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Production VPP", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Education VPP", result.Results[1].Name)
}

// =============================================================================
// GetVPPAccountByID
// =============================================================================

func TestUnit_VppAccounts_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetVPPAccountByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Production VPP", result.Name)
	assert.Equal(t, "admin@example.com", result.Contact)
	assert.Equal(t, "US", result.Country)
	assert.True(t, result.PopulateCatalogFromVPPContent)
	assert.True(t, result.NotifyDisassociation)
}

func TestUnit_VppAccounts_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
}

func TestUnit_VppAccounts_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), -1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
}

func TestUnit_VppAccounts_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, _, err := svc.GetByID(context.Background(), 999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

// =============================================================================
// CreateVPPAccount
// =============================================================================

func TestUnit_VppAccounts_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateVPPAccountMock()

	req := &RequestVPPAccount{
		Name:    "Production VPP",
		Country: "US",
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Production VPP", result.Name)
}

func TestUnit_VppAccounts_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.Create(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_VppAccounts_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestVPPAccount{Name: "Production VPP"}
	_, _, err := svc.Create(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdateVPPAccountByID
// =============================================================================

func TestUnit_VppAccounts_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateVPPAccountByIDMock()

	req := &RequestVPPAccount{Name: "Production VPP Updated"}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Production VPP Updated", result.Name)
}

func TestUnit_VppAccounts_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 0, &RequestVPPAccount{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
}

func TestUnit_VppAccounts_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteVPPAccountByID
// =============================================================================

func TestUnit_VppAccounts_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteVPPAccountByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_VppAccounts_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
}
