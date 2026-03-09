package computers

import (
	"encoding/xml"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
)

// ListResponse is the response for List (GET /JSSResource/computers).
type ListResponse struct {
	XMLName xml.Name           `xml:"computers"`
	Size    int                `xml:"size"`
	Results []ComputersListItem `xml:"computer"`
}

// ComputersListItem represents a single computer item in the list.
type ComputersListItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name"`
}

// ResponseComputer is the main computer resource structure.
type ResponseComputer struct {
	XMLName               xml.Name                            `xml:"computer"`
	General               ComputerSubsetGeneral                `xml:"general"`
	Location              ComputerSubsetLocation               `xml:"location"`
	Purchasing            ComputerSubsetPurchasing             `xml:"purchasing"`
	Peripherals           ComputerContainerPeripherals        `xml:"peripherals"`
	Hardware              ComputerSubsetHardware              `xml:"hardware"`
	Certificates          []ComputerSubsetCertificates        `xml:"certificates>certificate"`
	Security              ComputerSubsetSecurity              `xml:"security"`
	Software              ComputerSubsetSoftware              `xml:"software"`
	ExtensionAttributes   []ComputerSubsetExtensionAttributes `xml:"extension_attributes>extension_attribute"`
	GroupsAccounts        ComputerSubsetGroupsAccounts        `xml:"groups_accounts"`
	ConfigurationProfiles []ComputerSubsetConfigurationProfiles `xml:"configuration_profiles>configuration_profile"`
}

// ComputerSubsetGeneral contains general computer information.
type ComputerSubsetGeneral struct {
	ID                         int                                    `xml:"id"`
	Name                       string                                 `xml:"name"`
	MacAddress                 string                                 `xml:"mac_address"`
	NetworkAdapterType         string                                 `xml:"network_adapter_type"`
	AltMacAddress              string                                 `xml:"alt_mac_address"`
	AltNetworkAdapterType      string                                 `xml:"alt_network_adapter_type"`
	IPAddress                  string                                 `xml:"ip_address"`
	LastReportedIP             string                                 `xml:"last_reported_ip"`
	SerialNumber               string                                 `xml:"serial_number"`
	UDID                       string                                 `xml:"udid"`
	JamfVersion                string                                 `xml:"jamf_version"`
	Platform                   string                                 `xml:"platform"`
	Barcode1                   string                                 `xml:"barcode_1"`
	Barcode2                   string                                 `xml:"barcode_2"`
	AssetTag                   string                                 `xml:"asset_tag"`
	RemoteManagement          ComputerSubsetGeneralRemoteManagement  `xml:"remote_management"`
	MdmCapable                 bool                                   `xml:"mdm_capable"`
	MdmCapableUsers            ComputerSubsetGeneralMdmCapableUsers   `xml:"mdm_capable_users"`
	MdmProfileExpirationEpoch  int64                                  `xml:"mdm_profile_expiration_epoch"`
	MdmProfileExpirationUtc   string                                 `xml:"mdm_profile_expiration_utc"`
	ManagementStatus           ComputerSubsetGeneralManagementStatus `xml:"management_status"`
	ReportDate                 string                                 `xml:"report_date"`
	ReportDateEpoch            int64                                  `xml:"report_date_epoch"`
	ReportDateUtc              string                                 `xml:"report_date_utc"`
	LastContactTime            string                                 `xml:"last_contact_time"`
	LastContactTimeEpoch       int64                                  `xml:"last_contact_time_epoch"`
	LastContactTimeUtc         string                                 `xml:"last_contact_time_utc"`
	InitialEntryDate           string                                 `xml:"initial_entry_date"`
	InitialEntryDateEpoch      int64                                  `xml:"initial_entry_date_epoch"`
	InitialEntryDateUtc        string                                 `xml:"initial_entry_date_utc"`
	LastCloudBackupDateEpoch   int64                                  `xml:"last_cloud_backup_date_epoch"`
	LastCloudBackupDateUtc     string                                 `xml:"last_cloud_backup_date_utc"`
	LastEnrolledDateEpoch      int64                                  `xml:"last_enrolled_date_epoch"`
	LastEnrolledDateUtc        string                                 `xml:"last_enrolled_date_utc"`
	DistributionPoint         string                                 `xml:"distribution_point"`
	Sus                        string                                 `xml:"sus"`
	Supervised                 bool                                   `xml:"supervised"`
	Site                       shared.SharedResourceSite              `xml:"site"`
	ItunesStoreAccountIsActive bool                                   `xml:"itunes_store_account_is_active"`
}

// ComputerSubsetGeneralRemoteManagement contains remote management info.
type ComputerSubsetGeneralRemoteManagement struct {
	Managed            bool   `xml:"managed"`
	ManagementUsername string `xml:"management_username"`
}

// ComputerSubsetGeneralMdmCapableUsers contains MDM capable user info.
type ComputerSubsetGeneralMdmCapableUsers struct {
	MdmCapableUser string `xml:"mdm_capable_user"`
}

// ComputerSubsetGeneralManagementStatus contains management status info.
type ComputerSubsetGeneralManagementStatus struct {
	EnrolledViaDep         bool `xml:"enrolled_via_dep"`
	UserApprovedEnrollment bool `xml:"user_approved_enrollment"`
	UserApprovedMdm        bool `xml:"user_approved_mdm"`
}

// ComputerSubsetLocation contains location information.
type ComputerSubsetLocation struct {
	Username     string `xml:"username"`
	RealName     string `xml:"realname"`
	EmailAddress string `xml:"email_address"`
	Position     string `xml:"position"`
	Phone        string `xml:"phone"`
	PhoneNumber  string `xml:"phone_number"`
	Department   string `xml:"department"`
	Building     string `xml:"building"`
	Room         string `xml:"room"`
}

// ComputerSubsetPurchasing contains purchasing information.
type ComputerSubsetPurchasing struct {
	IsPurchased          bool   `xml:"is_purchased"`
	IsLeased             bool   `xml:"is_leased"`
	PoNumber             string `xml:"po_number"`
	Vendor               string `xml:"vendor"`
	ApplecareID          string `xml:"applecare_id"`
	PurchasePrice        string `xml:"purchase_price"`
	PurchasingAccount    string `xml:"purchasing_account"`
	PoDate               string `xml:"po_date"`
	PoDateEpoch          int64  `xml:"po_date_epoch"`
	PoDateUtc            string `xml:"po_date_utc"`
	WarrantyExpires      string `xml:"warranty_expires"`
	WarrantyExpiresEpoch int64  `xml:"warranty_expires_epoch"`
	WarrantyExpiresUtc   string `xml:"warranty_expires_utc"`
	LeaseExpires         string `xml:"lease_expires"`
	LeaseExpiresEpoch    int64  `xml:"lease_expires_epoch"`
	LeaseExpiresUtc      string `xml:"lease_expires_utc"`
	LifeExpectancy       int    `xml:"life_expectancy"`
	PurchasingContact    string `xml:"purchasing_contact"`
	OSAppleCareID        string `xml:"os_applecare_id,omitempty"`
	OSMaintenanceExpires string `xml:"os_maintenance_expires,omitempty"`
}

// ComputerContainerPeripherals contains peripherals container.
type ComputerContainerPeripherals struct {
	Size        int                         `xml:"size"`
	Peripherals []ComputerSubsetPeripherals `xml:"peripheral"`
}

// ComputerSubsetPeripherals contains peripheral information.
type ComputerSubsetPeripherals struct {
	ID          int                                      `xml:"id"`
	BarCode1    string                                   `xml:"bar_code_1"`
	BarCode2    string                                   `xml:"bar_code_2"`
	Type        string                                   `xml:"type"`
	Fields      ComputerSubsetPeripheralsContainerFields `xml:"fields"`
	Purchasing  ComputerSubsetPeripheralsPurchasing      `xml:"purchasing"`
	Attachments []ComputerSubsetPeripheralsAttachments   `xml:"attachments>attachment"`
}

// ComputerSubsetPeripheralsContainerFields contains peripheral fields.
type ComputerSubsetPeripheralsContainerFields struct {
	Field []ComputerSubsetPeripheralsField `xml:"field"`
}

// ComputerSubsetPeripheralsField contains a single peripheral field.
type ComputerSubsetPeripheralsField struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

// ComputerSubsetPeripheralsPurchasing contains peripheral purchasing info.
type ComputerSubsetPeripheralsPurchasing struct {
	IsPurchased          bool   `xml:"is_purchased"`
	IsLeased             bool   `xml:"is_leased"`
	PoNumber             string `xml:"po_number"`
	Vendor               string `xml:"vendor"`
	ApplecareID          string `xml:"applecare_id"`
	PurchasePrice        string `xml:"purchase_price"`
	PurchasingAccount    string `xml:"purchasing_account"`
	PoDate               string `xml:"po_date"`
	PoDateEpoch          int64  `xml:"po_date_epoch"`
	PoDateUtc            string `xml:"po_date_utc"`
	WarrantyExpires      string `xml:"warranty_expires"`
	WarrantyExpiresEpoch int64  `xml:"warranty_expires_epoch"`
	WarrantyExpiresUtc   string `xml:"warranty_expires_utc"`
	LeaseExpires         string `xml:"lease_expires"`
	LeaseExpiresEpoch    int64  `xml:"lease_expires_epoch"`
	LeaseExpiresUtc      string `xml:"lease_expires_utc"`
	LifeExpectancy       int    `xml:"life_expectancy"`
	PurchasingContact    string `xml:"purchasing_contact"`
}

// ComputerSubsetPeripheralsAttachments contains peripheral attachment info.
type ComputerSubsetPeripheralsAttachments struct {
	Size     int    `xml:"size"`
	ID       int    `xml:"id"`
	Filename string `xml:"filename"`
	URI      string `xml:"uri"`
}

// ComputerSubsetHardware contains hardware information.
type ComputerSubsetHardware struct {
	Make                        string                                  `xml:"make"`
	Model                       string                                  `xml:"model"`
	ModelIdentifier             string                                  `xml:"model_identifier"`
	OsName                      string                                  `xml:"os_name"`
	OsVersion                   string                                  `xml:"os_version"`
	OsBuild                     string                                  `xml:"os_build"`
	MasterPasswordSet           bool                                    `xml:"master_password_set"`
	ActiveDirectoryStatus       string                                  `xml:"active_directory_status"`
	ServicePack                 string                                  `xml:"service_pack"`
	ProcessorType               string                                  `xml:"processor_type"`
	ProcessorArchitecture       string                                  `xml:"processor_architecture"`
	ProcessorSpeed              int                                     `xml:"processor_speed"`
	ProcessorSpeedMhz           int                                     `xml:"processor_speed_mhz"`
	NumberProcessors            int                                     `xml:"number_processors"`
	NumberCores                 int                                     `xml:"number_cores"`
	TotalRam                    int                                     `xml:"total_ram"`
	TotalRamMb                  int                                     `xml:"total_ram_mb"`
	BootRom                     string                                  `xml:"boot_rom"`
	BusSpeed                    int                                     `xml:"bus_speed"`
	BusSpeedMhz                 int                                     `xml:"bus_speed_mhz"`
	BatteryCapacity             int                                     `xml:"battery_capacity"`
	CacheSize                   int                                     `xml:"cache_size"`
	CacheSizeKb                 int                                     `xml:"cache_size_kb"`
	AvailableRamSlots           int                                     `xml:"available_ram_slots"`
	OpticalDrive                string                                  `xml:"optical_drive"`
	NicSpeed                    string                                  `xml:"nic_speed"`
	SmcVersion                  string                                  `xml:"smc_version"`
	BleCapable                  bool                                    `xml:"ble_capable"`
	SipStatus                   string                                  `xml:"sip_status"`
	GatekeeperStatus            string                                  `xml:"gatekeeper_status"`
	XprotectVersion             string                                  `xml:"xprotect_version"`
	InstitutionalRecoveryKey    string                                  `xml:"institutional_recovery_key"`
	DiskEncryptionConfiguration string                                  `xml:"disk_encryption_configuration"`
	SoftwareUpdateDeviceID      string                                  `xml:"software_update_device_id,omitempty"`
	IsAppleSilicon              bool                                    `xml:"is_apple_silicon,omitempty"`
	SupportsIosAppInstalls      bool                                    `xml:"supports_ios_app_installs,omitempty"`
	Filevault2Users             []ComputerSubsetHardwareFileVault2Users `xml:"filevault2_users>user"`
	Storage                     []ComputerSubsetHardwareStorage        `xml:"storage>device"`
	MappedPrinters              []ComputerSubsetHardwareMappedPrinters  `xml:"mapped_printers>printer"`
}

// ComputerSubsetHardwareFileVault2Users contains FileVault 2 user info.
type ComputerSubsetHardwareFileVault2Users struct {
	User string `xml:"user"`
}

// ComputerSubsetHardwareStorage contains storage device info.
type ComputerSubsetHardwareStorage struct {
	Disk            string                                    `xml:"disk"`
	Model           string                                    `xml:"model"`
	Revision        string                                    `xml:"revision"`
	SerialNumber    string                                    `xml:"serial_number"`
	Size            int                                       `xml:"size"`
	DriveCapacityMb int                                       `xml:"drive_capacity_mb"`
	ConnectionType  string                                    `xml:"connection_type"`
	SmartStatus     string                                    `xml:"smart_status"`
	Partitions      []ComputerSubsetHardwareStoragePartitions `xml:"partition"`
}

// ComputerSubsetHardwareStoragePartitions contains partition info.
type ComputerSubsetHardwareStoragePartitions struct {
	Name                 string `xml:"name"`
	Size                 int    `xml:"size"`
	Type                 string `xml:"type"`
	PartitionCapacityMb  int    `xml:"partition_capacity_mb"`
	PercentageFull       int    `xml:"percentage_full"`
	FilevaultStatus      string `xml:"filevault_status"`
	FilevaultPercent     int    `xml:"filevault_percent"`
	Filevault2Status     string `xml:"filevault2_status"`
	Filevault2Percent    int    `xml:"filevault2_percent"`
	BootDriveAvailableMb int    `xml:"boot_drive_available_mb"`
	LvgUUID              string `xml:"lvgUUID"`
	LvUUID               string `xml:"lvUUID"`
	PvUUID               string `xml:"pvUUID"`
}

// ComputerSubsetHardwareMappedPrinters contains mapped printer info.
type ComputerSubsetHardwareMappedPrinters struct {
	Name     string `xml:"name"`
	URI      string `xml:"uri"`
	Type     string `xml:"type"`
	Location string `xml:"location"`
}

// ComputerSubsetCertificates contains certificate info.
type ComputerSubsetCertificates struct {
	CommonName   string `xml:"common_name"`
	Identity     bool   `xml:"identity"`
	ExpiresUtc   string `xml:"expires_utc"`
	ExpiresEpoch int64  `xml:"expires_epoch"`
	Name         string `xml:"name"`
}

// ComputerSubsetSecurity contains security information.
type ComputerSubsetSecurity struct {
	ActivationLock       bool   `xml:"activation_lock"`
	RecoveryLockEnabled bool   `xml:"recovery_lock_enabled"`
	SecureBootLevel      string `xml:"secure_boot_level"`
	ExternalBootLevel    string `xml:"external_boot_level"`
	FirewallEnabled      bool   `xml:"firewall_enabled"`
}

// ComputerSubsetSoftware contains software information.
type ComputerSubsetSoftware struct {
	UnixExecutables          []string                                 `xml:"unix_executables>string"`
	LicensedSoftware         []string                                 `xml:"licensed_software>string"`
	InstalledByCasper        []string                                 `xml:"installed_by_casper>package"`
	InstalledByInstallerSwu  []string                                 `xml:"installed_by_installer_swu>package"`
	CachedByCasper           []string                                 `xml:"cached_by_casper>package"`
	AvailableSoftwareUpdates []string                                 `xml:"available_software_updates>name"`
	AvailableUpdates         []ComputerSubsetSoftwareAvailableUpdates `xml:"available_updates>update"`
	RunningServices          []string                                 `xml:"running_services>name"`
	Applications             []ComputerSubsetSoftwareApplications     `xml:"applications>application"`
	Fonts                    []ComputerSubsetSoftwareFonts           `xml:"fonts>font"`
	Plugins                  []ComputerSubsetSoftwarePlugins         `xml:"plugins>plugin"`
}

// ComputerSubsetSoftwareAvailableUpdates contains available update info.
type ComputerSubsetSoftwareAvailableUpdates struct {
	Name        string `xml:"name"`
	PackageName string `xml:"package_name"`
	Version     string `xml:"version"`
}

// ComputerSubsetSoftwareApplications contains application info.
type ComputerSubsetSoftwareApplications struct {
	Name    string `xml:"name"`
	Path    string `xml:"path"`
	Version string `xml:"version"`
}

// ComputerSubsetSoftwareFonts contains font info.
type ComputerSubsetSoftwareFonts struct {
	Name    string `xml:"name"`
	Path    string `xml:"path"`
	Version string `xml:"version"`
}

// ComputerSubsetSoftwarePlugins contains plugin info.
type ComputerSubsetSoftwarePlugins struct {
	Name    string `xml:"name"`
	Path    string `xml:"path"`
	Version string `xml:"version"`
}

// ComputerSubsetExtensionAttributes contains extension attribute info.
type ComputerSubsetExtensionAttributes struct {
	ID    int    `xml:"id"`
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
}

// ComputerSubsetGroupsAccounts contains groups and accounts info.
type ComputerSubsetGroupsAccounts struct {
	ComputerGroupMemberships []ComputerSubsetGroupsAccountsComputerGroupMemberships `xml:"computer_group_memberships>group"`
	LocalAccounts            []ComputerSubsetGroupsAccountsLocalAccounts            `xml:"local_accounts>user"`
}

// ComputerSubsetGroupsAccountsComputerGroupMemberships contains group membership.
type ComputerSubsetGroupsAccountsComputerGroupMemberships struct {
	Group string `xml:"group"`
}

// ComputerSubsetGroupsAccountsLocalAccounts contains local account info.
type ComputerSubsetGroupsAccountsLocalAccounts struct {
	Name             string `xml:"name"`
	RealName         string `xml:"realname"`
	UID              string `xml:"uid"`
	Home             string `xml:"home"`
	HomeSize         string `xml:"home_size"`
	HomeSizeMb       int    `xml:"home_size_mb"`
	Administrator    bool   `xml:"administrator"`
	FilevaultEnabled bool   `xml:"filevault_enabled"`
}

// ComputerSubsetConfigurationProfiles contains configuration profile info.
type ComputerSubsetConfigurationProfiles struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	UUID        string `xml:"uuid"`
	IsRemovable bool   `xml:"is_removable"`
}
