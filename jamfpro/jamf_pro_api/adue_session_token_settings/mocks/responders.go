package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AdueSessionTokenSettingsMock struct {
	*mocks.GenericMock
}

func NewAdueSessionTokenSettingsMock() *AdueSessionTokenSettingsMock {
	return &AdueSessionTokenSettingsMock{
		GenericMock: mocks.NewJSONMock("AdueSessionTokenSettingsMock"),
	}
}

func (m *AdueSessionTokenSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/adue-session-token-settings", 200, "validate_get.json")
}

func (m *AdueSessionTokenSettingsMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/adue-session-token-settings", 200, "validate_get.json")
}
