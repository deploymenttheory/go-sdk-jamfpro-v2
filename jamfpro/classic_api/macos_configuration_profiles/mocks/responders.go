package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MacOSConfigurationProfilesMock struct {
	*mocks.GenericMock
}

func NewMacOSConfigurationProfilesMock() *MacOSConfigurationProfilesMock {
	return &MacOSConfigurationProfilesMock{
		GenericMock: mocks.NewXMLMock("MacOSConfigurationProfilesMock"),
	}
}

func (m *MacOSConfigurationProfilesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

func (m *MacOSConfigurationProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *MacOSConfigurationProfilesMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/osxconfigurationprofiles", 200, "validate_list_osx_configuration_profiles.xml")
}

func (m *MacOSConfigurationProfilesMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/osxconfigurationprofiles/id/1", 200, "validate_get_osx_configuration_profile.xml")
}

func (m *MacOSConfigurationProfilesMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/osxconfigurationprofiles/name/Wi-Fi Profile", 200, "validate_get_osx_configuration_profile.xml")
}

func (m *MacOSConfigurationProfilesMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/osxconfigurationprofiles/id/0", 201, "validate_create_osx_configuration_profile.xml")
}

func (m *MacOSConfigurationProfilesMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/osxconfigurationprofiles/id/1", 200, "validate_update_osx_configuration_profile.xml")
}

func (m *MacOSConfigurationProfilesMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/osxconfigurationprofiles/name/Wi-Fi Profile", 200, "validate_update_osx_configuration_profile.xml")
}

func (m *MacOSConfigurationProfilesMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/osxconfigurationprofiles/id/1", 200, "")
}

func (m *MacOSConfigurationProfilesMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/osxconfigurationprofiles/name/Wi-Fi Profile", 200, "")
}

func (m *MacOSConfigurationProfilesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/osxconfigurationprofiles/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *MacOSConfigurationProfilesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/osxconfigurationprofiles/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A configuration profile with that name already exists")
}

