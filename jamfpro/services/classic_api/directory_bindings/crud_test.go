package directory_bindings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/directory_bindings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.DirectoryBindingsMock) {
	t.Helper()
	mock := mocks.NewDirectoryBindingsMock()
	return NewService(mock), mock
}

// =============================================================================
// ListDirectoryBindings
// =============================================================================

func TestUnit_DirectoryBindings_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListDirectoryBindingsMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "AD Binding", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "LDAP Binding", result.Results[1].Name)
}

// =============================================================================
// GetDirectoryBindingByID
// =============================================================================

func TestUnit_DirectoryBindings_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDirectoryBindingByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AD Binding", result.Name)
	assert.Equal(t, "example.com", result.Domain)
	assert.Equal(t, "Active Directory", result.Type)
}

func TestUnit_DirectoryBindings_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
}

func TestUnit_DirectoryBindings_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), -1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
}

func TestUnit_DirectoryBindings_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, _, err := svc.GetByID(context.Background(), 999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

// =============================================================================
// GetDirectoryBindingByName
// =============================================================================

func TestUnit_DirectoryBindings_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDirectoryBindingByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "AD Binding")
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AD Binding", result.Name)
}

func TestUnit_DirectoryBindings_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding name is required")
}

// =============================================================================
// CreateDirectoryBinding
// =============================================================================

func TestUnit_DirectoryBindings_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateDirectoryBindingMock()

	req := &RequestDirectoryBinding{
		Name:   "AD Binding",
		Domain: "example.com",
		Type:   "Active Directory",
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AD Binding", result.Name)
}

func TestUnit_DirectoryBindings_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.Create(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_DirectoryBindings_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestDirectoryBinding{Name: "AD Binding"}
	_, _, err := svc.Create(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdateDirectoryBindingByID
// =============================================================================

func TestUnit_DirectoryBindings_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDirectoryBindingByIDMock()

	req := &RequestDirectoryBinding{Name: "AD Binding Updated"}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AD Binding Updated", result.Name)
}

func TestUnit_DirectoryBindings_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 0, &RequestDirectoryBinding{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
}

func TestUnit_DirectoryBindings_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateDirectoryBindingByName
// =============================================================================

func TestUnit_DirectoryBindings_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDirectoryBindingByNameMock()

	req := &RequestDirectoryBinding{Name: "AD Binding Updated"}
	result, resp, err := svc.UpdateByName(context.Background(), "AD Binding", req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
}

func TestUnit_DirectoryBindings_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "", &RequestDirectoryBinding{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding name is required")
}

func TestUnit_DirectoryBindings_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "AD Binding", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteDirectoryBindingByID
// =============================================================================

func TestUnit_DirectoryBindings_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDirectoryBindingByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_DirectoryBindings_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
}

// =============================================================================
// DeleteDirectoryBindingByName
// =============================================================================

func TestUnit_DirectoryBindings_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDirectoryBindingByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "AD Binding")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_DirectoryBindings_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding name is required")
}
