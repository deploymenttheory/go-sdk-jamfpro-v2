package users_inventory

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/users_inventory/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*UsersInventory, *mocks.UsersInventoryMock) {
	t.Helper()
	mock := mocks.NewUsersInventoryMock()
	return NewUsersInventory(mock), mock
}

func TestUnit_UsersInventory_NewService(t *testing.T) {
	mock := mocks.NewUsersInventoryMock()
	svc := NewUsersInventory(mock)
	require.NotNil(t, svc)
}

func TestUnit_UsersInventory_ListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListUsersMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "jsmith", result.Results[0].Username)
	assert.Equal(t, "John Smith", result.Results[0].Realname)
	assert.Equal(t, "john.smith@example.com", result.Results[0].Email)
	assert.Equal(t, "IT Manager", result.Results[0].Position)
}

func TestUnit_UsersInventory_ListV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListUsersErrorMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "failed to list users")
}

func TestUnit_UsersInventory_GetByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetUserMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "jsmith", result.Username)
	assert.Equal(t, "John Smith", result.Realname)
	assert.Equal(t, "john.smith@example.com", result.Email)
}

func TestUnit_UsersInventory_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "user ID is required")
}

func TestUnit_UsersInventory_GetByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_UsersInventory_CreateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateUserMock()

	req := &RequestUserInventory{
		Username: "jnewuser",
		Realname: "New User",
		Email:    "new.user@example.com",
		Position: "Engineer",
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_UsersInventory_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_UsersInventory_CreateV1_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateUserConflictMock()

	req := &RequestUserInventory{Username: "jsmith"}
	result, resp, err := svc.CreateV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

func TestUnit_UsersInventory_UpdateByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateUserMock()

	req := &RequestUserInventory{
		Username: "jsmith",
		Realname: "John Smith Updated",
		Email:    "jsmith.updated@example.com",
	}
	resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_UsersInventory_UpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateByIDV1(context.Background(), "", &RequestUserInventory{Username: "jsmith"})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "user ID is required")
}

func TestUnit_UsersInventory_UpdateByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_UsersInventory_DeleteByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteUserMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_UsersInventory_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "user ID is required")
}
