package app_installers

// ResourceJamfAppCatalogAppInstaller represents a Jamf App Catalog title (app installer) resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-titles-id
type ResourceJamfAppCatalogAppInstaller struct {
	ID                       string                                    `json:"id"`
	BundleId                 string                                    `json:"bundleId,omitempty"`
	TitleName                string                                    `json:"titleName,omitempty"`
	Publisher                string                                    `json:"publisher,omitempty"`
	IconUrl                  string                                    `json:"iconUrl,omitempty"`
	Version                  string                                    `json:"version,omitempty"`
	SizeInBytes              int                                      `json:"sizeInBytes,omitempty"`
	MinimumOsVersion         string                                    `json:"minimumOsVersion,omitempty"`
	Language                 string                                    `json:"language,omitempty"`
	AvailabilityDate         string                                    `json:"availabilityDate,omitempty"`
	PackageSigningIdentity   string                                    `json:"packageSigningIdentity,omitempty"`
	InstallerPackageHashType string                                    `json:"installerPackageHashType,omitempty"`
	InstallerPackageHash     string                                    `json:"installerPackageHash,omitempty"`
	ShortVersion             string                                    `json:"shortVersion,omitempty"`
	Architecture             string                                    `json:"architecture,omitempty"`
	OriginalMediaSources     []AppInstallerSubsetMediaSource           `json:"originalMediaSources,omitempty"`
	LaunchDaemonIncluded     *bool                                     `json:"launchDaemonIncluded"`
	NotificationAvailable    *bool                                     `json:"notificationAvailable"`
	SuppressAutoUpdate       *bool                                     `json:"suppressAutoUpdate"`
}

// AppInstallerSubsetMediaSource represents a media source for an app installer.
type AppInstallerSubsetMediaSource struct {
	HashType string `json:"hashType,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Url      string `json:"url,omitempty"`
}

// ListTitlesResponse is the response for ListTitlesV1.
type ListTitlesResponse struct {
	TotalCount int                              `json:"totalCount"`
	Results    []ResourceJamfAppCatalogAppInstaller `json:"results"`
}

// ResourceJamfAppCatalogDeployment represents a Jamf App Catalog deployment resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-installers-deployments-id
type ResourceJamfAppCatalogDeployment struct {
	ID                              string                                    `json:"id"`
	Name                            string                                    `json:"name,omitempty"`
	Enabled                         *bool                                     `json:"enabled"`
	AppTitleId                      string                                    `json:"appTitleId,omitempty"`
	DeploymentType                  string                                    `json:"deploymentType,omitempty"`
	UpdateBehavior                  string                                    `json:"updateBehavior,omitempty"`
	CategoryId                      string                                    `json:"categoryId,omitempty"`
	SiteId                          string                                    `json:"siteId,omitempty"`
	SmartGroupId                    string                                    `json:"smartGroupId,omitempty"`
	InstallPredefinedConfigProfiles *bool                                     `json:"installPredefinedConfigProfiles"`
	TitleAvailableInAis             *bool                                     `json:"titleAvailableInAis"`
	TriggerAdminNotifications       *bool                                     `json:"triggerAdminNotifications"`
	NotificationSettings            DeploymentSubsetNotificationSettings      `json:"notificationSettings,omitempty"`
	SelfServiceSettings             DeploymentSubsetSelfServiceSettings       `json:"selfServiceSettings,omitempty"`
	SelectedVersion                 string                                    `json:"selectedVersion,omitempty"`
	LatestAvailableVersion          string                                    `json:"latestAvailableVersion,omitempty"`
	VersionRemoved                  *bool                                    `json:"versionRemoved"`
}

// DeploymentSubsetNotificationSettings represents notification settings for a deployment.
type DeploymentSubsetNotificationSettings struct {
	NotificationMessage  string `json:"notificationMessage,omitempty"`
	NotificationInterval int    `json:"notificationInterval,omitempty"`
	DeadlineMessage      string `json:"deadlineMessage,omitempty"`
	Deadline             int    `json:"deadline,omitempty"`
	QuitDelay            int    `json:"quitDelay,omitempty"`
	CompleteMessage      string `json:"completeMessage,omitempty"`
	Relaunch             *bool  `json:"relaunch"`
	Suppress             *bool  `json:"suppress,omitempty"`
}

// DeploymentSubsetSelfServiceSettings represents self-service settings for a deployment.
type DeploymentSubsetSelfServiceSettings struct {
	IncludeInFeaturedCategory   *bool                          `json:"includeInFeaturedCategory"`
	IncludeInComplianceCategory *bool                          `json:"includeInComplianceCategory"`
	ForceViewDescription        *bool                          `json:"forceViewDescription"`
	Description                 string                         `json:"description,omitempty"`
	Categories                  []DeploymentSubsetCategory    `json:"categories,omitempty"`
}

// DeploymentSubsetCategory represents a category in self-service settings.
type DeploymentSubsetCategory struct {
	ID       string `json:"id,omitempty"`
	Featured *bool  `json:"featured,omitempty"`
}

// ListDeploymentsResponse is the response for ListDeploymentsV1.
type ListDeploymentsResponse struct {
	TotalCount int                                `json:"totalCount"`
	Results    []ResourceJamfAppCatalogDeployment `json:"results"`
}

// RequestDeployment is the body for creating and updating deployments.
type RequestDeployment struct {
	Name                            string                               `json:"name,omitempty"`
	Enabled                         *bool                                 `json:"enabled"`
	AppTitleId                      string                               `json:"appTitleId,omitempty"`
	DeploymentType                  string                               `json:"deploymentType,omitempty"`
	UpdateBehavior                  string                               `json:"updateBehavior,omitempty"`
	CategoryId                      string                               `json:"categoryId,omitempty"`
	SiteId                          string                               `json:"siteId,omitempty"`
	SmartGroupId                    string                               `json:"smartGroupId,omitempty"`
	InstallPredefinedConfigProfiles *bool                                `json:"installPredefinedConfigProfiles"`
	TitleAvailableInAis             *bool                                `json:"titleAvailableInAis"`
	TriggerAdminNotifications       *bool                                `json:"triggerAdminNotifications"`
	NotificationSettings            *DeploymentSubsetNotificationSettings `json:"notificationSettings,omitempty"`
	SelfServiceSettings             *DeploymentSubsetSelfServiceSettings  `json:"selfServiceSettings,omitempty"`
	SelectedVersion                 string                               `json:"selectedVersion,omitempty"`
}

// CreateDeploymentResponse is the response for CreateDeploymentV1.
type CreateDeploymentResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
