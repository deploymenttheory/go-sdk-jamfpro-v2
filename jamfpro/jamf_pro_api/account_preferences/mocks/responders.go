package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AccountPreferencesMock struct {
	*mocks.GenericMock
}

func NewAccountPreferencesMock() *AccountPreferencesMock {
	return &AccountPreferencesMock{
		GenericMock: mocks.NewJSONMock("AccountPreferencesMock"),
	}
}

func (m *AccountPreferencesMock) RegisterGetAccountPreferencesMock() {
	m.Register("GET", "/api/v3/account-preferences", 200, "validate_get.json")
}

func (m *AccountPreferencesMock) RegisterUpdateAccountPreferencesMock() {
	m.Register("PATCH", "/api/v3/account-preferences", 200, "validate_get.json")
}

func (m *AccountPreferencesMock) RegisterGetAccountPreferencesErrorMock() {
	m.RegisterError("GET", "/api/v3/account-preferences", 500, "error_internal.json", "")
}

func (m *AccountPreferencesMock) RegisterUpdateAccountPreferencesErrorMock() {
	m.RegisterError("PATCH", "/api/v3/account-preferences", 500, "error_internal.json", "")
}
