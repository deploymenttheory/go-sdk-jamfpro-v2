package advanced_user_content_searches

// CriteriaJamfProAPI represents a criterion for an advanced user content search (Jamf Pro API shape).
type CriteriaJamfProAPI struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority,omitempty"` // optional
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen *bool  `json:"openingParen,omitempty"`
	ClosingParen *bool  `json:"closingParen,omitempty"`
}

// ResourceAdvancedUserContentSearch represents a single advanced user content search.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-user-content-searches-id
type ResourceAdvancedUserContentSearch struct {
	ID            string               `json:"id,omitempty"`
	Name          string               `json:"name"`
	Criteria      []CriteriaJamfProAPI `json:"criteria,omitempty"`      // optional
	DisplayFields []string             `json:"displayFields,omitempty"` // optional
	SiteId        *string              `json:"siteId,omitempty"`
}

// ListResponse is the response for ListV1.
type ListResponse struct {
	TotalCount int                                 `json:"totalCount"`
	Results    []ResourceAdvancedUserContentSearch `json:"results"`
}

// CreateResponse is the response for CreateV1.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
