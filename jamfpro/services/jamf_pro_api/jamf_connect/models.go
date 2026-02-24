package jamf_connect

// ResourceJamfConnect represents the Jamf Connect settings.
type ResourceJamfConnect struct {
	ID             string `json:"id,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
	Description    string `json:"description,omitempty"`
	Enabled        bool   `json:"enabled"`
	Settings       string `json:"settings,omitempty"`
	Version        string `json:"version,omitempty"`
	LastModified   string `json:"lastModified,omitempty"`
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
}

// ResourceJamfConnectConfigProfile represents a Jamf Connect config profile.
type ResourceJamfConnectConfigProfile struct {
	UUID               string `json:"uuid"`
	ProfileID          int    `json:"profileId"`
	ProfileName        string `json:"profileName"`
	ScopeDescription   string `json:"scopeDescription"`
	SiteID             string `json:"siteId"`
	Version            string `json:"version"`
	AutoDeploymentType string `json:"autoDeploymentType"`
}

// ResourceJamfConnectConfigProfileUpdate represents the updateable fields for a Jamf Connect profile.
type ResourceJamfConnectConfigProfileUpdate struct {
	JamfConnectVersion string `json:"version,omitempty"`
	AutoDeploymentType string `json:"autoDeploymentType,omitempty"`
}

// ResourceJamfConnectTaskRetry represents the request structure for task retry.
type ResourceJamfConnectTaskRetry struct {
	IDs []string `json:"ids"`
}

// ListResponse is the response for ListConfigProfilesV1.
type ListResponse struct {
	TotalCount int                                `json:"totalCount"`
	Results    []ResourceJamfConnectConfigProfile `json:"results"`
}

// JamfConnectError represents a single error in the Jamf Connect error response.
type JamfConnectError struct {
	Code        string `json:"code"`
	Field       string `json:"field"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

// ErrorResponse represents the error response structure for Jamf Connect API errors.
type ErrorResponse struct {
	HTTPStatus int                `json:"httpStatus"`
	Errors     []JamfConnectError `json:"errors"`
}

// DeploymentTask represents a single Jamf Connect deployment task.
type DeploymentTask struct {
	Status  string `json:"status"`
	Updated string `json:"updated"`
	Version string `json:"version"`
}

// DeploymentTasksResponse is the response for GetDeploymentTasksByIDV1.
type DeploymentTasksResponse struct {
	TotalCount int              `json:"totalCount"`
	Results    []DeploymentTask `json:"results"`
}

// HistoryItem represents a single Jamf Connect history entry.
type HistoryItem struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// HistoryResponse is the response for GetHistoryV1.
type HistoryResponse struct {
	TotalCount int           `json:"totalCount"`
	Results    []HistoryItem `json:"results"`
}
