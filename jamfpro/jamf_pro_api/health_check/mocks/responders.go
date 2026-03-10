package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type HealthCheckMock struct {
	*mocks.GenericMock
}

func NewHealthCheckMock() *HealthCheckMock {
	return &HealthCheckMock{
		GenericMock: mocks.NewJSONMock("HealthCheckMock"),
	}
}

func (m *HealthCheckMock) RegisterMocks() {
	m.Register("GET", "/api/v1/health-check", 200, "validate_get.json")
	m.Register("GET", "/api/v1/health-status", 200, "validate_health_status.json")
}

func (m *HealthCheckMock) RegisterErrorMock() {
	m.RegisterError("GET", "/api/v1/health-check", 503, "", "Jamf Pro API error (503): service unavailable")
}
