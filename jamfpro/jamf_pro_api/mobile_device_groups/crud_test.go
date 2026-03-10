package mobile_device_groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*MobileDeviceGroups, *mocks.MobileDeviceGroupsMock) {
	t.Helper()
	mock := mocks.NewMobileDeviceGroupsMock()
	return NewMobileDeviceGroups(mock), mock
}

func TestUnit_MobileDeviceGroups_NewService(t *testing.T) {
	mock := mocks.NewMobileDeviceGroupsMock()
	svc := NewMobileDeviceGroups(mock)
	require.NotNil(t, svc)
}

func TestUnit_MobileDeviceGroups_ListSmartV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListSmartMock()

	result, resp, err := svc.ListSmartV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "iPhones", result.Results[0].Name)
}

func TestUnit_MobileDeviceGroups_GetSmartByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSmartMock()

	result, resp, err := svc.GetSmartByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "iPhones", result.Name)
	require.Len(t, result.Criteria, 1)
	assert.Equal(t, "Model", result.Criteria[0].Name)
}

func TestUnit_MobileDeviceGroups_CreateSmartV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateSmartMock()

	req := &RequestSmartMobileDeviceGroup{
		Name:        "New Smart",
		Description: "Desc",
		Criteria:    []CriteriaJamfProAPI{{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "iPhone"}},
	}
	result, resp, err := svc.CreateSmartV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "2", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_MobileDeviceGroups_UpdateSmartByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSmartMock()

	req := &RequestSmartMobileDeviceGroup{
		Name:        "iPhones Updated",
		Description: "Updated",
		Criteria:    []CriteriaJamfProAPI{{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "iPhone"}},
	}
	result, resp, err := svc.UpdateSmartByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "iPhones Updated", result.Name)
}

func TestUnit_MobileDeviceGroups_DeleteSmartByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteSmartMock()

	resp, err := svc.DeleteSmartByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_ListStaticV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListStaticMock()

	result, resp, err := svc.ListStaticV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "10", result.Results[0].ID)
	assert.Equal(t, "Static Devices", result.Results[0].Name)
}

func TestUnit_MobileDeviceGroups_ListStaticV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListStaticErrorMock()

	result, resp, err := svc.ListStaticV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_GetStaticByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetStaticMock()

	result, resp, err := svc.GetStaticByIDV1(context.Background(), "10")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "10", result.ID)
	assert.Equal(t, "Static Devices", result.Name)
}

func TestUnit_MobileDeviceGroups_CreateStaticV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateStaticMock()

	req := &RequestStaticMobileDeviceGroup{
		Name:        "New Static",
		Description: "Desc",
		SiteId:      "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{},
	}
	result, resp, err := svc.CreateStaticV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "11", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_MobileDeviceGroups_UpdateStaticByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateStaticMock()

	req := &RequestStaticMobileDeviceGroup{
		Name:        "Static Devices Updated",
		Description: "Updated",
		SiteId:      "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{},
	}
	result, resp, err := svc.UpdateStaticByIDV1(context.Background(), "10", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "10", result.ID)
	assert.Equal(t, "Static Devices Updated", result.Name)
}

func TestUnit_MobileDeviceGroups_DeleteStaticByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteStaticMock()

	resp, err := svc.DeleteStaticByIDV1(context.Background(), "10")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_ListSmartV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListSmartErrorMock()
	result, resp, err := svc.ListSmartV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_ListSmartV1_NoMock(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListSmartNoResponseErrorMock()
	result, resp, err := svc.ListSmartV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_MobileDeviceGroups_CreateSmartV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateSmartV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceGroups_GetSmartByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()
	result, resp, err := svc.GetSmartByIDV1(context.Background(), "999")
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_CreateStaticV1_NilAssignments(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateStaticMock()
	req := &RequestStaticMobileDeviceGroup{
		Name:   "New Static",
		SiteId: "-1",
	}
	result, resp, err := svc.CreateStaticV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_UpdateStaticByIDV1_NilAssignments(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateStaticMock()
	req := &RequestStaticMobileDeviceGroup{
		Name:   "Static Updated",
		SiteId: "-1",
	}
	result, resp, err := svc.UpdateStaticByIDV1(context.Background(), "10", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_GetSmartByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSmartByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart mobile device group ID is required")
}

func TestUnit_MobileDeviceGroups_UpdateSmartByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestSmartMobileDeviceGroup{Name: "Test"}
	result, resp, err := svc.UpdateSmartByIDV1(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDeviceGroups_UpdateSmartByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateSmartByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceGroups_DeleteSmartByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteSmartByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart mobile device group ID is required")
}

func TestUnit_MobileDeviceGroups_GetStaticByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetStaticNotFoundErrorMock()
	result, resp, err := svc.GetStaticByIDV1(context.Background(), "999")
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_GetStaticByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetStaticByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static mobile device group ID is required")
}

func TestUnit_MobileDeviceGroups_CreateStaticV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateStaticV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceGroups_UpdateStaticByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestStaticMobileDeviceGroup{Name: "Test"}
	result, resp, err := svc.UpdateStaticByIDV1(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDeviceGroups_UpdateStaticByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateStaticByIDV1(context.Background(), "10", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceGroups_DeleteStaticByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteStaticByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static mobile device group ID is required")
}

// -----------------------------------------------------------------------------
// ListAllV1, Membership, Erase
// -----------------------------------------------------------------------------

func TestUnit_MobileDeviceGroups_ListAllV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAllMock()

	result, resp, err := svc.ListAllV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "iPhones", result[0].Name)
	assert.True(t, result[0].IsSmartGroup)
	assert.Equal(t, 10, result[1].ID)
	assert.Equal(t, "Static Devices", result[1].Name)
	assert.False(t, result[1].IsSmartGroup)
}

func TestUnit_MobileDeviceGroups_ListAllV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAllErrorMock()

	result, resp, err := svc.ListAllV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_GetStaticGroupMembershipV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetStaticGroupMembershipMock()

	result, resp, err := svc.GetStaticGroupMembershipV1(context.Background(), "10", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "101", result.Results[0].MobileDeviceID)
	assert.Equal(t, "iPhone 1", result.Results[0].DisplayName)
	assert.Equal(t, "102", result.Results[1].MobileDeviceID)
}

func TestUnit_MobileDeviceGroups_GetStaticGroupMembershipV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetStaticGroupMembershipV1(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static group ID is required")
}

func TestUnit_MobileDeviceGroups_GetStaticGroupMembershipV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetStaticGroupMembershipErrorMock()

	result, resp, err := svc.GetStaticGroupMembershipV1(context.Background(), "10", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_GetSmartGroupMembershipV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSmartGroupMembershipMock()

	result, resp, err := svc.GetSmartGroupMembershipV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "201", result.Results[0].MobileDeviceID)
	assert.Equal(t, "iPad Pro", result.Results[0].DisplayName)
}

func TestUnit_MobileDeviceGroups_GetSmartGroupMembershipV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSmartGroupMembershipV1(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart group ID is required")
}

func TestUnit_MobileDeviceGroups_GetSmartGroupMembershipV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSmartGroupMembershipErrorMock()

	result, resp, err := svc.GetSmartGroupMembershipV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_EraseDevicesByGroupIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterEraseDevicesMock()

	req := &RequestEraseDevices{
		PreserveDataPlan:       boolPtr(false),
		DisallowProximitySetup: boolPtr(true),
	}
	resp, err := svc.EraseDevicesByGroupIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_EraseDevicesByGroupIDV1_NilRequest(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterEraseDevicesMock()

	resp, err := svc.EraseDevicesByGroupIDV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_EraseDevicesByGroupIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.EraseDevicesByGroupIDV1(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "group ID is required")
}

func TestUnit_MobileDeviceGroups_EraseDevicesByGroupIDV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterEraseDevicesErrorMock()

	resp, err := svc.EraseDevicesByGroupIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func boolPtr(b bool) *bool { return &b }
