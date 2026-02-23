package user_extension_attributes

import (
	"encoding/xml"
)

// ListResponse is the response for List (GET /JSSResource/userextensionattributes).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finduserextensionattributes
type ListResponse struct {
	XMLName                 xml.Name                    `xml:"user_extension_attributes"`
	Size                    int                         `xml:"size"`
	UserExtensionAttributes []UserExtensionAttributeItem `xml:"user_extension_attribute"`
}

// UserExtensionAttributeItem represents a single user extension attribute in the list.
type UserExtensionAttributeItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResourceUserExtensionAttribute represents a Jamf Pro Classic API user extension attribute resource.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finduserextensionattributesbyid
type ResourceUserExtensionAttribute struct {
	XMLName     xml.Name                          `xml:"user_extension_attribute"`
	ID          int                               `xml:"id,omitempty"`
	Name        string                            `xml:"name"`
	Description string                            `xml:"description"`
	DataType    string                            `xml:"data_type"`
	InputType   ResourceUserExtensionAttributeInputType `xml:"input_type"`
}

// ResourceUserExtensionAttributeInputType represents the input type subset.
type ResourceUserExtensionAttributeInputType struct {
	Type string `xml:"type"`
}

// RequestUserExtensionAttribute is the body for creating or updating a user extension attribute.
// The ID field is not included; the target is specified via the URL path.
type RequestUserExtensionAttribute struct {
	XMLName     xml.Name                          `xml:"user_extension_attribute"`
	Name        string                            `xml:"name"`
	Description string                            `xml:"description"`
	DataType    string                            `xml:"data_type"`
	InputType   ResourceUserExtensionAttributeInputType `xml:"input_type"`
}
