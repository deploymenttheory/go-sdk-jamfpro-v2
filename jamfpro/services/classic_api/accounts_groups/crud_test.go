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

func TestUnitGetAccountGroupByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountGroupByIDMock()

	result, resp, err := svc.GetAccountGroupByID(context.Background(), 1)
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

func TestUnitGetAccountGroupByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAccountGroupByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

func TestUnitGetAccountGroupByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAccountGroupByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

func TestUnitGetAccountGroupByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetAccountGroupByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetAccountGroupByName
// =============================================================================

func TestUnitGetAccountGroupByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountGroupByNameMock()

	result, resp, err := svc.GetAccountGroupByName(context.Background(), "testgroup1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "testgroup1", result.Name)
	assert.Equal(t, "Full Access", result.AccessLevel)
}

func TestUnitGetAccountGroupByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAccountGroupByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group name is required")
}

// =============================================================================
// CreateAccountGroup
// =============================================================================

func TestUnitCreateAccountGroup_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateAccountGroupMock()

	req := &RequestAccountGroup{
		Name:         "testgroup1",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
	}
	result, resp, err := svc.CreateAccountGroup(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
}

func TestUnitCreateAccountGroup_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateAccountGroup(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateAccountGroup_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestAccountGroup{Name: "testgroup1"}
	result, resp, err := svc.CreateAccountGroup(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateAccountGroupByID
// =============================================================================

func TestUnitUpdateAccountGroupByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountGroupByIDMock()

	req := &RequestAccountGroup{
		Name:         "updatedgroup",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
	}
	result, resp, err := svc.UpdateAccountGroupByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "updatedgroup", result.Name)
}

func TestUnitUpdateAccountGroupByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAccountGroup{Name: "updatedgroup"}
	result, resp, err := svc.UpdateAccountGroupByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

func TestUnitUpdateAccountGroupByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAccountGroupByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateAccountGroupByName
// =============================================================================

func TestUnitUpdateAccountGroupByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountGroupByNameMock()

	req := &RequestAccountGroup{
		Name:         "updatedgroup",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
	}
	result, resp, err := svc.UpdateAccountGroupByName(context.Background(), "testgroup1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "updatedgroup", result.Name)
}

func TestUnitUpdateAccountGroupByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestAccountGroup{Name: "updatedgroup"}
	result, resp, err := svc.UpdateAccountGroupByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group name is required")
}

func TestUnitUpdateAccountGroupByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAccountGroupByName(context.Background(), "testgroup1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteAccountGroupByID
// =============================================================================

func TestUnitDeleteAccountGroupByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccountGroupByIDMock()

	resp, err := svc.DeleteAccountGroupByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteAccountGroupByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAccountGroupByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

func TestUnitDeleteAccountGroupByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAccountGroupByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group ID must be a positive integer")
}

// =============================================================================
// DeleteAccountGroupByName
// =============================================================================

func TestUnitDeleteAccountGroupByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccountGroupByNameMock()

	resp, err := svc.DeleteAccountGroupByName(context.Background(), "testgroup1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteAccountGroupByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAccountGroupByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "account group name is required")
}
