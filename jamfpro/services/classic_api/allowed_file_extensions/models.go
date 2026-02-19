package allowed_file_extensions

import "encoding/xml"

// ResourceAllowedFileExtension represents a Jamf Pro Classic API allowed file extension resource.
type ResourceAllowedFileExtension struct {
	XMLName   xml.Name `xml:"allowed_file_extension"`
	ID        int      `xml:"id"`
	Extension string   `xml:"extension"`
}

// ListResponse is the response for ListAllowedFileExtensions (GET /JSSResource/allowedfileextensions).
type ListResponse struct {
	XMLName xml.Name                        `xml:"allowed_file_extensions"`
	Size    int                             `xml:"size"`
	Results []ResourceAllowedFileExtension  `xml:"allowed_file_extension"`
}

// RequestAllowedFileExtension is the body for creating an allowed file extension.
// The ID field is not included; the target is specified via the URL path.
type RequestAllowedFileExtension struct {
	XMLName   xml.Name `xml:"allowed_file_extension"`
	Extension string   `xml:"extension"`
}
