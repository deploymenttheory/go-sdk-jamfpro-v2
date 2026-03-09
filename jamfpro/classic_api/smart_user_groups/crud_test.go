package smart_user_groups_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/smart_user_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/smart_user_groups/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_SmartUserGroups_List(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterListSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "All Users", resp.Results[0].Name)
	assert.True(t, resp.Results[0].IsSmart)
	assert.Equal(t, "Smart Test Group", resp.Results[1].Name)
	assert.True(t, resp.Results[1].IsSmart)
}

func TestUnit_SmartUserGroups_GetByID(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterGetSmartUserGroupByIDMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

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

func TestUnit_SmartUserGroups_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group ID must be a positive integer")
}

func TestUnit_SmartUserGroups_GetByName(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterGetSmartUserGroupByNameMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "All Users")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "All Users", resp.Name)
	assert.True(t, resp.IsSmart)
}

func TestUnit_SmartUserGroups_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_SmartUserGroups_Create(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterCreateSmartUserGroupMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	req := &smart_user_groups.RequestSmartUserGroup{
		Name:             "Test Group",
		IsSmart:          true,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &smart_user_groups.CriteriaContainer{
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

func TestUnit_SmartUserGroups_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SmartUserGroups_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	req := &smart_user_groups.RequestSmartUserGroup{
		Name: "",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_SmartUserGroups_UpdateByID(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterUpdateSmartUserGroupByIDMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	req := &smart_user_groups.RequestSmartUserGroup{
		Name:    "Updated Group",
		IsSmart: true,
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_SmartUserGroups_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	req := &smart_user_groups.RequestSmartUserGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group ID must be a positive integer")
}

func TestUnit_SmartUserGroups_UpdateByName(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterUpdateSmartUserGroupByNameMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	req := &smart_user_groups.RequestSmartUserGroup{
		Name:    "Updated Group",
		IsSmart: true,
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "All Users", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_SmartUserGroups_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	req := &smart_user_groups.RequestSmartUserGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_SmartUserGroups_DeleteByID(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterDeleteSmartUserGroupByIDMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_SmartUserGroups_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group ID must be a positive integer")
}

func TestUnit_SmartUserGroups_DeleteByName(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterDeleteSmartUserGroupByNameMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	_, err := svc.DeleteByName(context.Background(), "All Users")

	require.NoError(t, err)
}

func TestUnit_SmartUserGroups_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_SmartUserGroups_NotFound(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_SmartUserGroups_Conflict(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	mockClient.RegisterConflictErrorMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)

	req := &smart_user_groups.RequestSmartUserGroup{
		Name:    "Duplicate Group",
		IsSmart: true,
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group with that name already exists")
}

func TestUnit_SmartUserGroups_List_Error(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_SmartUserGroups_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, _, err := svc.GetByName(context.Background(), "All Users")
	require.Error(t, err)
}

func TestUnit_SmartUserGroups_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SmartUserGroups_UpdateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &smart_user_groups.RequestSmartUserGroup{Name: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_SmartUserGroups_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, &smart_user_groups.RequestSmartUserGroup{Name: "Test"})
	require.Error(t, err)
}

func TestUnit_SmartUserGroups_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "All Users", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SmartUserGroups_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "All Users", &smart_user_groups.RequestSmartUserGroup{Name: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestUnit_SmartUserGroups_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "All Users", &smart_user_groups.RequestSmartUserGroup{Name: "Updated"})
	require.Error(t, err)
}

func TestUnit_SmartUserGroups_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, err := svc.DeleteByID(context.Background(), 1)
	require.Error(t, err)
}

func TestUnit_SmartUserGroups_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewSmartUserGroupsMock()
	svc := smart_user_groups.NewSmartUserGroups(mockClient)
	_, err := svc.DeleteByName(context.Background(), "All Users")
	require.Error(t, err)
}
