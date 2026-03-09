package printers

import "encoding/xml"

// ResourcePrinter represents a Jamf Pro Classic API printer resource.
type ResourcePrinter struct {
	XMLName     xml.Name `xml:"printer"`
	ID          int      `xml:"id,omitempty"`
	Name        string   `xml:"name,omitempty"`
	Category    string   `xml:"category,omitempty"`
	URI         string   `xml:"uri,omitempty"`
	CUPSName    string   `xml:"CUPS_name,omitempty"`
	Location    string   `xml:"location,omitempty"`
	Model       string   `xml:"model,omitempty"`
	Info        string   `xml:"info,omitempty"`
	Notes       string   `xml:"notes,omitempty"`
	MakeDefault bool     `xml:"make_default"`
	UseGeneric  bool     `xml:"use_generic"`
	PPD         string   `xml:"ppd,omitempty"`
	PPDPath     string   `xml:"ppd_path,omitempty"`
	PPDContents string   `xml:"ppd_contents,omitempty"`
}

// ListItemPrinter is the slim representation returned in list responses.
type ListItemPrinter struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListPrinters (GET /JSSResource/printers).
type ListResponse struct {
	XMLName xml.Name          `xml:"printers"`
	Size    int               `xml:"size"`
	Results []ListItemPrinter `xml:"printer"`
}

// CreateUpdateResponse is the response body for create and update operations,
// which return only the assigned resource ID.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"printer"`
	ID      int      `xml:"id"`
}

// RequestPrinter is the body for creating or updating a printer.
// The ID field is not included; the target is specified via the URL path.
type RequestPrinter struct {
	XMLName     xml.Name `xml:"printer"`
	Name        string   `xml:"name"`
	Category    string   `xml:"category,omitempty"`
	URI         string   `xml:"uri,omitempty"`
	CUPSName    string   `xml:"CUPS_name,omitempty"`
	Location    string   `xml:"location,omitempty"`
	Model       string   `xml:"model,omitempty"`
	Info        string   `xml:"info,omitempty"`
	Notes       string   `xml:"notes,omitempty"`
	MakeDefault bool     `xml:"make_default"`
	UseGeneric  bool     `xml:"use_generic"`
	PPD         string   `xml:"ppd,omitempty"`
	PPDPath     string   `xml:"ppd_path,omitempty"`
	PPDContents string   `xml:"ppd_contents,omitempty"`
}
