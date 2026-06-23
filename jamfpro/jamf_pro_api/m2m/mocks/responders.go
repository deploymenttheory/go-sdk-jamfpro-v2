package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type M2MMock struct {
	*mocks.GenericMock
}

func NewM2MMock() *M2MMock {
	return &M2MMock{
		GenericMock: mocks.NewJSONMock("M2MMock"),
	}
}

func (m *M2MMock) RegisterMocks() {
	m.RegisterGetTenantIdMock()
}

func (m *M2MMock) RegisterGetTenantIdMock() {
	m.Register("GET", "/api/v1/m2m/tenant-id", 200, "validate_get_tenant_id.json")
}
