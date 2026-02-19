package bookmarks

// ResourceBookmark represents a bookmark resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-bookmarks-id
type ResourceBookmark struct {
	ID               string `json:"id"`
	SiteID           string `json:"siteId"`
	Priority         int    `json:"priority"`
	DisplayInBrowser *bool  `json:"displayInBrowser"`
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	ScopeDescription string `json:"scopeDescription,omitempty"`
	URL              string `json:"url"`
	IconID           string `json:"iconId"`
}

// ListResponse is the response for ListV1.
type ListResponse struct {
	TotalCount int                `json:"totalCount"`
	Results    []ResourceBookmark `json:"results"`
}

// CreateResponse is the response for CreateV1.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
