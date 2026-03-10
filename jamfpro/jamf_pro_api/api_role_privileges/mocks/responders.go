package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type APIRolePrivilegesMock struct {
	*mocks.GenericMock
}

func NewAPIRolePrivilegesMock() *APIRolePrivilegesMock {
	return &APIRolePrivilegesMock{
		GenericMock: mocks.NewJSONMock("APIRolePrivilegesMock"),
	}
}

func (m *APIRolePrivilegesMock) RegisterMocks() {
	m.Register("GET", "/api/v1/api-role-privileges", 200, "validate_list.json")
	m.Register("GET", "/api/v1/api-role-privileges/search?name=Read&limit=10", 200, "validate_list.json")
}

func (m *APIRolePrivilegesMock) RegisterSearchDefaultLimitMock() {
	m.Register("GET", "/api/v1/api-role-privileges/search?name=Read&limit=15", 200, "validate_list.json")
}

func (m *APIRolePrivilegesMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v1/api-role-privileges", 500, "error_internal.json", "")
}
