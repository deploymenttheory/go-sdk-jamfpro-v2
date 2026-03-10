package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AccessManagementSettingsMock struct {
	*mocks.GenericMock
}

func NewAccessManagementSettingsMock() *AccessManagementSettingsMock {
	return &AccessManagementSettingsMock{
		GenericMock: mocks.NewJSONMock("AccessManagementSettingsMock"),
	}
}

func (m *AccessManagementSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v4/enrollment/access-management", 200, "validate_get.json")
}

func (m *AccessManagementSettingsMock) RegisterPostMock() {
	m.Register("POST", "/api/v4/enrollment/access-management", 200, "validate_get.json")
}

func (m *AccessManagementSettingsMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v4/enrollment/access-management", 500, "error_internal.json", "")
}

func (m *AccessManagementSettingsMock) RegisterPostErrorMock() {
	m.RegisterError("POST", "/api/v4/enrollment/access-management", 500, "error_internal.json", "")
}
