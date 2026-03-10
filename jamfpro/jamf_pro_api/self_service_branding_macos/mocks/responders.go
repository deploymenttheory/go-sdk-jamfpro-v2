package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SelfServiceBrandingMacOSMock struct {
	*mocks.GenericMock
}

func NewSelfServiceBrandingMacOSMock() *SelfServiceBrandingMacOSMock {
	return &SelfServiceBrandingMacOSMock{
		GenericMock: mocks.NewJSONMock("SelfServiceBrandingMacOSMock"),
	}
}

func (m *SelfServiceBrandingMacOSMock) RegisterListMock() {
	m.Register("GET", "/api/v1/self-service/branding/macos", 200, "validate_list.json")
}

func (m *SelfServiceBrandingMacOSMock) RegisterGetByIDMock() {
	m.Register("GET", "/api/v1/self-service/branding/macos/1", 200, "validate_get_by_id.json")
}

func (m *SelfServiceBrandingMacOSMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/self-service/branding/macos", 201, "validate_create.json")
}

func (m *SelfServiceBrandingMacOSMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/self-service/branding/macos/1", 200, "validate_update.json")
}

func (m *SelfServiceBrandingMacOSMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v1/self-service/branding/macos/1", 204, "")
}

func (m *SelfServiceBrandingMacOSMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/self-service/branding/macos/999", 404, "error_not_found.json", "")
}

func (m *SelfServiceBrandingMacOSMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v1/self-service/branding/macos", 409, "error_conflict.json", "")
}
