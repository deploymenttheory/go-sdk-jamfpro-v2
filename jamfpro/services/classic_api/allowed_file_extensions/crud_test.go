package allowed_file_extensions

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/allowed_file_extensions/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh AllowedFileExtensionsMock.
func setupMockService(t *testing.T) (*Service, *mocks.AllowedFileExtensionsMock) {
	t.Helper()
	mock := mocks.NewAllowedFileExtensionsMock()
	return NewService(mock), mock
}

// =============================================================================
// ListAllowedFileExtensions
// =============================================================================

func TestUnitListAllowedFileExtensions_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAllowedFileExtensionsMock()

	result, resp, err := svc.ListAllowedFileExtensions(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "dmg", result.Results[0].Extension)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "pkg", result.Results[1].Extension)
}

// =============================================================================
// GetAllowedFileExtensionByID
// =============================================================================

func TestUnitGetAllowedFileExtensionByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAllowedFileExtensionByIDMock()

	result, resp, err := svc.GetAllowedFileExtensionByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "dmg", result.Extension)
}

func TestUnitGetAllowedFileExtensionByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAllowedFileExtensionByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "allowed file extension ID must be a positive integer")
}

func TestUnitGetAllowedFileExtensionByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAllowedFileExtensionByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "allowed file extension ID must be a positive integer")
}

func TestUnitGetAllowedFileExtensionByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetAllowedFileExtensionByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetAllowedFileExtensionByExtension
// =============================================================================

func TestUnitGetAllowedFileExtensionByExtension_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAllowedFileExtensionByExtensionMock()

	result, resp, err := svc.GetAllowedFileExtensionByExtension(context.Background(), "dmg")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "dmg", result.Extension)
}

func TestUnitGetAllowedFileExtensionByExtension_Empty(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAllowedFileExtensionByExtension(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "extension is required")
}

// =============================================================================
// CreateAllowedFileExtension
// =============================================================================

func TestUnitCreateAllowedFileExtension_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateAllowedFileExtensionMock()

	req := &RequestAllowedFileExtension{Extension: "zip"}
	result, resp, err := svc.CreateAllowedFileExtension(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 3, result.ID)
	assert.Equal(t, "zip", result.Extension)
}

func TestUnitCreateAllowedFileExtension_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateAllowedFileExtension(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateAllowedFileExtension_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestAllowedFileExtension{Extension: "dmg"}
	result, resp, err := svc.CreateAllowedFileExtension(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// DeleteAllowedFileExtensionByID
// =============================================================================

func TestUnitDeleteAllowedFileExtensionByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAllowedFileExtensionByIDMock()

	resp, err := svc.DeleteAllowedFileExtensionByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteAllowedFileExtensionByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAllowedFileExtensionByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "allowed file extension ID must be a positive integer")
}
