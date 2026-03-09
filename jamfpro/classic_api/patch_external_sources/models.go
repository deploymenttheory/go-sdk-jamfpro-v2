package patch_external_sources

import "encoding/xml"

// ResourcePatchExternalSource represents a Jamf Pro Classic API patch external source resource.
type ResourcePatchExternalSource struct {
	XMLName    xml.Name `xml:"patch_external_source"`
	ID         int      `xml:"id"`
	Name       string   `xml:"name"`
	HostName   string   `xml:"host_name,omitempty"`
	SSLEnabled bool     `xml:"ssl_enabled"`
	Port       int      `xml:"port,omitempty"`
}

// ListItemPatchExternalSource is the slim representation returned in list responses.
type ListItemPatchExternalSource struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListPatchExternalSources (GET /JSSResource/patchexternalsources).
type ListResponse struct {
	XMLName xml.Name                      `xml:"patch_external_sources"`
	Size    int                           `xml:"size"`
	Results []ListItemPatchExternalSource `xml:"patch_external_source"`
}

// RequestPatchExternalSource is the body for creating or updating a patch external source.
// The ID field is not included; the target is specified via the URL path.
type RequestPatchExternalSource struct {
	XMLName    xml.Name `xml:"patch_external_source"`
	Name       string   `xml:"name"`
	HostName   string   `xml:"host_name,omitempty"`
	SSLEnabled bool     `xml:"ssl_enabled"`
	Port       int      `xml:"port,omitempty"`
}
