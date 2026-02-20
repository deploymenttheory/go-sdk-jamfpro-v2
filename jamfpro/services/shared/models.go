package shared

// SharedResourceSite represents a site reference used across Classic API resources.
type SharedResourceSite struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
