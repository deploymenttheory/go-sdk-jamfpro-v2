package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type UserSessionsMock struct {
	*mocks.GenericMock
}

func NewUserSessionsMock() *UserSessionsMock {
	return &UserSessionsMock{
		GenericMock: mocks.NewJSONMock("UserSessionsMock"),
	}
}

func (m *UserSessionsMock) RegisterMocks() {
	m.RegisterGetActiveSessionsMock()
	m.RegisterGetCountMock()
}

func (m *UserSessionsMock) RegisterGetActiveSessionsMock() {
	m.Register("GET", "/api/v1/user-sessions/active", 200, "validate_active.json")
}

func (m *UserSessionsMock) RegisterGetCountMock() {
	m.Register("GET", "/api/v1/user-sessions/count", 200, "validate_count.json")
}
