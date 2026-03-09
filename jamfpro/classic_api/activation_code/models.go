package activation_code

import "encoding/xml"

// ResourceActivationCode represents the structure of an activation code resource.
type ResourceActivationCode struct {
	XMLName          xml.Name `xml:"activation_code"`
	OrganizationName string   `xml:"organization_name,omitempty"`
	Code             string   `xml:"code,omitempty"`
}

// RequestActivationCode represents the request body for updating an activation code.
type RequestActivationCode struct {
	XMLName          xml.Name `xml:"activation_code"`
	OrganizationName string   `xml:"organization_name,omitempty"`
	Code             string   `xml:"code,omitempty"`
}
