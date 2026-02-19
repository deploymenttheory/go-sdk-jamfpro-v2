package dock_items

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dock_items/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.DockItemsMock) {
	t.Helper()
	mock := mocks.NewDockItemsMock()
	return NewService(mock), mock
}

func TestUnitGetDockItemByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDockItemMock()

	result, resp, err := svc.GetDockItemByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Safari", result.Name)
	assert.Equal(t, "/Applications/Safari.app", result.Path)
	assert.Equal(t, "App", result.Type)
}

func TestUnitGetDockItemByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDockItemByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "dock item ID is required")
}

func TestUnitGetDockItemByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetDockItemByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnitCreateDockItem_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateDockItemMock()

	req := &RequestDockItem{
		Name: "Safari",
		Path: "/Applications/Safari.app",
		Type: TypeApp,
	}
	result, resp, err := svc.CreateDockItemV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/dock-items/3")
}

func TestUnitCreateDockItem_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateDockItemV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateDockItem_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestDockItem{
		Name: "Duplicate",
		Path: "/Applications/App.app",
		Type: TypeApp,
	}
	result, resp, err := svc.CreateDockItemV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

func TestUnitUpdateDockItemByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDockItemMock()

	req := &RequestDockItem{
		Name: "Safari Updated",
		Path: "/Applications/Safari.app",
		Type: TypeApp,
	}
	result, resp, err := svc.UpdateDockItemByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Safari Updated", result.Name)
}

func TestUnitUpdateDockItemByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateDockItemByIDV1(context.Background(), "", &RequestDockItem{
		Name: "x",
		Path: "/path",
		Type: TypeApp,
	})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnitUpdateDockItemByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateDockItemByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeleteDockItemByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDockItemMock()

	resp, err := svc.DeleteDockItemByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteDockItemByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteDockItemByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "dock item ID is required")
}
