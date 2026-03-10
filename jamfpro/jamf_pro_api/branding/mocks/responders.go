package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type BrandingMock struct {
	*mocks.GenericMock
}

func NewBrandingMock() *BrandingMock {
	return &BrandingMock{
		GenericMock: mocks.NewJSONMock("BrandingMock"),
	}
}

func (m *BrandingMock) RegisterMocks() {
	m.Register("GET", "/api/v1/branding-images/download/test-id", 200, "validate_download.bin")
}
