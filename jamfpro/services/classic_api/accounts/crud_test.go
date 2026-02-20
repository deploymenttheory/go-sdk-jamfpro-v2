package accounts

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh AccountsMock.
func setupMockService(t *testing.T) (*Service, *mocks.AccountsMock) {
	t.Helper()
	mock := mocks.NewAccountsMock()
	return NewService(mock), mock
}

// =============================================================================
// ListAccounts
// =============================================================================

func TestUnitListAccounts_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccountsMock()

	result, resp, err := svc.ListAccounts(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.Users, 2)
	assert.Equal(t, 1, result.Users[0].ID)
	assert.Equal(t, "testuser1", result.Users[0].Name)
	assert.Equal(t, 2, result.Users[1].ID)
	assert.Equal(t, "testuser2", result.Users[1].Name)
	require.Len(t, result.Groups, 1)
	assert.Equal(t, 3, result.Groups[0].ID)
	assert.Equal(t, "testgroup1", result.Groups[0].Name)
}

// =============================================================================
// GetAccountByID
// =============================================================================

func TestUnitGetAccountByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountByIDMock()

	result, resp, err := svc.GetAccountByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testuser1", result.Name)
	assert.Equal(t, "Test User One", result.FullName)
	assert.Equal(t, "testuser1@example.com", result.Email)
}

func TestUnitGetAccountByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAccountByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

func TestUnitGetAccountByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAccountByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

func TestUnitGetAccountByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetAccountByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetAccountByName
// =============================================================================

func TestUnitGetAccountByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountByNameMock()

	result, resp, err := svc.GetAccountByName(context.Background(), "testuser1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testuser1", result.Name)
	assert.Equal(t, "Test User One", result.FullName)
}

func TestUnitGetAccountByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAccountByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account name is required")
}

// =============================================================================
// CreateAccount
// =============================================================================

func TestUnitCreateAccount_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateAccountMock()

	req := &RequestAccount{
		Name:     "testuser1",
		FullName: "Test User One",
		Email:    "testuser1@example.com",
		Enabled:  "Enabled",
	}
	result, resp, err := svc.CreateAccount(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
	assert.Equal(t, "testuser1", result.Name)
}

func TestUnitCreateAccount_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateAccount(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateAccount_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestAccount{Name: "testuser1"}
	result, resp, err := svc.CreateAccount(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateAccountByID
// =============================================================================

func TestUnitUpdateAccountByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountByIDMock()

	req := &RequestAccount{
		Name:     "updateduser",
		FullName: "Updated User",
		Email:    "updateduser@example.com",
	}
	result, resp, err := svc.UpdateAccountByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "updateduser", result.Name)
	assert.Equal(t, "Updated User", result.FullName)
}

func TestUnitUpdateAccountByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAccount{Name: "updateduser"}
	result, resp, err := svc.UpdateAccountByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

func TestUnitUpdateAccountByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAccountByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateAccountByName
// =============================================================================

func TestUnitUpdateAccountByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountByNameMock()

	req := &RequestAccount{
		Name:     "updateduser",
		FullName: "Updated User",
		Email:    "updateduser@example.com",
	}
	result, resp, err := svc.UpdateAccountByName(context.Background(), "testuser1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "updateduser", result.Name)
}

func TestUnitUpdateAccountByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAccount{Name: "updateduser"}
	result, resp, err := svc.UpdateAccountByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account name is required")
}

func TestUnitUpdateAccountByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAccountByName(context.Background(), "testuser1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteAccountByID
// =============================================================================

func TestUnitDeleteAccountByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccountByIDMock()

	resp, err := svc.DeleteAccountByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteAccountByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAccountByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

func TestUnitDeleteAccountByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAccountByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

// =============================================================================
// DeleteAccountByName
// =============================================================================

func TestUnitDeleteAccountByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccountByNameMock()

	resp, err := svc.DeleteAccountByName(context.Background(), "testuser1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteAccountByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAccountByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account name is required")
}
