package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CloudLdapMock struct {
	*mocks.GenericMock
}

func NewCloudLdapMock() *CloudLdapMock {
	return &CloudLdapMock{
		GenericMock: mocks.NewJSONMock("CloudLdapMock"),
	}
}

func (m *CloudLdapMock) RegisterGetDefaultMappingsMock(providerName string) {
	m.Register("GET", "/api/v2/cloud-ldaps/defaults/"+providerName+"/mappings", 200, "validate_default_mappings.json")
}

func (m *CloudLdapMock) RegisterGetDefaultServerConfigurationMock(providerName string) {
	m.Register("GET", "/api/v2/cloud-ldaps/defaults/"+providerName+"/server-configuration", 200, "validate_default_server.json")
}

func (m *CloudLdapMock) RegisterCreateMock() {
	m.Register("POST", "/api/v2/cloud-ldaps", 201, "validate_create.json")
}

func (m *CloudLdapMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v2/cloud-ldaps/"+id, 200, "validate_get.json")
}

func (m *CloudLdapMock) RegisterUpdateByIDMock(id string) {
	m.Register("PUT", "/api/v2/cloud-ldaps/"+id, 200, "validate_get.json")
}

func (m *CloudLdapMock) RegisterDeleteByIDMock(id string) {
	m.Register("DELETE", "/api/v2/cloud-ldaps/"+id, 204, "")
}

func (m *CloudLdapMock) RegisterGetBindConnectionPoolStatsMock(id string) {
	m.Register("GET", "/api/v2/cloud-ldaps/"+id+"/connection/bind", 200, "validate_connection_pool.json")
}

func (m *CloudLdapMock) RegisterGetSearchConnectionPoolStatsMock(id string) {
	m.Register("GET", "/api/v2/cloud-ldaps/"+id+"/connection/search", 200, "validate_connection_pool.json")
}

func (m *CloudLdapMock) RegisterTestConnectionMock(id string) {
	m.Register("GET", "/api/v2/cloud-ldaps/"+id+"/connection/status", 200, "validate_connection_status.json")
}

func (m *CloudLdapMock) RegisterGetMappingsByIDMock(id string) {
	m.Register("GET", "/api/v2/cloud-ldaps/"+id+"/mappings", 200, "validate_default_mappings.json")
}

func (m *CloudLdapMock) RegisterUpdateMappingsByIDMock(id string) {
	m.Register("PUT", "/api/v2/cloud-ldaps/"+id+"/mappings", 200, "validate_default_mappings.json")
}

func (m *CloudLdapMock) RegisterGetByIDErrorMock(id string) {
	m.RegisterError("GET", "/api/v2/cloud-ldaps/"+id, 404, "validate_get.json", "Jamf Pro API error (404): not found")
}

func (m *CloudLdapMock) RegisterCreateErrorMock() {
	m.RegisterError("POST", "/api/v2/cloud-ldaps", 500, "validate_create.json", "Jamf Pro API error (500): server error")
}
