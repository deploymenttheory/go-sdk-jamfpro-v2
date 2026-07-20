package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MacOSConfigProfilesMock struct {
	*mocks.GenericMock
}

func NewMacOSConfigProfilesMock() *MacOSConfigProfilesMock {
	return &MacOSConfigProfilesMock{
		GenericMock: mocks.NewJSONMock("MacOSConfigProfilesMock"),
	}
}

func (m *MacOSConfigProfilesMock) RegisterGetSchemaListMock() {
	m.Register("GET", "/api/config-profiles/macos/custom-settings/v1/schema-list", 200, "validate_list_schema.json")
}

func (m *MacOSConfigProfilesMock) RegisterGetByPayloadUUIDMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.Register("GET", path, 200, "validate_get.json")
}

func (m *MacOSConfigProfilesMock) RegisterCreateMock() {
	m.Register("POST", "/api/config-profiles/macos", 201, "validate_create.json")
}

func (m *MacOSConfigProfilesMock) RegisterUpdateByPayloadUUIDMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.Register("PUT", path, 200, "validate_update.json")
}

func (m *MacOSConfigProfilesMock) RegisterDeleteByPayloadUUIDMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.Register("DELETE", path, 204, "")
}

func (m *MacOSConfigProfilesMock) RegisterGetSchemaListErrorMock() {
	m.RegisterError("GET", "/api/config-profiles/macos/custom-settings/v1/schema-list", 500, "", "request failed: 500 Internal Server Error")
}

func (m *MacOSConfigProfilesMock) RegisterGetByPayloadUUIDErrorMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.RegisterError("GET", path, 404, "", "request failed: 404 Not Found")
}

func (m *MacOSConfigProfilesMock) RegisterCreateErrorMock() {
	m.RegisterError("POST", "/api/config-profiles/macos", 500, "", "request failed: 500 Internal Server Error")
}

func (m *MacOSConfigProfilesMock) RegisterUpdateByPayloadUUIDErrorMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.RegisterError("PUT", path, 404, "", "request failed: 404 Not Found")
}

func (m *MacOSConfigProfilesMock) RegisterDeleteByPayloadUUIDErrorMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.RegisterError("DELETE", path, 404, "", "request failed: 404 Not Found")
}

func (m *MacOSConfigProfilesMock) RegisterGetSchemaListNoResponseErrorMock() {
	m.RegisterError("GET", "/api/config-profiles/macos/custom-settings/v1/schema-list", 500, "error_internal.json", "no response registered")
}

func (m *MacOSConfigProfilesMock) RegisterGetByPayloadUUIDNoResponseErrorMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.RegisterError("GET", path, 500, "error_internal.json", "no response registered")
}

func (m *MacOSConfigProfilesMock) RegisterCreateNoResponseErrorMock() {
	m.RegisterError("POST", "/api/config-profiles/macos", 500, "error_internal.json", "no response registered")
}

func (m *MacOSConfigProfilesMock) RegisterUpdateByPayloadUUIDNoResponseErrorMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.RegisterError("PUT", path, 500, "error_internal.json", "no response registered")
}

func (m *MacOSConfigProfilesMock) RegisterDeleteByPayloadUUIDNoResponseErrorMock(id string) {
	path := "/api/config-profiles/macos/" + id
	m.RegisterError("DELETE", path, 500, "error_internal.json", "no response registered")
}
