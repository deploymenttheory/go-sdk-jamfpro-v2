package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type StartupStatusMock struct {
	*mocks.GenericMock
}

func NewStartupStatusMock() *StartupStatusMock {
	return &StartupStatusMock{
		GenericMock: mocks.NewJSONMock("StartupStatusMock"),
	}
}

func (m *StartupStatusMock) RegisterGetStartupStatusMock() {
	m.Register("GET", "/api/startup-status", 200, "validate_get.json")
}

func (m *StartupStatusMock) RegisterGetStartupStatusErrorMock() {
	m.RegisterError("GET", "/api/startup-status", 500, "error_internal.json", "no response")
}
