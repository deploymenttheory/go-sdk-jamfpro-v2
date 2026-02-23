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
