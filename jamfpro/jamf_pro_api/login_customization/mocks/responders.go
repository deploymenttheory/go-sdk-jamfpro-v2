package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type LoginCustomizationMock struct {
	*mocks.GenericMock
}

func NewLoginCustomizationMock() *LoginCustomizationMock {
	return &LoginCustomizationMock{
		GenericMock: mocks.NewJSONMock("LoginCustomizationMock"),
	}
}

func (m *LoginCustomizationMock) RegisterGetLoginCustomizationMock() {
	m.Register("GET", "/api/v1/login-customization", 200, "validate_get.json")
}

func (m *LoginCustomizationMock) RegisterUpdateLoginCustomizationMock() {
	m.Register("PUT", "/api/v1/login-customization", 200, "validate_update.json")
}

func (m *LoginCustomizationMock) RegisterGetLoginCustomizationErrorMock() {
	m.RegisterError("GET", "/api/v1/login-customization", 404, "error_not_found.json", "")
}
