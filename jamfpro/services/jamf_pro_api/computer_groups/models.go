package computer_groups

// Criterion represents a smart group criterion.
type Criterion struct {
	Name       string `json:"name"`
	Priority   int    `json:"priority"`
	AndOr      string `json:"andOr"`
	SearchType string `json:"searchType"`
	Value      string `json:"value"`
}

// ResourceSmartGroup represents a smart computer group resource.
type ResourceSmartGroup struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	IsSmart  bool         `json:"isSmart"`
	Criteria []Criterion  `json:"criteria"`
}

// ResourceStaticGroup represents a static computer group resource.
type ResourceStaticGroup struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	IsSmart     bool     `json:"isSmart"`
	ComputerIds []string `json:"computerIds"`
}

// ListSmartResponse is the response for ListSmartGroupsV2.
type ListSmartResponse struct {
	TotalCount int                  `json:"totalCount"`
	Results    []ResourceSmartGroup `json:"results"`
}

// ListStaticResponse is the response for ListStaticGroupsV2.
type ListStaticResponse struct {
	TotalCount int                   `json:"totalCount"`
	Results    []ResourceStaticGroup `json:"results"`
}

// RequestSmartGroup is the body for creating and updating smart groups.
type RequestSmartGroup struct {
	Name     string      `json:"name"`
	Criteria []Criterion `json:"criteria"`
	Site     *SiteRef    `json:"site,omitempty"`
}

// SiteRef is an optional site reference.
type SiteRef struct {
	ID string `json:"id"`
}

// RequestStaticGroup is the body for creating and updating static groups.
type RequestStaticGroup struct {
	Name        string   `json:"name"`
	ComputerIds []string `json:"computerIds"`
}

// CreateSmartResponse is the response for CreateSmartGroupV2.
type CreateSmartResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CreateStaticResponse is the response for CreateStaticGroupV2.
type CreateStaticResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
