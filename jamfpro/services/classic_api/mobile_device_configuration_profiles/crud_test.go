package mobile_device_configuration_profiles_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_configuration_profiles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_configuration_profiles/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDeviceConfigurationProfiles_List(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterListMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Len(t, resp.ConfigurationProfiles, 2)
	assert.Equal(t, "Wi-Fi Profile", resp.ConfigurationProfiles[0].Name)
	assert.Equal(t, 1, resp.ConfigurationProfiles[0].ID)
	assert.Equal(t, "VPN Profile", resp.ConfigurationProfiles[1].Name)
	assert.Equal(t, 2, resp.ConfigurationProfiles[1].ID)
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterGetByIDMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
	assert.Equal(t, "Test Wi-Fi configuration", resp.General.Description)
	assert.NotNil(t, resp.Scope)
	assert.True(t, resp.Scope.AllMobileDevices)
	assert.NotNil(t, resp.SelfService)
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile ID must be a positive integer")
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterGetByNameMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Wi-Fi Profile")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile name cannot be empty")
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByIDWithSubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterGetByIDWithSubsetMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	resp, _, err := svc.GetByIDWithSubset(context.Background(), 1, "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByIDWithSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	_, _, err := svc.GetByIDWithSubset(context.Background(), 1, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByNameWithSubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterGetByNameWithSubsetMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	resp, _, err := svc.GetByNameWithSubset(context.Background(), "Wi-Fi Profile", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
}

func TestUnit_MobileDeviceConfigurationProfiles_Create(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterCreateMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name: "Test Profile",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: &mobile_device_configuration_profiles.SubsetScope{
			AllMobileDevices: true,
			AllJSSUsers:      false,
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_MobileDeviceConfigurationProfiles_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceConfigurationProfiles_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile name is required")
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterUpdateByIDMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name: "Updated Profile",
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile ID must be a positive integer")
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterUpdateByNameMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name: "Updated Profile",
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{Name: "Test"},
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile name cannot be empty")
}

func TestUnit_MobileDeviceConfigurationProfiles_DeleteByID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterDeleteByIDMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile ID must be a positive integer")
}

func TestUnit_MobileDeviceConfigurationProfiles_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterDeleteByNameMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Wi-Fi Profile")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile name cannot be empty")
}

func TestUnit_MobileDeviceConfigurationProfiles_NotFound(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_MobileDeviceConfigurationProfiles_Conflict(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterConflictErrorMock()
	svc := mobile_device_configuration_profiles.NewService(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name: "Duplicate Profile",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile with that name already exists")
}
