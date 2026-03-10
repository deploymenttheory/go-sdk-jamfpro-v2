package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type OnboardingMock struct {
	*mocks.GenericMock
}

func NewOnboardingMock() *OnboardingMock {
	return &OnboardingMock{
		GenericMock: mocks.NewJSONMock("OnboardingMock"),
	}
}

func (m *OnboardingMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/onboarding", 200, "validate_get.json")
}

func (m *OnboardingMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/onboarding", 200, "validate_get.json")
}

func (m *OnboardingMock) RegisterGetEligibleAppsMock() {
	m.Register("GET", "/api/v1/onboarding/eligible-apps", 200, "validate_eligible_apps.json")
}

func (m *OnboardingMock) RegisterGetEligibleConfigurationProfilesMock() {
	m.Register("GET", "/api/v1/onboarding/eligible-configuration-profiles", 200, "validate_eligible_apps.json")
}

func (m *OnboardingMock) RegisterGetEligiblePoliciesMock() {
	m.Register("GET", "/api/v1/onboarding/eligible-policies", 200, "validate_eligible_apps.json")
}

func (m *OnboardingMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v1/onboarding/history", 200, "validate_history.json")
}

func (m *OnboardingMock) RegisterAddHistoryNotesMock() {
	m.Register("POST", "/api/v1/onboarding/history", 201, "validate_add_history_notes.json")
}

func (m *OnboardingMock) RegisterExportHistoryMock() {
	m.Register("GET", "/api/v1/onboarding/history/export", 200, "validate_export_history.csv")
}
