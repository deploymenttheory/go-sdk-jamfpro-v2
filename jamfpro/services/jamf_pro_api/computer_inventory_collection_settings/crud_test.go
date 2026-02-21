package computer_inventory_collection_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory_collection_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetV2(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	mock.RegisterGetMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetV2(ctx)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.ComputerInventoryCollectionPreferences.MonitorApplicationUsage)
	assert.True(t, result.ComputerInventoryCollectionPreferences.IncludePackages)
	assert.Len(t, result.ApplicationPaths, 2)
	assert.Equal(t, "/Applications", result.ApplicationPaths[0].Path)
}

func TestUpdateV2(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	mock.RegisterUpdateMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceComputerInventoryCollectionSettings{
		ComputerInventoryCollectionPreferences: Preferences{
			MonitorApplicationUsage: false,
			IncludePackages:         true,
		},
	}

	resp, err := svc.UpdateV2(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUpdateV2_NilRequest(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.UpdateV2(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestCreateCustomPathV2(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	mock.RegisterCreateCustomPathMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &CustomPathRequest{
		Scope: "USER_LIBRARY",
		Path:  "/Library/Custom",
	}

	result, resp, err := svc.CreateCustomPathV2(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "/api/v2/computer-inventory-collection-settings/custom-path/3", result.Href)
}

func TestCreateCustomPathV2_NilRequest(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.CreateCustomPathV2(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestDeleteCustomPathByIDV2(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	mock.RegisterDeleteCustomPathMock("3")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteCustomPathByIDV2(ctx, "3")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestDeleteCustomPathByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteCustomPathByIDV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}
