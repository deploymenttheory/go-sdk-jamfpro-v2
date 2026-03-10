package mobile_device_configuration_profiles_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_configuration_profiles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_configuration_profiles/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDeviceConfigurationProfiles_List(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterListMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

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
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

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
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile ID must be a positive integer")
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterGetByNameMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Wi-Fi Profile")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile name cannot be empty")
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByIDWithSubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterGetByIDWithSubsetMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	resp, _, err := svc.GetByIDWithSubset(context.Background(), 1, "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByIDWithSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByIDWithSubset(context.Background(), 1, "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByNameWithSubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterGetByNameWithSubsetMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	resp, _, err := svc.GetByNameWithSubset(context.Background(), "Wi-Fi Profile", "General")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
}

func TestUnit_MobileDeviceConfigurationProfiles_Create(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterCreateMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name: "Test Profile",
			Site: &models.SharedResourceSite{
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
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceConfigurationProfiles_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

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
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

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
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

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
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

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
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

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
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile ID must be a positive integer")
}

func TestUnit_MobileDeviceConfigurationProfiles_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterDeleteByNameMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Wi-Fi Profile")

	require.NoError(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile name cannot be empty")
}

func TestUnit_MobileDeviceConfigurationProfiles_NotFound(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_MobileDeviceConfigurationProfiles_Conflict(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	mockClient.RegisterConflictErrorMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name: "Duplicate Profile",
		},
	}

	_, _, err := svc.Create(context.Background(), req)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile with that name already exists")
}

func TestUnit_MobileDeviceConfigurationProfiles_List_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByName(context.Background(), "Unknown")
	require.Error(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByIDWithSubset_ZeroID(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByIDWithSubset(context.Background(), 0, "General")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile ID must be a positive integer")
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByIDWithSubset_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByIDWithSubset(context.Background(), 2, "General")
	require.Error(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByNameWithSubset_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByNameWithSubset(context.Background(), "", "General")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile name cannot be empty")
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByNameWithSubset_EmptySubset(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByNameWithSubset(context.Background(), "Test", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "subset cannot be empty")
}

func TestUnit_MobileDeviceConfigurationProfiles_GetByNameWithSubset_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.GetByNameWithSubset(context.Background(), "Unknown", "General")
	require.Error(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.UpdateByID(context.Background(), 1, req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile name is required")
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.UpdateByID(context.Background(), 2, req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, _, err := svc.UpdateByName(context.Background(), "Test", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{Name: ""},
	}
	_, _, err := svc.UpdateByName(context.Background(), "Test", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mobile device configuration profile name is required in request")
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{Name: "Test"},
	}
	_, _, err := svc.UpdateByName(context.Background(), "Unknown", req)
	require.Error(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, err := svc.DeleteByID(context.Background(), 999)
	require.Error(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Unknown")
	require.Error(t, err)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByID_WithPayloadsUUIDPreservation(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	mockClient.RegisterGetByIDMock()
	mockClient.RegisterUpdateByIDMock()

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadContent</key>
	<array/>
	<key>PayloadDisplayName</key>
	<string>New Profile</string>
	<key>PayloadIdentifier</key>
	<string>new-identifier</string>
	<key>PayloadType</key>
	<string>Configuration</string>
	<key>PayloadUUID</key>
	<string>new-uuid</string>
	<key>PayloadVersion</key>
	<integer>1</integer>
</dict>
</plist>`

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: newPlist,
		},
	}

	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByID_EmptyPayloads(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	mockClient.RegisterUpdateByIDMock()

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: "",
		},
	}

	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByID_GetExistingProfileError(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>new-uuid</string>
</dict>
</plist>`

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: newPlist,
		},
	}

	_, _, err := svc.UpdateByID(context.Background(), 999, req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get existing profile for UUID preservation")
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByName_WithPayloadsUUIDPreservation(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	mockClient.RegisterGetByNameMock()
	mockClient.RegisterUpdateByNameMock()

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadContent</key>
	<array/>
	<key>PayloadDisplayName</key>
	<string>New Profile</string>
	<key>PayloadIdentifier</key>
	<string>new-identifier</string>
	<key>PayloadType</key>
	<string>Configuration</string>
	<key>PayloadUUID</key>
	<string>new-uuid</string>
	<key>PayloadVersion</key>
	<integer>1</integer>
</dict>
</plist>`

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:     "Wi-Fi Profile",
			Payloads: newPlist,
		},
	}

	result, resp, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByName_EmptyPayloads(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	mockClient.RegisterUpdateByNameMockTestProfile()

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: "",
		},
	}

	result, resp, err := svc.UpdateByName(context.Background(), "Test Profile", req)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByName_GetExistingProfileError(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>new-uuid</string>
</dict>
</plist>`

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: newPlist,
		},
	}

	_, _, err := svc.UpdateByName(context.Background(), "Unknown", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get existing profile for UUID preservation")
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByID_ExistingProfileEmptyPayloads(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	mockClient.RegisterGetByIDMockEmptyPayloads()
	mockClient.RegisterUpdateByIDMock()

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>new-uuid</string>
</dict>
</plist>`

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: newPlist,
		},
	}

	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_MobileDeviceConfigurationProfiles_UpdateByName_ExistingProfileEmptyPayloads(t *testing.T) {
	mockClient := mocks.NewMobileDeviceConfigurationProfilesMock()
	svc := mobile_device_configuration_profiles.NewMobileDeviceConfigurationProfiles(mockClient)

	mockClient.RegisterGetByNameMockEmptyPayloads()
	mockClient.RegisterUpdateByNameMockTestProfile()

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>new-uuid</string>
</dict>
</plist>`

	req := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: newPlist,
		},
	}

	result, resp, err := svc.UpdateByName(context.Background(), "Test Profile", req)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, result.ID)
}
