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
// ListPatchExternalSources
// =============================================================================

func TestUnitListPatchExternalSources_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPatchExternalSourcesMock()

	result, resp, err := svc.ListPatchExternalSources(context.Background())
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
// GetPatchExternalSourceByID
// =============================================================================

func TestUnitGetPatchExternalSourceByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPatchExternalSourceByIDMock()

	result, resp, err := svc.GetPatchExternalSourceByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary Patch Source", result.Name)
	assert.Equal(t, "patches.example.com", result.HostName)
	assert.True(t, result.SSLEnabled)
	assert.Equal(t, 443, result.Port)
}

func TestUnitGetPatchExternalSourceByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetPatchExternalSourceByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
}

func TestUnitGetPatchExternalSourceByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetPatchExternalSourceByID(context.Background(), -1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
}

func TestUnitGetPatchExternalSourceByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, _, err := svc.GetPatchExternalSourceByID(context.Background(), 999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

// =============================================================================
// GetPatchExternalSourceByName
// =============================================================================

func TestUnitGetPatchExternalSourceByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPatchExternalSourceByNameMock()

	result, resp, err := svc.GetPatchExternalSourceByName(context.Background(), "Primary Patch Source")
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary Patch Source", result.Name)
}

func TestUnitGetPatchExternalSourceByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetPatchExternalSourceByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source name is required")
}

// =============================================================================
// CreatePatchExternalSource
// =============================================================================

func TestUnitCreatePatchExternalSource_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreatePatchExternalSourceMock()

	req := &RequestPatchExternalSource{
		Name:       "Primary Patch Source",
		HostName:   "patches.example.com",
		SSLEnabled: true,
		Port:       443,
	}
	result, resp, err := svc.CreatePatchExternalSource(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary Patch Source", result.Name)
}

func TestUnitCreatePatchExternalSource_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.CreatePatchExternalSource(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreatePatchExternalSource_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestPatchExternalSource{Name: "Primary Patch Source"}
	_, _, err := svc.CreatePatchExternalSource(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdatePatchExternalSourceByID
// =============================================================================

func TestUnitUpdatePatchExternalSourceByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdatePatchExternalSourceByIDMock()

	req := &RequestPatchExternalSource{Name: "Primary Patch Source Updated"}
	result, resp, err := svc.UpdatePatchExternalSourceByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary Patch Source Updated", result.Name)
}

func TestUnitUpdatePatchExternalSourceByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdatePatchExternalSourceByID(context.Background(), 0, &RequestPatchExternalSource{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
}

func TestUnitUpdatePatchExternalSourceByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdatePatchExternalSourceByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdatePatchExternalSourceByName
// =============================================================================

func TestUnitUpdatePatchExternalSourceByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdatePatchExternalSourceByNameMock()

	req := &RequestPatchExternalSource{Name: "Primary Patch Source Updated"}
	result, resp, err := svc.UpdatePatchExternalSourceByName(context.Background(), "Primary Patch Source", req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdatePatchExternalSourceByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdatePatchExternalSourceByName(context.Background(), "", &RequestPatchExternalSource{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source name is required")
}

func TestUnitUpdatePatchExternalSourceByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdatePatchExternalSourceByName(context.Background(), "Primary Patch Source", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeletePatchExternalSourceByID
// =============================================================================

func TestUnitDeletePatchExternalSourceByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePatchExternalSourceByIDMock()

	resp, err := svc.DeletePatchExternalSourceByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeletePatchExternalSourceByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeletePatchExternalSourceByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "patch external source ID must be a positive integer")
}
