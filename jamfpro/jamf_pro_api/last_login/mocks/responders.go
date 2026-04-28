package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type LastLoginMock struct {
	*mocks.GenericMock
}

func NewLastLoginMock() *LastLoginMock {
	return &LastLoginMock{
		GenericMock: mocks.NewJSONMock("LastLoginMock"),
	}
}

func (m *LastLoginMock) RegisterMocks() {
	m.RegisterGetLastLoginMock()
}

func (m *LastLoginMock) RegisterGetLastLoginMock() {
	m.Register("GET", "/api/v1/last-login", 200, "validate_get.json")
}
