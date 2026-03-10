package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CloudAzureMock struct {
	*mocks.GenericMock
}

func NewCloudAzureMock() *CloudAzureMock {
	return &CloudAzureMock{
		GenericMock: mocks.NewJSONMock("CloudAzureMock"),
	}
}

func (m *CloudAzureMock) RegisterGetDefaultServerConfigurationMock() {
	m.Register("GET", "/api/v1/cloud-azure/defaults/server-configuration", 200, "validate_default_server.json")
}

func (m *CloudAzureMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v1/cloud-azure/"+id, 200, "validate_get.json")
}

func (m *CloudAzureMock) RegisterListMock() {
	m.Register("GET", "/api/v1/cloud-azure", 200, "validate_list.json")
}

func (m *CloudAzureMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/cloud-azure", 201, "validate_create.json")
}

func (m *CloudAzureMock) RegisterUpdateByIDMock(id string) {
	m.Register("PUT", "/api/v1/cloud-azure/"+id, 200, "validate_get.json")
}

func (m *CloudAzureMock) RegisterDeleteByIDMock(id string) {
	m.Register("DELETE", "/api/v1/cloud-azure/"+id, 204, "")
}

func (m *CloudAzureMock) RegisterGetDefaultMappingsMock() {
	m.Register("GET", "/api/v1/cloud-azure/defaults/mappings", 200, "validate_default_mappings.json")
}

func (m *CloudAzureMock) RegisterErrorMocks() {
	m.RegisterError("GET", "/api/v1/cloud-azure/defaults/server-configuration", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/cloud-azure/1", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/cloud-azure", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/cloud-azure", 500, "error_internal.json", "")
	m.RegisterError("PUT", "/api/v1/cloud-azure/1", 500, "error_internal.json", "")
	m.RegisterError("DELETE", "/api/v1/cloud-azure/1", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/cloud-azure/defaults/mappings", 500, "error_internal.json", "")
}
