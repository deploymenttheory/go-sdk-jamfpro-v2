package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SelfServicePlusSettingsMock struct {
	*mocks.GenericMock
}

func NewSelfServicePlusSettingsMock() *SelfServicePlusSettingsMock {
	return &SelfServicePlusSettingsMock{
		GenericMock: mocks.NewJSONMock("SelfServicePlusSettingsMock"),
	}
}

func (m *SelfServicePlusSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/self-service-plus/settings", 200, "validate_get.json")
}

func (m *SelfServicePlusSettingsMock) RegisterFeatureToggleMock() {
	m.Register("GET", "/api/v1/self-service-plus/feature-toggle/enabled", 200, "validate_feature_toggle.json")
}

func (m *SelfServicePlusSettingsMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/self-service-plus/settings", 204, "")
}

func (m *SelfServicePlusSettingsMock) RegisterUpdateNon204Mock() {
	m.Register("PUT", "/api/v1/self-service-plus/settings", 200, "")
}
