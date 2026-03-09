package smart_computer_groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/smart_computer_groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*SmartComputerGroups, *mocks.SmartComputerGroupsMock) {
	t.Helper()
	mock := mocks.NewSmartComputerGroupsMock()
	return NewSmartComputerGroups(mock), mock
}

func TestUnit_SmartComputerGroups_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "All Macs", result.Results[0].Name)
	assert.Equal(t, 42, result.Results[0].MembershipCount)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Test Smart Group", result.Results[1].Name)
}

func TestUnit_SmartComputerGroups_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "All Macs", result.Name)
	assert.Equal(t, "All managed Mac computers", result.Description)
	assert.Equal(t, 42, result.MembershipCount)
	require.Len(t, result.Criteria, 1)
	assert.Equal(t, "Computer Name", result.Criteria[0].Name)
	assert.Equal(t, "is", result.Criteria[0].SearchType)
}

func TestUnit_SmartComputerGroups_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart computer group ID is required")
}

func TestUnit_SmartComputerGroups_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_SmartComputerGroups_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByName(context.Background(), "All Macs")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "All Macs", result.Name)
	assert.Equal(t, 42, result.MembershipCount)
}

func TestUnit_SmartComputerGroups_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart computer group name is required")
}

func TestUnit_SmartComputerGroups_GetByName_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListEmptyMock()

	result, resp, err := svc.GetByName(context.Background(), "NonExistent")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "not found")
}

func TestUnit_SmartComputerGroups_GetMembership_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMembershipMock()

	result, resp, err := svc.GetMembership(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Len(t, result.Members, 5)
	assert.Equal(t, []int{101, 102, 103, 104, 105}, result.Members)
}

func TestUnit_SmartComputerGroups_GetMembership_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetMembership(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart computer group ID is required")
}

func TestUnit_SmartComputerGroups_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestSmartGroup{
		Name:        "New Smart Group",
		Description: "Created via API",
		Criteria: []SubsetCriteria{
			{Name: "Computer Name", Priority: 1, AndOr: "and", SearchType: "is", Value: "*"},
		},
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
	assert.Equal(t, "/api/v2/computer-groups/smart-groups/3", result.Href)
}

func TestUnit_SmartComputerGroups_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SmartComputerGroups_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestSmartGroup{Name: "Duplicate"}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

func TestUnit_SmartComputerGroups_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestSmartGroup{
		Name:        "All Macs Updated",
		Description: "Updated description",
		Criteria: []SubsetCriteria{
			{Name: "Computer Name", Priority: 1, AndOr: "and", SearchType: "is", Value: "*"},
		},
	}
	result, resp, err := svc.UpdateByID(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "All Macs Updated", result.Name)
	assert.Equal(t, "Updated description", result.Description)
}

func TestUnit_SmartComputerGroups_UpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), "", &RequestSmartGroup{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_SmartComputerGroups_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SmartComputerGroups_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_SmartComputerGroups_DeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart computer group ID is required")
}
