package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type LDAPServersMock struct {
	*mocks.GenericMock
}

func NewLDAPServersMock() *LDAPServersMock {
	return &LDAPServersMock{
		GenericMock: mocks.NewXMLMock("LDAPServersMock"),
	}
}

func (m *LDAPServersMock) RegisterMocks() {
	m.RegisterListLDAPServersMock()
	m.RegisterGetLDAPServerByIDMock()
	m.RegisterGetLDAPServerByNameMock()
	m.RegisterCreateLDAPServerMock()
	m.RegisterUpdateLDAPServerByIDMock()
	m.RegisterUpdateLDAPServerByNameMock()
	m.RegisterDeleteLDAPServerByIDMock()
	m.RegisterDeleteLDAPServerByNameMock()
}

func (m *LDAPServersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *LDAPServersMock) RegisterListLDAPServersMock() {
	m.Register("GET", "/JSSResource/ldapservers", 200, "validate_list_ldap_servers.xml")
}

func (m *LDAPServersMock) RegisterGetLDAPServerByIDMock() {
	m.Register("GET", "/JSSResource/ldapservers/id/1", 200, "validate_get_ldap_server.xml")
}

func (m *LDAPServersMock) RegisterGetLDAPServerByNameMock() {
	m.Register("GET", "/JSSResource/ldapservers/name/Test LDAP Server", 200, "validate_get_ldap_server.xml")
}

func (m *LDAPServersMock) RegisterCreateLDAPServerMock() {
	m.Register("POST", "/JSSResource/ldapservers/id/0", 201, "validate_create_ldap_server.xml")
}

func (m *LDAPServersMock) RegisterUpdateLDAPServerByIDMock() {
	m.Register("PUT", "/JSSResource/ldapservers/id/1", 200, "validate_update_ldap_server.xml")
}

func (m *LDAPServersMock) RegisterUpdateLDAPServerByNameMock() {
	m.Register("PUT", "/JSSResource/ldapservers/name/Test LDAP Server", 200, "validate_update_ldap_server.xml")
}

func (m *LDAPServersMock) RegisterDeleteLDAPServerByIDMock() {
	m.Register("DELETE", "/JSSResource/ldapservers/id/1", 200, "")
}

func (m *LDAPServersMock) RegisterDeleteLDAPServerByNameMock() {
	m.Register("DELETE", "/JSSResource/ldapservers/name/Test LDAP Server", 200, "")
}

func (m *LDAPServersMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/ldapservers/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *LDAPServersMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/ldapservers/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): An LDAP server with that name already exists")
}

