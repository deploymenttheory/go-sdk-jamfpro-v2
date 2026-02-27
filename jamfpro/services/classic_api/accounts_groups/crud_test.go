package accounts_groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts_groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh AccountGroupsMock.
func setupMockService(t *testing.T) (*Service, *mocks.AccountGroupsMock) {
	t.Helper()
	mock := mocks.NewAccountGroupsMock()
	return NewService(mock), mock
}

// =============================================================================
// GetAccountGroupByID
// =============================================================================

func TestUnit_AccountsGroups_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountGroupByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testgroup1", result.Name)
	assert.Equal(t, "Full Access", result.AccessLevel)
	assert.Equal(t, "Administrator", result.PrivilegeSet)
	require.Len(t, result.Members, 2)
	assert.Equal(t, 1, result.Members[0].ID)
	assert.Equal(t, "testuser1", result.Members[0].Name)
}

func TestUnit_AccountsGroups_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

func TestUnit_AccountsGroups_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

func TestUnit_AccountsGroups_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetAccountGroupByName
// =============================================================================

func TestUnit_AccountsGroups_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountGroupByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "testgroup1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testgroup1", result.Name)
	assert.Equal(t, "Full Access", result.AccessLevel)
}

func TestUnit_AccountsGroups_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group name is required")
}

// =============================================================================
// CreateAccountGroup
// =============================================================================

func TestUnit_AccountsGroups_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateAccountGroupMock()

	req := &RequestAccountGroup{
		Name:         "testgroup1",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
}

func TestUnit_AccountsGroups_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_AccountsGroups_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestAccountGroup{Name: "testgroup1"}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateAccountGroupByID
// =============================================================================

func TestUnit_AccountsGroups_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountGroupByIDMock()

	req := &RequestAccountGroup{
		Name:         "updatedgroup",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
	}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_AccountsGroups_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAccountGroup{Name: "updatedgroup"}
	result, resp, err := svc.UpdateByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

func TestUnit_AccountsGroups_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateAccountGroupByName
// =============================================================================

func TestUnit_AccountsGroups_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountGroupByNameMock()

	req := &RequestAccountGroup{
		Name:         "updatedgroup",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
	}
	result, resp, err := svc.UpdateByName(context.Background(), "testgroup1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_AccountsGroups_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAccountGroup{Name: "updatedgroup"}
	result, resp, err := svc.UpdateByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group name is required")
}

func TestUnit_AccountsGroups_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "testgroup1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteAccountGroupByID
// =============================================================================

func TestUnit_AccountsGroups_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccountGroupByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_AccountsGroups_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

func TestUnit_AccountsGroups_DeleteByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

// =============================================================================
// DeleteAccountGroupByName
// =============================================================================

func TestUnit_AccountsGroups_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccountGroupByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "testgroup1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_AccountsGroups_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group name is required")
}
