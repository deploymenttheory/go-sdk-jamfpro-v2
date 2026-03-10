package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type APIRolesMock struct {
	*mocks.GenericMock
}

func NewAPIRolesMock() *APIRolesMock {
	return &APIRolesMock{
		GenericMock: mocks.NewJSONMock("APIRolesMock"),
	}
}

func (m *APIRolesMock) RegisterMocks() {
	m.Register("GET", "/api/v1/api-roles", 200, "validate_list.json")
	m.Register("GET", "/api/v1/api-roles/1", 200, "validate_get.json")
	m.Register("POST", "/api/v1/api-roles", 200, "validate_create.json")
	m.Register("PUT", "/api/v1/api-roles/1", 200, "validate_get.json")
	m.Register("DELETE", "/api/v1/api-roles/1", 204, "")
}

func (m *APIRolesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/api-roles/999", 404, "error_not_found.json", "")
}
