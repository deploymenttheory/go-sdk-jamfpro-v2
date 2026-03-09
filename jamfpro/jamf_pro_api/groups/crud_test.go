package groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/groups/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUnit_Groups_ListV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.ListV1(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, 3, result.TotalCount)
	assert.Len(t, result.Results, 3)
}

func TestUnit_Groups_ListV1_WithPagination(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
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

func TestUnit_Groups_GetByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterGetByIDMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetByIDV1(context.Background(), "platform-1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "platform-1", result.GroupPlatformId)
	assert.Equal(t, "Test Computer Group", result.GroupName)
}

func TestUnit_Groups_GetByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)

	result, resp, err := svc.GetByIDV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group ID is required")
}

func TestUnit_Groups_GetComputerGroupByNameV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "Test Computer Group")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Computer Group", result.GroupName)
	assert.Equal(t, "COMPUTER", result.GroupType)
}

func TestUnit_Groups_GetComputerGroupByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)

	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group name is required")
}

func TestUnit_Groups_GetComputerGroupByNameV1_NotFound(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "Nonexistent Group")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "computer group with name")
}

func TestUnit_Groups_GetMobileGroupByNameV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetMobileGroupByNameV1(context.Background(), "Test Mobile Group")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Mobile Group", result.GroupName)
	assert.Equal(t, "MOBILE", result.GroupType)
}

func TestUnit_Groups_GetMobileGroupByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)

	result, resp, err := svc.GetMobileGroupByNameV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group name is required")
}

func TestUnit_Groups_GetMobileGroupByNameV1_NotFound(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetMobileGroupByNameV1(context.Background(), "Nonexistent Group")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mobile group with name")
}

func TestUnit_Groups_GetComputerGroupByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetComputerGroupByIDV1(context.Background(), "101")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "101", result.GroupJamfProId)
	assert.Equal(t, "COMPUTER", result.GroupType)
}

func TestUnit_Groups_GetComputerGroupByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)

	result, resp, err := svc.GetComputerGroupByIDV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group Jamf Pro ID is required")
}

func TestUnit_Groups_GetComputerGroupByIDV1_NotFound(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetComputerGroupByIDV1(context.Background(), "999")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "computer group with Jamf Pro ID")
}

func TestUnit_Groups_GetMobileGroupByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetMobileGroupByIDV1(context.Background(), "102")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "102", result.GroupJamfProId)
	assert.Equal(t, "MOBILE", result.GroupType)
}

func TestUnit_Groups_GetMobileGroupByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)

	result, resp, err := svc.GetMobileGroupByIDV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group Jamf Pro ID is required")
}

func TestUnit_Groups_GetMobileGroupByIDV1_NotFound(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetMobileGroupByIDV1(context.Background(), "999")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mobile group with Jamf Pro ID")
}

func TestUnit_Groups_UpdateByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewGroups(mock)
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

func TestUnit_Groups_UpdateByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)

	req := &RequestUpdateGroup{GroupName: "Test"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "group ID is required")
}

func TestUnit_Groups_UpdateByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request body is required")
}

func TestUnit_Groups_DeleteByIDV1(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterDeleteMock()

	svc := NewGroups(mock)
	resp, err := svc.DeleteByIDV1(context.Background(), "platform-1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestUnit_Groups_DeleteByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)

	resp, err := svc.DeleteByIDV1(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "group ID is required")
}

// Additional comprehensive tests

func TestUnit_Groups_ListV1_WithRSQLFilter(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	rsqlQuery := map[string]string{
		"filter": `groupType=="COMPUTER"`,
	}
	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	// Mock returns all groups, but in real scenario this would filter
}

func TestUnit_Groups_ListV1_WithSorting(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	rsqlQuery := map[string]string{
		"sort": "groupName:asc,groupType:desc",
	}
	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUnit_Groups_ListV1_WithComplexRSQL(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
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

func TestUnit_Groups_GetByIDV1_SmartGroupWithCriteria(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterGetByIDMock()

	svc := NewGroups(mock)
	result, resp, err := svc.GetByIDV1(context.Background(), "platform-1")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "platform-1", result.GroupPlatformId)
	// Verify group structure
	assert.NotEmpty(t, result.GroupName)
	assert.NotEmpty(t, result.GroupType)
}

func TestUnit_Groups_UpdateByIDV1_SmartGroupCriteria(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewGroups(mock)
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

func TestUnit_Groups_UpdateByIDV1_StaticGroupAssignments(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewGroups(mock)
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

func TestUnit_Groups_UpdateByIDV1_OnlyName(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewGroups(mock)
	req := &RequestUpdateGroup{
		GroupName: "Only Name Updated",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUnit_Groups_UpdateByIDV1_OnlyDescription(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterUpdateMock()

	svc := NewGroups(mock)
	req := &RequestUpdateGroup{
		GroupDescription: "Only description updated",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
}

func TestUnit_Groups_GetComputerGroupByNameV1_MultipleGroups(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	// Should find "Another Computer Group" from the mock data
	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "Another Computer Group")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "Another Computer Group", result.GroupName)
	assert.Equal(t, "COMPUTER", result.GroupType)
}

func TestUnit_Groups_GetMobileGroupByNameV1_WrongType(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	// Try to get a computer group as mobile - should not find it
	result, resp, err := svc.GetMobileGroupByNameV1(context.Background(), "Test Computer Group")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mobile group with name")
}

func TestUnit_Groups_GetComputerGroupByNameV1_WrongType(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	// Try to get a mobile group as computer - should not find it
	result, resp, err := svc.GetComputerGroupByNameV1(context.Background(), "Test Mobile Group")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "computer group with name")
}

func TestUnit_Groups_GetComputerGroupByIDV1_WrongType(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	// Try to get a mobile group ID as computer - should not find it
	result, resp, err := svc.GetComputerGroupByIDV1(context.Background(), "102")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "computer group with Jamf Pro ID")
}

func TestUnit_Groups_GetMobileGroupByIDV1_WrongType(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListMock()

	svc := NewGroups(mock)
	// Try to get a computer group ID as mobile - should not find it
	result, resp, err := svc.GetMobileGroupByIDV1(context.Background(), "101")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "mobile group with Jamf Pro ID")
}

func TestUnit_Groups_ListV1_EmptyResults(t *testing.T) {
	mock := mocks.NewGroupsMock()
	// Register a mock that returns empty results
	mock.RegisterEmptyListMock()

	svc := NewGroups(mock)
	result, resp, err := svc.ListV1(context.Background(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, 0, result.TotalCount)
	assert.Len(t, result.Results, 0)
}

// Error case: no mock registered for GetPaginated
func TestUnit_Groups_ListV1_ClientError(t *testing.T) {
	mock := mocks.NewGroupsMock()
	// Do not register any mock - GetPaginated will return error
	svc := NewGroups(mock)
	result, resp, err := svc.ListV1(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to list groups")
}

// Error case: invalid JSON in list response triggers mergePage unmarshal error
func TestUnit_Groups_ListV1_InvalidJSON(t *testing.T) {
	mock := mocks.NewGroupsMock()
	mock.RegisterListInvalidJSONMock()
	svc := NewGroups(mock)
	result, _, err := svc.ListV1(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to list groups")
}

// Error case: no mock registered for Get
func TestUnit_Groups_GetByIDV1_ClientError(t *testing.T) {
	mock := mocks.NewGroupsMock()
	// Do not register GetByID - Get will return error
	svc := NewGroups(mock)
	result, resp, err := svc.GetByIDV1(context.Background(), "platform-1")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no mock registered")
}

// Error case: ListV1 fails when GetComputerGroupByNameV1 calls it
func TestUnit_Groups_GetComputerGroupByNameV1_ListError(t *testing.T) {
	mock := mocks.NewGroupsMock()
	// Do not register ListMock - ListV1 will fail
	svc := NewGroups(mock)
	result, _, err := svc.GetComputerGroupByNameV1(context.Background(), "Test Computer Group")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to list groups")
}

// Error case: ListV1 fails when GetMobileGroupByNameV1 calls it
func TestUnit_Groups_GetMobileGroupByNameV1_ListError(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)
	result, _, err := svc.GetMobileGroupByNameV1(context.Background(), "Test Mobile Group")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to list groups")
}

// Error case: ListV1 fails when GetComputerGroupByIDV1 calls it
func TestUnit_Groups_GetComputerGroupByIDV1_ListError(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)
	result, _, err := svc.GetComputerGroupByIDV1(context.Background(), "101")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to list groups")
}

// Error case: ListV1 fails when GetMobileGroupByIDV1 calls it
func TestUnit_Groups_GetMobileGroupByIDV1_ListError(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)
	result, _, err := svc.GetMobileGroupByIDV1(context.Background(), "102")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to list groups")
}

// Error case: no mock registered for Patch
func TestUnit_Groups_UpdateByIDV1_ClientError(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)
	req := &RequestUpdateGroup{GroupName: "Test"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "platform-1", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no mock registered")
}

// Error case: no mock registered for Delete
func TestUnit_Groups_DeleteByIDV1_ClientError(t *testing.T) {
	mock := mocks.NewGroupsMock()
	svc := NewGroups(mock)
	resp, err := svc.DeleteByIDV1(context.Background(), "platform-1")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}
