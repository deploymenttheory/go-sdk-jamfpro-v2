package ibeacons

import "encoding/xml"

// ResourceIBeacon represents a Jamf Pro Classic API iBeacon resource.
type ResourceIBeacon struct {
	XMLName xml.Name `xml:"ibeacon"`
	ID      int      `xml:"id,omitempty"`
	Name    string   `xml:"name,omitempty"`
	UUID    string   `xml:"uuid"`
	UDID    string   `xml:"udid,omitempty"`
	Major   int      `xml:"major,omitempty"`
	Minor   int      `xml:"minor,omitempty"`
}

// ListItemIBeacon is the slim representation returned in list responses.
type ListItemIBeacon struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListIBeacons (GET /JSSResource/ibeacons).
type ListResponse struct {
	XMLName xml.Name          `xml:"ibeacons"`
	Size    int               `xml:"size"`
	Results []ListItemIBeacon `xml:"ibeacon"`
}

// RequestIBeacon is the body for creating or updating an iBeacon.
// The ID field is not included; the target is specified via the URL path.
type RequestIBeacon struct {
	XMLName xml.Name `xml:"ibeacon"`
	Name    string   `xml:"name"`
	UUID    string   `xml:"uuid"`
	Major   int      `xml:"major,omitempty"`
	Minor   int      `xml:"minor,omitempty"`
}
