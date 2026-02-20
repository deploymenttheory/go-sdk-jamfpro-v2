package restricted_software

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/restricted_software/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh RestrictedSoftwareMock.
func setupMockService(t *testing.T) (*Service, *mocks.RestrictedSoftwareMock) {
	t.Helper()
	mock := mocks.NewRestrictedSoftwareMock()
	return NewService(mock), mock
}

// =============================================================================
// ListRestrictedSoftware
// =============================================================================

func TestUnitListRestrictedSoftware_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListRestrictedSoftwareMock()

	result, resp, err := svc.ListRestrictedSoftware(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Calculator", result.Results[0].Name)
}

// =============================================================================
// GetRestrictedSoftwareByID
// =============================================================================

func TestUnitGetRestrictedSoftwareByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetRestrictedSoftwareByIDMock()

	result, resp, err := svc.GetRestrictedSoftwareByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.General.ID)
	assert.Equal(t, "Calculator", result.General.Name)
	assert.Equal(t, "Calculator.app", result.General.ProcessName)
	assert.True(t, result.General.MatchExactProcessName)
}

func TestUnitGetRestrictedSoftwareByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetRestrictedSoftwareByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software ID must be a positive integer")
}

func TestUnitGetRestrictedSoftwareByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetRestrictedSoftwareByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetRestrictedSoftwareByName
// =============================================================================

func TestUnitGetRestrictedSoftwareByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetRestrictedSoftwareByNameMock()

	result, resp, err := svc.GetRestrictedSoftwareByName(context.Background(), "Calculator")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Calculator", result.General.Name)
}

func TestUnitGetRestrictedSoftwareByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetRestrictedSoftwareByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software name is required")
}

// =============================================================================
// CreateRestrictedSoftware
// =============================================================================

func TestUnitCreateRestrictedSoftware_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateRestrictedSoftwareMock()

	req := &RequestRestrictedSoftware{
		General: RequestGeneral{
			Name:                  "Calculator",
			ProcessName:           "Calculator.app",
			MatchExactProcessName: true,
			SendNotification:      true,
			KillProcess:           true,
			DeleteExecutable:      false,
			DisplayMessage:        "Calculator is restricted",
			Site:                  &shared.SharedResourceSite{ID: -1, Name: "None"},
		},
		Scope: Scope{
			AllComputers: false,
		},
	}
	result, resp, err := svc.CreateRestrictedSoftware(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
}

func TestUnitCreateRestrictedSoftware_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateRestrictedSoftware(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateRestrictedSoftwareByID
// =============================================================================

func TestUnitUpdateRestrictedSoftwareByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateRestrictedSoftwareByIDMock()

	req := &RequestRestrictedSoftware{
		General: RequestGeneral{
			Name:        "Updated Calculator",
			ProcessName: "Calculator.app",
		},
		Scope: Scope{AllComputers: true},
	}
	result, resp, err := svc.UpdateRestrictedSoftwareByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateRestrictedSoftwareByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestRestrictedSoftware{}
	result, resp, err := svc.UpdateRestrictedSoftwareByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software ID must be a positive integer")
}

// =============================================================================
// UpdateRestrictedSoftwareByName
// =============================================================================

func TestUnitUpdateRestrictedSoftwareByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateRestrictedSoftwareByNameMock()

	req := &RequestRestrictedSoftware{
		General: RequestGeneral{Name: "Updated"},
	}
	result, resp, err := svc.UpdateRestrictedSoftwareByName(context.Background(), "Calculator", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdateRestrictedSoftwareByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestRestrictedSoftware{}
	result, resp, err := svc.UpdateRestrictedSoftwareByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software name is required")
}

// =============================================================================
// DeleteRestrictedSoftwareByID
// =============================================================================

func TestUnitDeleteRestrictedSoftwareByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteRestrictedSoftwareByIDMock()

	resp, err := svc.DeleteRestrictedSoftwareByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteRestrictedSoftwareByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteRestrictedSoftwareByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software ID must be a positive integer")
}

// =============================================================================
// DeleteRestrictedSoftwareByName
// =============================================================================

func TestUnitDeleteRestrictedSoftwareByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteRestrictedSoftwareByNameMock()

	resp, err := svc.DeleteRestrictedSoftwareByName(context.Background(), "Calculator")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteRestrictedSoftwareByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteRestrictedSoftwareByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software name is required")
}
