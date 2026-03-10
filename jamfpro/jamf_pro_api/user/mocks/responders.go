package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type UserMock struct {
	*mocks.GenericMock
}

func NewUserMock() *UserMock {
	return &UserMock{
		GenericMock: mocks.NewJSONMock("UserMock"),
	}
}

func (m *UserMock) RegisterGetMock() {
	m.Register("GET", "/api/user", 200, "validate_get.json")
}

func (m *UserMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/user", 404, "error_not_found.json", "")
}

func (m *UserMock) RegisterInvalidJSONMock() {
	m.RegisterRawBody("GET", "/api/user", 200, []byte(`{invalid json`))
}

func (m *UserMock) RegisterChangePasswordMock() {
	m.Register("POST", "/api/v1/user/change-password", 204, "")
}

func (m *UserMock) RegisterChangePasswordErrorMock() {
	m.RegisterError("POST", "/api/v1/user/change-password", 400, "error_not_found.json", "")
}

func (m *UserMock) RegisterUpdateSessionMock() {
	m.Register("POST", "/api/user/updateSession", 204, "")
}

func (m *UserMock) RegisterUpdateSessionErrorMock() {
	m.RegisterError("POST", "/api/user/updateSession", 500, "error_not_found.json", "")
}

func (m *UserMock) RegisterGetNoResponseErrorMock() {
	m.RegisterError("GET", "/api/user", 500, "error_internal.json", "no response registered")
}
