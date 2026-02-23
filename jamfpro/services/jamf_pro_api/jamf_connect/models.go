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
