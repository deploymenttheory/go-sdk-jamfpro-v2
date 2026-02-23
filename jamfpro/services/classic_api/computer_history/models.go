package computer_history

import "encoding/xml"

// ResourceComputerHistory represents the root structure of the computer history resource.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerhistory
type ResourceComputerHistory struct {
	XMLName           xml.Name              `xml:"computer_history"`
	General           SubsetGeneralInfo     `xml:"general"`
	ComputerUsageLogs []SubsetUsageLog      `xml:"computer_usage_logs>usage_log,omitempty"`
	Audits            []SubsetAudit         `xml:"audits>audit,omitempty"`
	PolicyLogs        []SubsetPolicyLog     `xml:"policy_logs>policy_log,omitempty"`
	CasperRemoteLogs  []SubsetCasperRemote  `xml:"casper_remote_logs>casper_remote_log,omitempty"`
	ScreenSharingLogs []SubsetScreenSharing `xml:"screen_sharing_logs>screen_sharing_log,omitempty"`
	CasperImagingLogs []SubsetCasperImaging `xml:"casper_imaging_logs>casper_imaging_log,omitempty"`
	Commands          *SubsetCommands       `xml:"commands,omitempty"`
	UserLocation      []SubsetLocation      `xml:"user_location>location,omitempty"`
	MacAppStoreApps   *SubsetAppStoreApps   `xml:"mac_app_store_applications,omitempty"`
}

// SubsetGeneralInfo stores general information about the computer.
type SubsetGeneralInfo struct {
	ID           int    `xml:"id,omitempty"`
	Name         string `xml:"name,omitempty"`
	UDID         string `xml:"udid,omitempty"`
	SerialNumber string `xml:"serial_number,omitempty"`
	MacAddress   string `xml:"mac_address,omitempty"`
}

// SubsetUsageLog stores logs related to computer usage.
type SubsetUsageLog struct {
	SubsetEventDetails
}

// SubsetAudit stores audit logs.
type SubsetAudit struct {
	SubsetEventDetails
}

// SubsetPolicyLog stores logs related to policies.
type SubsetPolicyLog struct {
	SubsetPolicyDetails
}

// SubsetCasperRemote stores logs for Casper remote actions.
type SubsetCasperRemote struct {
	SubsetEventStatus
}

// SubsetScreenSharing stores logs related to screen sharing.
type SubsetScreenSharing struct {
	SubsetScreenSharingDetails
}

// SubsetCasperImaging stores logs for Casper imaging actions.
type SubsetCasperImaging struct {
	SubsetEventStatus
}

// SubsetCommands groups completed, pending, and failed commands.
type SubsetCommands struct {
	Completed []SubsetCommand `xml:"completed>command,omitempty"`
	Pending   []SubsetCommand `xml:"pending>command,omitempty"`
	Failed    []SubsetCommand `xml:"failed>command,omitempty"`
}

// SubsetLocation stores location data related to a user.
type SubsetLocation struct {
	SubsetUserLocation
}

// SubsetAppStoreApps groups installed, pending, and failed applications from the Mac App Store.
type SubsetAppStoreApps struct {
	Installed []SubsetApp `xml:"installed>app,omitempty"`
	Pending   []SubsetApp `xml:"pending>app,omitempty"`
	Failed    []SubsetApp `xml:"failed>app,omitempty"`
}

// SubsetEventDetails defines the structure for logging events with timestamps and user information.
type SubsetEventDetails struct {
	Event         string `xml:"event,omitempty"`
	Username      string `xml:"username,omitempty"`
	DateTime      string `xml:"date_time,omitempty"`
	DateTimeEpoch int64  `xml:"date_time_epoch,omitempty"`
	DateTimeUTC   string `xml:"date_time_utc,omitempty"`
}

// SubsetPolicyDetails defines the details for policy logs.
type SubsetPolicyDetails struct {
	PolicyID      int    `xml:"policy_id,omitempty"`
	PolicyName    string `xml:"policy_name,omitempty"`
	Username      string `xml:"username,omitempty"`
	DateTime      string `xml:"date_time,omitempty"`
	DateTimeEpoch int64  `xml:"date_time_epoch,omitempty"`
	DateTimeUTC   string `xml:"date_time_utc,omitempty"`
	Status        string `xml:"status,omitempty"`
}

// SubsetEventStatus defines a simple structure for logs with status and timestamps.
type SubsetEventStatus struct {
	DateTime      string `xml:"date_time,omitempty"`
	DateTimeEpoch int64  `xml:"date_time_epoch,omitempty"`
	DateTimeUTC   string `xml:"date_time_utc,omitempty"`
	Status        string `xml:"status,omitempty"`
}

// SubsetScreenSharingDetails extends event status with details specific to screen sharing.
type SubsetScreenSharingDetails struct {
	SubsetEventStatus
	Details string `xml:"details,omitempty"`
}

// SubsetCommand details a command with its issue and completion status.
type SubsetCommand struct {
	Name           string `xml:"name,omitempty"`
	Status         string `xml:"status,omitempty"`
	Issued         string `xml:"issued,omitempty"`
	IssuedEpoch    int64  `xml:"issued_epoch,omitempty"`
	IssuedUTC      string `xml:"issued_utc,omitempty"`
	LastPush       string `xml:"last_push,omitempty"`
	LastPushEpoch  int64  `xml:"last_push_epoch,omitempty"`
	LastPushUTC    string `xml:"last_push_utc,omitempty"`
	Username       string `xml:"username,omitempty"`
	Completed      string `xml:"completed,omitempty"`
	CompletedEpoch int64  `xml:"completed_epoch,omitempty"`
	CompletedUTC   string `xml:"completed_utc,omitempty"`
	Failed         string `xml:"failed,omitempty"`
	FailedEpoch    int64  `xml:"failed_epoch,omitempty"`
	FailedUTC      string `xml:"failed_utc,omitempty"`
}

// SubsetUserLocation defines the detailed information about a user's location.
type SubsetUserLocation struct {
	DateTime      string `xml:"date_time,omitempty"`
	DateTimeEpoch int64  `xml:"date_time_epoch,omitempty"`
	DateTimeUTC   string `xml:"date_time_utc,omitempty"`
	Username      string `xml:"username,omitempty"`
	FullName      string `xml:"full_name,omitempty"`
	EmailAddress  string `xml:"email_address,omitempty"`
	PhoneNumber   string `xml:"phone_number,omitempty"`
	Department    string `xml:"department,omitempty"`
	Building      string `xml:"building,omitempty"`
	Room          int    `xml:"room,omitempty"`
	Position      string `xml:"position,omitempty"`
}

// SubsetApp defines the structure for application details in the Mac App Store context.
type SubsetApp struct {
	Name            string `xml:"name,omitempty"`
	Version         string `xml:"version,omitempty"`
	SizeMB          int    `xml:"size_mb,omitempty"`
	Status          string `xml:"status,omitempty"`
	Deployed        string `xml:"deployed,omitempty"`
	DeployedEpoch   int64  `xml:"deployed_epoch,omitempty"`
	DeployedUTC     string `xml:"deployed_utc,omitempty"`
	LastUpdate      string `xml:"last_update,omitempty"`
	LastUpdateEpoch int64  `xml:"last_update_epoch,omitempty"`
	LastUpdateUTC   string `xml:"last_update_utc,omitempty"`
}
