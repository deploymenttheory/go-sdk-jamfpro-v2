package smart_mobile_device_groups

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/smart_mobile_device_groups/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SmartMobileDeviceGroupsMock) {
	t.Helper()
	mock := mocks.NewSmartMobileDeviceGroupsMock()
	return NewService(mock), mock
}

func TestUnit_SmartMobileDeviceGroups_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].GroupID)
	assert.Equal(t, "iPhones", result.Results[0].GroupName)
}

func TestUnit_SmartMobileDeviceGroups_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.GroupID)
	assert.Equal(t, "iPhones", result.GroupName)
	require.Len(t, result.Criteria, 1)
	assert.Equal(t, "Model", result.Criteria[0].Name)
}

func TestUnit_SmartMobileDeviceGroups_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByName(context.Background(), "iPhones")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.GroupID)
	assert.Equal(t, "iPhones", result.GroupName)
}

func TestUnit_SmartMobileDeviceGroups_GetByName_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListEmptyMock()

	_, _, err := svc.GetByName(context.Background(), "NonExistent")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestUnit_SmartMobileDeviceGroups_GetMembership_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMembershipMock()

	result, resp, err := svc.GetMembership(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].MobileDeviceId)
	assert.Equal(t, "iPhone 14", result.Results[0].DisplayName)
}

func TestUnit_SmartMobileDeviceGroups_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestSmartMobileDeviceGroup{
		GroupName:        "New Smart",
		GroupDescription: "Desc",
		Criteria: []SharedSubsetCriteriaJamfProAPI{
			{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "iPhone"},
		},
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "2", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_SmartMobileDeviceGroups_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &RequestSmartMobileDeviceGroup{
		GroupName:        "iPhones Updated",
		GroupDescription: "Updated",
		Criteria: []SharedSubsetCriteriaJamfProAPI{
			{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "iPhone"},
		},
	}
	result, resp, err := svc.UpdateByID(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.GroupID)
	assert.Equal(t, "iPhones Updated", result.GroupName)
}

func TestUnit_SmartMobileDeviceGroups_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_SmartMobileDeviceGroups_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetByID(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_SmartMobileDeviceGroups_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestUnit_SmartMobileDeviceGroups_GetMembership_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetMembership(context.Background(), "", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_SmartMobileDeviceGroups_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.Create(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SmartMobileDeviceGroups_UpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestSmartMobileDeviceGroup{GroupName: "Test"}

	_, _, err := svc.UpdateByID(context.Background(), "", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_SmartMobileDeviceGroups_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.UpdateByID(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SmartMobileDeviceGroups_DeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	_, err := svc.DeleteByID(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnit_SmartMobileDeviceGroups_List_ClientError(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered - dispatch returns (nil, err)

	result, resp, err := svc.List(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_SmartMobileDeviceGroups_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, resp, err := svc.GetByID(context.Background(), "999")
	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Contains(t, err.Error(), "Jamf Pro API error")
}

func TestUnit_SmartMobileDeviceGroups_GetByName_ListError(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered - List returns error

	_, resp, err := svc.GetByName(context.Background(), "iPhones")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_SmartMobileDeviceGroups_GetMembership_ClientError(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered for membership endpoint

	result, resp, err := svc.GetMembership(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_SmartMobileDeviceGroups_Create_ClientError(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	req := &RequestSmartMobileDeviceGroup{GroupName: "New", GroupDescription: "Desc"}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_SmartMobileDeviceGroups_DeleteByID_ClientError(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	resp, err := svc.DeleteByID(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_SmartMobileDeviceGroups_UpdateByID_ClientError(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	req := &RequestSmartMobileDeviceGroup{GroupName: "Updated", GroupDescription: "Desc"}
	result, resp, err := svc.UpdateByID(context.Background(), "1", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}
