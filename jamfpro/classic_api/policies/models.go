package policies

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ListResponse represents the response for listing policies.
type ListResponse struct {
	XMLName xml.Name         `xml:"policies"`
	Size    int              `xml:"size"`
	Results []PolicyListItem `xml:"policy"`
}

// PolicyListItem represents a single policy in a list response.
type PolicyListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// CreateUpdateResponse represents the response structure for creating or updating a policy.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"policy"`
	ID      int      `xml:"id"`
}

// ResourcePolicy represents the complete structure of a policy.
type ResourcePolicy struct {
	XMLName              xml.Name                     `xml:"policy"`
	General              PolicySubsetGeneral              `xml:"general"`
	Scope                PolicySubsetScope                `xml:"scope,omitempty"`
	SelfService          PolicySubsetSelfService          `xml:"self_service,omitempty"`
	PackageConfiguration PolicySubsetPackageConfiguration `xml:"package_configuration"`
	Scripts              []PolicySubsetScript             `xml:"scripts>script"`
	Printers             PolicySubsetPrinters             `xml:"printers"`
	DockItems            []PolicySubsetDockItem           `xml:"dock_items>dock_item"`
	AccountMaintenance   PolicySubsetAccountMaintenance   `xml:"account_maintenance,omitempty"`
	Maintenance          PolicySubsetMaintenance          `xml:"maintenance"`
	FilesProcesses       PolicySubsetFilesProcesses       `xml:"files_processes,omitempty"`
	UserInteraction      PolicySubsetUserInteraction      `xml:"user_interaction,omitempty"`
	DiskEncryption       PolicySubsetDiskEncryption       `xml:"disk_encryption,omitempty"`
	Reboot               PolicySubsetReboot               `xml:"reboot,omitempty"`
}

// PolicySubsetGeneral represents the general information of a policy.
type PolicySubsetGeneral struct {
	ID                         int                                     `xml:"id"`
	Name                       string                                  `xml:"name"`
	Enabled                    bool                                    `xml:"enabled"`
	Trigger                    string                                  `xml:"trigger,omitempty"`
	TriggerCheckin             bool                                    `xml:"trigger_checkin"`
	TriggerEnrollmentComplete  bool                                    `xml:"trigger_enrollment_complete"`
	TriggerLogin               bool                                    `xml:"trigger_login"`
	TriggerLogout              bool                                    `xml:"trigger_logout"`
	TriggerNetworkStateChanged bool                                    `xml:"trigger_network_state_changed"`
	TriggerStartup             bool                                    `xml:"trigger_startup"`
	TriggerOther               string                                  `xml:"trigger_other"`
	Frequency                  string                                  `xml:"frequency,omitempty"`
	RetryEvent                 string                                  `xml:"retry_event,omitempty"`
	RetryAttempts              int                                     `xml:"retry_attempts,omitempty"`
	NotifyOnEachFailedRetry    bool                                    `xml:"notify_on_each_failed_retry"`
	LocationUserOnly           bool                                    `xml:"location_user_only"`
	TargetDrive                string                                  `xml:"target_drive,omitempty"`
	Offline                    bool                                    `xml:"offline"`
	Category                   *models.SharedResourceCategory          `xml:"category,omitempty"`
	DateTimeLimitations        *PolicySubsetGeneralDateTimeLimitations `xml:"date_time_limitations,omitempty"`
	NetworkLimitations         *PolicySubsetGeneralNetworkLimitations  `xml:"network_limitations,omitempty"`
	OverrideDefaultSettings    *PolicySubsetGeneralOverrideSettings    `xml:"override_default_settings,omitempty"`
	NetworkRequirements        string                                  `xml:"network_requirements,omitempty"`
	Site                       *models.SharedResourceSite              `xml:"site,omitempty"`
}

// PolicySubsetGeneralDateTimeLimitations represents the date/time limitations for a policy.
type PolicySubsetGeneralDateTimeLimitations struct {
	ActivationDate      string   `xml:"activation_date"`
	ActivationDateEpoch int      `xml:"activation_date_epoch"`
	ActivationDateUTC   string   `xml:"activation_date_utc"`
	ExpirationDate      string   `xml:"expiration_date"`
	ExpirationDateEpoch int      `xml:"expiration_date_epoch"`
	ExpirationDateUTC   string   `xml:"expiration_date_utc"`
	NoExecuteOn         []string `xml:"no_execute_on>day,omitempty"`
	NoExecuteStart      string   `xml:"no_execute_start"`
	NoExecuteEnd        string   `xml:"no_execute_end"`
}

// PolicySubsetGeneralNetworkLimitations represents the network limitations for a policy.
type PolicySubsetGeneralNetworkLimitations struct {
	MinimumNetworkConnection string `xml:"minimum_network_connection"`
	AnyIPAddress             bool   `xml:"any_ip_address"`
	NetworkSegments          string `xml:"network_segments"`
}

// PolicySubsetGeneralOverrideSettings represents the override settings for a policy.
type PolicySubsetGeneralOverrideSettings struct {
	TargetDrive       string `xml:"target_drive"`
	DistributionPoint string `xml:"distribution_point"`
	ForceAfpSmb       bool   `xml:"force_afp_smb"`
	SUS               string `xml:"sus"`
	NetbootServer     string `xml:"netboot_server"`
}

// PolicySubsetScope represents the scope of the policy.
type PolicySubsetScope struct {
	AllComputers   bool                          `xml:"all_computers"`
	AllJSSUsers    bool                          `xml:"all_jss_users"`
	Computers      *[]PolicySubsetComputer       `xml:"computers>computer"`
	ComputerGroups *[]PolicySubsetComputerGroup  `xml:"computer_groups>computer_group"`
	JSSUsers       *[]PolicySubsetJSSUser        `xml:"jss_users>user"`
	JSSUserGroups  *[]PolicySubsetJSSUserGroup   `xml:"jss_user_groups>user_group"`
	Buildings      *[]PolicySubsetBuilding       `xml:"buildings>building"`
	Departments    *[]PolicySubsetDepartment     `xml:"departments>department"`
	LimitToUsers   PolicyLimitToUsers            `xml:"limit_to_users,omitempty"`
	Limitations    *PolicySubsetScopeLimitations `xml:"limitations"`
	Exclusions     *PolicySubsetScopeExclusions  `xml:"exclusions"`
}

// PolicySubsetScopeLimitations represents the limitations within a policy scope.
type PolicySubsetScopeLimitations struct {
	Users           *[]PolicySubsetUser           `xml:"users>user"`
	UserGroups      *[]PolicySubsetUserGroup      `xml:"user_groups>user_group"`
	NetworkSegments *[]PolicySubsetNetworkSegment `xml:"network_segments>network_segment"`
	IBeacons        *[]PolicySubsetIBeacon        `xml:"ibeacons>ibeacon"`
}

// PolicySubsetScopeExclusions represents the exclusions within a policy scope.
type PolicySubsetScopeExclusions struct {
	Computers       *[]PolicySubsetComputer       `xml:"computers>computer"`
	ComputerGroups  *[]PolicySubsetComputerGroup  `xml:"computer_groups>computer_group"`
	Users           *[]PolicySubsetUser           `xml:"users>user"`
	UserGroups      *[]PolicySubsetUserGroup      `xml:"user_groups>user_group"`
	Buildings       *[]PolicySubsetBuilding       `xml:"buildings>building"`
	Departments     *[]PolicySubsetDepartment     `xml:"departments>department"`
	NetworkSegments *[]PolicySubsetNetworkSegment `xml:"network_segments>network_segment"`
	JSSUsers        *[]PolicySubsetJSSUser        `xml:"jss_users>user"`
	JSSUserGroups   *[]PolicySubsetJSSUserGroup   `xml:"jss_user_groups>user_group"`
	IBeacons        *[]PolicySubsetIBeacon        `xml:"ibeacons>ibeacon"`
}

// PolicySubsetSelfService represents the self service settings of a policy.
type PolicySubsetSelfService struct {
	UseForSelfService           bool                              `xml:"use_for_self_service"`
	SelfServiceDisplayName      string                            `xml:"self_service_display_name"`
	InstallButtonText           string                            `xml:"install_button_text"`
	ReinstallButtonText         string                            `xml:"reinstall_button_text"`
	SelfServiceDescription      string                            `xml:"self_service_description"`
	ForceUsersToViewDescription bool                              `xml:"force_users_to_view_description"`
	SelfServiceIcon             *PolicySubsetSelfServiceIcon      `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                              `xml:"feature_on_main_page"`
	SelfServiceCategories       []PolicySubsetSelfServiceCategory `xml:"self_service_categories>category"`
	Notification                bool                              `xml:"notification"`
	NotificationType            string                            `xml:"notification_type"`
	NotificationSubject         string                            `xml:"notification_subject"`
	NotificationMessage         string                            `xml:"notification_message"`
}

// PolicySubsetSelfServiceIcon represents a self service icon.
type PolicySubsetSelfServiceIcon struct {
	ID       int    `xml:"id"`
	Filename string `xml:"filename"`
	URI      string `xml:"uri"`
}

// PolicySubsetSelfServiceCategory represents a self service category.
type PolicySubsetSelfServiceCategory struct {
	ID        int    `xml:"id"`
	Name      string `xml:"name"`
	DisplayIn bool   `xml:"display_in"`
	FeatureIn bool   `xml:"feature_in"`
}

// PolicySubsetPackageConfiguration represents the package configuration settings of a policy.
type PolicySubsetPackageConfiguration struct {
	Packages          []PolicySubsetPackageConfigurationPackage `xml:"packages>package"`
	DistributionPoint string                                    `xml:"distribution_point"`
}

// PolicySubsetPackageConfigurationPackage represents a package in a policy configuration.
type PolicySubsetPackageConfigurationPackage struct {
	ID                int    `xml:"id"`
	Name              string `xml:"name,omitempty"`
	Action            string `xml:"action"`
	FillUserTemplate  bool   `xml:"fut"`
	FillExistingUsers bool   `xml:"feu"`
	UpdateAutorun     bool   `xml:"update_autorun"`
}

// PolicySubsetScript represents a script in a policy.
type PolicySubsetScript struct {
	ID          string `xml:"id"`
	Name        string `xml:"name,omitempty"`
	Priority    string `xml:"priority"`
	Parameter4  string `xml:"parameter4,omitempty"`
	Parameter5  string `xml:"parameter5,omitempty"`
	Parameter6  string `xml:"parameter6,omitempty"`
	Parameter7  string `xml:"parameter7,omitempty"`
	Parameter8  string `xml:"parameter8,omitempty"`
	Parameter9  string `xml:"parameter9,omitempty"`
	Parameter10 string `xml:"parameter10,omitempty"`
	Parameter11 string `xml:"parameter11,omitempty"`
}

// PolicySubsetPrinters represents the printers settings of a policy.
type PolicySubsetPrinters struct {
	LeaveExistingDefault bool                  `xml:"leave_existing_default"`
	Printer              []PolicySubsetPrinter `xml:"printer"`
}

// PolicySubsetPrinter represents a printer in a policy.
type PolicySubsetPrinter struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Action      string `xml:"action"`
	MakeDefault bool   `xml:"make_default"`
}

// PolicySubsetDockItem represents a dock item in a policy.
type PolicySubsetDockItem struct {
	ID     int    `xml:"id"`
	Name   string `xml:"name"`
	Action string `xml:"action"`
}

// PolicySubsetAccountMaintenance represents the account maintenance settings of a policy.
type PolicySubsetAccountMaintenance struct {
	Accounts                *[]PolicySubsetAccountMaintenanceAccount               `xml:"accounts>account"`
	DirectoryBindings       *[]PolicySubsetAccountMaintenanceDirectoryBindings     `xml:"directory_bindings>binding"`
	ManagementAccount       *PolicySubsetAccountMaintenanceManagementAccount       `xml:"management_account"`
	OpenFirmwareEfiPassword *PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword `xml:"open_firmware_efi_password"`
}

// PolicySubsetAccountMaintenanceAccount represents an account in account maintenance.
type PolicySubsetAccountMaintenanceAccount struct {
	Action                 string `xml:"action"`
	Username               string `xml:"username"`
	Realname               string `xml:"realname"`
	Password               string `xml:"password"`
	ArchiveHomeDirectory   bool   `xml:"archive_home_directory"`
	ArchiveHomeDirectoryTo string `xml:"archive_home_directory_to"`
	Home                   string `xml:"home"`
	Hint                   string `xml:"hint"`
	Picture                string `xml:"picture"`
	Admin                  bool   `xml:"admin"`
	FilevaultEnabled       bool   `xml:"filevault_enabled"`
	PasswordSha256         string `xml:"password_sha256"`
}

// PolicySubsetAccountMaintenanceDirectoryBindings represents directory bindings in account maintenance.
type PolicySubsetAccountMaintenanceDirectoryBindings struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// PolicySubsetAccountMaintenanceManagementAccount represents the management account settings.
type PolicySubsetAccountMaintenanceManagementAccount struct {
	Action                string `xml:"action"`
	ManagedPassword       string `xml:"managed_password"`
	ManagedPasswordLength int    `xml:"managed_password_length"`
}

// PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword represents the OpenFirmware/EFI password settings.
type PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword struct {
	OfMode           string `xml:"of_mode"`
	OfPassword       string `xml:"of_password"`
	OfPasswordSHA256 string `xml:"of_password_sha256"`
}

// PolicySubsetMaintenance represents the maintenance settings of a policy.
type PolicySubsetMaintenance struct {
	Recon                    bool `xml:"recon"`
	ResetName                bool `xml:"reset_name"`
	InstallAllCachedPackages bool `xml:"install_all_cached_packages"`
	Heal                     bool `xml:"heal"`
	Prebindings              bool `xml:"prebindings"`
	Permissions              bool `xml:"permissions"`
	Byhost                   bool `xml:"byhost"`
	SystemCache              bool `xml:"system_cache"`
	UserCache                bool `xml:"user_cache"`
	Verify                   bool `xml:"verify"`
}

// PolicySubsetFilesProcesses represents the files and processes settings of a policy.
type PolicySubsetFilesProcesses struct {
	SearchByPath         string `xml:"search_by_path"`
	DeleteFile           bool   `xml:"delete_file"`
	LocateFile           string `xml:"locate_file"`
	UpdateLocateDatabase bool   `xml:"update_locate_database"`
	SpotlightSearch      string `xml:"spotlight_search"`
	SearchForProcess     string `xml:"search_for_process"`
	KillProcess          bool   `xml:"kill_process"`
	RunCommand           string `xml:"run_command"`
}

// PolicySubsetUserInteraction represents the user interaction settings of a policy.
type PolicySubsetUserInteraction struct {
	MessageStart          string `xml:"message_start"`
	AllowUsersToDefer     bool   `xml:"allow_users_to_defer"`
	AllowDeferralUntilUtc string `xml:"allow_deferral_until_utc"`
	AllowDeferralMinutes  int    `xml:"allow_deferral_minutes"`
	MessageFinish         string `xml:"message_finish"`
}

// PolicySubsetDiskEncryption represents the disk encryption settings of a policy.
type PolicySubsetDiskEncryption struct {
	Action                                 string `xml:"action"`
	DiskEncryptionConfigurationID          int    `xml:"disk_encryption_configuration_id"`
	AuthRestart                            bool   `xml:"auth_restart"`
	RemediateKeyType                       string `xml:"remediate_key_type"`
	RemediateDiskEncryptionConfigurationID int    `xml:"remediate_disk_encryption_configuration_id"`
}

// PolicySubsetReboot represents the reboot settings of a policy.
type PolicySubsetReboot struct {
	Message                     string `xml:"message"`
	StartupDisk                 string `xml:"startup_disk"`
	SpecifyStartup              string `xml:"specify_startup"`
	NoUserLoggedIn              string `xml:"no_user_logged_in"`
	UserLoggedIn                string `xml:"user_logged_in"`
	MinutesUntilReboot          int    `xml:"minutes_until_reboot"`
	StartRebootTimerImmediately bool   `xml:"start_reboot_timer_immediately"`
	FileVault2Reboot            bool   `xml:"file_vault_2_reboot"`
}

// PolicySubsetComputer represents a computer in a policy scope.
type PolicySubsetComputer struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UDID string `xml:"udid,omitempty"`
}

// PolicySubsetComputerGroup represents a computer group in a policy scope.
type PolicySubsetComputerGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// PolicySubsetJSSUser represents a JSS user in a policy scope.
type PolicySubsetJSSUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// PolicySubsetJSSUserGroup represents a JSS user group in a policy scope.
type PolicySubsetJSSUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// PolicySubsetBuilding represents a building in a policy scope.
type PolicySubsetBuilding struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// PolicySubsetDepartment represents a department in a policy scope.
type PolicySubsetDepartment struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// PolicyLimitToUsers represents user group limitations in a policy scope.
type PolicyLimitToUsers struct {
	UserGroups []string `xml:"user_groups>user_group"`
}

// PolicySubsetUser represents a user in a policy scope.
type PolicySubsetUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// PolicySubsetUserGroup represents a user group in a policy scope.
type PolicySubsetUserGroup struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
}

// PolicySubsetNetworkSegment represents a network segment in a policy scope.
type PolicySubsetNetworkSegment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	UID  string `xml:"uid"`
}

// PolicySubsetIBeacon represents an iBeacon in a policy scope.
type PolicySubsetIBeacon struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}
