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
