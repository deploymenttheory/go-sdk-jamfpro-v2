package static_computer_groups

// RequestStaticGroup represents the request structure for creating or updating a static computer group.
// Assignments are computer IDs to include in the group.
type RequestStaticGroup struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Assignments []string `json:"assignments"`
	SiteID      *string  `json:"siteId,omitempty"`
}

// ResourceStaticGroup represents a static computer group in list and get responses.
type ResourceStaticGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SiteID      string `json:"siteId"`
	Count       int    `json:"count"`
}

// ListResponse is the paginated response for ListV2.
type ListResponse struct {
	TotalCount int                  `json:"totalCount"`
	Results    []ResourceStaticGroup `json:"results"`
}

// CreateResponse is the response for CreateV2.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
