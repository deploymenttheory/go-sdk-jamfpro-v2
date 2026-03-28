package sites

// ResourceSite represents a site resource.
type ResourceSite struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ResourceSiteObject represents an object within a site.
// Note: as of Jamf Pro 11.25, the "BYO Profile" object type was removed from the
// GET /v1/sites/{id}/objects filter options.
type ResourceSiteObject struct {
	SiteID     string `json:"siteId"`
	ObjectType string `json:"objectType"`
	ObjectID   string `json:"objectId"`
}

// ObjectsListResponse is the response for GetObjectsByIDV1 (paginated).
type ObjectsListResponse struct {
	TotalCount int                  `json:"totalCount"`
	Results    []ResourceSiteObject  `json:"results"`
}
