package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ServiceDiscoveryEnrollmentMock struct {
	*mocks.GenericMock
}

func NewServiceDiscoveryEnrollmentMock() *ServiceDiscoveryEnrollmentMock {
	return &ServiceDiscoveryEnrollmentMock{
		GenericMock: mocks.NewJSONMock("ServiceDiscoveryEnrollmentMock"),
	}
}

func (m *ServiceDiscoveryEnrollmentMock) RegisterGetWellKnownSettingsMock() {
	m.Register("GET", "/api/v1/service-discovery-enrollment/well-known-settings", 200, "validate_get.json")
}

func (m *ServiceDiscoveryEnrollmentMock) RegisterUpdateWellKnownSettingsMock() {
	m.Register("PUT", "/api/v1/service-discovery-enrollment/well-known-settings", 204, "validate_update_204.json")
}

func (m *ServiceDiscoveryEnrollmentMock) RegisterUpdateWellKnownSettingsNon204Mock() {
	m.Register("PUT", "/api/v1/service-discovery-enrollment/well-known-settings", 200, "")
}
