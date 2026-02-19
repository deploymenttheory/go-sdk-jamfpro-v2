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

func TestUnitListVPPAccounts_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListVPPAccountsMock()

	result, resp, err := svc.ListVPPAccounts(context.Background())
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

func TestUnitGetVPPAccountByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetVPPAccountByIDMock()

	result, resp, err := svc.GetVPPAccountByID(context.Background(), 1)
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

func TestUnitGetVPPAccountByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetVPPAccountByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
}

func TestUnitGetVPPAccountByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetVPPAccountByID(context.Background(), -1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
}

func TestUnitGetVPPAccountByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, _, err := svc.GetVPPAccountByID(context.Background(), 999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

// =============================================================================
// CreateVPPAccount
// =============================================================================

func TestUnitCreateVPPAccount_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateVPPAccountMock()

	req := &RequestVPPAccount{
		Name:    "Production VPP",
		Country: "US",
	}
	result, resp, err := svc.CreateVPPAccount(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Production VPP", result.Name)
}

func TestUnitCreateVPPAccount_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.CreateVPPAccount(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateVPPAccount_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestVPPAccount{Name: "Production VPP"}
	_, _, err := svc.CreateVPPAccount(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdateVPPAccountByID
// =============================================================================

func TestUnitUpdateVPPAccountByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateVPPAccountByIDMock()

	req := &RequestVPPAccount{Name: "Production VPP Updated"}
	result, resp, err := svc.UpdateVPPAccountByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Production VPP Updated", result.Name)
}

func TestUnitUpdateVPPAccountByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateVPPAccountByID(context.Background(), 0, &RequestVPPAccount{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
}

func TestUnitUpdateVPPAccountByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateVPPAccountByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteVPPAccountByID
// =============================================================================

func TestUnitDeleteVPPAccountByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteVPPAccountByIDMock()

	resp, err := svc.DeleteVPPAccountByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteVPPAccountByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteVPPAccountByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "VPP account ID must be a positive integer")
}
