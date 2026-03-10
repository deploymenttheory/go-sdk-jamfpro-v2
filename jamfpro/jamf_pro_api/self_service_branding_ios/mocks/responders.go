package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SelfServiceBrandingMobileMock struct {
	*mocks.GenericMock
}

func NewSelfServiceBrandingMobileMock() *SelfServiceBrandingMobileMock {
	return &SelfServiceBrandingMobileMock{
		GenericMock: mocks.NewJSONMock("SelfServiceBrandingMobileMock"),
	}
}

func (m *SelfServiceBrandingMobileMock) RegisterListMock() {
	m.Register("GET", "/api/v1/self-service/branding/ios", 200, "validate_list.json")
}

func (m *SelfServiceBrandingMobileMock) RegisterGetByIDMock() {
	m.Register("GET", "/api/v1/self-service/branding/ios/1", 200, "validate_get.json")
}

func (m *SelfServiceBrandingMobileMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/self-service/branding/ios", 201, "validate_create.json")
}

func (m *SelfServiceBrandingMobileMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/self-service/branding/ios/1", 200, "validate_update.json")
}

func (m *SelfServiceBrandingMobileMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v1/self-service/branding/ios/1", 204, "")
}

func (m *SelfServiceBrandingMobileMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/self-service/branding/ios/999", 404, "error_not_found.json", "")
}

func (m *SelfServiceBrandingMobileMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v1/self-service/branding/ios", 409, "error_conflict.json", "")
}
