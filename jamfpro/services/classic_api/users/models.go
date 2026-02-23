package users

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ListResponse is the response for List (GET /JSSResource/users).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/users
type ListResponse struct {
	XMLName xml.Name   `xml:"users"`
	Size    int        `xml:"size"`
	Results []UserItem `xml:"user"`
}

// UserItem represents a single user item in the list.
type UserItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResourceUser represents a Jamf Pro Classic API user resource.
type ResourceUser struct {
	XMLName             xml.Name                    `xml:"user"`
	ID                  int                         `xml:"id,omitempty"`
	Name                string                     `xml:"name"`
	FullName             string                     `xml:"full_name,omitempty"`
	Email                string                     `xml:"email,omitempty"`
	EmailAddress         string                     `xml:"email_address,omitempty"`
	PhoneNumber          string                     `xml:"phone_number,omitempty"`
	Position             string                     `xml:"position,omitempty"`
	EnableCustomPhoto    bool                       `xml:"enable_custom_photo_url,omitempty"`
	CustomPhotoURL       string                     `xml:"custom_photo_url,omitempty"`
	LDAPServer           UserSubsetLDAPServer       `xml:"ldap_server,omitempty"`
	ExtensionAttributes   UserSubsetExtensionAttributes `xml:"extension_attributes,omitempty"`
	Sites                []shared.SharedResourceSite `xml:"sites>site,omitempty"`
	Links                UserSubsetLinks            `xml:"links,omitempty"`
}

// RequestUser is the body for creating or updating a user.
// The target is specified via the URL path (id, name, or email).
type RequestUser struct {
	XMLName             xml.Name                    `xml:"user"`
	Name                string                     `xml:"name"`
	FullName             string                     `xml:"full_name,omitempty"`
	Email                string                     `xml:"email,omitempty"`
	EmailAddress         string                     `xml:"email_address,omitempty"`
	PhoneNumber          string                     `xml:"phone_number,omitempty"`
	Position             string                     `xml:"position,omitempty"`
	EnableCustomPhoto    bool                       `xml:"enable_custom_photo_url,omitempty"`
	CustomPhotoURL       string                     `xml:"custom_photo_url,omitempty"`
	LDAPServer           UserSubsetLDAPServer       `xml:"ldap_server,omitempty"`
	ExtensionAttributes   UserSubsetExtensionAttributes `xml:"extension_attributes,omitempty"`
	Sites                []shared.SharedResourceSite `xml:"sites>site,omitempty"`
	Links                *UserSubsetLinks           `xml:"links,omitempty"`
}

// UserSubsetLDAPServer represents the LDAP server subset.
type UserSubsetLDAPServer struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// UserSubsetExtensionAttributes represents extension attributes.
type UserSubsetExtensionAttributes struct {
	Attributes []UserSubsetExtensionAttribute `xml:"extension_attribute,omitempty"`
}

// UserSubsetExtensionAttribute represents a single extension attribute.
type UserSubsetExtensionAttribute struct {
	ID    int    `xml:"id,omitempty"`
	Name  string `xml:"name,omitempty"`
	Type  string `xml:"type,omitempty"`
	Value string `xml:"value,omitempty"`
}

// UserSubsetLinks represents linked computers, peripherals, mobile devices, and VPP assignments.
type UserSubsetLinks struct {
	Computers         []UserSubsetLinksListItem `xml:"computers>computer,omitempty"`
	Peripherals       []UserSubsetLinksListItem `xml:"peripherals>peripheral,omitempty"`
	MobileDevices     []UserSubsetLinksListItem `xml:"mobile_devices>mobile_device,omitempty"`
	VPPAssignments   []UserSubsetLinksListItem `xml:"vpp_assignments>vpp_assignment,omitempty"`
	TotalVPPCodeCount int                       `xml:"total_vpp_code_count,omitempty"`
}

// UserSubsetLinksListItem represents a linked item.
type UserSubsetLinksListItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
