package advanced_computer_searches

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ListResponse is the response for ListAdvancedComputerSearches (GET /JSSResource/advancedcomputersearches).
type ListResponse struct {
	XMLName xml.Name   `xml:"advanced_computer_searches"`
	Size    int        `xml:"size"`
	Results []ListItem `xml:"advanced_computer_search"`
}

// ListItem represents a single advanced computer search item in the list.
type ListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResourceAdvancedComputerSearch represents a Jamf Pro Classic API advanced computer search resource.
// This is returned by GET operations and includes all fields including computed results.
type ResourceAdvancedComputerSearch struct {
	XMLName       xml.Name               `xml:"advanced_computer_search"`
	ID            int                    `xml:"id"`
	Name          string                 `xml:"name"`
	ViewAs        string                 `xml:"view_as,omitempty"`
	Sort1         string                 `xml:"sort_1,omitempty"`
	Sort2         string                 `xml:"sort_2,omitempty"`
	Sort3         string                 `xml:"sort_3,omitempty"`
	Criteria      CriteriaContainer      `xml:"criteria"`
	DisplayFields []DisplayField         `xml:"display_fields>display_field,omitempty"`
	Computers     []ComputerResult       `xml:"computers>computer,omitempty"`
	Site          *shared.SharedResourceSite `xml:"site,omitempty"`
}

// RequestAdvancedComputerSearch is the body for creating or updating an advanced computer search.
// The ID and Computers fields are not included; ID is specified via the URL path, and Computers are computed.
type RequestAdvancedComputerSearch struct {
	XMLName       xml.Name                   `xml:"advanced_computer_search"`
	Name          string                     `xml:"name"`
	ViewAs        string                     `xml:"view_as,omitempty"`
	Sort1         string                     `xml:"sort_1,omitempty"`
	Sort2         string                     `xml:"sort_2,omitempty"`
	Sort3         string                     `xml:"sort_3,omitempty"`
	Criteria      CriteriaContainer          `xml:"criteria"`
	DisplayFields []DisplayField             `xml:"display_fields>display_field,omitempty"`
	Site          *shared.SharedResourceSite `xml:"site,omitempty"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
// The Classic API returns only the ID for these operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"advanced_computer_search"`
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

// ComputerResult represents a computer in the search results.
// This is a read-only field returned by the API.
type ComputerResult struct {
	ID           int    `xml:"id"`
	Name         string `xml:"name"`
	UDID         string `xml:"udid,omitempty"`
	ComputerName string `xml:"Computer_Name,omitempty"`
}
