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

func TestUnitListDirectoryBindings_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListDirectoryBindingsMock()

	result, resp, err := svc.ListDirectoryBindings(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
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

func TestUnitGetDirectoryBindingByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDirectoryBindingByIDMock()

	result, resp, err := svc.GetDirectoryBindingByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AD Binding", result.Name)
	assert.Equal(t, "example.com", result.Domain)
	assert.Equal(t, "Active Directory", result.Type)
}

func TestUnitGetDirectoryBindingByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetDirectoryBindingByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
}

func TestUnitGetDirectoryBindingByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetDirectoryBindingByID(context.Background(), -1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
}

func TestUnitGetDirectoryBindingByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, _, err := svc.GetDirectoryBindingByID(context.Background(), 999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

// =============================================================================
// GetDirectoryBindingByName
// =============================================================================

func TestUnitGetDirectoryBindingByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDirectoryBindingByNameMock()

	result, resp, err := svc.GetDirectoryBindingByName(context.Background(), "AD Binding")
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AD Binding", result.Name)
}

func TestUnitGetDirectoryBindingByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetDirectoryBindingByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding name is required")
}

// =============================================================================
// CreateDirectoryBinding
// =============================================================================

func TestUnitCreateDirectoryBinding_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateDirectoryBindingMock()

	req := &RequestDirectoryBinding{
		Name:   "AD Binding",
		Domain: "example.com",
		Type:   "Active Directory",
	}
	result, resp, err := svc.CreateDirectoryBinding(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AD Binding", result.Name)
}

func TestUnitCreateDirectoryBinding_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.CreateDirectoryBinding(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateDirectoryBinding_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestDirectoryBinding{Name: "AD Binding"}
	_, _, err := svc.CreateDirectoryBinding(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdateDirectoryBindingByID
// =============================================================================

func TestUnitUpdateDirectoryBindingByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDirectoryBindingByIDMock()

	req := &RequestDirectoryBinding{Name: "AD Binding Updated"}
	result, resp, err := svc.UpdateDirectoryBindingByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "AD Binding Updated", result.Name)
}

func TestUnitUpdateDirectoryBindingByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateDirectoryBindingByID(context.Background(), 0, &RequestDirectoryBinding{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
}

func TestUnitUpdateDirectoryBindingByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateDirectoryBindingByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateDirectoryBindingByName
// =============================================================================

func TestUnitUpdateDirectoryBindingByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDirectoryBindingByNameMock()

	req := &RequestDirectoryBinding{Name: "AD Binding Updated"}
	result, resp, err := svc.UpdateDirectoryBindingByName(context.Background(), "AD Binding", req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateDirectoryBindingByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateDirectoryBindingByName(context.Background(), "", &RequestDirectoryBinding{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding name is required")
}

func TestUnitUpdateDirectoryBindingByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateDirectoryBindingByName(context.Background(), "AD Binding", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteDirectoryBindingByID
// =============================================================================

func TestUnitDeleteDirectoryBindingByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDirectoryBindingByIDMock()

	resp, err := svc.DeleteDirectoryBindingByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteDirectoryBindingByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteDirectoryBindingByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding ID must be a positive integer")
}

// =============================================================================
// DeleteDirectoryBindingByName
// =============================================================================

func TestUnitDeleteDirectoryBindingByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDirectoryBindingByNameMock()

	resp, err := svc.DeleteDirectoryBindingByName(context.Background(), "AD Binding")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteDirectoryBindingByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteDirectoryBindingByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "directory binding name is required")
}
