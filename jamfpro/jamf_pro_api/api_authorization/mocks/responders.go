package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ApiAuthorizationMock struct {
	*mocks.GenericMock
}

func NewApiAuthorizationMock() *ApiAuthorizationMock {
	return &ApiAuthorizationMock{
		GenericMock: mocks.NewJSONMock("ApiAuthorizationMock"),
	}
}

func (m *ApiAuthorizationMock) RegisterGetV1Mock() {
	m.Register("GET", "/api/v1/auth", 200, "validate_get.json")
}
