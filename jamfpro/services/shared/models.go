package shared

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

// SharedHistoryItem represents a history entry from Jamf Pro API.
// All history endpoints return the same structure with integer IDs.
type SharedHistoryItem struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Date     string  `json:"date"`
	Note     string  `json:"note"`
	Details  *string `json:"details,omitempty"`
}

// SharedHistoryResponse represents a paginated history response.
type SharedHistoryResponse struct {
	TotalCount int                 `json:"totalCount"`
	Results    []SharedHistoryItem `json:"results"`
}

// SharedHistoryNoteRequest represents a request to add a history note.
type SharedHistoryNoteRequest struct {
	Note string `json:"note"`
}
