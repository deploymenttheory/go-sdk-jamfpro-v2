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
	ID          string               `json:"groupId" mapstructure:"groupId"`
	Name        string               `json:"groupName" mapstructure:"groupName"`
	Description string               `json:"groupDescription" mapstructure:"groupDescription"`
	SiteId      string               `json:"siteId" mapstructure:"siteId"`
	Count       int                  `json:"count" mapstructure:"count"`
	Criteria    []CriteriaJamfProAPI `json:"criteria,omitempty" mapstructure:"criteria"`
}

// ResourceStaticMobileDeviceGroup represents a static mobile device group resource.
type ResourceStaticMobileDeviceGroup struct {
	ID          string `json:"groupId" mapstructure:"groupId"`
	Name        string `json:"groupName" mapstructure:"groupName"`
	Description string `json:"groupDescription" mapstructure:"groupDescription"`
	SiteId      string `json:"siteId" mapstructure:"siteId"`
	Count       int    `json:"count" mapstructure:"count"`
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

// ResourceMobileDeviceGroupSummary represents a mobile device group in the list-all response.
// Used by ListAllV1 (GET /api/v1/mobile-device-groups).
type ResourceMobileDeviceGroupSummary struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsSmartGroup bool   `json:"isSmartGroup"`
}

// ResourceMobileDeviceMember represents a mobile device in group membership responses.
// Used by GetStaticGroupMembershipV1 and GetSmartGroupMembershipV1.
// Note: as of Jamf Pro 11.25, the "personalDeviceProfileCurrent" field was removed from
// the filter and sort options for GET /v1/mobile-device-groups/static-group-membership/{id}.
type ResourceMobileDeviceMember struct {
	MobileDeviceID   string `json:"mobileDeviceId"`
	UDID             string `json:"udid"`
	DisplayName      string `json:"displayName"`
	SerialNumber     string `json:"serialNumber"`
	WifiMacAddress   string `json:"wifiMacAddress"`
	DevicePhoneNumber string `json:"devicePhoneNumber"`
	Model            string `json:"model"`
	ModelIdentifier  string `json:"modelIdentifier"`
	OSVersion        string `json:"osVersion"`
	Managed          bool   `json:"managed"`
	Supervised       bool   `json:"supervised"`
}

// GroupMembershipResponse is the response for GetStaticGroupMembershipV1 and GetSmartGroupMembershipV1.
type GroupMembershipResponse struct {
	TotalCount int                        `json:"totalCount"`
	Results    []ResourceMobileDeviceMember `json:"results"`
}

// RequestEraseDevices is the optional request body for EraseDevicesByGroupIDV1.
type RequestEraseDevices struct {
	PreserveDataPlan        *bool `json:"preserveDataPlan,omitempty"`
	DisallowProximitySetup  *bool `json:"disallowProximitySetup,omitempty"`
	ClearActivationLock     *bool `json:"clearActivationLock,omitempty"`
	ReturnToService        *bool `json:"returnToService,omitempty"`
}
