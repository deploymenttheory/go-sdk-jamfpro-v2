package patch_policies

// ResourcePatchPolicy represents a patch policy resource returned by the Jamf Pro API.
type ResourcePatchPolicy struct {
	ID                           string `json:"id"`
	Name                         string `json:"name"`
	Enabled                      bool   `json:"enabled"`
	TargetPatchVersion           string `json:"targetPatchVersion"`
	DeploymentMethod             string `json:"deploymentMethod"`
	SoftwareTitleId              string `json:"softwareTitleId"`
	SoftwareTitleConfigurationId string `json:"softwareTitleConfigurationId"`
	KillAppsDelayMinutes         int    `json:"killAppsDelayMinutes"`
	KillAppsMessage              string `json:"killAppsMessage"`
	Downgrade                    bool   `json:"downgrade"`
	PatchUnknownVersion          bool   `json:"patchUnknownVersion"`
	NotificationHeader           string `json:"notificationHeader"`
	SelfServiceEnforceDeadline   bool   `json:"selfServiceEnforceDeadline"`
	SelfServiceDeadline          int    `json:"selfServiceDeadline"`
	InstallButtonText            string `json:"installButtonText"`
	SelfServiceDescription       string `json:"selfServiceDescription"`
	IconId                       string `json:"iconId"`
	ReminderFrequency            int    `json:"reminderFrequency"`
	ReminderEnabled              bool   `json:"reminderEnabled"`
}

// ListResponse is the paginated response for ListV2.
type ListResponse struct {
	TotalCount int                   `json:"totalCount"`
	Results    []ResourcePatchPolicy `json:"results"`
}

// ResourcePatchPolicySummary represents a patch policy summary (from GET /api/v2/patch-policies).
// This is a lighter-weight representation than ResourcePatchPolicy which uses /policy-details.
type ResourcePatchPolicySummary struct {
	ID                           string `json:"id"`
	PolicyName                   string `json:"policyName"`
	PolicyEnabled                bool   `json:"policyEnabled"`
	PolicyTargetVersion          string `json:"policyTargetVersion"`
	PolicyDeploymentMethod       string `json:"policyDeploymentMethod"`
	SoftwareTitle                string `json:"softwareTitle"`
	SoftwareTitleConfigurationID string `json:"softwareTitleConfigurationId"`
	Pending                      int    `json:"pending"`
	Completed                    int    `json:"completed"`
	Deferred                     int    `json:"deferred"`
	Failed                       int    `json:"failed"`
}

// ListSummaryResponse is the paginated response for ListSummaryV2.
type ListSummaryResponse struct {
	TotalCount int                         `json:"totalCount"`
	Results    []ResourcePatchPolicySummary `json:"results"`
}

// DashboardStatusResponse represents the response for checking if a patch policy is on the dashboard.
type DashboardStatusResponse struct {
	OnDashboard bool `json:"onDashboard"`
}
