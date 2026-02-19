package software_update_servers

import "encoding/xml"

// ResourceSoftwareUpdateServer represents a Jamf Pro Classic API software update server resource.
type ResourceSoftwareUpdateServer struct {
	XMLName       xml.Name `xml:"software_update_server"`
	ID            int      `xml:"id"`
	Name          string   `xml:"name"`
	IPAddress     string   `xml:"ip_address,omitempty"`
	Port          int      `xml:"port,omitempty"`
	SetSystemWide bool     `xml:"set_system_wide"`
}

// ListItemSoftwareUpdateServer is the slim representation returned in list responses.
type ListItemSoftwareUpdateServer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListSoftwareUpdateServers (GET /JSSResource/softwareupdateservers).
type ListResponse struct {
	XMLName xml.Name                       `xml:"software_update_servers"`
	Size    int                            `xml:"size"`
	Results []ListItemSoftwareUpdateServer `xml:"software_update_server"`
}

// RequestSoftwareUpdateServer is the body for creating or updating a software update server.
// The ID field is not included; the target is specified via the URL path.
type RequestSoftwareUpdateServer struct {
	XMLName       xml.Name `xml:"software_update_server"`
	Name          string   `xml:"name"`
	IPAddress     string   `xml:"ip_address,omitempty"`
	Port          int      `xml:"port,omitempty"`
	SetSystemWide bool     `xml:"set_system_wide"`
}
