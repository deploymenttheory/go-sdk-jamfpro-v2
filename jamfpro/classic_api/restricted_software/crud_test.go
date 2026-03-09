package restricted_software

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/restricted_software/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh RestrictedSoftwareMock.
func setupMockService(t *testing.T) (*RestrictedSoftware, *mocks.RestrictedSoftwareMock) {
	t.Helper()
	mock := mocks.NewRestrictedSoftwareMock()
	return NewRestrictedSoftware(mock), mock
}

// =============================================================================
// ListRestrictedSoftware
// =============================================================================

func TestUnit_RestrictedSoftware_List_Success(t *testing.T) {
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
	assert.Equal(t, "Calculator", result.Results[0].Name)
}

// =============================================================================
// GetRestrictedSoftwareByID
// =============================================================================

func TestUnit_RestrictedSoftware_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.General.ID)
	assert.Equal(t, "Calculator", result.General.Name)
	assert.Equal(t, "Calculator.app", result.General.ProcessName)
	assert.True(t, result.General.MatchExactProcessName)
}

func TestUnit_RestrictedSoftware_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software ID must be a positive integer")
}

func TestUnit_RestrictedSoftware_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// =============================================================================
// GetRestrictedSoftwareByName
// =============================================================================

func TestUnit_RestrictedSoftware_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "Calculator")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Calculator", result.General.Name)
}

func TestUnit_RestrictedSoftware_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software name is required")
}

// =============================================================================
// CreateRestrictedSoftware
// =============================================================================

func TestUnit_RestrictedSoftware_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

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
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 100, result.ID)
}

func TestUnit_RestrictedSoftware_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateRestrictedSoftwareByID
// =============================================================================

func TestUnit_RestrictedSoftware_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDMock()

	req := &RequestRestrictedSoftware{
		General: RequestGeneral{
			Name:        "Updated Calculator",
			ProcessName: "Calculator.app",
		},
		Scope: Scope{AllComputers: true},
	}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
}

func TestUnit_RestrictedSoftware_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestRestrictedSoftware{}
	result, resp, err := svc.UpdateByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software ID must be a positive integer")
}

// =============================================================================
// UpdateRestrictedSoftwareByName
// =============================================================================

func TestUnit_RestrictedSoftware_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByNameMock()

	req := &RequestRestrictedSoftware{
		General: RequestGeneral{Name: "Updated"},
	}
	result, resp, err := svc.UpdateByName(context.Background(), "Calculator", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_RestrictedSoftware_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestRestrictedSoftware{}
	result, resp, err := svc.UpdateByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software name is required")
}

// =============================================================================
// DeleteRestrictedSoftwareByID
// =============================================================================

func TestUnit_RestrictedSoftware_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_RestrictedSoftware_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software ID must be a positive integer")
}

// =============================================================================
// DeleteRestrictedSoftwareByName
// =============================================================================

func TestUnit_RestrictedSoftware_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "Calculator")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_RestrictedSoftware_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "restricted software name is required")
}

func TestUnit_RestrictedSoftware_List_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_RestrictedSoftware_GetByName_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByName(context.Background(), "Calculator")
	require.Error(t, err)
}

func TestUnit_RestrictedSoftware_Create_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.Create(context.Background(), &RequestRestrictedSoftware{})
	require.Error(t, err)
}

func TestUnit_RestrictedSoftware_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_RestrictedSoftware_UpdateByID_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 1, &RequestRestrictedSoftware{})
	require.Error(t, err)
}

func TestUnit_RestrictedSoftware_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "Calculator", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_RestrictedSoftware_UpdateByName_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "Calculator", &RequestRestrictedSoftware{})
	require.Error(t, err)
}

func TestUnit_RestrictedSoftware_DeleteByID_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByID(context.Background(), 1)
	require.Error(t, err)
}

func TestUnit_RestrictedSoftware_DeleteByName_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByName(context.Background(), "Calculator")
	require.Error(t, err)
}
