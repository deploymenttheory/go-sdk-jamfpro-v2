package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CSAMock struct {
	*mocks.GenericMock
}

func NewCSAMock() *CSAMock {
	return &CSAMock{
		GenericMock: mocks.NewJSONMock("CSAMock"),
	}
}

func (m *CSAMock) RegisterGetTokenExchangeDetailsMock() {
	m.Register("GET", "/api/v1/csa/token", 200, "validate_get_token.json")
}

func (m *CSAMock) RegisterGetTenantIDMock() {
	m.Register("GET", "/api/v1/csa/token/tenant-id", 200, "validate_get_tenant.json")
}

func (m *CSAMock) RegisterDeleteTokenExchangeMock() {
	m.Register("DELETE", "/api/v1/csa/token", 204, "")
}

