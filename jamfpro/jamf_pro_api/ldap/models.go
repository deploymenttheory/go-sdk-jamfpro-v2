package ldap

// ResourceLdapGroupV1 represents a Jamf Pro LDAP group definition.
type ResourceLdapGroupV1 struct {
	ID                string `json:"id"`
	UUID              string `json:"uuid"`
	LdapServerID      int    `json:"ldapServerId"`
	Name              string `json:"name"`
	DistinguishedName string `json:"distinguishedName"`
}

// ListGroupsResponseV1 is the response for GetLdapGroupsV1.
type ListGroupsResponseV1 struct {
	TotalCount int                   `json:"totalCount"`
	Results    []ResourceLdapGroupV1 `json:"results"`
}

// ResourceLdapServerV1 represents a Jamf Pro LDAP server summary.
type ResourceLdapServerV1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
