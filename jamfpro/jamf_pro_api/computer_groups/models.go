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
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	IsSmart  bool        `json:"isSmart"`
	Criteria []Criterion `json:"criteria"`
}

// ResourceStaticGroup represents a static computer group resource.
type ResourceStaticGroup struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	IsSmart     bool     `json:"isSmart"`
	ComputerIds []string `json:"computerIds"`
}

// ListSmartResponse is the response for ListSmartV2.
type ListSmartResponse struct {
	TotalCount int                  `json:"totalCount"`
	Results    []ResourceSmartGroup `json:"results"`
}

// ListStaticResponse is the response for ListStaticV2.
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
// ComputerIds is omitted when empty so create (POST) accepts name-only; API may not accept computerIds on create.
type RequestStaticGroup struct {
	Name        string   `json:"name"`
	ComputerIds []string `json:"computerIds,omitempty"`
}

// CreateSmartResponse is the response for CreateSmartV2.
type CreateSmartResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CreateStaticResponse is the response for CreateStaticV2.
type CreateStaticResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResourceGroupV1 represents a computer group in the v1 list response.
// Used by ListAllV1 (GET /api/v1/computer-groups).
type ResourceGroupV1 struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SmartGroup  bool   `json:"smartGroup"`
}

// SmartGroupMembershipResponse is the response for GetSmartGroupMembershipByIDV2.
// Contains the computer IDs that are members of the smart group.
type SmartGroupMembershipResponse struct {
	Members []int `json:"members"`
}

// -----------------------------------------------------------------------------
// V3 models (Jamf Pro 11.28 computer-groups v3 surface).
//
// The v3 smart-group criteria use the ComputerSmartGroupCriteriaV2 shape, which
// adds openingParen/closingParen to the legacy criterion.
// -----------------------------------------------------------------------------

// CriterionV3 is a smart computer group criterion in the v3 (ComputerSmartGroupCriteriaV2) shape.
type CriterionV3 struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen bool   `json:"openingParen"`
	ClosingParen bool   `json:"closingParen"`
}

// ResourceSmartGroupV3 is a smart computer group as returned by the v3 endpoints.
// The list response omits criteria; GetSmartByIDV3 includes them.
type ResourceSmartGroupV3 struct {
	ID              string        `json:"id,omitempty"`
	Name            string        `json:"name"`
	Description     string        `json:"description,omitempty"`
	SiteID          string        `json:"siteId,omitempty"`
	MembershipCount int           `json:"membershipCount,omitempty"`
	Criteria        []CriterionV3 `json:"criteria,omitempty"`
}

// ListSmartV3Response is the response for ListSmartV3.
type ListSmartV3Response struct {
	TotalCount int                    `json:"totalCount"`
	Results    []ResourceSmartGroupV3 `json:"results"`
}

// RequestSmartGroupV3 is the SmartComputerGroupV3 create/update body.
type RequestSmartGroupV3 struct {
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	SiteID      string        `json:"siteId,omitempty"`
	Criteria    []CriterionV3 `json:"criteria"`
}

// ResourceStaticGroupV3 is a static computer group as returned by the v3 list/get endpoints.
type ResourceStaticGroupV3 struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	SiteID      string `json:"siteId,omitempty"`
	Count       int    `json:"count,omitempty"`
}

// ListStaticV3Response is the response for ListStaticV3.
type ListStaticV3Response struct {
	TotalCount int                     `json:"totalCount"`
	Results    []ResourceStaticGroupV3 `json:"results"`
}

// RequestStaticGroupV3 is the static computer group create/update body.
// Assignments is the set of computer IDs to assign; the API treats it as a set
// (uniqueItems), and the SDK deduplicates it before sending.
type RequestStaticGroupV3 struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	SiteID      string `json:"siteId,omitempty"`
	// Assignments must always be serialized (the API rejects a missing
	// assignments key with a 500); the create/update methods normalise a nil
	// slice to an empty array.
	Assignments []string `json:"assignments"`
}
