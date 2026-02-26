package software_update_servers

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/software_update_servers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SoftwareUpdateServersMock) {
	t.Helper()
	mock := mocks.NewSoftwareUpdateServersMock()
	return NewService(mock), mock
}

// =============================================================================
// ListSoftwareUpdateServers
// =============================================================================

func TestUnit_SoftwareUpdateServers_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListSoftwareUpdateServersMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Primary SUS", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Secondary SUS", result.Results[1].Name)
}

// =============================================================================
// GetSoftwareUpdateServerByID
// =============================================================================

func TestUnit_SoftwareUpdateServers_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSoftwareUpdateServerByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary SUS", result.Name)
	assert.Equal(t, "192.168.1.50", result.IPAddress)
	assert.Equal(t, 8088, result.Port)
	assert.True(t, result.SetSystemWide)
}

func TestUnit_SoftwareUpdateServers_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "software update server ID must be a positive integer")
}

func TestUnit_SoftwareUpdateServers_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), -1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "software update server ID must be a positive integer")
}

func TestUnit_SoftwareUpdateServers_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, _, err := svc.GetByID(context.Background(), 999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

// =============================================================================
// GetSoftwareUpdateServerByName
// =============================================================================

func TestUnit_SoftwareUpdateServers_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSoftwareUpdateServerByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "Primary SUS")
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary SUS", result.Name)
}

func TestUnit_SoftwareUpdateServers_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "software update server name is required")
}

// =============================================================================
// CreateSoftwareUpdateServer
// =============================================================================

func TestUnit_SoftwareUpdateServers_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateSoftwareUpdateServerMock()

	req := &RequestSoftwareUpdateServer{
		Name:          "Primary SUS",
		IPAddress:     "192.168.1.50",
		Port:          8088,
		SetSystemWide: true,
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary SUS", result.Name)
}

func TestUnit_SoftwareUpdateServers_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.Create(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SoftwareUpdateServers_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestSoftwareUpdateServer{Name: "Primary SUS"}
	_, _, err := svc.Create(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdateSoftwareUpdateServerByID
// =============================================================================

func TestUnit_SoftwareUpdateServers_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSoftwareUpdateServerByIDMock()

	req := &RequestSoftwareUpdateServer{Name: "Primary SUS Updated"}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary SUS Updated", result.Name)
}

func TestUnit_SoftwareUpdateServers_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 0, &RequestSoftwareUpdateServer{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "software update server ID must be a positive integer")
}

func TestUnit_SoftwareUpdateServers_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateSoftwareUpdateServerByName
// =============================================================================

func TestUnit_SoftwareUpdateServers_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSoftwareUpdateServerByNameMock()

	req := &RequestSoftwareUpdateServer{Name: "Primary SUS Updated"}
	result, resp, err := svc.UpdateByName(context.Background(), "Primary SUS", req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_SoftwareUpdateServers_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "", &RequestSoftwareUpdateServer{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "software update server name is required")
}

func TestUnit_SoftwareUpdateServers_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "Primary SUS", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteSoftwareUpdateServerByID
// =============================================================================

func TestUnit_SoftwareUpdateServers_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteSoftwareUpdateServerByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_SoftwareUpdateServers_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "software update server ID must be a positive integer")
}

// =============================================================================
// DeleteSoftwareUpdateServerByName
// =============================================================================

func TestUnit_SoftwareUpdateServers_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteSoftwareUpdateServerByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "Primary SUS")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_SoftwareUpdateServers_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "software update server name is required")
}
