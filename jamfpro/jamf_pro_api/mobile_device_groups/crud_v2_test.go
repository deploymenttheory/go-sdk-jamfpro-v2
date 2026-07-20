package mobile_device_groups

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// -----------------------------------------------------------------------------
// Smart Groups (V2)
// -----------------------------------------------------------------------------

func TestUnit_MobileDeviceGroups_ListSmartV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListSmartV2Mock()

	result, resp, err := svc.ListSmartV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "iPhones", result.Results[0].Name)
}

func TestUnit_MobileDeviceGroups_GetSmartByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSmartV2Mock()

	result, resp, err := svc.GetSmartByIDV2(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	require.Len(t, result.Criteria, 1)
	assert.Equal(t, "Model", result.Criteria[0].Name)
}

func TestUnit_MobileDeviceGroups_GetSmartByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSmartByIDV2(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart mobile device group ID is required")
}

func TestUnit_MobileDeviceGroups_GetSmartByIDV2_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSmartV2NotFoundErrorMock()
	result, resp, err := svc.GetSmartByIDV2(context.Background(), "999")
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_CreateSmartV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateSmartV2Mock()

	req := &RequestSmartMobileDeviceGroup{
		Name:     "New Smart",
		Criteria: []CriteriaJamfProAPI{{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "iPhone"}},
	}
	result, resp, err := svc.CreateSmartV2(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "2", result.ID)
}

func TestUnit_MobileDeviceGroups_CreateSmartV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateSmartV2(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceGroups_CreateSmartV2_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestSmartMobileDeviceGroup{Name: ""}
	result, resp, err := svc.CreateSmartV2(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "group name is required")
}

func TestUnit_MobileDeviceGroups_CreateSmartV2_InvalidAndOr(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestSmartMobileDeviceGroup{
		Name:     "Bad AndOr",
		Criteria: []CriteriaJamfProAPI{{Name: "Model", AndOr: "nand", SearchType: "is", Value: "iPhone"}},
	}
	result, resp, err := svc.CreateSmartV2(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "invalid andOr")
}

func TestUnit_MobileDeviceGroups_UpdateSmartByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSmartV2Mock()

	req := &RequestSmartMobileDeviceGroup{
		Name:     "iPhones Updated",
		Criteria: []CriteriaJamfProAPI{{Name: "Model", AndOr: "and", SearchType: "is", Value: "iPhone"}},
	}
	result, resp, err := svc.UpdateSmartByIDV2(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "iPhones Updated", result.Name)
}

func TestUnit_MobileDeviceGroups_UpdateSmartByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestSmartMobileDeviceGroup{Name: "Test"}
	result, resp, err := svc.UpdateSmartByIDV2(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDeviceGroups_UpdateSmartByIDV2_InvalidAndOr(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestSmartMobileDeviceGroup{
		Name:     "Bad",
		Criteria: []CriteriaJamfProAPI{{Name: "Model", AndOr: "xor", SearchType: "is", Value: "iPhone"}},
	}
	result, resp, err := svc.UpdateSmartByIDV2(context.Background(), "1", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "invalid andOr")
}

func TestUnit_MobileDeviceGroups_DeleteSmartByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteSmartV2Mock()

	resp, err := svc.DeleteSmartByIDV2(context.Background(), "1")
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_DeleteSmartByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteSmartByIDV2(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart mobile device group ID is required")
}

// -----------------------------------------------------------------------------
// Static Groups (V2)
// -----------------------------------------------------------------------------

func TestUnit_MobileDeviceGroups_ListStaticV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListStaticV2Mock()

	result, resp, err := svc.ListStaticV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result.Results, 1)
	assert.Equal(t, "10", result.Results[0].ID)
}

func TestUnit_MobileDeviceGroups_ListStaticV2_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListStaticV2ErrorMock()

	result, resp, err := svc.ListStaticV2(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_GetStaticByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetStaticV2Mock()

	result, resp, err := svc.GetStaticByIDV2(context.Background(), "10")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "10", result.ID)
	assert.Equal(t, "Static Devices", result.Name)
}

func TestUnit_MobileDeviceGroups_GetStaticByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetStaticByIDV2(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static mobile device group ID is required")
}

func TestUnit_MobileDeviceGroups_CreateStaticV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateStaticV2Mock()

	req := &RequestStaticMobileDeviceGroup{
		Name:        "New Static",
		SiteId:      "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{},
	}
	result, resp, err := svc.CreateStaticV2(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "11", result.ID)
}

func TestUnit_MobileDeviceGroups_CreateStaticV2_NilAssignmentsSerializeAsArray(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateStaticV2Mock()

	req := &RequestStaticMobileDeviceGroup{Name: "New Static", SiteId: "-1"}
	result, resp, err := svc.CreateStaticV2(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode())
	// The API 500s when "assignments" is missing; ensure it is a non-nil array.
	require.NotNil(t, req.Assignments)
	assert.Empty(t, req.Assignments)
}

func TestUnit_MobileDeviceGroups_CreateStaticV2_DedupesAssignments(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateStaticV2Mock()

	req := &RequestStaticMobileDeviceGroup{
		Name:   "Dup",
		SiteId: "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{
			{MobileDeviceID: "1", Selected: true},
			{MobileDeviceID: "1", Selected: true},
			{MobileDeviceID: "2", Selected: true},
		},
	}
	_, _, err := svc.CreateStaticV2(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, req.Assignments, 2)
	assert.Equal(t, "1", req.Assignments[0].MobileDeviceID)
	assert.Equal(t, "2", req.Assignments[1].MobileDeviceID)
}

func TestUnit_MobileDeviceGroups_CreateStaticV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateStaticV2(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceGroups_CreateStaticV2_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateStaticV2(context.Background(), &RequestStaticMobileDeviceGroup{Name: ""})
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "group name is required")
}

func TestUnit_MobileDeviceGroups_UpdateStaticByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateStaticV2Mock()

	req := &RequestStaticMobileDeviceGroup{
		Name:   "Static Devices Updated",
		SiteId: "-1",
	}
	result, resp, err := svc.UpdateStaticByIDV2(context.Background(), "10", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Static Devices Updated", result.Name)
	require.NotNil(t, req.Assignments)
}

func TestUnit_MobileDeviceGroups_UpdateStaticByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestStaticMobileDeviceGroup{Name: "Test"}
	result, resp, err := svc.UpdateStaticByIDV2(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDeviceGroups_UpdateStaticByIDV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateStaticByIDV2(context.Background(), "10", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceGroups_DeleteStaticByIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteStaticV2Mock()

	resp, err := svc.DeleteStaticByIDV2(context.Background(), "10")
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_DeleteStaticByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteStaticByIDV2(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static mobile device group ID is required")
}

// -----------------------------------------------------------------------------
// List All, Membership, Erase (V2)
// -----------------------------------------------------------------------------

func TestUnit_MobileDeviceGroups_ListAllV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAllV2Mock()

	result, resp, err := svc.ListAllV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	assert.Equal(t, 1, result[0].ID)
	assert.True(t, result[0].IsSmartGroup)
}

func TestUnit_MobileDeviceGroups_GetStaticGroupMembershipV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetStaticGroupMembershipV2Mock()

	result, resp, err := svc.GetStaticGroupMembershipV2(context.Background(), "10", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "101", result.Results[0].MobileDeviceID)
	// Additive 11.29 fields.
	assert.Equal(t, "jdoe", result.Results[0].LastLoggedInUsernameMdm)
	assert.Equal(t, "2026-06-01T12:00:00Z", result.Results[0].LastLoggedInUsernameMdmTimestamp)
	// Additive 11.30 field.
	assert.Equal(t, "2022-10-17T11:48:56.307Z", result.Results[0].LastContactDate)
}

func TestUnit_MobileDeviceGroups_GetStaticGroupMembershipV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetStaticGroupMembershipV2(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "static group ID is required")
}

func TestUnit_MobileDeviceGroups_GetSmartGroupMembershipV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSmartGroupMembershipV2Mock()

	result, resp, err := svc.GetSmartGroupMembershipV2(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result.Results, 1)
	assert.Equal(t, "201", result.Results[0].MobileDeviceID)
	// Additive 11.30 field.
	assert.Equal(t, "2022-10-17T11:48:56.307Z", result.Results[0].LastContactDate)
}

func TestUnit_MobileDeviceGroups_GetSmartGroupMembershipV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSmartGroupMembershipV2(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "smart group ID is required")
}

func TestUnit_MobileDeviceGroups_EraseDevicesByGroupIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterEraseDevicesV2Mock()

	resp, err := svc.EraseDevicesByGroupIDV2(context.Background(), "1", nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_MobileDeviceGroups_EraseDevicesByGroupIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.EraseDevicesByGroupIDV2(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "group ID is required")
}
