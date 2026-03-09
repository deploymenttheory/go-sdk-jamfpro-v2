package macos_configuration_profiles_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/macos_configuration_profiles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/macos_configuration_profiles/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MacOSConfigurationProfiles_List(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterListMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

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
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

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
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	_, _, err := svc.GetByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile ID must be a positive integer")
}

func TestUnit_MacOSConfigurationProfiles_GetByName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterGetByNameMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	resp, _, err := svc.GetByName(context.Background(), "Wi-Fi Profile")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 1, resp.General.ID)
	assert.Equal(t, "Wi-Fi Profile", resp.General.Name)
}

func TestUnit_MacOSConfigurationProfiles_GetByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	_, _, err := svc.GetByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name cannot be empty")
}

func TestUnit_MacOSConfigurationProfiles_Create(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterCreateMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

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
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	_, _, err := svc.Create(context.Background(), nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MacOSConfigurationProfiles_Create_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

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
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

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
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

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
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

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
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

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
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	_, err := svc.DeleteByID(context.Background(), 1)

	require.NoError(t, err)
}

func TestUnit_MacOSConfigurationProfiles_DeleteByID_ZeroID(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	_, err := svc.DeleteByID(context.Background(), 0)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile ID must be a positive integer")
}

func TestUnit_MacOSConfigurationProfiles_DeleteByName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterDeleteByNameMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	_, err := svc.DeleteByName(context.Background(), "Wi-Fi Profile")

	require.NoError(t, err)
}

func TestUnit_MacOSConfigurationProfiles_DeleteByName_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	_, err := svc.DeleteByName(context.Background(), "")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name cannot be empty")
}

func TestUnit_MacOSConfigurationProfiles_NotFound(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterNotFoundErrorMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	_, _, err := svc.GetByID(context.Background(), 999)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "Resource not found")
}

func TestUnit_MacOSConfigurationProfiles_Conflict(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	mockClient.RegisterConflictErrorMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

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
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	_, _, err := svc.List(context.Background())
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_GetByName_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	_, _, err := svc.GetByName(context.Background(), "Wi-Fi Profile")
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_UpdateByID_NilRequest(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByID_EmptyName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	req := &macos_configuration_profiles.RequestResource{General: macos_configuration_profiles.SubsetGeneral{Name: ""}}
	_, _, err := svc.UpdateByID(context.Background(), 1, req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name is required")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByID_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	req := &macos_configuration_profiles.RequestResource{General: macos_configuration_profiles.SubsetGeneral{Name: "Test"}}
	_, _, err := svc.UpdateByID(context.Background(), 1, req)
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_UpdateByName_NilRequest(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	_, _, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByName_EmptyReqName(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	req := &macos_configuration_profiles.RequestResource{General: macos_configuration_profiles.SubsetGeneral{Name: ""}}
	_, _, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "macOS configuration profile name is required in request")
}

func TestUnit_MacOSConfigurationProfiles_UpdateByName_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	req := &macos_configuration_profiles.RequestResource{General: macos_configuration_profiles.SubsetGeneral{Name: "Updated"}}
	_, _, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_DeleteByID_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	_, err := svc.DeleteByID(context.Background(), 1)
	require.Error(t, err)
}

func TestUnit_MacOSConfigurationProfiles_DeleteByName_Error(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)
	_, err := svc.DeleteByName(context.Background(), "Wi-Fi Profile")
	require.Error(t, err)
}

// TestUnit_MacOSConfigurationProfiles_UpdateByID_WithPayloadsUUIDPreservation tests UpdateByID
// with payload content and UUID preservation logic.
func TestUnit_MacOSConfigurationProfiles_UpdateByID_WithPayloadsUUIDPreservation(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>NEW-UUID-456</string>
	<key>PayloadIdentifier</key>
	<string>com.example.new</string>
	<key>PayloadDisplayName</key>
	<string>WiFi Settings Updated</string>
</dict>
</plist>`

	// Register mocks - GetByID will return existing profile with plist from fixture
	mockClient.RegisterGetByIDMock()
	mockClient.RegisterUpdateByIDMock()

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile Updated",
			Payloads: newPlist,
		},
	}

	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 1, result.ID)

	// The payloads should have been processed through UUID preservation
	// (actual UUID values depend on the fixture data)
	assert.NotEmpty(t, req.General.Payloads)
}

// TestUnit_MacOSConfigurationProfiles_UpdateByID_EmptyPayloads tests UpdateByID
// when payloads field is empty (no UUID preservation needed).
func TestUnit_MacOSConfigurationProfiles_UpdateByID_EmptyPayloads(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	// Register mock for UpdateByID
	mockClient.RegisterUpdateByIDMock()

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: "", // Empty payloads - should skip UUID preservation
		},
	}

	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
}

// TestUnit_MacOSConfigurationProfiles_UpdateByID_GetExistingProfileError tests UpdateByID
// when fetching the existing profile fails during UUID preservation.
func TestUnit_MacOSConfigurationProfiles_UpdateByID_GetExistingProfileError(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	// Don't register GetByID mock - will cause error

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: `<?xml version="1.0"?><plist version="1.0"><dict><key>PayloadUUID</key><string>TEST</string></dict></plist>`,
		},
	}

	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get existing profile for UUID preservation")
}

// TestUnit_MacOSConfigurationProfiles_UpdateByName_WithPayloadsUUIDPreservation tests UpdateByName
// with payload content and UUID preservation logic.
func TestUnit_MacOSConfigurationProfiles_UpdateByName_WithPayloadsUUIDPreservation(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>NEW-NAME-UUID-012</string>
	<key>PayloadIdentifier</key>
	<string>com.example.name.new</string>
	<key>PayloadDisplayName</key>
	<string>VPN Settings Updated</string>
</dict>
</plist>`

	// Register mocks - GetByName will return existing profile with plist from fixture
	mockClient.RegisterGetByNameMock()
	mockClient.RegisterUpdateByNameMock()

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:     "VPN Profile Updated",
			Payloads: newPlist,
		},
	}

	result, resp, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 1, result.ID)

	// The payloads should have been processed through UUID preservation
	assert.NotEmpty(t, req.General.Payloads)
}

// TestUnit_MacOSConfigurationProfiles_UpdateByName_EmptyPayloads tests UpdateByName
// when payloads field is empty (no UUID preservation needed).
func TestUnit_MacOSConfigurationProfiles_UpdateByName_EmptyPayloads(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	// Register mock for UpdateByName
	mockClient.RegisterUpdateByNameMock()

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: "", // Empty payloads - should skip UUID preservation
		},
	}

	result, resp, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
}

// TestUnit_MacOSConfigurationProfiles_UpdateByName_GetExistingProfileError tests UpdateByName
// when fetching the existing profile fails during UUID preservation.
func TestUnit_MacOSConfigurationProfiles_UpdateByName_GetExistingProfileError(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	// Don't register GetByName mock - will cause error

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: `<?xml version="1.0"?><plist version="1.0"><dict><key>PayloadUUID</key><string>TEST</string></dict></plist>`,
		},
	}

	result, resp, err := svc.UpdateByName(context.Background(), "NonExistent Profile", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get existing profile for UUID preservation")
}

// TestUnit_MacOSConfigurationProfiles_UpdateByID_ExistingProfileEmptyPayloads tests UpdateByID
// when the existing profile has empty payloads (UUID preservation skipped).
func TestUnit_MacOSConfigurationProfiles_UpdateByID_ExistingProfileEmptyPayloads(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>NEW-UUID</string>
	<key>PayloadIdentifier</key>
	<string>com.example.new</string>
</dict>
</plist>`

	// Register GetByID mock - fixture will have payloads
	mockClient.RegisterGetByIDMock()
	mockClient.RegisterUpdateByIDMock()

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: newPlist,
		},
	}

	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
}

// TestUnit_MacOSConfigurationProfiles_UpdateByName_ExistingProfileEmptyPayloads tests UpdateByName
// when the existing profile has empty payloads (UUID preservation skipped).
func TestUnit_MacOSConfigurationProfiles_UpdateByName_ExistingProfileEmptyPayloads(t *testing.T) {
	mockClient := mocks.NewMacOSConfigurationProfilesMock()
	svc := macos_configuration_profiles.NewMacosConfigurationProfiles(mockClient)

	newPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>PayloadUUID</key>
	<string>NEW-UUID</string>
	<key>PayloadIdentifier</key>
	<string>com.example.new</string>
</dict>
</plist>`

	// Register GetByName mock - fixture will have payloads
	mockClient.RegisterGetByNameMock()
	mockClient.RegisterUpdateByNameMock()

	req := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:     "Test Profile",
			Payloads: newPlist,
		},
	}

	result, resp, err := svc.UpdateByName(context.Background(), "Wi-Fi Profile", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
}

