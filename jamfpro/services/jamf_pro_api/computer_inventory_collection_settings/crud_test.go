package computer_inventory_collection_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory_collection_settings/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUnit_ComputerInventoryCollectionSettings_GetV2_Success(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	mock.RegisterGetMock()

	svc := NewService(mock)
	result, resp, err := svc.GetV2(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.True(t, result.ComputerInventoryCollectionPreferences.MonitorApplicationUsage)
	assert.Len(t, result.ApplicationPaths, 2)
}

func TestUnit_ComputerInventoryCollectionSettings_GetV2_ClientError(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	// No mock registered - dispatch returns (nil, err)

	svc := NewService(mock)
	result, resp, err := svc.GetV2(context.Background())

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnit_ComputerInventoryCollectionSettings_UpdateV2_Success(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	mock.RegisterUpdateMock()

	svc := NewService(mock)
	settings := &ResourceComputerInventoryCollectionSettings{
		ComputerInventoryCollectionPreferences: SubsetPreferences{
			MonitorApplicationUsage: false,
		},
	}
	resp, err := svc.UpdateV2(context.Background(), settings)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_ComputerInventoryCollectionSettings_UpdateV2_NilSettings(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	svc := NewService(mock)

	resp, err := svc.UpdateV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "settings is required")
}

func TestUnit_ComputerInventoryCollectionSettings_UpdateV2_ClientError(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	// No mock registered

	svc := NewService(mock)
	settings := &ResourceComputerInventoryCollectionSettings{
		ComputerInventoryCollectionPreferences: SubsetPreferences{MonitorApplicationUsage: false},
	}
	resp, err := svc.UpdateV2(context.Background(), settings)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnit_ComputerInventoryCollectionSettings_CreateCustomPathV2_Success(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	mock.RegisterCreateCustomPathMock()

	svc := NewService(mock)
	req := &RequestCustomPath{
		Scope: "USER",
		Path:  "/Users/Shared/CustomApp",
	}
	result, resp, err := svc.CreateCustomPathV2(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, result)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "custom-path/3")
}

func TestUnit_ComputerInventoryCollectionSettings_CreateCustomPathV2_NilRequest(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	svc := NewService(mock)

	result, resp, err := svc.CreateCustomPathV2(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request body is required")
}

func TestUnit_ComputerInventoryCollectionSettings_CreateCustomPathV2_EmptyPath(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	svc := NewService(mock)

	req := &RequestCustomPath{Scope: "USER", Path: ""}
	result, resp, err := svc.CreateCustomPathV2(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "path is required")
}

func TestUnit_ComputerInventoryCollectionSettings_CreateCustomPathV2_ClientError(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	// No mock registered

	svc := NewService(mock)
	req := &RequestCustomPath{Scope: "USER", Path: "/Users/Shared/CustomApp"}
	result, resp, err := svc.CreateCustomPathV2(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnit_ComputerInventoryCollectionSettings_DeleteCustomPathByIDV2_Success(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	mock.RegisterDeleteCustomPathMock()

	svc := NewService(mock)
	resp, err := svc.DeleteCustomPathByIDV2(context.Background(), "3")

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_ComputerInventoryCollectionSettings_DeleteCustomPathByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	svc := NewService(mock)

	resp, err := svc.DeleteCustomPathByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "custom path ID is required")
}

func TestUnit_ComputerInventoryCollectionSettings_DeleteCustomPathByIDV2_ClientError(t *testing.T) {
	mock := mocks.NewComputerInventoryCollectionSettingsMock()
	// No mock registered

	svc := NewService(mock)
	resp, err := svc.DeleteCustomPathByIDV2(context.Background(), "3")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no mock registered")
}
