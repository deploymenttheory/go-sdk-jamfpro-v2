package dock_items

import "encoding/xml"

// ListResponse is the response for List (GET /JSSResource/dockitems).
// API reference: https://developer.jamf.com/jamf-pro/reference/finddockitems
type ListResponse struct {
	XMLName   xml.Name      `xml:"dock_items"`
	Size      int           `xml:"size"`
	DockItems []ListItem    `xml:"dock_item"`
}

// ListItem represents a single dock item in the list.
type ListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource represents a single dock item resource.
// API reference: https://developer.jamf.com/jamf-pro/reference/finddockitemsbyid
type Resource struct {
	XMLName   xml.Name `xml:"dock_item"`
	ID        int      `xml:"id"`
	Name      string   `xml:"name"`
	Type      string   `xml:"type"`
	Path      string   `xml:"path"`
	Contents  string   `xml:"contents"`
}

// Request is the body for creating or updating a dock item.
// The ID field is not included; the target is specified via the URL path.
type Request struct {
	XMLName  xml.Name `xml:"dock_item"`
	Name     string   `xml:"name"`
	Type     string   `xml:"type"`
	Path     string   `xml:"path"`
	Contents string   `xml:"contents"`
}
