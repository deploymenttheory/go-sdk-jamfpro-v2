package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type JamfProVersionMock struct {
	*mocks.GenericMock
}

func NewJamfProVersionMock() *JamfProVersionMock {
	return &JamfProVersionMock{
		GenericMock: mocks.NewJSONMock("JamfProVersionMock"),
	}
}

func (m *JamfProVersionMock) RegisterMocks() {
	m.Register("GET", "/api/v1/jamf-pro-version", 200, "validate_get.json")
}

func (m *JamfProVersionMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v1/jamf-pro-version", 500, "error_internal.json", "mock client error")
}

func (m *JamfProVersionMock) RegisterGetNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/jamf-pro-version", 500, "error_internal.json", "no response for")
}
