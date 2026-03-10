package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type LdapMock struct {
	*mocks.GenericMock
}

func NewLdapMock() *LdapMock {
	return &LdapMock{
		GenericMock: mocks.NewJSONMock("LdapMock"),
	}
}

func (m *LdapMock) RegisterGetLdapGroupsMock() {
	m.Register("GET", "/api/v1/ldap/groups", 200, "validate_list_groups.json")
}

func (m *LdapMock) RegisterGetLdapServersMock() {
	m.Register("GET", "/api/v1/ldap/servers", 200, "validate_list_servers.json")
}

func (m *LdapMock) RegisterGetLdapServersOnlyMock() {
	m.Register("GET", "/api/v1/ldap/ldap-servers", 200, "validate_list_servers_only.json")
}

func (m *LdapMock) RegisterGetLdapGroupsErrorMock() {
	m.RegisterError("GET", "/api/v1/ldap/groups", 500, "validate_list_groups.json", "Jamf Pro API error (500): server error")
}

func (m *LdapMock) RegisterGetLdapServersErrorMock() {
	m.RegisterError("GET", "/api/v1/ldap/servers", 404, "validate_list_servers.json", "Jamf Pro API error (404): not found")
}

func (m *LdapMock) RegisterGetLdapServersOnlyErrorMock() {
	m.RegisterError("GET", "/api/v1/ldap/ldap-servers", 500, "validate_list_servers_only.json", "Jamf Pro API error (500): server error")
}

func (m *LdapMock) RegisterGetLdapGroupsNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/ldap/groups", 500, "error_internal.json", "")
}

func (m *LdapMock) RegisterGetLdapServersNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/ldap/servers", 500, "error_internal.json", "")
}

func (m *LdapMock) RegisterGetLdapServersOnlyNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/ldap/ldap-servers", 500, "error_internal.json", "")
}
