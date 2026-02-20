package computer_groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ComputerGroupsMock) {
	t.Helper()
	mock := mocks.NewComputerGroupsMock()
	return NewService(mock), mock
}

// --- Smart Groups ---

func TestUnitListSmartGroups_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListSmartGroupsMock()

	result, resp, err := svc.ListSmartV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "MacBooks", result.Results[0].Name)
	assert.True(t, result.Results[0].IsSmart)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Engineering Macs", result.Results[1].Name)
}

func TestUnitGetSmartGroupByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSmartGroupMock()

	result, resp, err := svc.GetSmartByIDV2(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "MacBooks", result.Name)
	assert.True(t, result.IsSmart)
	require.Len(t, result.Criteria, 1)
	assert.Equal(t, "Model", result.Criteria[0].Name)
	assert.Equal(t, "MacBook Pro", result.Criteria[0].Value)
}

func TestUnitGetSmartGroupByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetSmartByIDV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart group ID is required")
}

func TestUnitGetSmartGroupByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetSmartByIDV2(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnitCreateSmartGroup_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateSmartGroupMock()

	req := &RequestSmartGroup{
		Name: "Test Smart Group",
		Criteria: []Criterion{
			{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "MacBook Pro"},
		},
	}
	result, resp, err := svc.CreateSmartV2(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v2/computer-groups/smart-groups/3")
}

func TestUnitCreateSmartGroup_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateSmartV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateSmartGroup_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestSmartGroup{Name: "Duplicate"}
	result, resp, err := svc.CreateSmartV2(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

func TestUnitUpdateSmartGroup_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSmartGroupMock()

	req := &RequestSmartGroup{
		Name: "MacBooks Updated",
		Criteria: []Criterion{
			{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "MacBook Air"},
		},
	}
	result, resp, err := svc.UpdateSmartV2(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "MacBooks Updated", result.Name)
}

func TestUnitUpdateSmartGroup_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateSmartV2(context.Background(), "", &RequestSmartGroup{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnitUpdateSmartGroup_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateSmartV2(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeleteSmartGroup_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteSmartGroupMock()

	resp, err := svc.DeleteSmartV2(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteSmartGroup_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteSmartV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart group ID is required")
}

// --- Static Groups ---

func TestUnitListStaticGroups_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListStaticGroupsMock()

	result, resp, err := svc.ListStaticV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "10", result.Results[0].ID)
	assert.Equal(t, "Test Machines", result.Results[0].Name)
	assert.False(t, result.Results[0].IsSmart)
	assert.Equal(t, []string{"101", "102"}, result.Results[0].ComputerIds)
}

func TestUnitGetStaticGroupByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetStaticGroupMock()

	result, resp, err := svc.GetStaticByIDV2(context.Background(), "10")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "10", result.ID)
	assert.Equal(t, "Test Machines", result.Name)
	assert.False(t, result.IsSmart)
	assert.Equal(t, []string{"101", "102"}, result.ComputerIds)
}

func TestUnitGetStaticGroupByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetStaticByIDV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static group ID is required")
}

func TestUnitCreateStaticGroup_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateStaticGroupMock()

	req := &RequestStaticGroup{
		Name:        "New Static Group",
		ComputerIds: []string{"1", "2", "3"},
	}
	result, resp, err := svc.CreateStaticV2(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "12", result.ID)
	assert.Contains(t, result.Href, "/api/v2/computer-groups/static-groups/12")
}

func TestUnitCreateStaticGroup_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateStaticV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitUpdateStaticGroup_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateStaticGroupMock()

	req := &RequestStaticGroup{
		Name:        "Test Machines",
		ComputerIds: []string{"101", "102", "103"},
	}
	result, resp, err := svc.UpdateStaticByIDV2(context.Background(), "10", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "10", result.ID)
	assert.Equal(t, []string{"101", "102", "103"}, result.ComputerIds)
}

func TestUnitUpdateStaticGroup_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateStaticByIDV2(context.Background(), "", &RequestStaticGroup{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnitUpdateStaticGroup_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateStaticByIDV2(context.Background(), "10", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeleteStaticGroup_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteStaticGroupMock()

	resp, err := svc.DeleteStaticByIDV2(context.Background(), "10")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteStaticGroup_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteStaticByIDV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static group ID is required")
}
