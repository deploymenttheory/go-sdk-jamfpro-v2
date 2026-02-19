package mobile_device_groups

// CriteriaJamfProAPI represents a criterion for smart group search (Jamf Pro API shape).
// Same structure as SharedSubsetCriteriaJamfProAPI used in smart mobile device groups.
type CriteriaJamfProAPI struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen *bool  `json:"openingParen,omitempty"`
	ClosingParen *bool  `json:"closingParen,omitempty"`
}

// ResourceSmartMobileDeviceGroup represents a smart mobile device group resource.
type ResourceSmartMobileDeviceGroup struct {
	ID          string              `json:"groupId"`
	Name        string              `json:"groupName"`
	Description string              `json:"groupDescription"`
	SiteId      string              `json:"siteId"`
	Count       int                 `json:"count"`
	Criteria    []CriteriaJamfProAPI `json:"criteria,omitempty"`
}

// ResourceStaticMobileDeviceGroup represents a static mobile device group resource.
type ResourceStaticMobileDeviceGroup struct {
	ID          string `json:"groupId"`
	Name        string `json:"groupName"`
	Description string `json:"groupDescription"`
	SiteId      string `json:"siteId"`
	Count       int    `json:"count"`
}

// ListSmartResponse is the response for ListSmartV1.
type ListSmartResponse struct {
	TotalCount int                            `json:"totalCount"`
	Results    []ResourceSmartMobileDeviceGroup `json:"results"`
}

// ListStaticResponse is the response for ListStaticV1.
type ListStaticResponse struct {
	TotalCount int                             `json:"totalCount"`
	Results    []ResourceStaticMobileDeviceGroup `json:"results"`
}

// RequestSmartMobileDeviceGroup is the body for creating and updating smart mobile device groups.
type RequestSmartMobileDeviceGroup struct {
	Name        string              `json:"groupName"`
	Description string              `json:"groupDescription,omitempty"`
	Criteria    []CriteriaJamfProAPI `json:"criteria,omitempty"`
	SiteId      *string             `json:"siteId,omitempty"`
}

// RequestStaticMobileDeviceGroup is the body for creating and updating static mobile device groups.
type RequestStaticMobileDeviceGroup struct {
	Name        string                           `json:"groupName"`
	Description string                           `json:"groupDescription,omitempty"`
	Assignments []StaticMobileDeviceGroupAssignment `json:"assignments"`
	SiteId      string                           `json:"siteId"`
}

// StaticMobileDeviceGroupAssignment represents a device assignment for a static group.
type StaticMobileDeviceGroupAssignment struct {
	MobileDeviceID string `json:"mobileDeviceId"`
	Selected       bool   `json:"selected"`
}

// CreateSmartResponse is the response for CreateSmartV1.
type CreateSmartResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CreateStaticResponse is the response for CreateStaticV1.
type CreateStaticResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
