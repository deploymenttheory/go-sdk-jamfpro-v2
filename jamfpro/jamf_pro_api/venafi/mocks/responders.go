package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type VenafiMock struct {
	*mocks.GenericMock
}

func NewVenafiMock() *VenafiMock {
	return &VenafiMock{
		GenericMock: mocks.NewJSONMock("VenafiMock"),
	}
}

func (m *VenafiMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/pki/venafi", 201, "validate_create.json")
}

func (m *VenafiMock) RegisterGetByIDMock(id string) {
	path := "/api/v1/pki/venafi/" + id
	m.Register("GET", path, 200, "validate_get.json")
}

func (m *VenafiMock) RegisterUpdateByIDMock(id string) {
	path := "/api/v1/pki/venafi/" + id
	m.Register("PATCH", path, 200, "validate_get.json")
}

func (m *VenafiMock) RegisterDeleteByIDMock(id string) {
	path := "/api/v1/pki/venafi/" + id
	m.Register("DELETE", path, 204, "")
}

func (m *VenafiMock) RegisterGetConnectionStatusMock(id string) {
	path := "/api/v1/pki/venafi/" + id + "/connection-status"
	m.Register("GET", path, 200, "validate_connection_status.json")
}

func (m *VenafiMock) RegisterGetDependentProfilesMock(id string) {
	path := "/api/v1/pki/venafi/" + id + "/dependent-profiles"
	m.Register("GET", path, 200, "validate_dependent_profiles.json")
}

func (m *VenafiMock) RegisterGetHistoryMock(id string) {
	path := "/api/v1/pki/venafi/" + id + "/history"
	m.Register("GET", path, 200, "validate_history.json")
}

func (m *VenafiMock) RegisterAddHistoryNoteMock(id string) {
	path := "/api/v1/pki/venafi/" + id + "/history"
	m.Register("POST", path, 201, "validate_history_note.json")
}

func (m *VenafiMock) RegisterGetJamfPublicKeyMock(id string) {
	path := "/api/v1/pki/venafi/" + id + "/jamf-public-key"
	m.Register("GET", path, 200, "validate_jamf_public_key.pem")
}

func (m *VenafiMock) RegisterGetProxyTrustStoreMock(id string) {
	path := "/api/v1/pki/venafi/" + id + "/proxy-trust-store"
	m.Register("GET", path, 200, "validate_jamf_public_key.pem")
}

func (m *VenafiMock) RegisterRegenerateJamfPublicKeyMock(id string) {
	path := "/api/v1/pki/venafi/" + id + "/jamf-public-key/regenerate"
	m.Register("POST", path, 200, "")
}

func (m *VenafiMock) RegisterUploadProxyTrustStoreMock(id string) {
	path := "/api/v1/pki/venafi/" + id + "/proxy-trust-store"
	m.Register("POST", path, 200, "")
}

func (m *VenafiMock) RegisterDeleteProxyTrustStoreMock(id string) {
	path := "/api/v1/pki/venafi/" + id + "/proxy-trust-store"
	m.Register("DELETE", path, 204, "")
}

func (m *VenafiMock) RegisterCreateNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v1/pki/venafi", 500, "error_internal.json", "no response registered")
}

func (m *VenafiMock) RegisterGetByIDNoResponseErrorMock(id string) {
	path := "/api/v1/pki/venafi/" + id
	m.RegisterError("GET", path, 500, "error_internal.json", "no response registered")
}
