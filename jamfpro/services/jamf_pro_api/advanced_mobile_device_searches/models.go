package advanced_mobile_device_searches

// CriteriaJamfProAPI represents a criterion for an advanced mobile device search (Jamf Pro API shape).
type CriteriaJamfProAPI struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen *bool  `json:"openingParen,omitempty"`
	ClosingParen *bool  `json:"closingParen,omitempty"`
}

// ResourceAdvancedMobileDeviceSearch represents a single advanced mobile device search.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches-id
type ResourceAdvancedMobileDeviceSearch struct {
	ID            string               `json:"id,omitempty"`
	Name          string               `json:"name"`
	Criteria      []CriteriaJamfProAPI `json:"criteria"`
	DisplayFields []string             `json:"displayFields"`
	SiteId        *string              `json:"siteId,omitempty"`
}

// ListResponse is the response for ListV1.
type ListResponse struct {
	TotalCount int                               `json:"totalCount"`
	Results    []ResourceAdvancedMobileDeviceSearch `json:"results"`
}

// CreateResponse is the response for CreateV1.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ChoicesResponse is the response for GetChoicesV1.
type ChoicesResponse struct {
	Choices []string `json:"choices"`
}
