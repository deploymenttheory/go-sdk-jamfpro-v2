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

func TestUnitListBYOProfiles_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListBYOProfilesMock()

	result, resp, err := svc.ListBYOProfiles(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
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

func TestUnitGetBYOProfileByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetBYOProfileByIDMock()

	result, resp, err := svc.GetBYOProfileByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test BYO Profile 1", result.General.Name)
	assert.Equal(t, "Test BYO profile description", result.General.Description)
	assert.True(t, result.General.Enabled)
}

func TestUnitGetBYOProfileByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetBYOProfileByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

func TestUnitGetBYOProfileByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetBYOProfileByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

func TestUnitGetBYOProfileByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetBYOProfileByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetBYOProfileByName
// =============================================================================

func TestUnitGetBYOProfileByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetBYOProfileByNameMock()

	result, resp, err := svc.GetBYOProfileByName(context.Background(), "Test BYO Profile 1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test BYO Profile 1", result.General.Name)
}

func TestUnitGetBYOProfileByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetBYOProfileByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile name is required")
}

// =============================================================================
// CreateBYOProfile
// =============================================================================

func TestUnitCreateBYOProfile_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateBYOProfileMock()

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name:    "New Test BYO Profile",
			Enabled: true,
		},
	}
	result, resp, err := svc.CreateBYOProfile(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 100, result.ID)
}

func TestUnitCreateBYOProfile_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateBYOProfile(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateBYOProfile_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Duplicate BYO Profile",
		},
	}
	result, resp, err := svc.CreateBYOProfile(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdateBYOProfileByID
// =============================================================================

func TestUnitUpdateBYOProfileByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateBYOProfileByIDMock()

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Updated BYO Profile Name",
		},
	}
	result, resp, err := svc.UpdateBYOProfileByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateBYOProfileByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Updated BYO Profile Name",
		},
	}
	result, resp, err := svc.UpdateBYOProfileByID(context.Background(), 0, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

func TestUnitUpdateBYOProfileByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateBYOProfileByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateBYOProfileByName
// =============================================================================

func TestUnitUpdateBYOProfileByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateBYOProfileByNameMock()

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Updated BYO Profile Name",
		},
	}
	result, resp, err := svc.UpdateBYOProfileByName(context.Background(), "Test BYO Profile 1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateBYOProfileByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestBYOProfile{
		General: GeneralSettings{
			Name: "Updated BYO Profile Name",
		},
	}
	result, resp, err := svc.UpdateBYOProfileByName(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile name is required")
}

func TestUnitUpdateBYOProfileByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateBYOProfileByName(context.Background(), "Test BYO Profile 1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteBYOProfileByID
// =============================================================================

func TestUnitDeleteBYOProfileByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteBYOProfileByIDMock()

	resp, err := svc.DeleteBYOProfileByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteBYOProfileByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteBYOProfileByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

func TestUnitDeleteBYOProfileByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteBYOProfileByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile ID must be a positive integer")
}

// =============================================================================
// DeleteBYOProfileByName
// =============================================================================

func TestUnitDeleteBYOProfileByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteBYOProfileByNameMock()

	resp, err := svc.DeleteBYOProfileByName(context.Background(), "Test BYO Profile 1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteBYOProfileByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteBYOProfileByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "BYO profile name is required")
}
