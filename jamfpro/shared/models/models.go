package models

// SharedResourceSite represents a site reference used across Classic API resources.
type SharedResourceSite struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// SharedResourceCategory represents a category reference used across Classic API resources.
type SharedResourceCategory struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// SharedResourceSelfServiceIcon represents a self-service icon used across Classic API resources.
type SharedResourceSelfServiceIcon struct {
	ID       int    `xml:"id,omitempty"`
	URI      string `xml:"uri,omitempty"`
	Data     string `xml:"data,omitempty"`
	Filename string `xml:"filename,omitempty"`
}

// SharedResourceSelfServiceCategory represents a self-service category used across Classic API resources.
type SharedResourceSelfServiceCategory struct {
	ID       int    `xml:"id,omitempty"`
	Name     string `xml:"name,omitempty"`
	Priority int    `xml:"priority,omitempty"`
}

// SharedSubsetCriteria represents search criteria used across Classic API resources.
// Used by computer groups, mobile device groups, and other smart group resources.
type SharedSubsetCriteria struct {
	Name         string `xml:"name,omitempty"`
	Priority     int    `xml:"priority,omitempty"`
	AndOr        string `xml:"and_or,omitempty"`
	SearchType   string `xml:"search_type,omitempty"`
	Value        string `xml:"value,omitempty"`
	OpeningParen bool   `xml:"opening_paren,omitempty"`
	ClosingParen bool   `xml:"closing_paren,omitempty"`
}

// SharedHistoryItem represents a history entry from Jamf Pro API with integer IDs.
type SharedHistoryItem struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Date     string  `json:"date"`
	Note     string  `json:"note"`
	Details  *string `json:"details,omitempty"`
}

// SharedHistoryResponse represents a paginated history response with integer IDs.
type SharedHistoryResponse struct {
	TotalCount int                 `json:"totalCount"`
	Results    []SharedHistoryItem `json:"results"`
}

// SharedHistoryItemString represents a history entry from Jamf Pro API with string IDs.
type SharedHistoryItemString struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// SharedHistoryResponseString represents a paginated history response with string IDs.
type SharedHistoryResponseString struct {
	TotalCount int                       `json:"totalCount"`
	Results    []SharedHistoryItemString `json:"results"`
}

// SharedHistoryNoteRequest represents a request to add a history note.
// The Details field is optional and only used by some services.
type SharedHistoryNoteRequest struct {
	Note    string `json:"note"`
	Details string `json:"details,omitempty"`
}

// SharedHistoryNoteResponse represents a response after adding a history note.
// ID is always string to handle both int and string responses from the API.
// Fields are optional to accommodate different response formats across services.
type SharedHistoryNoteResponse struct {
	ID       string `json:"id"`
	Href     string `json:"href,omitempty"`
	Username string `json:"username,omitempty"`
	Date     string `json:"date,omitempty"`
	Note     string `json:"note,omitempty"`
	Details  string `json:"details,omitempty"`
}
