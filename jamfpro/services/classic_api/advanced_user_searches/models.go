package advanced_user_searches

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ListResponse is the response for ListAdvancedUserSearches (GET /JSSResource/advancedusersearches).
type ListResponse struct {
	XMLName xml.Name   `xml:"advanced_user_searches"`
	Size    int        `xml:"size"`
	Results []ListItem `xml:"advanced_user_search"`
}

// ListItem represents a single advanced user search item in the list.
type ListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResourceAdvancedUserSearch represents a Jamf Pro Classic API advanced user search resource.
// This is returned by GET operations and includes all fields including computed results.
type ResourceAdvancedUserSearch struct {
	XMLName       xml.Name               `xml:"advanced_user_search"`
	ID            int                    `xml:"id,omitempty"`
	Name          string                 `xml:"name,omitempty"`
	Criteria      CriteriaContainer      `xml:"criteria"`
	DisplayFields []DisplayField         `xml:"display_fields>display_field,omitempty"`
	Users         []UserResult           `xml:"users>user,omitempty"`
	Site          *shared.SharedResourceSite `xml:"site,omitempty"`
}

// RequestAdvancedUserSearch is the body for creating or updating an advanced user search.
// The ID and Users fields are not included; ID is specified via the URL path, and Users are computed.
type RequestAdvancedUserSearch struct {
	XMLName       xml.Name                   `xml:"advanced_user_search"`
	Name          string                     `xml:"name"`
	Criteria      CriteriaContainer          `xml:"criteria"`
	DisplayFields []DisplayField             `xml:"display_fields>display_field,omitempty"`
	Site          *shared.SharedResourceSite `xml:"site,omitempty"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
// The Classic API returns only the ID for these operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"advanced_user_search"`
	ID      int      `xml:"id"`
}

// CriteriaContainer holds the search criteria with size and criterion list.
type CriteriaContainer struct {
	Size      int         `xml:"size"`
	Criterion []Criterion `xml:"criterion,omitempty"`
}

// Criterion represents a single search criterion.
type Criterion struct {
	Name         string `xml:"name"`
	Priority     int    `xml:"priority"`
	AndOr        string `xml:"and_or,omitempty"`
	SearchType   string `xml:"search_type"`
	Value        string `xml:"value"`
	OpeningParen bool   `xml:"opening_paren,omitempty"`
	ClosingParen bool   `xml:"closing_paren,omitempty"`
}

// DisplayField represents a field to display in search results.
type DisplayField struct {
	Name string `xml:"name"`
}

// UserResult represents a user in the search results.
// This is a read-only field returned by the API.
type UserResult struct {
	ID       int    `xml:"id"`
	Name     string `xml:"name"`
	Username string `xml:"Username,omitempty"`
	Email    string `xml:"Email,omitempty"`
}
