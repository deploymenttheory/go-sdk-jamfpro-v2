package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type JamfProInformationMock struct {
	*mocks.GenericMock
}

func NewJamfProInformationMock() *JamfProInformationMock {
	return &JamfProInformationMock{
		GenericMock: mocks.NewJSONMock("JamfProInformationMock"),
	}
}

func (m *JamfProInformationMock) RegisterMocks() {
	m.Register("GET", "/api/v2/jamf-pro-information", 200, "validate_get.json")
}

func (m *JamfProInformationMock) RegisterGetV2ErrorMock() {
	m.RegisterError("GET", "/api/v2/jamf-pro-information", 500, "error_internal.json", "mock client error")
}

func (m *JamfProInformationMock) RegisterGetV2NoResponseErrorMock() {
	m.RegisterError("GET", "/api/v2/jamf-pro-information", 500, "error_internal.json", "no response for")
}
