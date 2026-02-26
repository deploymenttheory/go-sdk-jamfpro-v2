package ebooks

// ResourceEbook represents an ebook resource (Jamf Pro API v1).
// Aligned with Jamf Pro API v1 ebooks schema.
type ResourceEbook struct {
	ID                 string `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	Kind               string `json:"kind,omitempty"` // UNKNOWN, PDF, EPUB, IBOOKS
	URL                string `json:"url,omitempty"`
	Free               bool   `json:"free"`
	Version            string `json:"version,omitempty"`
	Author             string `json:"author,omitempty"`
	DeployAsManaged    bool   `json:"deployAsManaged"`
	InstallAutomatically bool `json:"installAutomatically"`
	CategoryID         string `json:"categoryId,omitempty"`
	SiteID             string `json:"siteId,omitempty"`
}

// ListResponse is the response for ListV1.
type ListResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []ResourceEbook `json:"results"`
}

// ResourceScope represents the scope of an ebook (assignments and exclusions).
// Returned by GET /api/v1/ebooks/{id}/scope.
type ResourceScope struct {
	AllComputers         bool                    `json:"allComputers"`
	AllMobileDevices     bool                    `json:"allMobileDevices"`
	AllUsers             bool                    `json:"allUsers"`
	ComputerIDs          []string                `json:"computerIds,omitempty"`
	ComputerGroupIDs     []string                `json:"computerGroupIds,omitempty"`
	MobileDeviceIDs      []string                `json:"mobileDeviceIds,omitempty"`
	MobileDeviceGroupIDs []string                `json:"mobileDeviceGroupIds,omitempty"`
	BuildingIDs          []string                `json:"buildingIds,omitempty"`
	DepartmentIDs        []string                `json:"departmentIds,omitempty"`
	UserIDs              []string                `json:"userIds,omitempty"`
	UserGroupIDs         []string                `json:"userGroupIds,omitempty"`
	ClassroomIDs         []string                `json:"classroomIds,omitempty"`
	Limitations          *ScopeLimitations       `json:"limitations,omitempty"`
	Exclusions           *ScopeExclusions        `json:"exclusions,omitempty"`
}

// ScopeLimitations represents limitations within the scope.
type ScopeLimitations struct {
	NetworkSegments []string           `json:"networkSegments,omitempty"`
	Users           []ScopeLimitationUser `json:"users,omitempty"`
	UserGroups      []string           `json:"userGroups,omitempty"`
}

// ScopeLimitationUser represents a user in scope limitations.
type ScopeLimitationUser struct {
	Name string `json:"name,omitempty"`
}

// ScopeExclusions represents exclusions within the scope.
type ScopeExclusions struct {
	ComputerIDs          []string           `json:"computerIds,omitempty"`
	ComputerGroupIDs     []string           `json:"computerGroupIds,omitempty"`
	MobileDeviceIDs      []string           `json:"mobileDeviceIds,omitempty"`
	MobileDeviceGroupIDs []string           `json:"mobileDeviceGroupIds,omitempty"`
	BuildingIDs          []string           `json:"buildingIds,omitempty"`
	DepartmentIDs        []string           `json:"departmentIds,omitempty"`
	UserIDs              []string           `json:"userIds,omitempty"`
	UserGroupIDs         []string           `json:"userGroupIds,omitempty"`
	Limitations          *ScopeLimitations  `json:"limitations,omitempty"`
}
