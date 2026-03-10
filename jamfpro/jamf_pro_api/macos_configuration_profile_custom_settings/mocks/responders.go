package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MacOSConfigProfileCustomSettingsMock struct {
	*mocks.GenericMock
}

func NewMacOSConfigProfileCustomSettingsMock() *MacOSConfigProfileCustomSettingsMock {
	return &MacOSConfigProfileCustomSettingsMock{
		GenericMock: mocks.NewJSONMock("MacOSConfigProfileCustomSettingsMock"),
	}
}

func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetSchemaListMock() {
	m.Register("GET", "/api/config-profiles/macos/custom-settings/v1/schema-list", 200, "validate_list_schema.json")
}

func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetByPayloadUUIDMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.Register("GET", path, 200, "validate_get.json")
}

func (m *MacOSConfigProfileCustomSettingsMock) RegisterCreateMock() {
	m.Register("POST", "/api/config-profiles/macos", 200, "validate_create.json")
}

func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetSchemaListErrorMock() {
	m.RegisterError("GET", "/api/config-profiles/macos/custom-settings/v1/schema-list", 500, "", "request failed: 500 Internal Server Error")
}

func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetByPayloadUUIDErrorMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.RegisterError("GET", path, 404, "", "request failed: 404 Not Found")
}

func (m *MacOSConfigProfileCustomSettingsMock) RegisterCreateErrorMock() {
	m.RegisterError("POST", "/api/config-profiles/macos", 500, "", "request failed: 500 Internal Server Error")
}

func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetSchemaListNoResponseErrorMock() {
	m.RegisterError("GET", "/api/config-profiles/macos/custom-settings/v1/schema-list", 500, "error_internal.json", "no response registered")
}

func (m *MacOSConfigProfileCustomSettingsMock) RegisterGetByPayloadUUIDNoResponseErrorMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.RegisterError("GET", path, 500, "error_internal.json", "no response registered")
}

func (m *MacOSConfigProfileCustomSettingsMock) RegisterCreateNoResponseErrorMock() {
	m.RegisterError("POST", "/api/config-profiles/macos", 500, "error_internal.json", "no response registered")
}
