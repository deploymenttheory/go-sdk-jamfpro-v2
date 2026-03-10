package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type OAuth2SessionTokensMock struct {
	*mocks.GenericMock
}

func NewOAuth2SessionTokensMock() *OAuth2SessionTokensMock {
	return &OAuth2SessionTokensMock{
		GenericMock: mocks.NewJSONMock("OAuth2SessionTokensMock"),
	}
}

func (m *OAuth2SessionTokensMock) RegisterMocks() {
	m.Register("GET", "/api/v1/oauth2/session-tokens", 200, "validate_get.json")
}

func (m *OAuth2SessionTokensMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v1/oauth2/session-tokens", 401, "", "Jamf Pro API error (401): unauthorized")
}
