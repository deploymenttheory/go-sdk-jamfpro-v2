package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type UserExtensionAttributesMock struct {
	*mocks.GenericMock
}

func NewUserExtensionAttributesMock() *UserExtensionAttributesMock {
	return &UserExtensionAttributesMock{
		GenericMock: mocks.NewXMLMock("UserExtensionAttributesMock"),
	}
}

func (m *UserExtensionAttributesMock) RegisterMocks() {
	m.RegisterListUserExtensionAttributesMock()
	m.RegisterGetUserExtensionAttributeByIDMock()
	m.RegisterGetUserExtensionAttributeByNameMock()
	m.RegisterCreateUserExtensionAttributeMock()
	m.RegisterUpdateUserExtensionAttributeByIDMock()
	m.RegisterUpdateUserExtensionAttributeByNameMock()
	m.RegisterDeleteUserExtensionAttributeByIDMock()
	m.RegisterDeleteUserExtensionAttributeByNameMock()
}

func (m *UserExtensionAttributesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *UserExtensionAttributesMock) RegisterListUserExtensionAttributesMock() {
	m.Register("GET", "/JSSResource/userextensionattributes", 200, "validate_list_user_extension_attributes.xml")
}

func (m *UserExtensionAttributesMock) RegisterGetUserExtensionAttributeByIDMock() {
	m.Register("GET", "/JSSResource/userextensionattributes/id/1", 200, "validate_get_user_extension_attribute.xml")
}

func (m *UserExtensionAttributesMock) RegisterGetUserExtensionAttributeByNameMock() {
	m.Register("GET", "/JSSResource/userextensionattributes/name/Department", 200, "validate_get_user_extension_attribute.xml")
}

func (m *UserExtensionAttributesMock) RegisterCreateUserExtensionAttributeMock() {
	m.Register("POST", "/JSSResource/userextensionattributes/id/0", 201, "validate_create_user_extension_attribute.xml")
}

func (m *UserExtensionAttributesMock) RegisterUpdateUserExtensionAttributeByIDMock() {
	m.Register("PUT", "/JSSResource/userextensionattributes/id/1", 200, "validate_update_user_extension_attribute.xml")
}

func (m *UserExtensionAttributesMock) RegisterUpdateUserExtensionAttributeByNameMock() {
	m.Register("PUT", "/JSSResource/userextensionattributes/name/Department", 200, "validate_update_user_extension_attribute.xml")
}

func (m *UserExtensionAttributesMock) RegisterDeleteUserExtensionAttributeByIDMock() {
	m.Register("DELETE", "/JSSResource/userextensionattributes/id/1", 200, "")
}

func (m *UserExtensionAttributesMock) RegisterDeleteUserExtensionAttributeByNameMock() {
	m.Register("DELETE", "/JSSResource/userextensionattributes/name/Department", 200, "")
}

func (m *UserExtensionAttributesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/userextensionattributes/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *UserExtensionAttributesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/userextensionattributes/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A user extension attribute with that name already exists")
}

