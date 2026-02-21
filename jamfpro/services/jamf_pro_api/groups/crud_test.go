package groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/groups/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.ListV1(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, 3, result.TotalCount)
	assert.Len(t, result.Results, 3)
}

func TestListV1_WithPagination(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "10",
		"sort":      "groupName:asc",
	}
	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, 3, result.TotalCount)
}

func TestGetByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterGetByIDMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByIDV1(context.Background(), "platform-1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "platform-1", result.GroupPlatformId)
	assert.Equal(t, "Test Computer Group", result.GroupName)
}

func TestGetByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetByIDV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group ID is required")
}

func TestGetComputerGroupByNameV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "Test Computer Group")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Computer Group", result.GroupName)
	assert.Equal(t, "COMPUTER", result.GroupType)
}

func TestGetComputerGroupByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group name is required")
}

func TestGetComputerGroupByNameV1_NotFound(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "Nonexistent Group")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "computer group with name")
}

func TestGetMobileGroupByNameV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetMobileGroupByNameV1(context.Background(), "Test Mobile Group")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Mobile Group", result.GroupName)
	assert.Equal(t, "MOBILE", result.GroupType)
}

func TestGetMobileGroupByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetMobileGroupByNameV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group name is required")
}

func TestGetMobileGroupByNameV1_NotFound(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetMobileGroupByNameV1(context.Background(), "Nonexistent Group")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mobile group with name")
}

func TestGetComputerGroupByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetComputerGroupByIDV1(context.Background(), "101")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "101", result.GroupJamfProId)
	assert.Equal(t, "COMPUTER", result.GroupType)
}

func TestGetComputerGroupByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetComputerGroupByIDV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group Jamf Pro ID is required")
}

func TestGetComputerGroupByIDV1_NotFound(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetComputerGroupByIDV1(context.Background(), "999")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "computer group with Jamf Pro ID")
}

func TestGetMobileGroupByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetMobileGroupByIDV1(context.Background(), "102")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "102", result.GroupJamfProId)
	assert.Equal(t, "MOBILE", result.GroupType)
}

func TestGetMobileGroupByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewService(mock)

	result, resp, err := svc.GetMobileGroupByIDV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group Jamf Pro ID is required")
}

func TestGetMobileGroupByIDV1_NotFound(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	result, resp, err := svc.GetMobileGroupByIDV1(context.Background(), "999")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mobile group with Jamf Pro ID")
}

func TestUpdateByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewService(mock)
	req := &RequestUpdateGroup{
		GroupName:        "Updated Group Name",
		GroupDescription: "Updated description",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Computer Group", result.GroupName)
}

func TestUpdateByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewService(mock)

	req := &RequestUpdateGroup{GroupName: "Test"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group ID is required")
}

func TestUpdateByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewService(mock)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request body is required")
}

func TestDeleteByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterDeleteMock()

	svc := NewService(mock)
	resp, err := svc.DeleteByIDV1(context.Background(), "platform-1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestDeleteByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewService(mock)

	resp, err := svc.DeleteByIDV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "group ID is required")
}

// Additional comprehensive tests

func TestListV1_WithRSQLFilter(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	rsqlQuery := map[string]string{
		"filter": `groupType=="COMPUTER"`,
	}
	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	// Mock returns all groups, but in real scenario this would filter
}

func TestListV1_WithSorting(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	rsqlQuery := map[string]string{
		"sort": "groupName:asc,groupType:desc",
	}
	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestListV1_WithComplexRSQL(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	rsqlQuery := map[string]string{
		"filter":    `groupName=="Managed" and isSmart=="true"`,
		"page":      "0",
		"page-size": "50",
		"sort":      "groupName:desc",
	}
	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestGetByIDV1_SmartGroupWithCriteria(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterGetByIDMock()

	svc := NewService(mock)
	result, resp, err := svc.GetByIDV1(context.Background(), "platform-1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "platform-1", result.GroupPlatformId)
	// Verify group structure
	assert.NotEmpty(t, result.GroupName)
	assert.NotEmpty(t, result.GroupType)
}

func TestUpdateByIDV1_SmartGroupCriteria(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewService(mock)
	req := &RequestUpdateGroup{
		GroupName:        "Updated Smart Group",
		GroupDescription: "Updated smart group description",
		Criteria: []SubsetCriterion{
			{
				Name:         "Operating System Version",
				Priority:     0,
				AndOr:        "and",
				SearchType:   "greater than or equal",
				Value:        "15.0",
				OpeningParen: false,
				ClosingParen: false,
			},
			{
				Name:         "Device Name",
				Priority:     1,
				AndOr:        "and",
				SearchType:   "like",
				Value:        "iPhone",
				OpeningParen: false,
				ClosingParen: false,
			},
		},
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUpdateByIDV1_StaticGroupAssignments(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewService(mock)
	req := &RequestUpdateGroup{
		GroupName:        "Updated Static Group",
		GroupDescription: "Updated static group description",
		Assignments: []SubsetAssignment{
			{
				DeviceID: "device-1",
				Selected: true,
			},
			{
				DeviceID: "device-2",
				Selected: true,
			},
			{
				DeviceID: "device-3",
				Selected: false,
			},
		},
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUpdateByIDV1_OnlyName(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewService(mock)
	req := &RequestUpdateGroup{
		GroupName: "Only Name Updated",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUpdateByIDV1_OnlyDescription(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewService(mock)
	req := &RequestUpdateGroup{
		GroupDescription: "Only description updated",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestGetComputerGroupByNameV1_MultipleGroups(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	// Should find "Another Computer Group" from the mock data
	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "Another Computer Group")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "Another Computer Group", result.GroupName)
	assert.Equal(t, "COMPUTER", result.GroupType)
}

func TestGetMobileGroupByNameV1_WrongType(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	// Try to get a computer group as mobile - should not find it
	result, resp, err := svc.GetMobileGroupByNameV1(context.Background(), "Test Computer Group")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mobile group with name")
}

func TestGetComputerGroupByNameV1_WrongType(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	// Try to get a mobile group as computer - should not find it
	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "Test Mobile Group")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "computer group with name")
}

func TestGetComputerGroupByIDV1_WrongType(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	// Try to get a mobile group ID as computer - should not find it
	result, resp, err := svc.GetComputerGroupByIDV1(context.Background(), "102")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "computer group with Jamf Pro ID")
}

func TestGetMobileGroupByIDV1_WrongType(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	// Try to get a computer group ID as mobile - should not find it
	result, resp, err := svc.GetMobileGroupByIDV1(context.Background(), "101")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mobile group with Jamf Pro ID")
}

func TestListV1_EmptyResults(t *testing.T) {
	mock := mocks.NewGroupsMock()
	// Register a mock that returns empty results
	mock.RegisterEmptyListMock()

	svc := NewService(mock)
	result, resp, err := svc.ListV1(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, 0, result.TotalCount)
	assert.Len(t, result.Results, 0)
}
