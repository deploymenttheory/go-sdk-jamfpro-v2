package accounts_groups

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ResourceAccountGroup represents a Jamf Pro Classic API account group resource.
// Note: GET operations return <group>, while CREATE returns <user_group>
// This struct is used for GET operations
type ResourceAccountGroup struct {
	XMLName      xml.Name                   `xml:"group"`
	ID           int                        `xml:"id,omitempty"`
	Name         string                     `xml:"name,omitempty"`
	AccessLevel  string                     `xml:"access_level,omitempty"`
	PrivilegeSet string                     `xml:"privilege_set,omitempty"`
	Site         *shared.SharedResourceSite `xml:"site,omitempty"`
	Privileges   PrivilegesSubset           `xml:"privileges,omitempty"`
	Members      []MemberUser               `xml:"members>user,omitempty"`
	LdapServer   LdapServerSubset           `xml:"ldap_server,omitempty"`
}

// RequestAccountGroup is the body for creating or updating an account group.
// The ID field is not included; the target is specified via the URL path.
type RequestAccountGroup struct {
	XMLName      xml.Name                   `xml:"group"`
	Name         string                     `xml:"name,omitempty"`
	AccessLevel  string                     `xml:"access_level,omitempty"`
	PrivilegeSet string                     `xml:"privilege_set,omitempty"`
	Site         *shared.SharedResourceSite `xml:"site,omitempty"`
	Privileges   PrivilegesSubset           `xml:"privileges,omitempty"`
	Members      []MemberUser               `xml:"members>user,omitempty"`
	LdapServer   LdapServerSubset           `xml:"ldap_server,omitempty"`
}

// CreateResponse is the response from creating an account group.
// The Classic API returns only the ID when creating an account group.
// Note: CREATE/UPDATE operations return <user_group>, while GET returns <group>
type CreateResponse struct {
	XMLName xml.Name `xml:"user_group"`
	ID      int      `xml:"id,omitempty"`
}

// UpdateResponse is the response from updating an account group.
// The Classic API returns only the ID when updating an account group.
// Note: UPDATE operations return <user_group>, while GET returns <group>
type UpdateResponse struct {
	XMLName xml.Name `xml:"user_group"`
	ID      int      `xml:"id,omitempty"`
}

// MemberUser represents a user member of an account group.
type MemberUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// PrivilegesSubset represents the privileges assigned to an account group.
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
