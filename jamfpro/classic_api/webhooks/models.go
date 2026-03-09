package webhooks

import "encoding/xml"

// ResourceWebhook represents a Jamf Pro Classic API webhook resource.
type ResourceWebhook struct {
	XMLName                     xml.Name       `xml:"webhook"`
	ID                          int            `xml:"id,omitempty"`
	Name                        string         `xml:"name,omitempty"`
	Enabled                     bool           `xml:"enabled"`
	URL                         string         `xml:"url,omitempty"`
	ContentType                 string         `xml:"content_type,omitempty"`
	Event                       string         `xml:"event,omitempty"`
	ConnectionTimeout           int            `xml:"connection_timeout,omitempty"`
	ReadTimeout                 int            `xml:"read_timeout,omitempty"`
	AuthenticationType          string         `xml:"authentication_type,omitempty"`
	Username                    string         `xml:"username,omitempty"`
	Password                    string         `xml:"password,omitempty"`
	EnableDisplayFieldsForGroup bool           `xml:"enable_display_fields_for_group_object,omitempty"`
	DisplayFields               []DisplayField `xml:"display_fields>display_field,omitempty"`
	Header                      string         `xml:"header,omitempty"`
	SmartGroupID                int            `xml:"smart_group_id,omitempty"`
}

// DisplayField represents a single display field included in webhook requests.
type DisplayField struct {
	Name string `xml:"name"`
}

// ListItemWebhook is the slim representation returned in list responses.
type ListItemWebhook struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListWebhooks (GET /JSSResource/webhooks).
type ListResponse struct {
	XMLName xml.Name          `xml:"webhooks"`
	Size    int               `xml:"size"`
	Results []ListItemWebhook `xml:"webhook"`
}

// RequestWebhook is the body for creating or updating a webhook.
// The ID field is not included; the target is specified via the URL path.
type RequestWebhook struct {
	XMLName                     xml.Name       `xml:"webhook"`
	Name                        string         `xml:"name"`
	Enabled                     bool           `xml:"enabled"`
	URL                         string         `xml:"url,omitempty"`
	ContentType                 string         `xml:"content_type,omitempty"`
	Event                       string         `xml:"event,omitempty"`
	ConnectionTimeout           int            `xml:"connection_timeout,omitempty"`
	ReadTimeout                 int            `xml:"read_timeout,omitempty"`
	AuthenticationType          string         `xml:"authentication_type,omitempty"`
	Username                    string         `xml:"username,omitempty"`
	Password                    string         `xml:"password,omitempty"`
	EnableDisplayFieldsForGroup bool           `xml:"enable_display_fields_for_group_object,omitempty"`
	DisplayFields               []DisplayField `xml:"display_fields>display_field,omitempty"`
	Header                      string         `xml:"header,omitempty"`
	SmartGroupID                int            `xml:"smart_group_id,omitempty"`
}
