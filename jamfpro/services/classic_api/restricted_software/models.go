package restricted_software

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ResourceRestrictedSoftware represents a Jamf Pro Classic API restricted software resource.
type ResourceRestrictedSoftware struct {
	XMLName xml.Name `xml:"restricted_software"`
	General General  `xml:"general"`
	Scope   Scope    `xml:"scope"`
}

// ListItem represents a restricted software item in a list response.
type ListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListRestrictedSoftware (GET /JSSResource/restrictedsoftware).
type ListResponse struct {
	XMLName xml.Name   `xml:"restricted_software_titles"`
	Size    int        `xml:"size"`
	Results []ListItem `xml:"restricted_software_title"`
}

// RequestRestrictedSoftware is the body for creating or updating restricted software.
// The ID field is not included in General; it's specified via the URL path.
type RequestRestrictedSoftware struct {
	XMLName xml.Name       `xml:"restricted_software"`
	General RequestGeneral `xml:"general"`
	Scope   Scope          `xml:"scope"`
}

// CreateUpdateResponse is the response for create/update operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"restricted_software"`
	ID      int      `xml:"id"`
}

// General contains the general settings for restricted software.
type General struct {
	ID                    int                        `xml:"id"`
	Name                  string                     `xml:"name"`
	ProcessName           string                     `xml:"process_name"`
	MatchExactProcessName bool                       `xml:"match_exact_process_name"`
	SendNotification      bool                       `xml:"send_notification"`
	KillProcess           bool                       `xml:"kill_process"`
	DeleteExecutable      bool                       `xml:"delete_executable"`
	DisplayMessage        string                     `xml:"display_message"`
	Site                  *shared.SharedResourceSite `xml:"site,omitempty"`
}

// RequestGeneral is the general section for create/update requests (excludes ID).
type RequestGeneral struct {
	Name                  string                     `xml:"name"`
	ProcessName           string                     `xml:"process_name"`
	MatchExactProcessName bool                       `xml:"match_exact_process_name"`
	SendNotification      bool                       `xml:"send_notification"`
	KillProcess           bool                       `xml:"kill_process"`
	DeleteExecutable      bool                       `xml:"delete_executable"`
	DisplayMessage        string                     `xml:"display_message"`
	Site                  *shared.SharedResourceSite `xml:"site,omitempty"`
}

// Scope defines the targeting and exclusions for restricted software.
type Scope struct {
	AllComputers   bool        `xml:"all_computers"`
	Computers      []ScopeItem `xml:"computers>computer,omitempty"`
	ComputerGroups []ScopeItem `xml:"computer_groups>computer_group,omitempty"`
	Buildings      []ScopeItem `xml:"buildings>building,omitempty"`
	Departments    []ScopeItem `xml:"departments>department,omitempty"`
	Exclusions     Exclusions  `xml:"exclusions"`
}

// Exclusions defines the scope exclusions for restricted software.
type Exclusions struct {
	Computers      []ScopeItem `xml:"computers>computer,omitempty"`
	ComputerGroups []ScopeItem `xml:"computer_groups>computer_group,omitempty"`
	Buildings      []ScopeItem `xml:"buildings>building,omitempty"`
	Departments    []ScopeItem `xml:"departments>department,omitempty"`
	Users          []ScopeItem `xml:"users>user,omitempty"`
}

// ScopeItem represents an entity in the scope or exclusions.
type ScopeItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}
