package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type EnrollmentSettingsMock struct {
	*mocks.GenericMock
}

func NewEnrollmentSettingsMock() *EnrollmentSettingsMock {
	return &EnrollmentSettingsMock{
		GenericMock: mocks.NewJSONMock("EnrollmentSettingsMock"),
	}
}

func (m *EnrollmentSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v4/enrollment", 200, "validate_get.json")
}

func (m *EnrollmentSettingsMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v4/enrollment", 500, "error_not_found.json", "")
}

