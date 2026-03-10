package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DigicertMock struct {
	*mocks.GenericMock
}

func NewDigicertMock() *DigicertMock {
	return &DigicertMock{
		GenericMock: mocks.NewJSONMock("DigicertMock"),
	}
}

func (m *DigicertMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/pki/digicert/trust-lifecycle-manager", 201, "validate_create.json")
}

func (m *DigicertMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 200, "validate_get.json")
}

func (m *DigicertMock) RegisterUpdateByIDMock(id string) {
	m.Register("PATCH", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 200, "")
}

func (m *DigicertMock) RegisterDeleteByIDMock(id string) {
	m.Register("DELETE", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 204, "")
}

func (m *DigicertMock) RegisterValidateClientCertificateMock() {
	m.Register("POST", "/api/v1/pki/digicert/trust-lifecycle-manager/validate-client-certificate", 200, "")
}

func (m *DigicertMock) RegisterGetConnectionStatusMock(id string) {
	m.Register("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/connection-status", 200, "validate_connection_status.json")
}

func (m *DigicertMock) RegisterGetDependenciesMock(id string) {
	m.Register("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/dependencies", 200, "validate_dependencies.json")
}

func (m *DigicertMock) RegisterNotFoundErrorMock(id string) {
	m.RegisterError("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 404, "error_not_found.json", "")
}

func (m *DigicertMock) RegisterConnectionStatusNotFoundErrorMock(id string) {
	m.RegisterError("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/connection-status", 404, "error_not_found.json", "")
}

func (m *DigicertMock) RegisterDependenciesNotFoundErrorMock(id string) {
	m.RegisterError("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/dependencies", 404, "error_not_found.json", "")
}

func (m *DigicertMock) RegisterCreateErrorMock() {
	m.RegisterError("POST", "/api/v1/pki/digicert/trust-lifecycle-manager", 500, "error_internal.json", "no response for")
}

func (m *DigicertMock) RegisterGetByIDErrorMock(id string) {
	m.RegisterError("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 500, "error_internal.json", "no response for")
}

func (m *DigicertMock) RegisterUpdateByIDErrorMock(id string) {
	m.RegisterError("PATCH", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 500, "error_internal.json", "no response for")
}

func (m *DigicertMock) RegisterDeleteByIDErrorMock(id string) {
	m.RegisterError("DELETE", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id, 500, "error_internal.json", "no response for")
}

func (m *DigicertMock) RegisterValidateClientCertificateErrorMock() {
	m.RegisterError("POST", "/api/v1/pki/digicert/trust-lifecycle-manager/validate-client-certificate", 500, "error_internal.json", "no response for")
}

func (m *DigicertMock) RegisterGetConnectionStatusErrorMock(id string) {
	m.RegisterError("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/connection-status", 500, "error_internal.json", "no response for")
}

func (m *DigicertMock) RegisterGetDependenciesErrorMock(id string) {
	m.RegisterError("GET", "/api/v1/pki/digicert/trust-lifecycle-manager/"+id+"/dependencies", 500, "error_internal.json", "no response for")
}
