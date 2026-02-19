package ibeacons

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ibeacons/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh IBeaconsMock.
func setupMockService(t *testing.T) (*Service, *mocks.IBeaconsMock) {
	t.Helper()
	mock := mocks.NewIBeaconsMock()
	return NewService(mock), mock
}

// =============================================================================
// ListIBeacons
// =============================================================================

func TestUnitListIBeacons_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListIBeaconsMock()

	result, resp, err := svc.ListIBeacons(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Lobby Beacon", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Cafeteria Beacon", result.Results[1].Name)
}

// =============================================================================
// GetIBeaconByID
// =============================================================================

func TestUnitGetIBeaconByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetIBeaconByIDMock()

	result, resp, err := svc.GetIBeaconByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Lobby Beacon", result.Name)
	assert.Equal(t, "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0", result.UUID)
	assert.Equal(t, 1, result.Major)
	assert.Equal(t, 1, result.Minor)
}

func TestUnitGetIBeaconByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetIBeaconByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "iBeacon ID must be a positive integer")
}

func TestUnitGetIBeaconByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetIBeaconByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "iBeacon ID must be a positive integer")
}

func TestUnitGetIBeaconByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetIBeaconByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetIBeaconByName
// =============================================================================

func TestUnitGetIBeaconByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetIBeaconByNameMock()

	result, resp, err := svc.GetIBeaconByName(context.Background(), "Lobby Beacon")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Lobby Beacon", result.Name)
}

func TestUnitGetIBeaconByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetIBeaconByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "iBeacon name is required")
}

// =============================================================================
// CreateIBeacon
// =============================================================================

func TestUnitCreateIBeacon_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateIBeaconMock()

	req := &RequestIBeacon{
		Name:  "New Beacon",
		UUID:  "F3D67EC6-EGGA-59E3-C171-E1G6B82107F1",
		Major: 1,
		Minor: 3,
	}
	result, resp, err := svc.CreateIBeacon(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 3, result.ID)
	assert.Equal(t, "New Beacon", result.Name)
}

func TestUnitCreateIBeacon_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateIBeacon(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateIBeacon_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestIBeacon{Name: "Lobby Beacon", UUID: "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0"}
	result, resp, err := svc.CreateIBeacon(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateIBeaconByID
// =============================================================================

func TestUnitUpdateIBeaconByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateIBeaconByIDMock()

	req := &RequestIBeacon{
		Name:  "Lobby Beacon Updated",
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 2,
	}
	result, resp, err := svc.UpdateIBeaconByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Lobby Beacon Updated", result.Name)
}

func TestUnitUpdateIBeaconByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateIBeaconByID(context.Background(), 0, &RequestIBeacon{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "iBeacon ID must be a positive integer")
}

func TestUnitUpdateIBeaconByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateIBeaconByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateIBeaconByName
// =============================================================================

func TestUnitUpdateIBeaconByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateIBeaconByNameMock()

	req := &RequestIBeacon{
		Name:  "Lobby Beacon Updated",
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 2,
	}
	result, resp, err := svc.UpdateIBeaconByName(context.Background(), "Lobby Beacon", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Lobby Beacon Updated", result.Name)
}

func TestUnitUpdateIBeaconByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateIBeaconByName(context.Background(), "", &RequestIBeacon{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "iBeacon name is required")
}

func TestUnitUpdateIBeaconByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateIBeaconByName(context.Background(), "Lobby Beacon", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteIBeaconByID
// =============================================================================

func TestUnitDeleteIBeaconByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteIBeaconByIDMock()

	resp, err := svc.DeleteIBeaconByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteIBeaconByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteIBeaconByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "iBeacon ID must be a positive integer")
}

// =============================================================================
// DeleteIBeaconByName
// =============================================================================

func TestUnitDeleteIBeaconByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteIBeaconByNameMock()

	resp, err := svc.DeleteIBeaconByName(context.Background(), "Lobby Beacon")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteIBeaconByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteIBeaconByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "iBeacon name is required")
}
