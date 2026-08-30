package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type EnvironmentTypeMock struct {
	*mocks.GenericMock
}

func NewEnvironmentTypeMock() *EnvironmentTypeMock {
	return &EnvironmentTypeMock{
		GenericMock: mocks.NewJSONMock("EnvironmentTypeMock"),
	}
}

func (m *EnvironmentTypeMock) RegisterMocks() {
	m.Register("GET", "/api/v2/environment-type", 200, "validate_get.json")
}

func (m *EnvironmentTypeMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v2/environment-type", 500, "error_internal.json", "mock client error")
}

func (m *EnvironmentTypeMock) RegisterGetNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v2/environment-type", 500, "error_internal.json", "no response for")
}
