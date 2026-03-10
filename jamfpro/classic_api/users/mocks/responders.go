package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type UsersMock struct {
	*mocks.GenericMock
}

func NewUsersMock() *UsersMock {
	return &UsersMock{
		GenericMock: mocks.NewXMLMock("UsersMock"),
	}
}

func (m *UsersMock) RegisterMocks() {
	m.RegisterListUsersMock()
	m.RegisterGetUserByIDMock()
	m.RegisterGetUserByNameMock()
	m.RegisterGetUserByEmailMock()
	m.RegisterCreateUserMock()
	m.RegisterUpdateUserByIDMock()
	m.RegisterUpdateUserByNameMock()
	m.RegisterUpdateUserByEmailMock()
	m.RegisterDeleteUserByIDMock()
	m.RegisterDeleteUserByNameMock()
	m.RegisterDeleteUserByEmailMock()
}

func (m *UsersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *UsersMock) RegisterListUsersMock() {
	m.Register("GET", "/JSSResource/users", 200, "validate_list_users.xml")
}

func (m *UsersMock) RegisterGetUserByIDMock() {
	m.Register("GET", "/JSSResource/users/id/1", 200, "validate_get_user.xml")
}

func (m *UsersMock) RegisterGetUserByNameMock() {
	m.Register("GET", "/JSSResource/users/name/admin", 200, "validate_get_user.xml")
}

func (m *UsersMock) RegisterGetUserByEmailMock() {
	m.Register("GET", "/JSSResource/users/email/admin@example.com", 200, "validate_list_users.xml")
	m.Register("GET", "/JSSResource/users/email/admin%40example.com", 200, "validate_list_users.xml")
}

func (m *UsersMock) RegisterCreateUserMock() {
	m.Register("POST", "/JSSResource/users/id/0", 201, "validate_create_user.xml")
}

func (m *UsersMock) RegisterUpdateUserByIDMock() {
	m.Register("PUT", "/JSSResource/users/id/1", 200, "validate_update_user.xml")
}

func (m *UsersMock) RegisterUpdateUserByNameMock() {
	m.Register("PUT", "/JSSResource/users/name/admin", 200, "validate_update_user.xml")
}

func (m *UsersMock) RegisterUpdateUserByEmailMock() {
	m.Register("PUT", "/JSSResource/users/email/admin@example.com", 200, "validate_update_user.xml")
	m.Register("PUT", "/JSSResource/users/email/admin%40example.com", 200, "validate_update_user.xml")
}

func (m *UsersMock) RegisterDeleteUserByIDMock() {
	m.Register("DELETE", "/JSSResource/users/id/1", 200, "")
}

func (m *UsersMock) RegisterDeleteUserByNameMock() {
	m.Register("DELETE", "/JSSResource/users/name/admin", 200, "")
}

func (m *UsersMock) RegisterDeleteUserByEmailMock() {
	m.Register("DELETE", "/JSSResource/users/email/admin@example.com", 200, "")
	m.Register("DELETE", "/JSSResource/users/email/admin%40example.com", 200, "")
}

func (m *UsersMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/users/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *UsersMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/users/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A user with that name already exists")
}

