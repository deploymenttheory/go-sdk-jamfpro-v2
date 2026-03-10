package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SSOFailoverMock struct {
	*mocks.GenericMock
}

func NewSSOFailoverMock() *SSOFailoverMock {
	return &SSOFailoverMock{
		GenericMock: mocks.NewJSONMock("SSOFailoverMock"),
	}
}

func (m *SSOFailoverMock) RegisterMocks() {
	m.Register("GET", "/api/v1/sso/failover", 200, "validate_get.json")
	m.Register("POST", "/api/v1/sso/failover/generate", 200, "validate_get.json")
}

func (m *SSOFailoverMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v1/sso/failover", 500, "", "Jamf Pro API error (500): server error")
}
