package computer_groups_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computer_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/computer_groups/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_ComputerGroups_List(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterListComputerGroupsMock()
	svc := computer_groups.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "All Managed Clients", resp.Results[0].Name)
	assert.True(t, resp.Results[0].IsSmart)
	assert.Equal(t, "Static Test Group", resp.Results[1].Name)
	assert.False(t, resp.Results[1].IsSmart)
}

func TestUnit_ComputerGroups_GetByID(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterGetComputerGroupByIDMock()
	svc := computer_groups.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "All Managed Clients", resp.Name)
	assert.True(t, resp.IsSmart)
	assert.NotNil(t, resp.Site)
	assert.Equal(t, -1, resp.Site.ID)
	assert.NotNil(t, resp.Criteria)
	assert.Equal(t, 2, resp.Criteria.Size)
	assert.Len(t, resp.Criteria.Criterion, 2)
	assert.Len(t, resp.Computers, 1)
	assert.Equal(t, "test-computer-01", resp.Computers[0].Name)
}

func TestUnit_ComputerGroups_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	svc := computer_groups.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer group ID must be a positive integer")
}

func TestUnit_ComputerGroups_GetByName(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterGetComputerGroupByNameMock()
	svc := computer_groups.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "All Managed Clients")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "All Managed Clients", resp.Name)
	assert.True(t, resp.IsSmart)
}

func TestUnit_ComputerGroups_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	svc := computer_groups.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer group name cannot be empty")
}

func TestUnit_ComputerGroups_Create(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterCreateComputerGroupMock()
	svc := computer_groups.NewService(mockClient)

	req := &computer_groups.RequestComputerGroup{
		Name:    "Test Group",
		IsSmart: true,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &computer_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Operating System",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "macOS",
				},
			},
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_ComputerGroups_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	svc := computer_groups.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ComputerGroups_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	svc := computer_groups.NewService(mockClient)

	req := &computer_groups.RequestComputerGroup{
		Name: "",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer group name is required")
}

func TestUnit_ComputerGroups_UpdateByID(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterUpdateComputerGroupByIDMock()
	svc := computer_groups.NewService(mockClient)

	req := &computer_groups.RequestComputerGroup{
		Name:    "Updated Group",
		IsSmart: false,
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_ComputerGroups_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	svc := computer_groups.NewService(mockClient)

	req := &computer_groups.RequestComputerGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer group ID must be a positive integer")
}

func TestUnit_ComputerGroups_UpdateByName(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterUpdateComputerGroupByNameMock()
	svc := computer_groups.NewService(mockClient)

	req := &computer_groups.RequestComputerGroup{
		Name:    "Updated Group",
		IsSmart: false,
	}

	resp, _, err := svc.UpdateByName(context.Background(), "All Managed Clients", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_ComputerGroups_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	svc := computer_groups.NewService(mockClient)

	req := &computer_groups.RequestComputerGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer group name cannot be empty")
}

func TestUnit_ComputerGroups_DeleteByID(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterDeleteComputerGroupByIDMock()
	svc := computer_groups.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_ComputerGroups_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	svc := computer_groups.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer group ID must be a positive integer")
}

func TestUnit_ComputerGroups_DeleteByName(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterDeleteComputerGroupByNameMock()
	svc := computer_groups.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "All Managed Clients")

	require.NoError(t, err)
}

func TestUnit_ComputerGroups_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	svc := computer_groups.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "computer group name cannot be empty")
}

func TestUnit_ComputerGroups_NotFound(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := computer_groups.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Computer group not found")
}

func TestUnit_ComputerGroups_Conflict(t *testing.T) {
	mockClient := mocks.NewComputerGroupsMock()
	mockClient.RegisterConflictErrorMock()
	svc := computer_groups.NewService(mockClient)

	req := &computer_groups.RequestComputerGroup{
		Name:    "Duplicate Group",
		IsSmart: false,
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Computer group name already exists")
}
