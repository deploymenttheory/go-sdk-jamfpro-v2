package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ConditionalAccessMock struct {
	*mocks.GenericMock
}

func NewConditionalAccessMock() *ConditionalAccessMock {
	return &ConditionalAccessMock{
		GenericMock: mocks.NewJSONMock("ConditionalAccessMock"),
	}
}

func (m *ConditionalAccessMock) RegisterGetDeviceComplianceFeatureToggleMock() {
	m.Register("GET", "/api/v1/conditional-access/device-compliance/feature-toggle", 200, "validate_get.json")
}

func (m *ConditionalAccessMock) RegisterGetDeviceComplianceInformationComputerMock() {
	m.Register("GET", "/api/v1/conditional-access/device-compliance-information/computer/1", 200, "validate_compliance_computer.json")
}

func (m *ConditionalAccessMock) RegisterGetDeviceComplianceInformationMobileMock() {
	m.Register("GET", "/api/v1/conditional-access/device-compliance-information/mobile/1", 200, "validate_compliance_mobile.json")
}

