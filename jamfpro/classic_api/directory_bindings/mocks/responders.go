package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DirectoryBindingsMock struct {
	*mocks.GenericMock
}

func NewDirectoryBindingsMock() *DirectoryBindingsMock {
	return &DirectoryBindingsMock{
		GenericMock: mocks.NewXMLMock("DirectoryBindingsMock"),
	}
}

func (m *DirectoryBindingsMock) RegisterMocks() {
	m.RegisterListDirectoryBindingsMock()
	m.RegisterGetDirectoryBindingByIDMock()
	m.RegisterGetDirectoryBindingByNameMock()
	m.RegisterCreateDirectoryBindingMock()
	m.RegisterUpdateDirectoryBindingByIDMock()
	m.RegisterUpdateDirectoryBindingByNameMock()
	m.RegisterDeleteDirectoryBindingByIDMock()
	m.RegisterDeleteDirectoryBindingByNameMock()
}

func (m *DirectoryBindingsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *DirectoryBindingsMock) RegisterListDirectoryBindingsMock() {
	m.Register("GET", "/JSSResource/directorybindings", 200, "validate_list_directory_bindings.xml")
}

func (m *DirectoryBindingsMock) RegisterGetDirectoryBindingByIDMock() {
	m.Register("GET", "/JSSResource/directorybindings/id/1", 200, "validate_get_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterGetDirectoryBindingByNameMock() {
	m.Register("GET", "/JSSResource/directorybindings/name/AD%20Binding", 200, "validate_get_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterCreateDirectoryBindingMock() {
	m.Register("POST", "/JSSResource/directorybindings/id/0", 201, "validate_create_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterUpdateDirectoryBindingByIDMock() {
	m.Register("PUT", "/JSSResource/directorybindings/id/1", 200, "validate_update_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterUpdateDirectoryBindingByNameMock() {
	m.Register("PUT", "/JSSResource/directorybindings/name/AD%20Binding", 200, "validate_update_directory_binding.xml")
}

func (m *DirectoryBindingsMock) RegisterDeleteDirectoryBindingByIDMock() {
	m.Register("DELETE", "/JSSResource/directorybindings/id/1", 200, "")
}

func (m *DirectoryBindingsMock) RegisterDeleteDirectoryBindingByNameMock() {
	m.Register("DELETE", "/JSSResource/directorybindings/name/AD%20Binding", 200, "")
}

func (m *DirectoryBindingsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/directorybindings/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *DirectoryBindingsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/directorybindings/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A directory binding with that name already exists")
}

