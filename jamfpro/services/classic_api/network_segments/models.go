package network_segments

import "encoding/xml"

// ResourceNetworkSegment represents a Jamf Pro Classic API network segment resource.
type ResourceNetworkSegment struct {
	XMLName             xml.Name `xml:"network_segment"`
	ID                  int      `xml:"id"`
	Name                string   `xml:"name"`
	StartingAddress     string   `xml:"starting_address"`
	EndingAddress       string   `xml:"ending_address"`
	DistributionServer  string   `xml:"distribution_server,omitempty"`
	DistributionPoint   string   `xml:"distribution_point,omitempty"`
	URL                 string   `xml:"url,omitempty"`
	SWUServer           string   `xml:"swu_server,omitempty"`
	Building            string   `xml:"building,omitempty"`
	Department          string   `xml:"department,omitempty"`
	OverrideBuildings   bool     `xml:"override_buildings"`
	OverrideDepartments bool     `xml:"override_departments"`
}

// ListItemNetworkSegment is the slim representation returned in list responses.
type ListItemNetworkSegment struct {
	ID              int    `xml:"id"`
	Name            string `xml:"name"`
	StartingAddress string `xml:"starting_address"`
	EndingAddress   string `xml:"ending_address"`
}

// ListResponse is the response for ListNetworkSegments (GET /JSSResource/networksegments).
type ListResponse struct {
	XMLName xml.Name                 `xml:"network_segments"`
	Size    int                      `xml:"size"`
	Results []ListItemNetworkSegment `xml:"network_segment"`
}

// CreateUpdateResponse is the response body for create and update operations,
// which return only the assigned resource ID.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"network_segment"`
	ID      int      `xml:"id"`
}

// RequestNetworkSegment is the body for creating or updating a network segment.
// The ID field is not included; the target is specified via the URL path.
type RequestNetworkSegment struct {
	XMLName             xml.Name `xml:"network_segment"`
	Name                string   `xml:"name"`
	StartingAddress     string   `xml:"starting_address"`
	EndingAddress       string   `xml:"ending_address"`
	DistributionServer  string   `xml:"distribution_server,omitempty"`
	DistributionPoint   string   `xml:"distribution_point,omitempty"`
	URL                 string   `xml:"url,omitempty"`
	SWUServer           string   `xml:"swu_server,omitempty"`
	Building            string   `xml:"building,omitempty"`
	Department          string   `xml:"department,omitempty"`
	OverrideBuildings   bool     `xml:"override_buildings"`
	OverrideDepartments bool     `xml:"override_departments"`
}
