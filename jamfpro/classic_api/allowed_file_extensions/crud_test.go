package allowed_file_extensions

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/allowed_file_extensions/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh AllowedFileExtensionsMock.
func setupMockService(t *testing.T) (*AllowedFileExtensions, *mocks.AllowedFileExtensionsMock) {
	t.Helper()
	mock := mocks.NewAllowedFileExtensionsMock()
	return NewAllowedFileExtensions(mock), mock
}

// =============================================================================
// ListAllowedFileExtensions
// =============================================================================

func TestUnit_AllowedFileExtensions_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
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

func TestUnit_AllowedFileExtensions_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "dmg", result.Extension)
}

func TestUnit_AllowedFileExtensions_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "allowed file extension ID must be a positive integer")
}

func TestUnit_AllowedFileExtensions_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "allowed file extension ID must be a positive integer")
}

func TestUnit_AllowedFileExtensions_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// =============================================================================
// GetAllowedFileExtensionByExtension
// =============================================================================

func TestUnit_AllowedFileExtensions_GetByExtension_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByExtensionMock()

	result, resp, err := svc.GetByExtension(context.Background(), "dmg")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "dmg", result.Extension)
}

func TestUnit_AllowedFileExtensions_GetByExtension_Empty(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByExtension(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "extension is required")
}

// =============================================================================
// CreateAllowedFileExtension
// =============================================================================

func TestUnit_AllowedFileExtensions_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestAllowedFileExtension{Extension: "zip"}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 3, result.ID)
	assert.Equal(t, "zip", result.Extension)
}

func TestUnit_AllowedFileExtensions_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_AllowedFileExtensions_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestAllowedFileExtension{Extension: "dmg"}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

// =============================================================================
// DeleteAllowedFileExtensionByID
// =============================================================================

func TestUnit_AllowedFileExtensions_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_AllowedFileExtensions_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "allowed file extension ID must be a positive integer")
}
