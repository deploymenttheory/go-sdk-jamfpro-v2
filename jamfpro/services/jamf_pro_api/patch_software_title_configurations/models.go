package patch_software_title_configurations

// ListResponse represents a list of patch software title configurations.
type ListResponse []ResourcePatchSoftwareTitleConfiguration

// ResourcePatchSoftwareTitleConfiguration represents a patch software title configuration resource.
type ResourcePatchSoftwareTitleConfiguration struct {
	ID                     string                                                    `json:"id,omitempty"`
	DisplayName            string                                                    `json:"displayName"`
	SoftwareTitleID        string                                                    `json:"softwareTitleId"`
	CategoryID             string                                                    `json:"categoryId,omitempty"`
	SiteID                 string                                                    `json:"siteId,omitempty"`
	UINotifications        bool                                                      `json:"uiNotifications,omitempty"`
	EmailNotifications     bool                                                      `json:"emailNotifications,omitempty"`
	ExtensionAttributes    []SubsetExtensionAttribute                                `json:"extensionAttributes,omitempty"`
	SoftwareTitleName      string                                                    `json:"softwareTitleName,omitempty"`
	SoftwareTitleNameID    string                                                    `json:"softwareTitleNameId,omitempty"`
	SoftwareTitlePublisher string                                                    `json:"softwareTitlePublisher,omitempty"`
	JamfOfficial           bool                                                      `json:"jamfOfficial,omitempty"`
	PatchSourceName        string                                                    `json:"patchSourceName,omitempty"`
	PatchSourceEnabled     bool                                                      `json:"patchSourceEnabled,omitempty"`
	Packages               []SubsetPackage                                           `json:"packages,omitempty"`
}

// SubsetExtensionAttribute represents an extension attribute in a patch software title configuration.
type SubsetExtensionAttribute struct {
	Accepted bool   `json:"accepted,omitempty"`
	EAID     string `json:"eaId,omitempty"`
}

// SubsetPackage represents a package in a patch software title configuration.
type SubsetPackage struct {
	PackageID   string `json:"packageId,omitempty"`
	Version     string `json:"version,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// CreateResponse represents the response when creating a patch software title configuration.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResourceDashboardStatus represents whether a software title configuration is on the dashboard.
type ResourceDashboardStatus struct {
	OnDashboard bool `json:"onDashboard"`
}

// ResourceDefinition represents a patch software title definition.
type ResourceDefinition struct {
	Version                 string           `json:"version"`
	MinimumOperatingSystem  string           `json:"minimumOperatingSystem,omitempty"`
	ReleaseDate             string           `json:"releaseDate,omitempty"`
	RebootRequired          bool             `json:"rebootRequired,omitempty"`
	KillApps                []ResourceKillApp `json:"killApps,omitempty"`
	Standalone              bool             `json:"standalone,omitempty"`
	AbsoluteOrderID         string           `json:"absoluteOrderId,omitempty"`
}

// ResourceKillApp represents an app to kill during patch installation.
type ResourceKillApp struct {
	AppName string `json:"appName,omitempty"`
}

// DefinitionsResponse is the paginated response for GetDefinitionsByIDV2.
type DefinitionsResponse struct {
	TotalCount int                  `json:"totalCount"`
	Results    []ResourceDefinition `json:"results"`
}

// ResourceDependency represents a patch software title configuration dependency (smart group).
type ResourceDependency struct {
	SmartGroupID   string `json:"smartGroupId,omitempty"`
	SmartGroupName string `json:"smartGroupName,omitempty"`
}

// DependenciesResponse is the paginated response for GetDependenciesByIDV2.
type DependenciesResponse struct {
	TotalCount int                    `json:"totalCount"`
	Results    []ResourceDependency   `json:"results"`
}

// ResourceExtensionAttribute represents an extension attribute for a software title.
type ResourceExtensionAttribute struct {
	EAID           string `json:"eaId,omitempty"`
	Accepted       bool   `json:"accepted,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
	ScriptContents string `json:"scriptContents,omitempty"`
}

// ResourcePatchReportItem represents a single item in the patch report.
type ResourcePatchReportItem struct {
	ComputerName           string `json:"computerName,omitempty"`
	DeviceID               string `json:"deviceId,omitempty"`
	Username               string `json:"username,omitempty"`
	OperatingSystemVersion string `json:"operatingSystemVersion,omitempty"`
	LastContactTime        string `json:"lastContactTime,omitempty"`
	BuildingName           string `json:"buildingName,omitempty"`
	DepartmentName         string `json:"departmentName,omitempty"`
	SiteName               string `json:"siteName,omitempty"`
	Version                string `json:"version,omitempty"`
}

// PatchReportResponse is the paginated response for GetPatchReportByIDV2.
type PatchReportResponse struct {
	TotalCount int                      `json:"totalCount"`
	Results    []ResourcePatchReportItem `json:"results"`
}

// ResourcePatchSummary represents the patch summary for a software title configuration.
type ResourcePatchSummary struct {
	SoftwareTitleID              string `json:"softwareTitleId,omitempty"`
	Title                        string `json:"title,omitempty"`
	LatestVersion                string `json:"latestVersion,omitempty"`
	ReleaseDate                  string `json:"releaseDate,omitempty"`
	UpToDate                     int    `json:"upToDate,omitempty"`
	OutOfDate                    int    `json:"outOfDate,omitempty"`
	OnDashboard                  bool   `json:"onDashboard,omitempty"`
	SoftwareTitleConfigurationID string `json:"softwareTitleConfigurationId,omitempty"`
}

// ResourceHistoryItem represents a history entry for a patch software title configuration.
type ResourceHistoryItem struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Date     string `json:"date,omitempty"`
	Note     string `json:"note,omitempty"`
	Details  string `json:"details,omitempty"`
}

// HistoryResponse is the paginated response for GetHistoryByIDV2.
type HistoryResponse struct {
	TotalCount int                   `json:"totalCount"`
	Results    []ResourceHistoryItem `json:"results"`
}

// RequestAddHistoryNote is the request body for AddHistoryNoteByIDV2.
type RequestAddHistoryNote struct {
	Note string `json:"note"`
}

// ResponseAddHistoryNote is the response for AddHistoryNoteByIDV2.
type ResponseAddHistoryNote struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResourcePatchVersion represents a patch version.
type ResourcePatchVersion struct {
	AbsoluteOrderID string `json:"absoluteOrderId,omitempty"`
	Version         string `json:"version,omitempty"`
	OnVersion       int    `json:"onVersion,omitempty"`
}
