package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDeviceConfigurationProfilesMock struct {
	*mocks.GenericMock
}

func NewMobileDeviceConfigurationProfilesMock() *MobileDeviceConfigurationProfilesMock {
	return &MobileDeviceConfigurationProfilesMock{
		GenericMock: mocks.NewXMLMock("MobileDeviceConfigurationProfilesMock"),
	}
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterGetByIDWithSubsetMock()
	m.RegisterGetByNameWithSubsetMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/mobiledeviceconfigurationprofiles", 200, "validate_list_mobile_device_configuration_profiles.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/id/1", 200, "validate_get_mobile_device_configuration_profile.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile", 200, "validate_get_mobile_device_configuration_profile.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByIDWithSubsetMock() {
	m.Register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/id/1/subset/General", 200, "validate_get_mobile_device_configuration_profile.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByNameWithSubsetMock() {
	m.Register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile/subset/General", 200, "validate_get_mobile_device_configuration_profile.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/mobiledeviceconfigurationprofiles/id/0", 201, "validate_create_mobile_device_configuration_profile.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceconfigurationprofiles/id/1", 200, "validate_update_mobile_device_configuration_profile.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile", 200, "validate_update_mobile_device_configuration_profile.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceconfigurationprofiles/id/1", 200, "")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceconfigurationprofiles/name/Wi-Fi Profile", 200, "")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByIDMockEmptyPayloads() {
	m.Register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/id/1", 200, "validate_get_mobile_device_configuration_profile_empty_payloads.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterGetByNameMockEmptyPayloads() {
	m.Register("GET", "/JSSResource/mobiledeviceconfigurationprofiles/name/Test Profile", 200, "validate_get_mobile_device_configuration_profile_empty_payloads.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterUpdateByNameMockTestProfile() {
	m.Register("PUT", "/JSSResource/mobiledeviceconfigurationprofiles/name/Test Profile", 200, "validate_update_mobile_device_configuration_profile.xml")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/mobiledeviceconfigurationprofiles/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *MobileDeviceConfigurationProfilesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/mobiledeviceconfigurationprofiles/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A mobile device configuration profile with that name already exists")
}

