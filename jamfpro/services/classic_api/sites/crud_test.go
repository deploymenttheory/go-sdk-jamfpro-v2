package sites

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/sites/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh SitesMock.
func setupMockService(t *testing.T) (*Service, *mocks.SitesMock) {
	t.Helper()
	mock := mocks.NewSitesMock()
	return NewService(mock), mock
}

// =============================================================================
// ListSites
// =============================================================================

func TestUnitListSites_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListSitesMock()

	result, resp, err := svc.ListSites(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Main Campus", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Remote Office", result.Results[1].Name)
}

// =============================================================================
// GetSiteByID
// =============================================================================

func TestUnitGetSiteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSiteByIDMock()

	result, resp, err := svc.GetSiteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Main Campus", result.Name)
}

func TestUnitGetSiteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetSiteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "site ID must be a positive integer")
}

func TestUnitGetSiteByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetSiteByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "site ID must be a positive integer")
}

func TestUnitGetSiteByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetSiteByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetSiteByName
// =============================================================================

func TestUnitGetSiteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSiteByNameMock()

	result, resp, err := svc.GetSiteByName(context.Background(), "Main Campus")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Main Campus", result.Name)
}

func TestUnitGetSiteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetSiteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "site name is required")
}

// =============================================================================
// CreateSite
// =============================================================================

func TestUnitCreateSite_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateSiteMock()

	req := &RequestSite{Name: "New Site"}
	result, resp, err := svc.CreateSite(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 3, result.ID)
	assert.Equal(t, "New Site", result.Name)
}

func TestUnitCreateSite_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateSite(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateSite_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestSite{Name: "Main Campus"}
	result, resp, err := svc.CreateSite(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateSiteByID
// =============================================================================

func TestUnitUpdateSiteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSiteByIDMock()

	req := &RequestSite{Name: "Main Campus Updated"}
	result, resp, err := svc.UpdateSiteByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Main Campus Updated", result.Name)
}

func TestUnitUpdateSiteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateSiteByID(context.Background(), 0, &RequestSite{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "site ID must be a positive integer")
}

func TestUnitUpdateSiteByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateSiteByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateSiteByName
// =============================================================================

func TestUnitUpdateSiteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSiteByNameMock()

	req := &RequestSite{Name: "Main Campus Updated"}
	result, resp, err := svc.UpdateSiteByName(context.Background(), "Main Campus", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Main Campus Updated", result.Name)
}

func TestUnitUpdateSiteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateSiteByName(context.Background(), "", &RequestSite{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "site name is required")
}

func TestUnitUpdateSiteByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateSiteByName(context.Background(), "Main Campus", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteSiteByID
// =============================================================================

func TestUnitDeleteSiteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteSiteByIDMock()

	resp, err := svc.DeleteSiteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteSiteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteSiteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "site ID must be a positive integer")
}

// =============================================================================
// DeleteSiteByName
// =============================================================================

func TestUnitDeleteSiteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteSiteByNameMock()

	resp, err := svc.DeleteSiteByName(context.Background(), "Main Campus")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteSiteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteSiteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "site name is required")
}
