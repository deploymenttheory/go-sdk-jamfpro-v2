package byoprofiles

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/byoprofiles/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh BYOProfilesMock.
func setupMockService(t *testing.T) (*Service, *mocks.BYOProfilesMock) {
	t.Helper()
	mock := mocks.NewBYOProfilesMock()
	return NewService(mock), mock
}

// =============================================================================
// ListBYOProfiles
// =============================================================================

func TestUnit_Byoprofiles_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListBYOProfilesMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Test BYO Profile 1", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Test BYO Profile 2", result.Results[1].Name)
}

// =============================================================================
// GetBYOProfileByID
// =============================================================================

func TestUnit_Byoprofiles_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetBYOProfileByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test BYO Profile 1", result.General.Name)
	assert.Equal(t, "Test BYO profile description", result.General.Description)
	assert.True(t, result.General.Enabled)
}

func TestUnit_Byoprofiles_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

func TestUnit_Byoprofiles_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

func TestUnit_Byoprofiles_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// =============================================================================
// GetBYOProfileByName
// =============================================================================

func TestUnit_Byoprofiles_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetBYOProfileByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "Test BYO Profile 1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test BYO Profile 1", result.General.Name)
}

func TestUnit_Byoprofiles_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile name is required")
}

// =============================================================================
// CreateBYOProfile
// =============================================================================

func TestUnit_Byoprofiles_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateBYOProfileMock()

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name:    "New Test BYO Profile",
			Enabled: true,
		},
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 100, result.ID)
}

func TestUnit_Byoprofiles_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Byoprofiles_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Duplicate BYO Profile",
		},
	}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

// =============================================================================
// UpdateBYOProfileByID
// =============================================================================

func TestUnit_Byoprofiles_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateBYOProfileByIDMock()

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Updated BYO Profile Name",
		},
	}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
}

func TestUnit_Byoprofiles_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Updated BYO Profile Name",
		},
	}
	result, resp, err := svc.UpdateByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

func TestUnit_Byoprofiles_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateBYOProfileByName
// =============================================================================

func TestUnit_Byoprofiles_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateBYOProfileByNameMock()

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Updated BYO Profile Name",
		},
	}
	result, resp, err := svc.UpdateByName(context.Background(), "Test BYO Profile 1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
}

func TestUnit_Byoprofiles_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Updated BYO Profile Name",
		},
	}
	result, resp, err := svc.UpdateByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile name is required")
}

func TestUnit_Byoprofiles_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "Test BYO Profile 1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteBYOProfileByID
// =============================================================================

func TestUnit_Byoprofiles_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteBYOProfileByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Byoprofiles_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

func TestUnit_Byoprofiles_DeleteByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

// =============================================================================
// DeleteBYOProfileByName
// =============================================================================

func TestUnit_Byoprofiles_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteBYOProfileByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "Test BYO Profile 1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Byoprofiles_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile name is required")
}
