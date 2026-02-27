package static_user_groups_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/static_user_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/static_user_groups/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_StaticUserGroups_List(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterListStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "Static Test Group", resp.Results[0].Name)
	assert.False(t, resp.Results[0].IsSmart)
	assert.Equal(t, "Another Static Group", resp.Results[1].Name)
	assert.False(t, resp.Results[1].IsSmart)
}

func TestUnit_StaticUserGroups_GetByID(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterGetStaticUserGroupByIDMock()
	svc := static_user_groups.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Static Test Group", resp.Name)
	assert.False(t, resp.IsSmart)
	assert.NotNil(t, resp.Site)
	assert.Equal(t, -1, resp.Site.ID)
	assert.Len(t, resp.Users, 1)
	assert.Equal(t, "testuser", resp.Users[0].Username)
}

func TestUnit_StaticUserGroups_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group ID must be a positive integer")
}

func TestUnit_StaticUserGroups_GetByName(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterGetStaticUserGroupByNameMock()
	svc := static_user_groups.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Static Test Group")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Static Test Group", resp.Name)
	assert.False(t, resp.IsSmart)
}

func TestUnit_StaticUserGroups_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_StaticUserGroups_Create(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterCreateStaticUserGroupMock()
	svc := static_user_groups.NewService(mockClient)

	req := &static_user_groups.RequestStaticUserGroup{
		Name:             "Test Group",
		IsSmart:          false,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_StaticUserGroups_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_StaticUserGroups_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)

	req := &static_user_groups.RequestStaticUserGroup{
		Name: "",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_StaticUserGroups_UpdateByID(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterUpdateStaticUserGroupByIDMock()
	svc := static_user_groups.NewService(mockClient)

	req := &static_user_groups.RequestStaticUserGroup{
		Name:    "Updated Group",
		IsSmart: false,
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_StaticUserGroups_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)

	req := &static_user_groups.RequestStaticUserGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group ID must be a positive integer")
}

func TestUnit_StaticUserGroups_UpdateByName(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterUpdateStaticUserGroupByNameMock()
	svc := static_user_groups.NewService(mockClient)

	req := &static_user_groups.RequestStaticUserGroup{
		Name:    "Updated Group",
		IsSmart: false,
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Static Test Group", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_StaticUserGroups_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)

	req := &static_user_groups.RequestStaticUserGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_StaticUserGroups_DeleteByID(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterDeleteStaticUserGroupByIDMock()
	svc := static_user_groups.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_StaticUserGroups_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group ID must be a positive integer")
}

func TestUnit_StaticUserGroups_DeleteByName(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterDeleteStaticUserGroupByNameMock()
	svc := static_user_groups.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Static Test Group")

	require.NoError(t, err)
}

func TestUnit_StaticUserGroups_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_StaticUserGroups_NotFound(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := static_user_groups.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_StaticUserGroups_Conflict(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	mockClient.RegisterConflictErrorMock()
	svc := static_user_groups.NewService(mockClient)

	req := &static_user_groups.RequestStaticUserGroup{
		Name:    "Duplicate Group",
		IsSmart: false,
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group with that name already exists")
}

func TestUnit_StaticUserGroups_List_Error(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_StaticUserGroups_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, _, err := svc.GetByName(context.Background(), "Static Test Group")
	require.Error(t, err)
}

func TestUnit_StaticUserGroups_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_StaticUserGroups_UpdateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &static_user_groups.RequestStaticUserGroup{Name: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_StaticUserGroups_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &static_user_groups.RequestStaticUserGroup{Name: "Test"})
	require.Error(t, err)
}

func TestUnit_StaticUserGroups_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Static Test Group", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_StaticUserGroups_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Static Test Group", &static_user_groups.RequestStaticUserGroup{Name: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_StaticUserGroups_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Static Test Group", &static_user_groups.RequestStaticUserGroup{Name: "Updated"})
	require.Error(t, err)
}

func TestUnit_StaticUserGroups_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, err := svc.DeleteByID(context.Background(), 1)
	require.Error(t, err)
}

func TestUnit_StaticUserGroups_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewStaticUserGroupsMock()
	svc := static_user_groups.NewService(mockClient)
	_, err := svc.DeleteByName(context.Background(), "Static Test Group")
	require.Error(t, err)
}
