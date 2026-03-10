package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type PolicyPropertiesMock struct {
	*mocks.GenericMock
}

func NewPolicyPropertiesMock() *PolicyPropertiesMock {
	return &PolicyPropertiesMock{
		GenericMock: mocks.NewJSONMock("PolicyPropertiesMock"),
	}
}

func (m *PolicyPropertiesMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/policy-properties", 200, "validate_get.json")
}

func (m *PolicyPropertiesMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/policy-properties", 200, "validate_get.json")
}

func (m *PolicyPropertiesMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v1/policy-properties", 500, "error_internal.json", "no response registered")
}

func (m *PolicyPropertiesMock) RegisterUpdateErrorMock() {
	m.RegisterError("PUT", "/api/v1/policy-properties", 500, "error_internal.json", "no response registered")
}
