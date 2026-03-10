package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ApiIntegrationsMock struct {
	*mocks.GenericMock
}

func NewApiIntegrationsMock() *ApiIntegrationsMock {
	return &ApiIntegrationsMock{
		GenericMock: mocks.NewJSONMock("ApiIntegrationsMock"),
	}
}

func (m *ApiIntegrationsMock) RegisterListMock() {
	m.Register("GET", "/api/v1/api-integrations", 200, "validate_list.json")
}

func (m *ApiIntegrationsMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v1/api-integrations/"+id, 200, "validate_get.json")
}

func (m *ApiIntegrationsMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/api-integrations", 200, "validate_get.json")
}

func (m *ApiIntegrationsMock) RegisterUpdateByIDMock(id string) {
	m.Register("PUT", "/api/v1/api-integrations/"+id, 200, "validate_get.json")
}

func (m *ApiIntegrationsMock) RegisterDeleteByIDMock(id string) {
	m.Register("DELETE", "/api/v1/api-integrations/"+id, 200, "")
}

func (m *ApiIntegrationsMock) RegisterRefreshClientCredentialsMock(id string) {
	m.Register("POST", "/api/v1/api-integrations/"+id+"/client-credentials", 200, "validate_client_credentials.json")
}
