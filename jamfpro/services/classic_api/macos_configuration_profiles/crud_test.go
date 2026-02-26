package macos_configuration_profiles_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/macos_configuration_profiles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/macos_configuration_profiles/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MacOSConfigurationProfiles_List(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterListMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	resp, _, err := svc.List(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 2, resp.Size)
	assert.Len(t, resp.Results, 2)
	assert.Equal(t, "Wi-Fi Profile", resp.Results[0].Name)
	assert.Equal(t, 1, resp.Results[0].ID)
	assert.Equal(t, "Screen Saver Profile", resp.Results[1].Name)
	assert.Equal(t, 2, resp.Results[1].ID)
}

func TestUnit_MacOSConfigurationProfiles_GetByID(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterGetByIDMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	resp, _, err := svc.GetByID(context.Background(), 1)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
	assert.Equal(t, "Corporate Wi-Fi configuration", resp.General.Description)
	assert.NotNil(t, resp.General.Site)
	assert.Equal(t, -1, resp.General.Site.ID)
	assert.True(t, resp.Scope.AllComputers)
	assert.False(t, resp.Scope.AllJSSUsers)
	assert.Equal(t, "Install Wi-Fi", resp.SelfService.SelfServiceDisplayName)
}

func TestUnit_MacOSConfigurationProfiles_GetByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile ID must be a positive integer")
}

func TestUnit_MacOSConfigurationProfiles_GetByName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterGetByNameMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Wi-Fi Profile")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
}

func TestUnit_MacOSConfigurationProfiles_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name cannot be empty")
}

func TestUnit_MacOSConfigurationProfiles_Create(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterCreateMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:          "Test Profile",
			UserRemovable: false,
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: &macos_configuration_profiles.SubsetScope{
			AllComputers: true,
			AllJSSUsers:  false,
		},
	}

	resp, _, err := svc.Create(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 123, resp.ID)
}

func TestUnit_MacOSConfigurationProfiles_Create_NilRequest(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MacOSConfigurationProfiles_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name: "",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name is required")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByID(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterUpdateByIDMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:          "Updated Profile",
			UserRemovable: true,
		},
	}

	resp, _, err := svc.UpdateByID(context.Background(), 1, req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_MacOSConfigurationProfiles_UpdateByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name: "Test",
		},
	}

	_, _, err := svc.UpdateByID(context.Background(), 0, req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile ID must be a positive integer")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterUpdateByNameMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:          "Updated Profile",
			UserRemovable: true,
		},
	}

	resp, _, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
}

func TestUnit_MacOSConfigurationProfiles_UpdateByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name: "Test",
		},
	}

	_, _, err := svc.UpdateByName(context.Background(), "", req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name cannot be empty")
}

func TestUnit_MacOSConfigurationProfiles_DeleteByID(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterDeleteByIDMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_MacOSConfigurationProfiles_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile ID must be a positive integer")
}

func TestUnit_MacOSConfigurationProfiles_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterDeleteByNameMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Wi-Fi Profile")

	require.NoError(t, err)
}

func TestUnit_MacOSConfigurationProfiles_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name cannot be empty")
}

func TestUnit_MacOSConfigurationProfiles_NotFound(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_MacOSConfigurationProfiles_Conflict(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterConflictErrorMock()
	svc := macos_configuration_profiles.NewService(mockClient)

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name: "Duplicate Profile",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "configuration profile with that name already exists")
}

func TestUnit_MacOSConfigurationProfiles_List_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	_, _, err := svc.GetByName(context.Background(), "Wi-Fi Profile")
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	req := &macos_configuration_profiles.RequestResource{General: macos_configuration_profiles.SubsetGeneral{Name: ""}}
	_, _, err := svc.UpdateByID(context.Background(), 1, req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name is required")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	req := &macos_configuration_profiles.RequestResource{General: macos_configuration_profiles.SubsetGeneral{Name: "Test"}}
	_, _, err := svc.UpdateByID(context.Background(), 1, req)
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	req := &macos_configuration_profiles.RequestResource{General: macos_configuration_profiles.SubsetGeneral{Name: ""}}
	_, _, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name is required in request")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	req := &macos_configuration_profiles.RequestResource{General: macos_configuration_profiles.SubsetGeneral{Name: "Updated"}}
	_, _, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	_, err := svc.DeleteByID(context.Background(), 1)
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewService(mockClient)
	_, err := svc.DeleteByName(context.Background(), "Wi-Fi Profile")
	require.Error(t, err)
}
