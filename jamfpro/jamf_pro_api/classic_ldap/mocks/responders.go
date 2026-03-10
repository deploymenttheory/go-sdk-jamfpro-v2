package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ClassicLdapMock struct {
	*mocks.GenericMock
}

func NewClassicLdapMock() *ClassicLdapMock {
	return &ClassicLdapMock{
		GenericMock: mocks.NewJSONMock("ClassicLdapMock"),
	}
}

func (m *ClassicLdapMock) RegisterGetMappingsByIDMock(id string) {
	m.Register("GET", "/api/v1/classic-ldap/"+id, 200, "validate_get.json")
}

func (m *ClassicLdapMock) RegisterGetMappingsByIDErrorMock(id string) {
	m.RegisterError("GET", "/api/v1/classic-ldap/"+id, 500, "error_internal.json", "")
}
