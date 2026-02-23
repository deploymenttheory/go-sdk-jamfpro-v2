package usergroups_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/usergroups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/usergroups/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_UserGroups_List(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterListUserGroupsMock()
	svc := usergroups.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "All Users", resp.Results[0].Name)
	assert.True(t, resp.Results[0].IsSmart)
	assert.Equal(t, "Static Test Group", resp.Results[1].Name)
	assert.False(t, resp.Results[1].IsSmart)
}

func TestUnit_UserGroups_GetByID(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterGetUserGroupByIDMock()
	svc := usergroups.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "All Users", resp.Name)
	assert.True(t, resp.IsSmart)
	assert.NotNil(t, resp.Site)
	assert.Equal(t, -1, resp.Site.ID)
	assert.NotNil(t, resp.Criteria)
	assert.Equal(t, 2, resp.Criteria.Size)
	assert.Len(t, resp.Criteria.Criterion, 2)
	assert.Len(t, resp.Users, 1)
	assert.Equal(t, "admin", resp.Users[0].Username)
}

func TestUnit_UserGroups_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	svc := usergroups.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group ID must be a positive integer")
}

func TestUnit_UserGroups_GetByName(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterGetUserGroupByNameMock()
	svc := usergroups.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "All Users")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "All Users", resp.Name)
	assert.True(t, resp.IsSmart)
}

func TestUnit_UserGroups_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	svc := usergroups.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name cannot be empty")
}

func TestUnit_UserGroups_Create(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterCreateUserGroupMock()
	svc := usergroups.NewService(mockClient)

	req := &usergroups.RequestUserGroup{
		Name:             "Test Group",
		IsSmart:          true,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &usergroups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_UserGroups_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	svc := usergroups.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_UserGroups_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	svc := usergroups.NewService(mockClient)

	req := &usergroups.RequestUserGroup{
		Name: "",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_UserGroups_UpdateByID(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterUpdateUserGroupByIDMock()
	svc := usergroups.NewService(mockClient)

	req := &usergroups.RequestUserGroup{
		Name:    "Updated Group",
		IsSmart: false,
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_UserGroups_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	svc := usergroups.NewService(mockClient)

	req := &usergroups.RequestUserGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group ID must be a positive integer")
}

func TestUnit_UserGroups_UpdateByName(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterUpdateUserGroupByNameMock()
	svc := usergroups.NewService(mockClient)

	req := &usergroups.RequestUserGroup{
		Name:    "Updated Group",
		IsSmart: false,
	}

	resp, _, err := svc.UpdateByName(context.Background(), "All Users", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_UserGroups_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	svc := usergroups.NewService(mockClient)

	req := &usergroups.RequestUserGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name cannot be empty")
}

func TestUnit_UserGroups_DeleteByID(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterDeleteUserGroupByIDMock()
	svc := usergroups.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_UserGroups_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	svc := usergroups.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group ID must be a positive integer")
}

func TestUnit_UserGroups_DeleteByName(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterDeleteUserGroupByNameMock()
	svc := usergroups.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "All Users")

	require.NoError(t, err)
}

func TestUnit_UserGroups_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	svc := usergroups.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name cannot be empty")
}

func TestUnit_UserGroups_NotFound(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := usergroups.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_UserGroups_Conflict(t *testing.T) {
	mockClient := mocks.NewUserGroupsMock()
	mockClient.RegisterConflictErrorMock()
	svc := usergroups.NewService(mockClient)

	req := &usergroups.RequestUserGroup{
		Name:    "Duplicate Group",
		IsSmart: false,
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group with that name already exists")
}
