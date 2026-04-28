package account_groups

// ResourceAccountGroup represents a Jamf Pro account group.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-account-groups-id
type ResourceAccountGroup struct {
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	Description    string                   `json:"description,omitempty"`
	AccessLevel    string                   `json:"accessLevel,omitempty"`    // FullAccess, SiteAccess, GroupBasedAccess
	PrivilegeLevel string                   `json:"privilegeLevel,omitempty"` // ADMINISTRATOR, AUDITOR, ENROLLMENT, CUSTOM
	Site           *AccountGroupSite        `json:"site,omitempty"`
	LdapServer     *AccountGroupLdap        `json:"ldapServer,omitempty"`
	Users          []AccountGroupUserMember `json:"users,omitempty"`
}

// AccountGroupSite represents the site associated with an account group.
type AccountGroupSite struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// AccountGroupLdap represents the LDAP server associated with an account group.
type AccountGroupLdap struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// AccountGroupUserMember represents a user member within an account group.
type AccountGroupUserMember struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// ListAccountGroupsResponse is the paginated response for ListV1.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-account-groups
type ListAccountGroupsResponse struct {
	TotalCount int                    `json:"totalCount"`
	Results    []ResourceAccountGroup `json:"results"`
}
