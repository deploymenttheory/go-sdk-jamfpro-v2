package smart_mobile_device_groups

// SharedSubsetCriteriaJamfProAPI represents the criteria for a search item in the Jamf Pro API.
// Used by smart mobile device groups and other smart group resources.
type SharedSubsetCriteriaJamfProAPI struct {
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
	GroupID          string                           `json:"groupId"`
	GroupName        string                           `json:"groupName"`
	GroupDescription string                           `json:"groupDescription"`
	SiteID           string                           `json:"siteId"`
	Count            int                              `json:"count"`
	Criteria         []SharedSubsetCriteriaJamfProAPI `json:"criteria,omitempty"`
}

// ListItem represents an individual smart mobile device group in list responses.
type ListItem struct {
	GroupID          string `json:"groupId"`
	GroupName        string `json:"groupName"`
	GroupDescription string `json:"groupDescription"`
	SiteID           string `json:"siteId"`
	Count            int    `json:"count"`
}

// ListResponse is the response for List.
type ListResponse struct {
	TotalCount int        `json:"totalCount"`
	Results    []ListItem `json:"results"`
}

// RequestSmartMobileDeviceGroup is the body for creating and updating smart mobile device groups.
type RequestSmartMobileDeviceGroup struct {
	GroupName        string                           `json:"groupName"`
	GroupDescription string                           `json:"groupDescription,omitempty"`
	Criteria         []SharedSubsetCriteriaJamfProAPI `json:"criteria,omitempty"`
	SiteId           *string                          `json:"siteId,omitempty"`
}

// CreateResponse is the response for Create.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// MembershipResponse is the response for GetMembership.
type MembershipResponse struct {
	TotalCount int                      `json:"totalCount"`
	Results    []InventoryListMobileDeviceItem `json:"results"`
}

// InventoryListMobileDeviceItem represents a mobile device in the membership/inventory list.
type InventoryListMobileDeviceItem struct {
	MobileDeviceId     string `json:"mobileDeviceId,omitempty"`
	Udid               string `json:"udid,omitempty"`
	DisplayName        string `json:"displayName,omitempty"`
	Model              string `json:"model,omitempty"`
	ModelIdentifier    string `json:"modelIdentifier,omitempty"`
	SerialNumber       string `json:"serialNumber,omitempty"`
	OsVersion          string `json:"osVersion,omitempty"`
	Managed            bool   `json:"managed,omitempty"`
	Supervised         bool   `json:"supervised,omitempty"`
	DeviceOwnershipType string `json:"deviceOwnershipType,omitempty"`
}
