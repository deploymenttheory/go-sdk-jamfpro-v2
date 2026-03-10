package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type JamfAccountPreferencesMock struct {
	*mocks.GenericMock
}

func NewJamfAccountPreferencesMock() *JamfAccountPreferencesMock {
	return &JamfAccountPreferencesMock{
		GenericMock: mocks.NewJSONMock("JamfAccountPreferencesMock"),
	}
}

func (m *JamfAccountPreferencesMock) RegisterGetV3Mock() {
	m.Register("GET", "/api/v3/account-preferences", 200, "validate_get.json")
}

func (m *JamfAccountPreferencesMock) RegisterUpdateV3Mock() {
	m.Register("PATCH", "/api/v3/account-preferences", 200, "validate_update.json")
}

func (m *JamfAccountPreferencesMock) RegisterUpdateV3_204NoContentMock() {
	m.Register("PATCH", "/api/v3/account-preferences", 204, "")
}

func (m *JamfAccountPreferencesMock) RegisterGetV3ErrorMock() {
	m.RegisterError("GET", "/api/v3/account-preferences", 404, "error_not_found.json", "")
}

func (m *JamfAccountPreferencesMock) RegisterUpdateV3ErrorMock() {
	m.RegisterError("PATCH", "/api/v3/account-preferences", 500, "error_not_found.json", "")
}

func (m *JamfAccountPreferencesMock) RegisterInvalidJSONMock() {
	m.Register("GET", "/api/v3/account-preferences", 200, "validate_get_invalid.json")
}
