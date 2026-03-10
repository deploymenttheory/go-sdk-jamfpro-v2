package jamf_protect

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"

// ResourceJamfProtectSettings represents the Jamf Protect integration settings.
type ResourceJamfProtectSettings struct {
	ID             string `json:"id"`
	ProtectURL     string `json:"protectUrl"`
	SyncStatus     string `json:"syncStatus"`
	APIClientID    string `json:"apiClientId"`
	AutoInstall    bool   `json:"autoInstall"`
	LastSyncTime   string `json:"lastSyncTime"`
	APIClientName  string `json:"apiClientName"`
	RegistrationID string `json:"registrationId"`
}

// RequestJamfProtectRegistration represents a registration request for Jamf Protect.
type RequestJamfProtectRegistration struct {
	ProtectURL string `json:"protectUrl"`
	ClientID   string `json:"clientId"`
	Password   string `json:"password"`
}

// RequestJamfProtectSettings represents a settings update request for Jamf Protect.
type RequestJamfProtectSettings struct {
	AutoInstall bool `json:"autoInstall"`
}

// ResourceJamfProtectPlan represents a Jamf Protect deployment plan.
type ResourceJamfProtectPlan struct {
	UUID             string `json:"uuid"`
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ProfileName      string `json:"profileName"`
	ProfileID        int    `json:"profileId"`
	ScopeDescription string `json:"scopeDescription"`
}

// ListResponseJamfProtectPlans represents a paginated list of Jamf Protect plans.
type ListResponseJamfProtectPlans struct {
	TotalCount int                        `json:"totalCount"`
	Results    []ResourceJamfProtectPlan  `json:"results"`
}

// ResourceJamfProtectDeploymentTask represents a Jamf Protect deployment task.
type ResourceJamfProtectDeploymentTask struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	Updated      string `json:"updated"`
	Version      string `json:"version"`
	ComputerID   string `json:"computerId"`
	ComputerName string `json:"computerName"`
}

// ListResponseJamfProtectDeploymentTasks represents a paginated list of deployment tasks.
type ListResponseJamfProtectDeploymentTasks struct {
	TotalCount int                                 `json:"totalCount"`
	Results    []ResourceJamfProtectDeploymentTask `json:"results"`
}

// RequestJamfProtectDeployment represents a request to deploy Jamf Protect.
type RequestJamfProtectDeployment struct {
	PlanID          string   `json:"planId"`
	TargetComputers []string `json:"targetComputers"`
}

// ResourceJamfProtectHistory represents a Jamf Protect history entry from list operations.
type ResourceJamfProtectHistory struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// ResourceJamfProtectHistoryCreate is an alias to the shared history note response.
type ResourceJamfProtectHistoryCreate = models.SharedHistoryNoteResponse

// ListResponseJamfProtectHistory represents a paginated list of Jamf Protect history entries.
type ListResponseJamfProtectHistory struct {
	TotalCount int                          `json:"totalCount"`
	Results    []ResourceJamfProtectHistory `json:"results"`
}

// RequestJamfProtectHistoryNote is an alias to the shared history note request.
type RequestJamfProtectHistoryNote = models.SharedHistoryNoteRequest
