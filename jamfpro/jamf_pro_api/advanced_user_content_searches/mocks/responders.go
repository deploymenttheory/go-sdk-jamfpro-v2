package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AdvancedUserContentSearchesMock struct {
	*mocks.GenericMock
}

func NewAdvancedUserContentSearchesMock() *AdvancedUserContentSearchesMock {
	return &AdvancedUserContentSearchesMock{
		GenericMock: mocks.NewJSONMock("AdvancedUserContentSearchesMock"),
	}
}

func (m *AdvancedUserContentSearchesMock) RegisterMocks() {
	m.Register("GET", "/api/v1/advanced-user-content-searches", 200, "validate_list.json")
	m.Register("GET", "/api/v1/advanced-user-content-searches/1", 200, "validate_get.json")
	m.Register("POST", "/api/v1/advanced-user-content-searches", 201, "validate_create.json")
	m.Register("PUT", "/api/v1/advanced-user-content-searches/1", 200, "validate_get.json")
	m.Register("DELETE", "/api/v1/advanced-user-content-searches/1", 204, "")
}

func (m *AdvancedUserContentSearchesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/advanced-user-content-searches/999", 404, "error_not_found.json", "")
}

func (m *AdvancedUserContentSearchesMock) RegisterDeleteErrorMock() {
	m.RegisterError("DELETE", "/api/v1/advanced-user-content-searches/999", 500, "error_internal.json", "")
}
