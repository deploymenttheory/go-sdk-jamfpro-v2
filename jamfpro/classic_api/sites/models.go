package sites

import "encoding/xml"

// ResourceSite represents a Jamf Pro Classic API site resource.
type ResourceSite struct {
	XMLName xml.Name `xml:"site"`
	ID      int      `xml:"id"`
	Name    string   `xml:"name"`
}

// ListResponse is the response for ListSites (GET /JSSResource/sites).
type ListResponse struct {
	XMLName xml.Name       `xml:"sites"`
	Size    int            `xml:"size"`
	Results []ResourceSite `xml:"site"`
}

// RequestSite is the body for creating or updating a site.
// The ID field is not included; the target is specified via the URL path.
type RequestSite struct {
	XMLName xml.Name `xml:"site"`
	Name    string   `xml:"name"`
}
