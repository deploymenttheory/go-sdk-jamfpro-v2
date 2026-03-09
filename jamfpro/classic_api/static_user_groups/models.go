package static_user_groups

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
)

// ResourceStaticUserGroup represents a Jamf Pro Classic API static user group resource.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/usergroups
type ResourceStaticUserGroup struct {
	XMLName          xml.Name                   `xml:"user_group"`
	ID               int                        `xml:"id,omitempty"`
	Name             string                     `xml:"name,omitempty"`
	IsSmart          bool                       `xml:"is_smart"`
	IsNotifyOnChange bool                       `xml:"is_notify_on_change"`
	Site             *shared.SharedResourceSite `xml:"site,omitempty"`
	Users            []UserItem                 `xml:"users>user,omitempty"`
	UserAdditions    []UserItem                 `xml:"user_additions>user,omitempty"`
	UserDeletions    []UserItem                 `xml:"user_deletions>user,omitempty"`
}

// ListResponse is the response for List (GET /JSSResource/usergroups).
type ListResponse struct {
	XMLName xml.Name        `xml:"user_groups"`
	Size    int             `xml:"size"`
	Results []UserGroupItem `xml:"user_group"`
}

// UserGroupItem represents a single user group item in the list.
type UserGroupItem struct {
	ID               int    `xml:"id"`
	Name             string `xml:"name"`
	IsSmart          bool   `xml:"is_smart"`
	IsNotifyOnChange bool   `xml:"is_notify_on_change"`
}

// RequestStaticUserGroup is the body for creating or updating a static user group.
type RequestStaticUserGroup struct {
	XMLName          xml.Name                   `xml:"user_group"`
	Name             string                     `xml:"name"`
	IsSmart          bool                       `xml:"is_smart"`
	IsNotifyOnChange bool                       `xml:"is_notify_on_change"`
	Site             *shared.SharedResourceSite `xml:"site,omitempty"`
	Users            []UserItem                 `xml:"users>user,omitempty"`
	UserAdditions    []UserItem                 `xml:"user_additions>user,omitempty"`
	UserDeletions    []UserItem                 `xml:"user_deletions>user,omitempty"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
// The Classic API returns only the ID for these operations.
// Note: CREATE returns <static_user_group>, UPDATE returns <user_group>.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"static_user_group"`
	ID      int      `xml:"id"`
}

// UserItem represents a user in a user group.
type UserItem struct {
	ID           int    `xml:"id,omitempty"`
	Username     string `xml:"username,omitempty"`
	FullName     string `xml:"full_name,omitempty"`
	PhoneNumber  string `xml:"phone_number,omitempty"`
	EmailAddress string `xml:"email_address,omitempty"`
}
