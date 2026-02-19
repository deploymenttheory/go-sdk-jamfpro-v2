package directory_bindings

import "encoding/xml"

// ResourceDirectoryBinding represents a Jamf Pro Classic API directory binding resource.
type ResourceDirectoryBinding struct {
	XMLName    xml.Name `xml:"directory_binding"`
	ID         int      `xml:"id"`
	Name       string   `xml:"name"`
	Priority   int      `xml:"priority,omitempty"`
	Domain     string   `xml:"domain,omitempty"`
	Username   string   `xml:"username,omitempty"`
	Password   string   `xml:"password,omitempty"`
	ComputerOU string   `xml:"computer_ou,omitempty"`
	Type       string   `xml:"type,omitempty"`
}

// ListItemDirectoryBinding is the slim representation returned in list responses.
type ListItemDirectoryBinding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListDirectoryBindings (GET /JSSResource/directorybindings).
type ListResponse struct {
	XMLName xml.Name                   `xml:"directory_bindings"`
	Size    int                        `xml:"size"`
	Results []ListItemDirectoryBinding `xml:"directory_binding"`
}

// RequestDirectoryBinding is the body for creating or updating a directory binding.
// The ID field is not included; the target is specified via the URL path.
type RequestDirectoryBinding struct {
	XMLName    xml.Name `xml:"directory_binding"`
	Name       string   `xml:"name"`
	Priority   int      `xml:"priority,omitempty"`
	Domain     string   `xml:"domain,omitempty"`
	Username   string   `xml:"username,omitempty"`
	Password   string   `xml:"password,omitempty"`
	ComputerOU string   `xml:"computer_ou,omitempty"`
	Type       string   `xml:"type,omitempty"`
}
