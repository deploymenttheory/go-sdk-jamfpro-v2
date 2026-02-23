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

func TestUnitList_Success(t *testing.T) {
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

func TestUnitGetByID_Success(t *testing.T) {
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

func TestUnitGetByName_Success(t *testing.T) {
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

func TestUnitGetByName_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListEmptyMock()

	_, _, err := svc.GetByName(context.Background(), "NonExistent")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestUnitGetMembership_Success(t *testing.T) {
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

func TestUnitCreate_Success(t *testing.T) {
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

func TestUnitUpdateByID_Success(t *testing.T) {
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

func TestUnitDeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitGetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetByID(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnitGetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
}

func TestUnitGetMembership_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetMembership(context.Background(), "", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ID is required")
}

func TestUnitCreate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.Create(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitUpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestSmartMobileDeviceGroup{GroupName: "Test"}

	_, _, err := svc.UpdateByID(context.Background(), "", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnitUpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.UpdateByID(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	_, err := svc.DeleteByID(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "ID is required")
}
