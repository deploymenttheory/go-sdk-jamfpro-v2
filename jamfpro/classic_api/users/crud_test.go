package users_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/users"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/users/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Users_List(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterListUsersMock()
	svc := users.NewUsers(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "admin", resp.Results[0].Name)
	assert.Equal(t, 1, resp.Results[0].ID)
	assert.Equal(t, "testuser", resp.Results[1].Name)
	assert.Equal(t, 2, resp.Results[1].ID)
}

func TestUnit_Users_GetByID(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterGetUserByIDMock()
	svc := users.NewUsers(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "admin", resp.Name)
	assert.Equal(t, "Administrator", resp.FullName)
	assert.Equal(t, "admin@example.com", resp.Email)
	assert.NotNil(t, resp.LDAPServer)
	assert.Equal(t, -1, resp.LDAPServer.ID)
	assert.Len(t, resp.Sites, 1)
	assert.Equal(t, -1, resp.Sites[0].ID)
	assert.Equal(t, "None", resp.Sites[0].Name)
}

func TestUnit_Users_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user ID must be a positive integer")
}

func TestUnit_Users_GetByName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterGetUserByNameMock()
	svc := users.NewUsers(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "admin")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "admin", resp.Name)
}

func TestUnit_Users_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user name cannot be empty")
}

func TestUnit_Users_GetByEmail(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterGetUserByEmailMock()
	svc := users.NewUsers(mockClient)

	resp, _, err := svc.GetByEmail(context.Background(), "admin@example.com")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.Results[0].ID)
	assert.Equal(t, "admin", resp.Results[0].Name)
}

func TestUnit_Users_GetByEmail_EmptyEmail(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	_, _, err := svc.GetByEmail(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user email cannot be empty")
}

func TestUnit_Users_Create(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterCreateUserMock()
	svc := users.NewUsers(mockClient)

	req := &users.RequestUser{
		Name:     "newuser",
		FullName: "New User",
		Email:    "newuser@example.com",
		Sites: []shared.SharedResourceSite{
			{ID: -1, Name: "None"},
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
	assert.Equal(t, "newuser", resp.Name)
}

func TestUnit_Users_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Users_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	req := &users.RequestUser{
		Name: "",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user name is required")
}

func TestUnit_Users_UpdateByID(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterUpdateUserByIDMock()
	svc := users.NewUsers(mockClient)

	req := &users.RequestUser{
		Name:     "admin",
		FullName: "Administrator Updated",
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Administrator Updated", resp.FullName)
}

func TestUnit_Users_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	req := &users.RequestUser{
		Name: "admin",
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user ID must be a positive integer")
}

func TestUnit_Users_UpdateByName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterUpdateUserByNameMock()
	svc := users.NewUsers(mockClient)

	req := &users.RequestUser{
		Name:     "admin",
		FullName: "Administrator Updated",
	}

	resp, _, err := svc.UpdateByName(context.Background(), "admin", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_Users_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	req := &users.RequestUser{
		Name: "admin",
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user name cannot be empty")
}

func TestUnit_Users_UpdateByEmail(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterUpdateUserByEmailMock()
	svc := users.NewUsers(mockClient)

	req := &users.RequestUser{
		Name:     "admin",
		FullName: "Administrator Updated",
	}

	resp, _, err := svc.UpdateByEmail(context.Background(), "admin@example.com", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_Users_UpdateByEmail_EmptyEmail(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	req := &users.RequestUser{
		Name: "admin",
	}

	_, _, err := svc.UpdateByEmail(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user email cannot be empty")
}

func TestUnit_Users_DeleteByID(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterDeleteUserByIDMock()
	svc := users.NewUsers(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_Users_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user ID must be a positive integer")
}

func TestUnit_Users_DeleteByName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterDeleteUserByNameMock()
	svc := users.NewUsers(mockClient)

	_, err := svc.DeleteByName(context.Background(), "admin")

	require.NoError(t, err)
}

func TestUnit_Users_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user name cannot be empty")
}

func TestUnit_Users_DeleteByEmail(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterDeleteUserByEmailMock()
	svc := users.NewUsers(mockClient)

	_, err := svc.DeleteByEmail(context.Background(), "admin@example.com")

	require.NoError(t, err)
}

func TestUnit_Users_DeleteByEmail_EmptyEmail(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)

	_, err := svc.DeleteByEmail(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user email cannot be empty")
}

func TestUnit_Users_NotFound(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := users.NewUsers(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_Users_Conflict(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	mockClient.RegisterConflictErrorMock()
	svc := users.NewUsers(mockClient)

	req := &users.RequestUser{
		Name: "admin",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user with that name already exists")
}

func TestUnit_Users_List_Error(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_Users_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.GetByName(context.Background(), "admin")
	require.Error(t, err)
}

func TestUnit_Users_GetByEmail_Error(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.GetByEmail(context.Background(), "admin@example.com")
	require.Error(t, err)
}

func TestUnit_Users_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Users_UpdateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &users.RequestUser{Name: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user name is required")
}

func TestUnit_Users_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &users.RequestUser{Name: "admin"})
	require.Error(t, err)
}

func TestUnit_Users_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "admin", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Users_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "admin", &users.RequestUser{Name: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user name is required in request")
}

func TestUnit_Users_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "admin", &users.RequestUser{Name: "updated"})
	require.Error(t, err)
}

func TestUnit_Users_UpdateByEmail_NilRequest(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.UpdateByEmail(context.Background(), "admin@example.com", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Users_UpdateByEmail_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.UpdateByEmail(context.Background(), "admin@example.com", &users.RequestUser{Name: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user name is required in request")
}

func TestUnit_Users_UpdateByEmail_Error(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, _, err := svc.UpdateByEmail(context.Background(), "admin@example.com", &users.RequestUser{Name: "admin"})
	require.Error(t, err)
}

func TestUnit_Users_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, err := svc.DeleteByID(context.Background(), 1)
	require.Error(t, err)
}

func TestUnit_Users_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, err := svc.DeleteByName(context.Background(), "admin")
	require.Error(t, err)
}

func TestUnit_Users_DeleteByEmail_Error(t *testing.T) {
	mockClient := mocks.NewUsersMock()
	svc := users.NewUsers(mockClient)
	_, err := svc.DeleteByEmail(context.Background(), "admin@example.com")
	require.Error(t, err)
}
