package mobile_device_groups_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_groups/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDeviceGroups_List(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterListMobileDeviceGroupsMock()
	svc := mobile_device_groups.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "All Mobile Devices", resp.Results[0].Name)
	assert.True(t, resp.Results[0].IsSmart)
	assert.Equal(t, "Static Mobile Group", resp.Results[1].Name)
	assert.False(t, resp.Results[1].IsSmart)
}

func TestUnit_MobileDeviceGroups_GetByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterGetMobileDeviceGroupByIDMock()
	svc := mobile_device_groups.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "All Mobile Devices", resp.Name)
	assert.True(t, resp.IsSmart)
	assert.NotNil(t, resp.Site)
	assert.Equal(t, -1, resp.Site.ID)
	assert.NotNil(t, resp.Criteria)
	assert.Equal(t, 2, resp.Criteria.Size)
	assert.Len(t, resp.Criteria.Criterion, 2)
	assert.Len(t, resp.MobileDevices, 1)
	assert.Equal(t, "test-iphone-01", resp.MobileDevices[0].Name)
}

func TestUnit_MobileDeviceGroups_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	svc := mobile_device_groups.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device group ID must be a positive integer")
}

func TestUnit_MobileDeviceGroups_GetByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterGetMobileDeviceGroupByNameMock()
	svc := mobile_device_groups.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "All Mobile Devices")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "All Mobile Devices", resp.Name)
	assert.True(t, resp.IsSmart)
}

func TestUnit_MobileDeviceGroups_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	svc := mobile_device_groups.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device group name cannot be empty")
}

func TestUnit_MobileDeviceGroups_Create(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterCreateMobileDeviceGroupMock()
	svc := mobile_device_groups.NewService(mockClient)

	req := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    "Test Group",
		IsSmart: true,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &mobile_device_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Model",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "iPhone",
				},
			},
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_MobileDeviceGroups_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	svc := mobile_device_groups.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceGroups_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	svc := mobile_device_groups.NewService(mockClient)

	req := &mobile_device_groups.RequestMobileDeviceGroup{
		Name: "",
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device group name is required")
}

func TestUnit_MobileDeviceGroups_UpdateByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterUpdateMobileDeviceGroupByIDMock()
	svc := mobile_device_groups.NewService(mockClient)

	req := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    "Updated Group",
		IsSmart: false,
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_MobileDeviceGroups_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	svc := mobile_device_groups.NewService(mockClient)

	req := &mobile_device_groups.RequestMobileDeviceGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device group ID must be a positive integer")
}

func TestUnit_MobileDeviceGroups_UpdateByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterUpdateMobileDeviceGroupByNameMock()
	svc := mobile_device_groups.NewService(mockClient)

	req := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    "Updated Group",
		IsSmart: false,
	}

	resp, _, err := svc.UpdateByName(context.Background(), "All Mobile Devices", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_MobileDeviceGroups_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	svc := mobile_device_groups.NewService(mockClient)

	req := &mobile_device_groups.RequestMobileDeviceGroup{
		Name: "Test",
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device group name cannot be empty")
}

func TestUnit_MobileDeviceGroups_DeleteByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterDeleteMobileDeviceGroupByIDMock()
	svc := mobile_device_groups.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_MobileDeviceGroups_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	svc := mobile_device_groups.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device group ID must be a positive integer")
}

func TestUnit_MobileDeviceGroups_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterDeleteMobileDeviceGroupByNameMock()
	svc := mobile_device_groups.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "All Mobile Devices")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceGroups_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	svc := mobile_device_groups.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device group name cannot be empty")
}

func TestUnit_MobileDeviceGroups_NotFound(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := mobile_device_groups.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Mobile device group not found")
}

func TestUnit_MobileDeviceGroups_Conflict(t *testing.T) {
	mockClient := mocks.NewMobileDeviceGroupsMock()
	mockClient.RegisterConflictErrorMock()
	svc := mobile_device_groups.NewService(mockClient)

	req := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    "Duplicate Group",
		IsSmart: false,
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}
