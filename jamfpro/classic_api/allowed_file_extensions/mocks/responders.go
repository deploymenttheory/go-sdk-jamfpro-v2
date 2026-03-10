package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AllowedFileExtensionsMock struct {
	*mocks.GenericMock
}

func NewAllowedFileExtensionsMock() *AllowedFileExtensionsMock {
	return &AllowedFileExtensionsMock{
		GenericMock: mocks.NewXMLMock("AllowedFileExtensionsMock"),
	}
}

func (m *AllowedFileExtensionsMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByExtensionMock()
	m.RegisterCreateMock()
	m.RegisterDeleteByIDMock()
}

func (m *AllowedFileExtensionsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *AllowedFileExtensionsMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/allowedfileextensions", 200, "validate_list_allowed_file_extensions.xml")
}

func (m *AllowedFileExtensionsMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/allowedfileextensions/id/1", 200, "validate_get_allowed_file_extension.xml")
}

func (m *AllowedFileExtensionsMock) RegisterGetByExtensionMock() {
	m.Register("GET", "/JSSResource/allowedfileextensions/extension/dmg", 200, "validate_get_allowed_file_extension.xml")
}

func (m *AllowedFileExtensionsMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/allowedfileextensions/id/0", 201, "validate_create_allowed_file_extension.xml")
}

func (m *AllowedFileExtensionsMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/allowedfileextensions/id/1", 200, "")
}

func (m *AllowedFileExtensionsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/allowedfileextensions/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *AllowedFileExtensionsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/allowedfileextensions/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): An allowed file extension with that name already exists")
}

