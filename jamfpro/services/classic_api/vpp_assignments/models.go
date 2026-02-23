package vpp_assignments

import (
	"encoding/xml"
)

// ListResponse is the response for List (GET /JSSResource/vppassignments).
type ListResponse struct {
	XMLName       xml.Name           `xml:"vpp_assignments"`
	VPPAssignments []VPPAssignmentItem `xml:"vpp_assignment"`
}

// VPPAssignmentItem represents a single VPP assignment item in the list.
type VPPAssignmentItem struct {
	ID                int    `xml:"id"`
	VPPAdminAccountID int    `xml:"vpp_admin_account_id"`
	Name              string `xml:"name"`
}

// Resource represents the detailed VPP assignment resource.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/vppassignments
type Resource struct {
	XMLName xml.Name      `xml:"vpp_assignment"`
	General SubsetGeneral `xml:"general"`
	IOSApps  []VPPApp      `xml:"ios_apps>ios_app"`
	MacApps  []VPPApp      `xml:"mac_apps>mac_app"`
	EBooks   []VPPApp      `xml:"ebooks>ebook"`
	Scope    SubsetScope   `xml:"scope"`
}

// SubsetGeneral contains general VPP assignment attributes.
type SubsetGeneral struct {
	ID                  int    `xml:"id"`
	Name                string `xml:"name"`
	VPPAdminAccountID   int    `xml:"vpp_admin_account_id"`
	VPPAdminAccountName string `xml:"vpp_admin_account_name"`
}

// SubsetScope contains scope configuration for the VPP assignment.
type SubsetScope struct {
	AllJSSUsers   bool                    `xml:"all_jss_users"`
	JSSUsers      []VPPUser               `xml:"jss_users>user"`
	JSSUserGroups []VPPUserGroup          `xml:"jss_user_groups>user_group"`
	Limitations   SubsetScopeLimitations  `xml:"limitations"`
	Exclusions    SubsetScopeExclusions   `xml:"exclusions"`
}

// SubsetScopeLimitations contains scope limitations.
type SubsetScopeLimitations struct {
	UserGroups []VPPUserGroup `xml:"user_groups>user_group"`
}

// SubsetScopeExclusions contains scope exclusions.
type SubsetScopeExclusions struct {
	JSSUsers      []VPPUser     `xml:"jss_users>user"`
	UserGroups    []VPPUserGroup `xml:"user_groups>user_group"`
	JSSUserGroups []VPPUserGroup `xml:"jss_user_groups>user_group"`
}

// VPPApp represents an app (iOS, Mac, or eBook) in a VPP assignment.
type VPPApp struct {
	AdamID int    `xml:"adam_id"`
	Name   string `xml:"name"`
}

// VPPUser represents a user in scope.
type VPPUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// VPPUserGroup represents a user group in scope.
type VPPUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// RequestVPPAssignment is the request body for Create and Update operations.
// The root element must be vpp_assignment for the Classic API.
type RequestVPPAssignment struct {
	XMLName xml.Name `xml:"vpp_assignment"`
	*Resource
}

// CreateUpdateResponse represents the response from Create/Update operations.
// The Classic API may return only the ID for these operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"vpp_assignment"`
	ID      int      `xml:"id"`
}
