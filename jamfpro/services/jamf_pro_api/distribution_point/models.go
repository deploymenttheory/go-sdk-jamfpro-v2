package distribution_point

// ResourceDistributionPoint represents a file share distribution point in Jamf Pro.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points-id
type ResourceDistributionPoint struct {
	ID                         string `json:"id,omitempty"`
	Name                       string `json:"name"`
	ServerName                 string `json:"serverName"`
	Principal                  bool   `json:"principal,omitempty"`
	BackupDistributionPointID  string `json:"backupDistributionPointId,omitempty"`
	SSHUsername                string `json:"sshUsername,omitempty"`
	LocalPathToShare           string `json:"localPathToShare,omitempty"`
	FileSharingConnectionType  string `json:"fileSharingConnectionType"` // AFP, SMB, NONE
	ShareName                  string `json:"shareName,omitempty"`
	Workgroup                  string `json:"workgroup,omitempty"`
	Port                       int    `json:"port,omitempty"`
	ReadWriteUsername          string `json:"readWriteUsername,omitempty"`
	ReadOnlyUsername           string `json:"readOnlyUsername,omitempty"`
	HTTPSEnabled               bool   `json:"httpsEnabled,omitempty"`
	HTTPSPort                  int    `json:"httpsPort,omitempty"`
	HTTPSContext               string `json:"httpsContext,omitempty"`
	HTTPSSecurityType          string `json:"httpsSecurityType,omitempty"` // USERNAME_PASSWORD, NONE
	HTTPSUsername              string `json:"httpsUsername,omitempty"`
	EnableLoadBalancing        bool   `json:"enableLoadBalancing,omitempty"`
}

// RequestDistributionPoint represents the request body for creating or updating a distribution point.
// Includes password fields that are write-only.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points
type RequestDistributionPoint struct {
	Name                       string `json:"name"`
	ServerName                 string `json:"serverName"`
	Principal                  bool   `json:"principal,omitempty"`
	BackupDistributionPointID  string `json:"backupDistributionPointId,omitempty"`
	SSHUsername                string `json:"sshUsername,omitempty"`
	SSHPassword                string `json:"sshPassword,omitempty"`
	LocalPathToShare           string `json:"localPathToShare,omitempty"`
	FileSharingConnectionType  string `json:"fileSharingConnectionType"` // AFP, SMB, NONE
	ShareName                  string `json:"shareName,omitempty"`
	Workgroup                  string `json:"workgroup,omitempty"`
	Port                       int    `json:"port,omitempty"`
	ReadWriteUsername          string `json:"readWriteUsername,omitempty"`
	ReadWritePassword          string `json:"readWritePassword,omitempty"`
	ReadOnlyUsername           string `json:"readOnlyUsername,omitempty"`
	ReadOnlyPassword           string `json:"readOnlyPassword,omitempty"`
	HTTPSEnabled               bool   `json:"httpsEnabled,omitempty"`
	HTTPSPort                  int    `json:"httpsPort,omitempty"`
	HTTPSContext               string `json:"httpsContext,omitempty"`
	HTTPSSecurityType          string `json:"httpsSecurityType,omitempty"` // USERNAME_PASSWORD, NONE
	HTTPSUsername              string `json:"httpsUsername,omitempty"`
	HTTPSPassword              string `json:"httpsPassword,omitempty"`
	EnableLoadBalancing        bool   `json:"enableLoadBalancing,omitempty"`
}

// ListResponse is the response for List operations.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points
type ListResponse struct {
	TotalCount int                          `json:"totalCount"`
	Results    []ResourceDistributionPoint  `json:"results"`
}

// CreateResponse represents the response after creating a distribution point.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// DeleteMultipleRequest represents the request body for deleting multiple distribution points.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points-delete-multiple
type DeleteMultipleRequest struct {
	IDs []string `json:"ids"`
}

// HistoryEntry represents a single history entry for a distribution point.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points-id-history
type HistoryEntry struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Date     string  `json:"date"`
	Note     string  `json:"note"`
	Details  *string `json:"details"` // Can be null
}

// HistoryListResponse is the response for history list operations.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points-id-history
type HistoryListResponse struct {
	TotalCount int            `json:"totalCount"`
	Results    []HistoryEntry `json:"results"`
}

// CreateHistoryNoteRequest represents the request body for creating a history note.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points-id-history
type CreateHistoryNoteRequest struct {
	Note string `json:"note"`
}
