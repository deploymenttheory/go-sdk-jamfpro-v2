package patch_external_sources

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/patch_external_sources/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.PatchExternalSourcesMock) {
	t.Helper()
	mock := mocks.NewPatchExternalSourcesMock()
	return NewService(mock), mock
}

// =============================================================================
// List
// =============================================================================

func TestUnitList_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPatchExternalSourcesMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Primary Patch Source", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Secondary Patch Source", result.Results[1].Name)
}

// =============================================================================
// GetByID
// =============================================================================

func TestUnitGetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPatchExternalSourceByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary Patch Source", result.Name)
	assert.Equal(t, "patches.example.com", result.HostName)
	assert.True(t, result.SSLEnabled)
	assert.Equal(t, 443, result.Port)
}

func TestUnitGetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
}

func TestUnitGetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), -1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
}

func TestUnitGetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, _, err := svc.GetByID(context.Background(), 999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

// =============================================================================
// GetByName
// =============================================================================

func TestUnitGetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPatchExternalSourceByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "Primary Patch Source")
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary Patch Source", result.Name)
}

func TestUnitGetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source name is required")
}

// =============================================================================
// Create
// =============================================================================

func TestUnitCreate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreatePatchExternalSourceMock()

	req := &RequestPatchExternalSource{
		Name:       "Primary Patch Source",
		HostName:   "patches.example.com",
		SSLEnabled: true,
		Port:       443,
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary Patch Source", result.Name)
}

func TestUnitCreate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.Create(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreate_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestPatchExternalSource{Name: "Primary Patch Source"}
	_, _, err := svc.Create(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdateByID
// =============================================================================

func TestUnitUpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdatePatchExternalSourceByIDMock()

	req := &RequestPatchExternalSource{Name: "Primary Patch Source Updated"}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary Patch Source Updated", result.Name)
}

func TestUnitUpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 0, &RequestPatchExternalSource{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
}

func TestUnitUpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateByName
// =============================================================================

func TestUnitUpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdatePatchExternalSourceByNameMock()

	req := &RequestPatchExternalSource{Name: "Primary Patch Source Updated"}
	result, resp, err := svc.UpdateByName(context.Background(), "Primary Patch Source", req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "", &RequestPatchExternalSource{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source name is required")
}

func TestUnitUpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "Primary Patch Source", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteByID
// =============================================================================

func TestUnitDeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePatchExternalSourceByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
}
