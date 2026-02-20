package byoprofiles

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ResourceBYOProfile represents a Jamf Pro Classic API BYO profile resource.
type ResourceBYOProfile struct {
	XMLName xml.Name        `xml:"byoprofile"`
	ID      int             `xml:"id"`
	General GeneralSettings `xml:"general"`
}

// ListResponse is the response for ListBYOProfiles (GET /JSSResource/byoprofiles).
type ListResponse struct {
	XMLName xml.Name   `xml:"byoprofiles"`
	Size    int        `xml:"size"`
	Results []ListItem `xml:"byoprofile"`
}

// ListItem represents a single BYO profile item in the list.
type ListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// RequestBYOProfile is the body for creating or updating a BYO profile.
// The ID field is not included; the target is specified via the URL path.
type RequestBYOProfile struct {
	XMLName xml.Name        `xml:"byoprofile"`
	General GeneralSettings `xml:"general"`
}

// CreateUpdateResponse represents the response from Create/Update operations.
// The Classic API returns only the ID for these operations.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"byoprofile"`
	ID      int      `xml:"id"`
}

// GeneralSettings represents the general settings of a BYO profile.
type GeneralSettings struct {
	ID          int                        `xml:"id,omitempty"`
	Name        string                     `xml:"name"`
	Site        *shared.SharedResourceSite `xml:"site,omitempty"`
	Enabled     bool                       `xml:"enabled"`
	Description string                     `xml:"description,omitempty"`
}
