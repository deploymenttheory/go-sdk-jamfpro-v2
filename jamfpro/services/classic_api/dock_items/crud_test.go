package dock_items

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/dock_items/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_DockItems_List(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterListDockItemsMock()
	svc := NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.DockItems, 2)
	assert.Equal(t, "Safari", resp.DockItems[0].Name)
	assert.Equal(t, 1, resp.DockItems[0].ID)
	assert.Equal(t, "Finder", resp.DockItems[1].Name)
	assert.Equal(t, 2, resp.DockItems[1].ID)
}

func TestUnit_DockItems_GetByID(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterGetDockItemByIDMock()
	svc := NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Safari", resp.Name)
	assert.Equal(t, "App", resp.Type)
	assert.Equal(t, "/Applications/Safari.app", resp.Path)
}

func TestUnit_DockItems_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	svc := NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "dock item ID must be a positive integer")
}

func TestUnit_DockItems_GetByName(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterGetDockItemByNameMock()
	svc := NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Safari")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Safari", resp.Name)
	assert.Equal(t, "App", resp.Type)
}

func TestUnit_DockItems_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	svc := NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "dock item name cannot be empty")
}

func TestUnit_DockItems_Create(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterCreateDockItemMock()
	svc := NewService(mockClient)

	req := &Request{
		Name:     "Test Dock Item",
		Type:     "App",
		Path:     "/Applications/Test.app",
		Contents: "",
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
	assert.Equal(t, "Test Dock Item", resp.Name)
}

func TestUnit_DockItems_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	svc := NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_DockItems_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	svc := NewService(mockClient)

	req := &Request{
		Name: "",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "dock item name is required")
}

func TestUnit_DockItems_UpdateByID(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterUpdateDockItemByIDMock()
	svc := NewService(mockClient)

	req := &Request{
		Name:     "Updated Safari",
		Type:     "App",
		Path:     "/Applications/Safari.app",
		Contents: "",
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Updated Safari", resp.Name)
}

func TestUnit_DockItems_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	svc := NewService(mockClient)

	req := &Request{
		Name: "Test",
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "dock item ID must be a positive integer")
}

func TestUnit_DockItems_UpdateByName(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterUpdateDockItemByNameMock()
	svc := NewService(mockClient)

	req := &Request{
		Name:     "Updated Safari",
		Type:     "App",
		Path:     "/Applications/Safari.app",
		Contents: "",
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Safari", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Updated Safari", resp.Name)
}

func TestUnit_DockItems_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	svc := NewService(mockClient)

	req := &Request{
		Name: "Test",
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "dock item name cannot be empty")
}

func TestUnit_DockItems_DeleteByID(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterDeleteDockItemByIDMock()
	svc := NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_DockItems_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	svc := NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "dock item ID must be a positive integer")
}

func TestUnit_DockItems_DeleteByName(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterDeleteDockItemByNameMock()
	svc := NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Safari")

	require.NoError(t, err)
}

func TestUnit_DockItems_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	svc := NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "dock item name cannot be empty")
}

func TestUnit_DockItems_NotFound(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_DockItems_Conflict(t *testing.T) {
	mockClient := mocks.NewDockItemsMock()
	mockClient.RegisterConflictErrorMock()
	svc := NewService(mockClient)

	req := &Request{
		Name: "Duplicate Dock Item",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "dock item with that name already exists")
}
