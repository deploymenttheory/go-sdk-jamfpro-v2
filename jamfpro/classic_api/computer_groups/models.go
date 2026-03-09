package computer_groups

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
)

// ResourceComputerGroup represents a Jamf Pro Classic API computer group resource.
type ResourceComputerGroup struct {
	XMLName   xml.Name                 `xml:"computer_group"`
	ID        int                      `xml:"id,omitempty"`
	Name      string                   `xml:"name"`
	IsSmart   bool                     `xml:"is_smart"`
	Site      *shared.SharedResourceSite `xml:"site,omitempty"`
	Criteria  *CriteriaContainer       `xml:"criteria,omitempty"`
	Computers []Computer               `xml:"computers>computer,omitempty"`
}

// ListResponse is the response for List (GET /JSSResource/computergroups).
type ListResponse struct {
	XMLName xml.Name              `xml:"computer_groups"`
	Size    int                   `xml:"size"`
	Results []ComputerGroupItem   `xml:"computer_group"`
}

// ComputerGroupItem represents a single computer group item in the list.
type ComputerGroupItem struct {
	ID      int    `xml:"id"`
	Name    string `xml:"name"`
	IsSmart bool   `xml:"is_smart"`
}

// RequestComputerGroup is the body for creating or updating a computer group.
// The ID field is not included; the target is specified via the URL path.
type RequestComputerGroup struct {
	XMLName   xml.Name                 `xml:"computer_group"`
	Name      string                   `xml:"name"`
	IsSmart   bool                     `xml:"is_smart"`
	Site      *shared.SharedResourceSite `xml:"site,omitempty"`
	Criteria  *CriteriaContainer       `xml:"criteria,omitempty"`
	Computers []Computer               `xml:"computers>computer,omitempty"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
// The Classic API returns only the ID for these operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"computer_group"`
	ID      int      `xml:"id"`
}

// CriteriaContainer wraps the criteria for smart computer groups.
type CriteriaContainer struct {
	Size      int                            `xml:"size,omitempty"`
	Criterion []shared.SharedSubsetCriteria  `xml:"criterion,omitempty"`
}

// Computer represents a computer in a computer group.
type Computer struct {
	ID            int    `xml:"id,omitempty"`
	Name          string `xml:"name,omitempty"`
	SerialNumber  string `xml:"serial_number,omitempty"`
	MacAddress    string `xml:"mac_address,omitempty"`
	AltMacAddress string `xml:"alt_mac_address,omitempty"`
}
