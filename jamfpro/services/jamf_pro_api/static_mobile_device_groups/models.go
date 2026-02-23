package static_mobile_device_groups

// ResourceStaticMobileDeviceGroup represents a static mobile device group resource.
type ResourceStaticMobileDeviceGroup struct {
	ID          string `json:"groupId"`
	Name        string `json:"groupName"`
	Description string `json:"groupDescription"`
	SiteID      string `json:"siteId"`
	Count       int    `json:"count"`
}

// ListResponse is the response for List.
type ListResponse struct {
	TotalCount int                             `json:"totalCount"`
	Results    []ResourceStaticMobileDeviceGroup `json:"results"`
}

// RequestStaticMobileDeviceGroup is the body for creating and updating static mobile device groups.
type RequestStaticMobileDeviceGroup struct {
	Name        string                           `json:"groupName"`
	Description string                           `json:"groupDescription,omitempty"`
	Assignments []StaticMobileDeviceGroupAssignment `json:"assignments"`
	SiteID      string                           `json:"siteId"`
}

// StaticMobileDeviceGroupAssignment represents a device assignment for a static group.
type StaticMobileDeviceGroupAssignment struct {
	MobileDeviceID string `json:"mobileDeviceId"`
	Selected       bool   `json:"selected"`
}

// CreateResponse is the response for Create.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
