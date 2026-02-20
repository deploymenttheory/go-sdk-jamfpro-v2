package removeable_mac_addresses

import "encoding/xml"

// ResourceRemoveableMacAddress represents a Jamf Pro Classic API removeable MAC address resource.
type ResourceRemoveableMacAddress struct {
	XMLName xml.Name `xml:"removable_mac_address"`
	ID      int      `xml:"id"`
	Name    string   `xml:"name"`
}

// ListResponse is the response for ListRemoveableMacAddresses (GET /JSSResource/removablemacaddresses).
type ListResponse struct {
	XMLName xml.Name                       `xml:"removable_mac_addresses"`
	Size    int                            `xml:"size"`
	Results []ResourceRemoveableMacAddress `xml:"removable_mac_address"`
}

// RequestRemoveableMacAddress is the body for creating or updating a removeable MAC address.
// The ID field is not included; the target is specified via the URL path.
type RequestRemoveableMacAddress struct {
	XMLName xml.Name `xml:"removable_mac_address"`
	Name    string   `xml:"name"`
}
