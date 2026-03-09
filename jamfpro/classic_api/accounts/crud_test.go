package accounts

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/accounts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh AccountsMock.
func setupMockService(t *testing.T) (*Accounts, *mocks.AccountsMock) {
	t.Helper()
	mock := mocks.NewAccountsMock()
	return NewAccounts(mock), mock
}

// =============================================================================
// ListAccounts
// =============================================================================

func TestUnit_Accounts_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccountsMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
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

func TestUnit_Accounts_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testuser1", result.Name)
	assert.Equal(t, "Test User One", result.FullName)
	assert.Equal(t, "testuser1@example.com", result.Email)
}

func TestUnit_Accounts_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

func TestUnit_Accounts_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

func TestUnit_Accounts_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// =============================================================================
// GetAccountByName
// =============================================================================

func TestUnit_Accounts_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "testuser1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testuser1", result.Name)
	assert.Equal(t, "Test User One", result.FullName)
}

func TestUnit_Accounts_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account name is required")
}

// =============================================================================
// CreateAccount
// =============================================================================

func TestUnit_Accounts_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateAccountMock()

	req := &RequestAccount{
		Name:     "testuser1",
		FullName: "Test User One",
		Email:    "testuser1@example.com",
		Enabled:  "Enabled",
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 100, result.ID)
	assert.Equal(t, "testuser1", result.Name)
}

func TestUnit_Accounts_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Accounts_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestAccount{Name: "testuser1"}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

// =============================================================================
// UpdateAccountByID
// =============================================================================

func TestUnit_Accounts_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountByIDMock()

	req := &RequestAccount{
		Name:     "updateduser",
		FullName: "Updated User",
		Email:    "updateduser@example.com",
	}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "updateduser", result.Name)
	assert.Equal(t, "Updated User", result.FullName)
}

func TestUnit_Accounts_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAccount{Name: "updateduser"}
	result, resp, err := svc.UpdateByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

func TestUnit_Accounts_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateAccountByName
// =============================================================================

func TestUnit_Accounts_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountByNameMock()

	req := &RequestAccount{
		Name:     "updateduser",
		FullName: "Updated User",
		Email:    "updateduser@example.com",
	}
	result, resp, err := svc.UpdateByName(context.Background(), "testuser1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "updateduser", result.Name)
}

func TestUnit_Accounts_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAccount{Name: "updateduser"}
	result, resp, err := svc.UpdateByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account name is required")
}

func TestUnit_Accounts_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "testuser1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteAccountByID
// =============================================================================

func TestUnit_Accounts_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccountByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Accounts_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

func TestUnit_Accounts_DeleteByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account ID must be a positive integer")
}

// =============================================================================
// DeleteAccountByName
// =============================================================================

func TestUnit_Accounts_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccountByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "testuser1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Accounts_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account name is required")
}
