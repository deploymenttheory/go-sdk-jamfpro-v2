package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SelfServiceBrandingUploadMock struct {
	*mocks.GenericMock
}

func NewSelfServiceBrandingUploadMock() *SelfServiceBrandingUploadMock {
	return &SelfServiceBrandingUploadMock{
		GenericMock: mocks.NewJSONMock("SelfServiceBrandingUploadMock"),
	}
}

func (m *SelfServiceBrandingUploadMock) RegisterUploadMock() {
	m.Register("POST", "/api/self-service/branding/images", 200, "validate_upload.json")
}
