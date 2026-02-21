package computer_inventory

// SharedResourceSiteProAPI represents a site in the Jamf Pro API.
type SharedResourceSiteProAPI struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// List

// ResponseComputerInventoryList represents the top-level JSON response structure.
type ResponseComputerInventoryList struct {
	TotalCount int                         `json:"totalCount"`
	Results    []ResourceComputerInventory `json:"results"`
}

// Resource

// ResponseComputerInventory represents an individual computer from the inventory.
type ResourceComputerInventory struct {
	ID                    string                                        `json:"id"`
	UDID                  string                                        `json:"udid"`
	General               ComputerInventorySubsetGeneral                `json:"general"`
	DiskEncryption        ComputerInventorySubsetDiskEncryption         `json:"diskEncryption"`
	Purchasing            ComputerInventorySubsetPurchasing             `json:"purchasing"`
	Applications          []ComputerInventorySubsetApplication          `json:"applications"`
	Storage               ComputerInventorySubsetStorage                `json:"storage"`
	UserAndLocation       ComputerInventorySubsetUserAndLocation        `json:"userAndLocation"`
	ConfigurationProfiles []ComputerInventorySubsetConfigurationProfile `json:"configurationProfiles"`
	Printers              []ComputerInventorySubsetPrinter              `json:"printers"`
	Services              []ComputerInventorySubsetService              `json:"services"`
	Hardware              ComputerInventorySubsetHardware               `json:"hardware"`
	LocalUserAccounts     []ComputerInventorySubsetLocalUserAccount     `json:"localUserAccounts"`
	Certificates          []ComputerInventorySubsetCertificate          `json:"certificates"`
	Attachments           []ComputerInventorySubsetAttachment           `json:"attachments"`
	Plugins               []ComputerInventorySubsetPlugin               `json:"plugins"`
	PackageReceipts       ComputerInventorySubsetPackageReceipts        `json:"packageReceipts"`
	Fonts                 []ComputerInventorySubsetFont                 `json:"fonts"`
	Security              ComputerInventorySubsetSecurity               `json:"security"`
	OperatingSystem       ComputerInventorySubsetOperatingSystem        `json:"operatingSystem"`
	LicensedSoftware      []ComputerInventorySubsetLicensedSoftware     `json:"licensedSoftware"`
	Ibeacons              []ComputerInventorySubsetIBeacon              `json:"ibeacons"`
	SoftwareUpdates       []ComputerInventorySubsetSoftwareUpdate       `json:"softwareUpdates"`
	ExtensionAttributes   []ComputerInventorySubsetExtensionAttribute   `json:"extensionAttributes"`
	ContentCaching        ComputerInventorySubsetContentCaching         `json:"contentCaching"`
	GroupMemberships      []ComputerInventorySubsetGroupMembership      `json:"groupMemberships"`
}

// Subsets

// General

type ComputerInventorySubsetGeneral struct {
	Name                                     string                                         `json:"name"`
	LastIpAddress                            string                                         `json:"lastIpAddress"`
	LastReportedIp                           string                                         `json:"lastReportedIp"`
	LastReportedIpV4                         string                                         `json:"lastReportedIpV4"`
	LastReportedIpV6                         string                                         `json:"lastReportedIpV6"`
	JamfBinaryVersion                        string                                         `json:"jamfBinaryVersion"`
	Platform                                 string                                         `json:"platform"`
	Barcode1                                 string                                         `json:"barcode1"`
	Barcode2                                 string                                         `json:"barcode2"`
	AssetTag                                 string                                         `json:"assetTag"`
	RemoteManagement                         ComputerInventorySubsetGeneralRemoteManagement `json:"remoteManagement"`
	Supervised                               bool                                           `json:"supervised"`
	MdmCapable                               ComputerInventorySubsetGeneralMdmCapable       `json:"mdmCapable"`
	ReportDate                               string                                         `json:"reportDate"`
	LastContactTime                          string                                         `json:"lastContactTime"`
	LastCloudBackupDate                      string                                         `json:"lastCloudBackupDate"`
	LastEnrolledDate                         string                                         `json:"lastEnrolledDate"`
	MdmProfileExpiration                     string                                         `json:"mdmProfileExpiration"`
	InitialEntryDate                         string                                         `json:"initialEntryDate"`
	DistributionPoint                        string                                         `json:"distributionPoint"`
	EnrollmentMethod                         ComputerInventorySubsetGeneralEnrollmentMethod `json:"enrollmentMethod"`
	Site                                     SharedResourceSiteProAPI                       `json:"site"`
	ItunesStoreAccountActive                 bool                                           `json:"itunesStoreAccountActive"`
	EnrolledViaAutomatedDeviceEnrollment     bool                                           `json:"enrolledViaAutomatedDeviceEnrollment"`
	UserApprovedMdm                          bool                                           `json:"userApprovedMdm"`
	DeclarativeDeviceManagementEnabled       bool                                           `json:"declarativeDeviceManagementEnabled"`
	ExtensionAttributes                      []ComputerInventorySubsetExtensionAttribute    `json:"extensionAttributes"`
	ManagementId                             string                                         `json:"managementId"`
	LastLoggedInUsernameSelfService          string                                         `json:"lastLoggedInUsernameSelfService"`
	LastLoggedInUsernameSelfServiceTimestamp string                                         `json:"lastLoggedInUsernameSelfServiceTimestamp"`
	LastLoggedInUsernameBinary               string                                         `json:"lastLoggedInUsernameBinary"`
	LastLoggedInUsernameBinaryTimestamp      string                                         `json:"lastLoggedInUsernameBinaryTimestamp"`
}

type ComputerInventorySubsetGeneralRemoteManagement struct {
	Managed            bool   `json:"managed"`
	ManagementUsername string `json:"managementUsername"`
}

type ComputerInventorySubsetGeneralMdmCapable struct {
	Capable            bool     `json:"capable"`
	CapableUsers       []string `json:"capableUsers"`
	UserManagementInfo []struct {
		CapableUser  string `json:"capableUser"`
		ManagementId string `json:"managementId"`
	} `json:"userManagementInfo"`
}

type ComputerInventorySubsetGeneralEnrollmentMethod struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
	ObjectType string `json:"objectType"`
}

// Disk Encryption

type ComputerInventorySubsetDiskEncryption struct {
	BootPartitionEncryptionDetails      ComputerInventorySubsetBootPartitionEncryptionDetails `json:"bootPartitionEncryptionDetails"`
	IndividualRecoveryKeyValidityStatus string                                                `json:"individualRecoveryKeyValidityStatus"`
	InstitutionalRecoveryKeyPresent     bool                                                  `json:"institutionalRecoveryKeyPresent"`
	DiskEncryptionConfigurationName     string                                                `json:"diskEncryptionConfigurationName"`
	FileVault2Enabled                   bool                                                  `json:"fileVault2Enabled"`
	FileVault2EnabledUserNames          []string                                              `json:"fileVault2EnabledUserNames"`
	FileVault2EligibilityMessage        string                                                `json:"fileVault2EligibilityMessage"`
}

// Purchasing

type ComputerInventorySubsetPurchasing struct {
	Leased              bool                                        `json:"leased"`
	Purchased           bool                                        `json:"purchased"`
	PoNumber            string                                      `json:"poNumber"`
	PoDate              string                                      `json:"poDate"`
	Vendor              string                                      `json:"vendor"`
	WarrantyDate        string                                      `json:"warrantyDate"`
	AppleCareId         string                                      `json:"appleCareId"`
	LeaseDate           string                                      `json:"leaseDate"`
	PurchasePrice       string                                      `json:"purchasePrice"`
	LifeExpectancy      int                                         `json:"lifeExpectancy"`
	PurchasingAccount   string                                      `json:"purchasingAccount"`
	PurchasingContact   string                                      `json:"purchasingContact"`
	ExtensionAttributes []ComputerInventorySubsetExtensionAttribute `json:"extensionAttributes"`
}

// Applications

type ComputerInventorySubsetApplication struct {
	Name              string `json:"name"`
	Path              string `json:"path"`
	Version           string `json:"version"`
	MacAppStore       bool   `json:"macAppStore"`
	SizeMegabytes     int    `json:"sizeMegabytes"`
	BundleId          string `json:"bundleId"`
	UpdateAvailable   bool   `json:"updateAvailable"`
	ExternalVersionId string `json:"externalVersionId"`
}

// Storage

type ComputerInventorySubsetStorage struct {
	BootDriveAvailableSpaceMegabytes int                                  `json:"bootDriveAvailableSpaceMegabytes"`
	Disks                            []ComputerInventorySubsetStorageDisk `json:"disks"`
}

type ComputerInventorySubsetStorageDisk struct {
	ID            string                                        `json:"id"`
	Device        string                                        `json:"device"`
	Model         string                                        `json:"model"`
	Revision      string                                        `json:"revision"`
	SerialNumber  string                                        `json:"serialNumber"`
	SizeMegabytes int                                           `json:"sizeMegabytes"`
	SmartStatus   string                                        `json:"smartStatus"`
	Type          string                                        `json:"type"`
	Partitions    []ComputerInventorySubsetStorageDiskPartition `json:"partitions"`
}

type ComputerInventorySubsetStorageDiskPartition struct {
	Name                      string `json:"name"`
	SizeMegabytes             int    `json:"sizeMegabytes"`
	AvailableMegabytes        int    `json:"availableMegabytes"`
	PartitionType             string `json:"partitionType"`
	PercentUsed               int    `json:"percentUsed"`
	FileVault2State           string `json:"fileVault2State"`
	FileVault2ProgressPercent int    `json:"fileVault2ProgressPercent"`
	LvmManaged                bool   `json:"lvmManaged"`
}

// User and Location

type ComputerInventorySubsetUserAndLocation struct {
	Username            string                                      `json:"username"`
	Realname            string                                      `json:"realname"`
	Email               string                                      `json:"email"`
	Position            string                                      `json:"position"`
	Phone               string                                      `json:"phone"`
	DepartmentId        string                                      `json:"departmentId"`
	BuildingId          string                                      `json:"buildingId"`
	Room                string                                      `json:"room"`
	ExtensionAttributes []ComputerInventorySubsetExtensionAttribute `json:"extensionAttributes"`
}

// Configuration Profiles

type ComputerInventorySubsetConfigurationProfile struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	LastInstalled     string `json:"lastInstalled"`
	Removable         bool   `json:"removable"`
	DisplayName       string `json:"displayName"`
	ProfileIdentifier string `json:"profileIdentifier"`
}

// Printers

type ComputerInventorySubsetPrinter struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	URI      string `json:"uri"`
	Location string `json:"location"`
}

// Services

type ComputerInventorySubsetService struct {
	Name string `json:"name"`
}

// Hardware

type ComputerInventorySubsetHardware struct {
	Make                   string                                      `json:"make"`
	Model                  string                                      `json:"model"`
	ModelIdentifier        string                                      `json:"modelIdentifier"`
	SerialNumber           string                                      `json:"serialNumber"`
	ProcessorSpeedMhz      int                                         `json:"processorSpeedMhz"`
	ProcessorCount         int                                         `json:"processorCount"`
	CoreCount              int                                         `json:"coreCount"`
	ProcessorType          string                                      `json:"processorType"`
	ProcessorArchitecture  string                                      `json:"processorArchitecture"`
	BusSpeedMhz            int                                         `json:"busSpeedMhz"`
	CacheSizeKilobytes     int                                         `json:"cacheSizeKilobytes"`
	NetworkAdapterType     string                                      `json:"networkAdapterType"`
	MacAddress             string                                      `json:"macAddress"`
	AltNetworkAdapterType  string                                      `json:"altNetworkAdapterType"`
	AltMacAddress          string                                      `json:"altMacAddress"`
	TotalRamMegabytes      int                                         `json:"totalRamMegabytes"`
	OpenRamSlots           int                                         `json:"openRamSlots"`
	BatteryCapacityPercent int                                         `json:"batteryCapacityPercent"`
	BatteryHealth          string                                      `json:"batteryHealth"`
	SmcVersion             string                                      `json:"smcVersion"`
	NicSpeed               string                                      `json:"nicSpeed"`
	OpticalDrive           string                                      `json:"opticalDrive"`
	BootRom                string                                      `json:"bootRom"`
	BleCapable             bool                                        `json:"bleCapable"`
	SupportsIosAppInstalls bool                                        `json:"supportsIosAppInstalls"`
	AppleSilicon           bool                                        `json:"appleSilicon"`
	ProvisioningUdid       string                                      `json:"provisioningUdid"`
	ExtensionAttributes    []ComputerInventorySubsetExtensionAttribute `json:"extensionAttributes"`
}

// Local User Accounts

type ComputerInventorySubsetLocalUserAccount struct {
	UID                            string `json:"uid"`
	UserGuid                       string `json:"userGuid"`
	Username                       string `json:"username"`
	FullName                       string `json:"fullName"`
	Admin                          bool   `json:"admin"`
	HomeDirectory                  string `json:"homeDirectory"`
	HomeDirectorySizeMb            int    `json:"homeDirectorySizeMb"`
	FileVault2Enabled              bool   `json:"fileVault2Enabled"`
	UserAccountType                string `json:"userAccountType"`
	PasswordMinLength              int    `json:"passwordMinLength"`
	PasswordMaxAge                 int    `json:"passwordMaxAge"`
	PasswordMinComplexCharacters   int    `json:"passwordMinComplexCharacters"`
	PasswordHistoryDepth           int    `json:"passwordHistoryDepth"`
	PasswordRequireAlphanumeric    bool   `json:"passwordRequireAlphanumeric"`
	ComputerAzureActiveDirectoryId string `json:"computerAzureActiveDirectoryId"`
	UserAzureActiveDirectoryId     string `json:"userAzureActiveDirectoryId"`
	AzureActiveDirectoryId         string `json:"azureActiveDirectoryId"`
}

// Certificates

type ComputerInventorySubsetCertificate struct {
	CommonName        string `json:"commonName"`
	Identity          bool   `json:"identity"`
	ExpirationDate    string `json:"expirationDate"`
	Username          string `json:"username"`
	LifecycleStatus   string `json:"lifecycleStatus"`
	CertificateStatus string `json:"certificateStatus"`
	SubjectName       string `json:"subjectName"`
	SerialNumber      string `json:"serialNumber"`
	Sha1Fingerprint   string `json:"sha1Fingerprint"`
	IssuedDate        string `json:"issuedDate"`
}

// Attachments

type ComputerInventorySubsetAttachment struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	FileType  string `json:"fileType"`
	SizeBytes int    `json:"sizeBytes"`
}

// Plugins

type ComputerInventorySubsetPlugin struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

// Package Receipts

type ComputerInventorySubsetPackageReceipts struct {
	InstalledByJamfPro      []string `json:"installedByJamfPro"`
	InstalledByInstallerSwu []string `json:"installedByInstallerSwu"`
	Cached                  []string `json:"cached"`
}

// Fonts

type ComputerInventorySubsetFont struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

// Security

type ComputerInventorySubsetSecurity struct {
	SipStatus                    string `json:"sipStatus"`
	GatekeeperStatus             string `json:"gatekeeperStatus"`
	XprotectVersion              string `json:"xprotectVersion"`
	AutoLoginDisabled            bool   `json:"autoLoginDisabled"`
	RemoteDesktopEnabled         bool   `json:"remoteDesktopEnabled"`
	ActivationLockEnabled        bool   `json:"activationLockEnabled"`
	RecoveryLockEnabled          bool   `json:"recoveryLockEnabled"`
	FirewallEnabled              bool   `json:"firewallEnabled"`
	SecureBootLevel              string `json:"secureBootLevel"`
	ExternalBootLevel            string `json:"externalBootLevel"`
	BootstrapTokenAllowed        bool   `json:"bootstrapTokenAllowed"`
	BootstrapTokenEscrowedStatus string `json:"bootstrapTokenEscrowedStatus"`
	LastAttestationAttempt       string `json:"lastAttestationAttempt"`
	LastSuccessfulAttestation    string `json:"lastSuccessfulAttestation"`
	AttestationStatus            string `json:"attestationStatus"`
}

// Operating System

type ComputerInventorySubsetOperatingSystem struct {
	Name                     string                                      `json:"name"`
	Version                  string                                      `json:"version"`
	Build                    string                                      `json:"build"`
	SupplementalBuildVersion string                                      `json:"supplementalBuildVersion"`
	RapidSecurityResponse    string                                      `json:"rapidSecurityResponse"`
	ActiveDirectoryStatus    string                                      `json:"activeDirectoryStatus"`
	FileVault2Status         string                                      `json:"fileVault2Status"`
	SoftwareUpdateDeviceId   string                                      `json:"softwareUpdateDeviceId"`
	ExtensionAttributes      []ComputerInventorySubsetExtensionAttribute `json:"extensionAttributes"`
}

// Licensed Software

type ComputerInventorySubsetLicensedSoftware struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// IBeacon

type ComputerInventorySubsetIBeacon struct {
	Name string `json:"name"`
}

// Software Updates

type ComputerInventorySubsetSoftwareUpdate struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	PackageName string `json:"packageName"`
}

// Content Caching

type ComputerInventorySubsetContentCaching struct {
	ComputerContentCachingInformationId string                                                  `json:"computerContentCachingInformationId"`
	Parents                             []ComputerInventorySubsetContentCachingParent           `json:"parents"`
	Alerts                              []ComputerInventorySubsetContentCachingAlert            `json:"alerts"`
	Activated                           bool                                                    `json:"activated"`
	Active                              bool                                                    `json:"active"`
	ActualCacheBytesUsed                int                                                     `json:"actualCacheBytesUsed"`
	CacheDetails                        []ComputerInventorySubsetContentCachingCacheDetail      `json:"cacheDetails"`
	CacheBytesFree                      int                                                     `json:"cacheBytesFree"`
	CacheBytesLimit                     int                                                     `json:"cacheBytesLimit"`
	CacheStatus                         string                                                  `json:"cacheStatus"`
	CacheBytesUsed                      int                                                     `json:"cacheBytesUsed"`
	DataMigrationCompleted              bool                                                    `json:"dataMigrationCompleted"`
	DataMigrationProgressPercentage     int                                                     `json:"dataMigrationProgressPercentage"`
	DataMigrationError                  ComputerInventorySubsetContentCachingDataMigrationError `json:"dataMigrationError"`
	MaxCachePressureLast1HourPercentage int                                                     `json:"maxCachePressureLast1HourPercentage"`
	PersonalCacheBytesFree              int                                                     `json:"personalCacheBytesFree"`
	PersonalCacheBytesLimit             int                                                     `json:"personalCacheBytesLimit"`
	PersonalCacheBytesUsed              int                                                     `json:"personalCacheBytesUsed"`
	Port                                int                                                     `json:"port"`
	PublicAddress                       string                                                  `json:"publicAddress"`
	RegistrationError                   string                                                  `json:"registrationError"`
	RegistrationResponseCode            int                                                     `json:"registrationResponseCode"`
	RegistrationStarted                 string                                                  `json:"registrationStarted"`
	RegistrationStatus                  string                                                  `json:"registrationStatus"`
	RestrictedMedia                     bool                                                    `json:"restrictedMedia"`
	ServerGuid                          string                                                  `json:"serverGuid"`
	StartupStatus                       string                                                  `json:"startupStatus"`
	TetheratorStatus                    string                                                  `json:"tetheratorStatus"`
	TotalBytesAreSince                  string                                                  `json:"totalBytesAreSince"`
	TotalBytesDropped                   int64                                                   `json:"totalBytesDropped"`
	TotalBytesImported                  int64                                                   `json:"totalBytesImported"`
	TotalBytesReturnedToChildren        int64                                                   `json:"totalBytesReturnedToChildren"`
	TotalBytesReturnedToClients         int64                                                   `json:"totalBytesReturnedToClients"`
	TotalBytesReturnedToPeers           int64                                                   `json:"totalBytesReturnedToPeers"`
	TotalBytesStoredFromOrigin          int64                                                   `json:"totalBytesStoredFromOrigin"`
	TotalBytesStoredFromParents         int64                                                   `json:"totalBytesStoredFromParents"`
	TotalBytesStoredFromPeers           int64                                                   `json:"totalBytesStoredFromPeers"`
}

type ComputerInventorySubsetContentCachingParent struct {
	ContentCachingParentId string                                             `json:"contentCachingParentId"`
	Address                string                                             `json:"address"`
	Alerts                 ComputerInventorySubsetContentCachingAlert         `json:"alerts"`
	Details                ComputerInventorySubsetContentCachingParentDetails `json:"details"`
	Guid                   string                                             `json:"guid"`
	Healthy                bool                                               `json:"healthy"`
	Port                   int                                                `json:"port"`
	Version                string                                             `json:"version"`
}

type ComputerInventorySubsetContentCachingParentDetails struct {
	ContentCachingParentDetailsId string                                                           `json:"contentCachingParentDetailsId"`
	AcPower                       bool                                                             `json:"acPower"`
	CacheSizeBytes                int64                                                            `json:"cacheSizeBytes"`
	Capabilities                  ComputerInventorySubsetContentCachingParentDetailsCapabilities   `json:"capabilities"`
	Portable                      bool                                                             `json:"portable"`
	LocalNetwork                  []ComputerInventorySubsetContentCachingParentDetailsLocalNetwork `json:"localNetwork"`
}

type ComputerInventorySubsetContentCachingParentDetailsCapabilities struct {
	ContentCachingParentCapabilitiesId string `json:"contentCachingParentCapabilitiesId"`
	Imports                            bool   `json:"imports"`
	Namespaces                         bool   `json:"namespaces"`
	PersonalContent                    bool   `json:"personalContent"`
	QueryParameters                    bool   `json:"queryParameters"`
	SharedContent                      bool   `json:"sharedContent"`
	Prioritization                     bool   `json:"prioritization"`
}

type ComputerInventorySubsetContentCachingParentDetailsLocalNetwork struct {
	ContentCachingParentLocalNetworkId string `json:"contentCachingParentLocalNetworkId"`
	Speed                              int    `json:"speed"`
	Wired                              bool   `json:"wired"`
}

type ComputerInventorySubsetContentCachingCacheDetail struct {
	ComputerContentCachingCacheDetailsId string `json:"computerContentCachingCacheDetailsId"`
	CategoryName                         string `json:"categoryName"`
	DiskSpaceBytesUsed                   int64  `json:"diskSpaceBytesUsed"`
}

type ComputerInventorySubsetContentCachingDataMigrationError struct {
	Code     int                                                               `json:"code"`
	Domain   string                                                            `json:"domain"`
	UserInfo []ComputerInventorySubsetContentCachingDataMigrationErrorUserInfo `json:"userInfo"`
}

type ComputerInventorySubsetContentCachingDataMigrationErrorUserInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Group Memberships

type ComputerInventorySubsetGroupMembership struct {
	GroupId          string `json:"groupId"`
	GroupName        string `json:"groupName"`
	GroupDescription string `json:"groupDescription"`
	SmartGroup       bool   `json:"smartGroup"`
}

// Shared

// ExtensionAttribute represents a generic extension attribute.
type ComputerInventorySubsetExtensionAttribute struct {
	DefinitionId string   `json:"definitionId"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Enabled      bool     `json:"enabled"`
	MultiValue   bool     `json:"multiValue"`
	Values       []string `json:"values"`
	DataType     string   `json:"dataType"`
	Options      []string `json:"options"`
	InputType    string   `json:"inputType"`
}

// BootPartitionEncryptionDetails represents the details of disk encryption.
type ComputerInventorySubsetBootPartitionEncryptionDetails struct {
	PartitionName              string `json:"partitionName"`
	PartitionFileVault2State   string `json:"partitionFileVault2State"`
	PartitionFileVault2Percent int    `json:"partitionFileVault2Percent"`
}

// ContentCachingAlert represents an alert in the content caching details.
type ComputerInventorySubsetContentCachingAlert struct {
	ContentCachingParentAlertId string   `json:"contentCachingParentAlertId"`
	Addresses                   []string `json:"addresses"`
	ClassName                   string   `json:"className"`
	PostDate                    string   `json:"postDate"`
	CacheBytesLimit             int      `json:"cacheBytesLimit"`
	PathPreventingAccess        string   `json:"pathPreventingAccess"`
	ReservedVolumeBytes         int      `json:"reservedVolumeBytes"`
	Resource                    string   `json:"resource"`
}

// FileVaultInventoryList represents the paginated FileVault inventory response.
type FileVaultInventoryList struct {
	TotalCount int                  `json:"totalCount"`
	Results    []FileVaultInventory `json:"results"`
}

// FileVaultInventory represents the FileVault information for a single computer.
type FileVaultInventory struct {
	ComputerId                          string                                                `json:"computerId"`
	Name                                string                                                `json:"name"`
	PersonalRecoveryKey                 string                                                `json:"personalRecoveryKey"`
	BootPartitionEncryptionDetails      ComputerInventorySubsetBootPartitionEncryptionDetails `json:"bootPartitionEncryptionDetails"`
	IndividualRecoveryKeyValidityStatus string                                                `json:"individualRecoveryKeyValidityStatus"`
	InstitutionalRecoveryKeyPresent     bool                                                  `json:"institutionalRecoveryKeyPresent"`
	DiskEncryptionConfigurationName     string                                                `json:"diskEncryptionConfigurationName"`
}

// ResponseRecoveryLockPassword represents the response structure for a recovery lock password.
type ResponseRecoveryLockPassword struct {
	RecoveryLockPassword string `json:"recoveryLockPassword"`
}

// ResponseUploadAttachment represents the response structure for uploading an attachment.
type ResponseUploadAttachment struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResponseRemoveMDMProfile represents the response structure for removing an MDM profile.
type ResponseRemoveMDMProfile struct {
	DeviceID    string `json:"deviceId"`
	CommandUUID string `json:"commandUuid"`
}

// Request

// RequestEraseDeviceComputer represents the request structure for erasing a device.
type RequestEraseDeviceComputer struct {
	Pin *string `json:"pin,omitempty"`
}
