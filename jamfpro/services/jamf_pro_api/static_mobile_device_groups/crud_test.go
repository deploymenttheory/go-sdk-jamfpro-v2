package static_mobile_device_groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/static_mobile_device_groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.StaticMobileDeviceGroupsMock) {
	t.Helper()
	mock := mocks.NewStaticMobileDeviceGroupsMock()
	return NewService(mock), mock
}

func TestUnit_StaticMobileDeviceGroups_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "10", result.Results[0].ID)
	assert.Equal(t, "Static Devices", result.Results[0].Name)
}

func TestUnit_StaticMobileDeviceGroups_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetByID(context.Background(), "10")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "10", result.ID)
	assert.Equal(t, "Static Devices", result.Name)
}

func TestUnit_StaticMobileDeviceGroups_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_StaticMobileDeviceGroups_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestStaticMobileDeviceGroup{
		Name:        "New Static",
		Description: "Desc",
		SiteID:      "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{},
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "11", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_StaticMobileDeviceGroups_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_StaticMobileDeviceGroups_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestStaticMobileDeviceGroup{
		Name:        "Static Devices Updated",
		Description: "Updated",
		SiteID:      "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{},
	}
	result, resp, err := svc.UpdateByID(context.Background(), "10", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "10", result.ID)
	assert.Equal(t, "Static Devices Updated", result.Name)
}

func TestUnit_StaticMobileDeviceGroups_UpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestStaticMobileDeviceGroup{Name: "Test", SiteID: "-1"}

	result, resp, err := svc.UpdateByID(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_StaticMobileDeviceGroups_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), "10", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_StaticMobileDeviceGroups_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByID(context.Background(), "10")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_StaticMobileDeviceGroups_DeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_StaticMobileDeviceGroups_List_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.List(context.Background(), nil)
	require.Error(t, err)
}

func TestUnit_StaticMobileDeviceGroups_GetByID_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), "10")
	require.Error(t, err)
}

func TestUnit_StaticMobileDeviceGroups_Create_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestStaticMobileDeviceGroup{
		Name:        "test",
		SiteID:      "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{},
	}
	_, _, err := svc.Create(context.Background(), req)
	require.Error(t, err)
}

func TestUnit_StaticMobileDeviceGroups_UpdateByID_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestStaticMobileDeviceGroup{
		Name:        "test",
		SiteID:      "-1",
		Assignments: []StaticMobileDeviceGroupAssignment{},
	}
	_, _, err := svc.UpdateByID(context.Background(), "10", req)
	require.Error(t, err)
}

func TestUnit_StaticMobileDeviceGroups_DeleteByID_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByID(context.Background(), "10")
	require.Error(t, err)
}
