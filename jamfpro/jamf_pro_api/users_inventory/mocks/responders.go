package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type UsersInventoryMock struct {
	*mocks.GenericMock
}

func NewUsersInventoryMock() *UsersInventoryMock {
	return &UsersInventoryMock{
		GenericMock: mocks.NewJSONMock("UsersInventoryMock"),
	}
}

func (m *UsersInventoryMock) RegisterListUsersMock() {
	m.Register("GET", "/api/v1/users", 200, "validate_list_users.json")
}

func (m *UsersInventoryMock) RegisterGetUserMock() {
	m.Register("GET", "/api/v1/users/1", 200, "validate_get_user.json")
}

func (m *UsersInventoryMock) RegisterCreateUserMock() {
	m.Register("POST", "/api/v1/users", 201, "validate_create_user.json")
}

func (m *UsersInventoryMock) RegisterUpdateUserMock() {
	m.Register("PUT", "/api/v1/users/1", 204, "")
}

func (m *UsersInventoryMock) RegisterDeleteUserMock() {
	m.Register("DELETE", "/api/v1/users/1", 204, "")
}

func (m *UsersInventoryMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/users/999", 404, "error_not_found.json", "")
}

func (m *UsersInventoryMock) RegisterListUsersErrorMock() {
	m.RegisterError("GET", "/api/v1/users", 500, "error_internal.json", "")
}

func (m *UsersInventoryMock) RegisterCreateUserConflictMock() {
	m.RegisterError("POST", "/api/v1/users", 409, "error_conflict.json", "")
}
