package mdm

// BlankPushResponse represents the response structure for the blank push MDM command.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-blank-push
type BlankPushResponse struct {
	ErrorUUIDs []string `json:"errorUuids"`
}

// CommandRequest represents the overall request structure for the MDM command.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
type CommandRequest struct {
	CommandData CommandData  `json:"commandData"`
	ClientData  []ClientData `json:"clientData"`
}

// ClientData represents the client data structure in the request.
type ClientData struct {
	ManagementID string `json:"managementId"`
}

// CommandData represents the command data structure in the request.
type CommandData struct {
	CommandType string `json:"commandType"`
	// Delete_User
	UserName       string `json:"userName,omitempty"`
	ForceDeletion  bool   `json:"forceDeletion,omitempty"`
	DeleteAllUsers bool   `json:"deleteAllUsers,omitempty"`
	// Enable_Lost_Mode
	LostModeMessage  string `json:"lostModeMessage,omitempty"`
	LostModePhone    string `json:"lostModePhone,omitempty"`
	LostModeFootnote string `json:"lostModeFootnote,omitempty"`
	// Erase_Device
	ReturnToService        *ReturnToService `json:"returnToService,omitempty"`
	PreserveDataPlan       bool             `json:"preserveDataPlan,omitempty"`
	DisallowProximitySetup bool             `json:"disallowProximitySetup,omitempty"`
	PIN                    string           `json:"pin,omitempty"`
	ObliterationBehavior   string           `json:"obliterationBehavior,omitempty"`
	// Restart_Device
	RebuildKernelCache bool     `json:"rebuildKernelCache,omitempty"`
	KextPaths          []string `json:"kextPaths,omitempty"`
	NotifyUser         bool     `json:"notifyUser,omitempty"`
	// Settings
	ApplicationAttributes     *ApplicationAttributes     `json:"applicationAttributes,omitempty"`
	SharedDeviceConfiguration *SharedDeviceConfiguration `json:"sharedDeviceConfiguration,omitempty"`
	ApplicationConfiguration  *ApplicationConfiguration  `json:"applicationConfiguration,omitempty"`
	SoftwareUpdateSettings    *SoftwareUpdateSettings    `json:"softwareUpdateSettings,omitempty"`
	BootstrapTokenAllowed     bool                       `json:"bootstrapTokenAllowed,omitempty"`
	Bluetooth                 bool                       `json:"bluetooth,omitempty"`
	AppAnalytics              string                     `json:"appAnalytics,omitempty"`
	DiagnosticSubmission      string                     `json:"diagnosticSubmission,omitempty"`
	DataRoaming               string                     `json:"dataRoaming,omitempty"`
	VoiceRoaming              string                     `json:"voiceRoaming,omitempty"`
	PersonalHotspot           string                     `json:"personalHotspot,omitempty"`
	MaximumResidentUsers      int                        `json:"maximumResidentUsers,omitempty"`
	DeviceName                string                     `json:"deviceName,omitempty"`
	TimeZone                  string                     `json:"timeZone,omitempty"`
	PasscodeLockGracePeriod   int                        `json:"passcodeLockGracePeriod,omitempty"`
	// Set_Auto_Admin_Password
	GUID     string `json:"guid,omitempty"`
	Password string `json:"password,omitempty"`
}

// ReturnToService represents the return to service structure in the erase device command.
type ReturnToService struct {
	Enabled         bool   `json:"enabled"`
	MDMProfileData  string `json:"mdmProfileData,omitempty"`
	WifiProfileData string `json:"wifiProfileData,omitempty"`
}

// ApplicationAttributes represents the application attributes structure in the settings command.
type ApplicationAttributes struct {
	VpnUuid               string   `json:"vpnUuid"`
	AssociatedDomains     []string `json:"associatedDomains"`
	Removable             bool     `json:"removable"`
	EnableDirectDownloads bool     `json:"enableDirectDownloads"`
	Identifier            string   `json:"identifier"`
}

// SharedDeviceConfiguration represents the shared device configuration structure in the settings command.
type SharedDeviceConfiguration struct {
	QuotaSize     int `json:"quotaSize"`
	ResidentUsers int `json:"residentUsers"`
}

// ApplicationConfiguration represents the application configuration structure in the settings command.
type ApplicationConfiguration struct {
	Configuration string `json:"configuration"`
	Identifier    string `json:"identifier"`
}

// SoftwareUpdateSettings represents the software update settings structure in the settings command.
type SoftwareUpdateSettings struct {
	RecommendationCadence string `json:"recommendationCadence"`
}

// CommandResponse represents the response structure for the MDM command.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
type CommandResponse struct {
	ID   string `json:"id,omitempty"`
	Href string `json:"href"`
}

// DeployPackageRequest represents the request structure for deploying a package.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-deploy-package
type DeployPackageRequest struct {
	Manifest         PackageManifest `json:"manifest"`
	InstallAsManaged bool            `json:"installAsManaged"`
	Devices          []int           `json:"devices"`
	GroupID          string          `json:"groupId"`
}

// PackageManifest represents the package manifest structure in the deploy package command.
type PackageManifest struct {
	HashType         string `json:"hashType"`
	URL              string `json:"url"`
	Hash             string `json:"hash"`
	DisplayImageURL  string `json:"displayImageUrl"`
	FullSizeImageURL string `json:"fullSizeImageUrl"`
	BundleID         string `json:"bundleId"`
	BundleVersion    string `json:"bundleVersion"`
	Subtitle         string `json:"subtitle"`
	Title            string `json:"title"`
	SizeInBytes      int    `json:"sizeInBytes"`
}

// DeployPackageResponse represents the response structure for deploying a package.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-deploy-package
type DeployPackageResponse struct {
	QueuedCommands []QueuedCommand             `json:"queuedCommands"`
	Errors         []SharedResourceErrorDetail `json:"errors"`
}

// QueuedCommand represents the details of a queued command in the response.
type QueuedCommand struct {
	Device      int    `json:"device"`
	CommandUUID string `json:"commandUuid"`
}

// SharedResourceErrorDetail represents error details in shared resource responses.
type SharedResourceErrorDetail struct {
	Device      int    `json:"device"`
	Group       int    `json:"group"`
	Reason      string `json:"reason"`
	Code        string `json:"code"`
	Field       string `json:"field"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

// RenewProfileRequest represents the request structure for renewing MDM profiles.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mdm-renew-profile
type RenewProfileRequest struct {
	UDIDs []string `json:"udids"`
}

// RenewProfileResponse represents the response structure for renewing MDM profiles.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mdm-renew-profile
type RenewProfileResponse struct {
	UDIDsNotProcessed struct {
		UDIDs []string `json:"udids"`
	} `json:"udidsNotProcessed"`
}
