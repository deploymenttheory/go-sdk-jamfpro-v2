package smart_computer_groups

// SubsetCriteria represents a criterion for a smart computer group search.
// Shared structure used across Jamf Pro API smart group endpoints.
type SubsetCriteria struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen *bool  `json:"openingParen,omitempty"`
	ClosingParen *bool  `json:"closingParen,omitempty"`
}

// RequestSmartGroup represents the request body for creating or updating a smart computer group.
type RequestSmartGroup struct {
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Criteria    []SubsetCriteria `json:"criteria,omitempty"`
	SiteId      *string         `json:"siteId,omitempty"`
}

// ResourceSmartGroup represents a smart computer group resource (full detail from GET by ID).
type ResourceSmartGroup struct {
	ID              string          `json:"id"`
	SiteId          string          `json:"siteId"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	MembershipCount int             `json:"membershipCount"`
	Criteria        []SubsetCriteria `json:"criteria,omitempty"`
}

// ListItem represents a smart computer group in a list response.
type ListItem struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	SiteID          string `json:"siteId"`
	MembershipCount int    `json:"membershipCount"`
}

// ListResponse represents the paginated response for listing smart computer groups.
type ListResponse struct {
	TotalCount int        `json:"totalCount"`
	Results    []ListItem `json:"results"`
}

// MembershipResponse represents the membership response for a smart computer group.
type MembershipResponse struct {
	Members []int `json:"members"`
}

// CreateResponse represents the response from creating a smart computer group.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
