package accounts

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ResourceAccount represents a Jamf Pro Classic API account resource.
// This can be either a user or a group account.
type ResourceAccount struct {
	XMLName             xml.Name                 `xml:"account"`
	ID                  int                      `xml:"id,omitempty"`
	Name                string                   `xml:"name,omitempty"`
	DirectoryUser       bool                     `xml:"directory_user,omitempty"`
	FullName            string                   `xml:"full_name,omitempty"`
	Email               string                   `xml:"email,omitempty"`
	EmailAddress        string                   `xml:"email_address,omitempty"`
	Enabled             string                   `xml:"enabled,omitempty"`
	LdapServer          LdapServerSubset         `xml:"ldap_server,omitempty"`
	ForcePasswordChange bool                     `xml:"force_password_change,omitempty"`
	AccessLevel         string                   `xml:"access_level,omitempty"`
	Password            string                   `xml:"password,omitempty"`
	PrivilegeSet        string                   `xml:"privilege_set,omitempty"`
	Site                *shared.SharedResourceSite `xml:"site,omitempty"`
	Privileges          PrivilegesSubset         `xml:"privileges,omitempty"`
	Groups              []GroupSubset            `xml:"groups>group,omitempty"`
}

// ListResponse is the response for ListAccounts (GET /JSSResource/accounts).
// Accounts are divided into users and groups.
type ListResponse struct {
	XMLName xml.Name       `xml:"accounts"`
	Users   []UserSubset   `xml:"users>user,omitempty"`
	Groups  []GroupSubset  `xml:"groups>group,omitempty"`
}

// RequestAccount is the body for creating or updating an account.
// The ID field is not included; the target is specified via the URL path.
type RequestAccount struct {
	XMLName             xml.Name                 `xml:"account"`
	Name                string                   `xml:"name,omitempty"`
	DirectoryUser       bool                     `xml:"directory_user,omitempty"`
	FullName            string                   `xml:"full_name,omitempty"`
	Email               string                   `xml:"email,omitempty"`
	EmailAddress        string                   `xml:"email_address,omitempty"`
	Enabled             string                   `xml:"enabled,omitempty"`
	LdapServer          LdapServerSubset         `xml:"ldap_server,omitempty"`
	ForcePasswordChange bool                     `xml:"force_password_change,omitempty"`
	AccessLevel         string                   `xml:"access_level,omitempty"`
	Password            string                   `xml:"password,omitempty"`
	PrivilegeSet        string                   `xml:"privilege_set,omitempty"`
	Site                *shared.SharedResourceSite `xml:"site,omitempty"`
	Privileges          PrivilegesSubset         `xml:"privileges,omitempty"`
	Groups              []GroupSubset            `xml:"groups>group,omitempty"`
}

// UserSubset represents a user account in the list response.
type UserSubset struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// GroupSubset represents a group account in the list response or as a nested group.
type GroupSubset struct {
	ID         int              `xml:"id,omitempty"`
	Name       string           `xml:"name,omitempty"`
	Privileges PrivilegesSubset `xml:"privileges,omitempty"`
}

// PrivilegesSubset represents the privileges assigned to an account.
type PrivilegesSubset struct {
	JSSObjects    []string `xml:"jss_objects>privilege,omitempty"`
	JSSSettings   []string `xml:"jss_settings>privilege,omitempty"`
	JSSActions    []string `xml:"jss_actions>privilege,omitempty"`
	Recon         []string `xml:"recon>privilege,omitempty"`
	CasperAdmin   []string `xml:"casper_admin>privilege,omitempty"`
	CasperRemote  []string `xml:"casper_remote>privilege,omitempty"`
	CasperImaging []string `xml:"casper_imaging>privilege,omitempty"`
}

// LdapServerSubset represents an LDAP server reference.
type LdapServerSubset struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
