package static_computer_groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/static_computer_groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.StaticComputerGroupsMock) {
	t.Helper()
	mock := mocks.NewStaticComputerGroupsMock()
	return NewService(mock), mock
}

func TestUnit_StaticComputerGroups_ListV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Engineering Workstations", result.Results[0].Name)
	assert.Equal(t, "Computers used by engineering team", result.Results[0].Description)
	assert.Equal(t, 5, result.Results[0].Count)
}

func TestUnit_StaticComputerGroups_ListV2_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered - dispatch returns error

	result, resp, err := svc.ListV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to list static computer groups")
}

func TestUnit_StaticComputerGroups_GetByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByIDV2(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Engineering Workstations", result.Name)
	assert.Equal(t, "Computers used by engineering team", result.Description)
	assert.Equal(t, 5, result.Count)
}

func TestUnit_StaticComputerGroups_GetByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static computer group ID is required")
}

func TestUnit_StaticComputerGroups_GetByIDV2_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV2(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_StaticComputerGroups_GetByNameV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByNameV2(context.Background(), "Engineering Workstations")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Engineering Workstations", result.Name)
	assert.Equal(t, 5, result.Count)
}

func TestUnit_StaticComputerGroups_GetByNameV2_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByNameV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static computer group name is required")
}

func TestUnit_StaticComputerGroups_GetByNameV2_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListEmptyMock()

	result, resp, err := svc.GetByNameV2(context.Background(), "NonExistent")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "was not found")
}

func TestUnit_StaticComputerGroups_GetByNameV2_ListError(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered - ListV2 returns error

	result, resp, err := svc.GetByNameV2(context.Background(), "SomeName")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_StaticComputerGroups_CreateV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestStaticGroup{
		Name:        "New Static Group",
		Description: "Created via API",
		Assignments: []string{"1", "2"},
	}
	result, resp, err := svc.CreateV2(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v2/computer-groups/static-groups/3")
}

func TestUnit_StaticComputerGroups_CreateV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_StaticComputerGroups_CreateV2_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestStaticGroup{
		Name:        "Duplicate",
		Assignments: []string{},
	}
	result, resp, err := svc.CreateV2(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

func TestUnit_StaticComputerGroups_UpdateByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestStaticGroup{
		Name:        "Engineering Workstations Updated",
		Description: "Updated description",
		Assignments: []string{"1", "2", "3"},
	}
	siteID := "1"
	req.SiteID = &siteID

	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Engineering Workstations Updated", result.Name)
	assert.Equal(t, "Updated description", result.Description)
	require.Len(t, result.Assignments, 3)
}

func TestUnit_StaticComputerGroups_UpdateByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestStaticGroup{Name: "x", Assignments: []string{}}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_StaticComputerGroups_UpdateByIDV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_StaticComputerGroups_UpdateByIDV2_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateNotFoundErrorMock()

	req := &RequestStaticGroup{Name: "x", Assignments: []string{}}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "999", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_StaticComputerGroups_UpdateByIDV2_NilAssignments(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestStaticGroup{Name: "Updated", Assignments: nil}
	result, resp, err := svc.UpdateByIDV2(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_StaticComputerGroups_DeleteByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByIDV2(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_StaticComputerGroups_DeleteByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static computer group ID is required")
}

func TestUnit_StaticComputerGroups_DeleteByIDV2_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteNotFoundErrorMock()

	resp, err := svc.DeleteByIDV2(context.Background(), "999")
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_StaticComputerGroups_CreateV2_NilAssignments(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestStaticGroup{Name: "New Group", Assignments: nil}
	result, resp, err := svc.CreateV2(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
}
