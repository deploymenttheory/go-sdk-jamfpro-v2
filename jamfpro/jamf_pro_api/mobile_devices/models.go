package mobile_devices

// -----------------------------------------------------------------------------
// List wrappers
// -----------------------------------------------------------------------------

// MobileDeviceListResponse is the paginated response from GET /v2/mobile-devices.
type MobileDeviceListResponse struct {
	TotalCount int                    `json:"totalCount"`
	Results    []ResourceMobileDevice `json:"results"`
}

// MobileDeviceDetailListResponse is the paginated response from
// GET /v2/mobile-devices/detail (and /v2/mobile-devices/{id}/paired-devices).
type MobileDeviceDetailListResponse struct {
	TotalCount int                          `json:"totalCount"`
	Results    []ResourceMobileDeviceDetail `json:"results"`
}

// -----------------------------------------------------------------------------
// Basic Mobile Device (list item) — GET /v2/mobile-devices
// -----------------------------------------------------------------------------

// ResourceMobileDevice is the basic mobile device record returned by the
// non-detail list and get-by-id endpoints.
type ResourceMobileDevice struct {
	ID                     string `json:"id,omitempty"`
	Name                   string `json:"name,omitempty"`
	SerialNumber           string `json:"serialNumber,omitempty"`
	WifiMacAddress         string `json:"wifiMacAddress,omitempty"`
	Udid                   string `json:"udid,omitempty"`
	PhoneNumber            string `json:"phoneNumber,omitempty"`
	Model                  string `json:"model,omitempty"`
	ModelIdentifier        string `json:"modelIdentifier,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
	ManagementID           string `json:"managementId,omitempty"`
	SoftwareUpdateDeviceID string `json:"softwareUpdateDeviceId,omitempty"`
}

// -----------------------------------------------------------------------------
// Detailed Mobile Device (inventory detail) — GET /v2/mobile-devices/detail
//
// The detail object is built from named sections. Each section is only present
// when the corresponding `section` query parameter is requested, so every
// section is modelled as a pointer with omitempty.
// -----------------------------------------------------------------------------

// ResourceMobileDeviceDetail is a single full mobile device inventory record.
type ResourceMobileDeviceDetail struct {
	MobileDeviceID string `json:"mobileDeviceId,omitempty"`
	// DeviceType determines which device-type-specific data is populated.
	// One of ios, appleTv, watch, visionOS (see api_docs note on the enum).
	DeviceType string `json:"deviceType,omitempty"`

	General              *MobileDeviceSubsetGeneral              `json:"general,omitempty"`
	Hardware             *MobileDeviceSubsetHardware             `json:"hardware,omitempty"`
	UserAndLocation      *MobileDeviceSubsetUserAndLocation      `json:"userAndLocation,omitempty"`
	Purchasing           *MobileDeviceSubsetPurchasing           `json:"purchasing,omitempty"`
	Security             *MobileDeviceSubsetSecurity             `json:"security,omitempty"`
	Network              *MobileDeviceSubsetNetwork              `json:"network,omitempty"`
	Applications         []MobileDeviceSubsetApplication         `json:"applications,omitempty"`
	Ebooks               []MobileDeviceSubsetEbook               `json:"ebooks,omitempty"`
	Certificates         []MobileDeviceSubsetCertificate         `json:"certificates,omitempty"`
	Profiles             []MobileDeviceSubsetProfile             `json:"profiles,omitempty"`
	UserProfiles         []MobileDeviceSubsetUserProfile         `json:"userProfiles,omitempty"`
	ProvisioningProfiles []MobileDeviceSubsetProvisioningProfile `json:"provisioningProfiles,omitempty"`
	ServiceSubscriptions []MobileDeviceSubsetServiceSubscription `json:"serviceSubscriptions,omitempty"`
	SharedUsers          []MobileDeviceSubsetSharedUser          `json:"sharedUsers,omitempty"`
	Groups               []MobileDeviceSubsetGroup               `json:"groups,omitempty"`
	ExtensionAttributes  []MobileDeviceSubsetExtensionAttribute  `json:"extensionAttributes,omitempty"`
}

// MobileDeviceSubsetExtensionAttribute is a single inventory extension attribute.
type MobileDeviceSubsetExtensionAttribute struct {
	ID                                  string   `json:"id,omitempty"`
	Name                                string   `json:"name,omitempty"`
	Type                                string   `json:"type,omitempty"`
	Value                               []string `json:"value,omitempty"`
	ExtensionAttributeCollectionAllowed bool     `json:"extensionAttributeCollectionAllowed,omitempty"`
	InventoryDisplay                    string   `json:"inventoryDisplay,omitempty"`
}

// MobileDeviceSubsetHardware is the HARDWARE section.
type MobileDeviceSubsetHardware struct {
	CapacityMb                int                                    `json:"capacityMb,omitempty"`
	AvailableSpaceMb          int                                    `json:"availableSpaceMb,omitempty"`
	UsedSpacePercentage       int                                    `json:"usedSpacePercentage,omitempty"`
	BatteryLevel              int                                    `json:"batteryLevel,omitempty"`
	BatteryHealth             string                                 `json:"batteryHealth,omitempty"`
	SerialNumber              string                                 `json:"serialNumber,omitempty"`
	WifiMacAddress            string                                 `json:"wifiMacAddress,omitempty"`
	BluetoothMacAddress       string                                 `json:"bluetoothMacAddress,omitempty"`
	ModemFirmwareVersion      string                                 `json:"modemFirmwareVersion,omitempty"`
	Model                     string                                 `json:"model,omitempty"`
	ModelIdentifier           string                                 `json:"modelIdentifier,omitempty"`
	ModelNumber               string                                 `json:"modelNumber,omitempty"`
	BluetoothLowEnergyCapable bool                                   `json:"bluetoothLowEnergyCapable,omitempty"`
	DeviceID                  string                                 `json:"deviceId,omitempty"`
	ExtensionAttributes       []MobileDeviceSubsetExtensionAttribute `json:"extensionAttributes,omitempty"`
}

// MobileDeviceSubsetUserAndLocation is the USER_AND_LOCATION section.
type MobileDeviceSubsetUserAndLocation struct {
	Username            string                                 `json:"username,omitempty"`
	RealName            string                                 `json:"realName,omitempty"`
	EmailAddress        string                                 `json:"emailAddress,omitempty"`
	Position            string                                 `json:"position,omitempty"`
	PhoneNumber         string                                 `json:"phoneNumber,omitempty"`
	DepartmentID        string                                 `json:"departmentId,omitempty"`
	BuildingID          string                                 `json:"buildingId,omitempty"`
	Room                string                                 `json:"room,omitempty"`
	Building            string                                 `json:"building,omitempty"`
	Department          string                                 `json:"department,omitempty"`
	ExtensionAttributes []MobileDeviceSubsetExtensionAttribute `json:"extensionAttributes,omitempty"`

	// 11.29 additive MDM self-service login fields.
	LastLoggedInUsernameMdm          string `json:"lastLoggedInUsernameMdm,omitempty"`
	LastLoggedInUsernameMdmTimestamp string `json:"lastLoggedInUsernameMdmTimestamp,omitempty"`
}

// MobileDeviceSubsetGeneral is the GENERAL section.
type MobileDeviceSubsetGeneral struct {
	Udid                                        string                                      `json:"udid,omitempty"`
	DisplayName                                 string                                      `json:"displayName,omitempty"`
	AssetTag                                    string                                      `json:"assetTag,omitempty"`
	SiteID                                      string                                      `json:"siteId,omitempty"`
	LastInventoryUpdateDate                     string                                      `json:"lastInventoryUpdateDate,omitempty"`
	OsVersion                                   string                                      `json:"osVersion,omitempty"`
	OsRapidSecurityResponse                     string                                      `json:"osRapidSecurityResponse,omitempty"`
	OsBuild                                     string                                      `json:"osBuild,omitempty"`
	OsSupplementalBuildVersion                  string                                      `json:"osSupplementalBuildVersion,omitempty"`
	SoftwareUpdateDeviceID                      string                                      `json:"softwareUpdateDeviceId,omitempty"`
	IPAddress                                   string                                      `json:"ipAddress,omitempty"`
	Managed                                     bool                                        `json:"managed,omitempty"`
	Supervised                                  bool                                        `json:"supervised,omitempty"`
	DeviceOwnershipType                         string                                      `json:"deviceOwnershipType,omitempty"`
	EnrollmentMethodPrestage                    *MobileDeviceSubsetEnrollmentMethodPrestage `json:"enrollmentMethodPrestage,omitempty"`
	EnrollmentSessionTokenValid                 bool                                        `json:"enrollmentSessionTokenValid,omitempty"`
	LastEnrolledDate                            string                                      `json:"lastEnrolledDate,omitempty"`
	MdmProfileExpirationDate                    string                                      `json:"mdmProfileExpirationDate,omitempty"`
	TimeZone                                    string                                      `json:"timeZone,omitempty"`
	DeclarativeDeviceManagementEnabled          bool                                        `json:"declarativeDeviceManagementEnabled,omitempty"`
	ManagementID                                string                                      `json:"managementId,omitempty"`
	ExtensionAttributes                         []MobileDeviceSubsetExtensionAttribute      `json:"extensionAttributes,omitempty"`
	LastLoggedInUsernameSelfService             string                                      `json:"lastLoggedInUsernameSelfService,omitempty"`
	LastLoggedInUsernameSelfServiceTimestamp    string                                      `json:"lastLoggedInUsernameSelfServiceTimestamp,omitempty"`
	DiagnosticAndUsageReportingEnabled          bool                                        `json:"diagnosticAndUsageReportingEnabled,omitempty"`
	AppAnalyticsEnabled                         bool                                        `json:"appAnalyticsEnabled,omitempty"`
	DeviceLocatorServiceEnabled                 bool                                        `json:"deviceLocatorServiceEnabled,omitempty"`
	DoNotDisturbEnabled                         bool                                        `json:"doNotDisturbEnabled,omitempty"`
	LastCloudBackupDate                         string                                      `json:"lastCloudBackupDate,omitempty"`
	ItunesStoreAccountActive                    bool                                        `json:"itunesStoreAccountActive,omitempty"`
	AirPlayPassword                             string                                      `json:"airPlayPassword,omitempty"`
	Locales                                     string                                      `json:"locales,omitempty"`
	Languages                                   string                                      `json:"languages,omitempty"`
	SharedIpad                                  bool                                        `json:"sharedIpad,omitempty"`
	ResidentUsers                               int                                         `json:"residentUsers,omitempty"`
	QuotaSize                                   int                                         `json:"quotaSize,omitempty"`
	TemporarySessionOnly                        bool                                        `json:"temporarySessionOnly,omitempty"`
	TemporarySessionTimeout                     int                                         `json:"temporarySessionTimeout,omitempty"`
	UserSessionTimeout                          int                                         `json:"userSessionTimeout,omitempty"`
	SyncedToComputer                            int                                         `json:"syncedToComputer,omitempty"`
	MaximumSharediPadUsersStored                int                                         `json:"maximumSharediPadUsersStored,omitempty"`
	LastBackupDate                              string                                      `json:"lastBackupDate,omitempty"`
	CloudBackupEnabled                          bool                                        `json:"cloudBackupEnabled,omitempty"`
	LocationServicesForSelfServiceMobileEnabled bool                                        `json:"locationServicesForSelfServiceMobileEnabled,omitempty"`
	ExchangeDeviceID                            string                                      `json:"exchangeDeviceId,omitempty"`
	Tethered                                    bool                                        `json:"tethered,omitempty"`
}

// MobileDeviceSubsetEnrollmentMethodPrestage describes the prestage used at enrollment.
type MobileDeviceSubsetEnrollmentMethodPrestage struct {
	MobileDevicePrestageID string `json:"mobileDevicePrestageId,omitempty"`
	ProfileName            string `json:"profileName,omitempty"`
}

// MobileDeviceSubsetSecurity is the SECURITY section (iOS only).
type MobileDeviceSubsetSecurity struct {
	DataProtected                          bool                                `json:"dataProtected,omitempty"`
	BlockLevelEncryptionCapable            bool                                `json:"blockLevelEncryptionCapable,omitempty"`
	FileLevelEncryptionCapable             bool                                `json:"fileLevelEncryptionCapable,omitempty"`
	PasscodePresent                        bool                                `json:"passcodePresent,omitempty"`
	PasscodeCompliant                      bool                                `json:"passcodeCompliant,omitempty"`
	PasscodeCompliantWithProfile           bool                                `json:"passcodeCompliantWithProfile,omitempty"`
	HardwareEncryption                     int                                 `json:"hardwareEncryption,omitempty"`
	ActivationLockEnabled                  bool                                `json:"activationLockEnabled,omitempty"`
	JailBreakDetected                      bool                                `json:"jailBreakDetected,omitempty"`
	AttestationStatus                      string                              `json:"attestationStatus,omitempty"`
	LastAttestationAttemptDate             string                              `json:"lastAttestationAttemptDate,omitempty"`
	LastSuccessfulAttestationDate          string                              `json:"lastSuccessfulAttestationDate,omitempty"`
	PasscodeLockGracePeriodEnforcedSeconds int                                 `json:"passcodeLockGracePeriodEnforcedSeconds,omitempty"`
	PersonalDeviceProfileCurrent           bool                                `json:"personalDeviceProfileCurrent,omitempty"`
	LostModeEnabled                        bool                                `json:"lostModeEnabled,omitempty"`
	LostModePersistent                     bool                                `json:"lostModePersistent,omitempty"`
	LostModeMessage                        string                              `json:"lostModeMessage,omitempty"`
	LostModePhoneNumber                    string                              `json:"lostModePhoneNumber,omitempty"`
	LostModeFootnote                       string                              `json:"lostModeFootnote,omitempty"`
	LostModeLocation                       *MobileDeviceSubsetLostModeLocation `json:"lostModeLocation,omitempty"`
	BootstrapTokenEscrowed                 string                              `json:"bootstrapTokenEscrowed,omitempty"`
}

// MobileDeviceSubsetLostModeLocation describes the last lost-mode location.
type MobileDeviceSubsetLostModeLocation struct {
	LastLocationUpdate                       string  `json:"lastLocationUpdate,omitempty"`
	LostModeLocationHorizontalAccuracyMeters float64 `json:"lostModeLocationHorizontalAccuracyMeters,omitempty"`
	LostModeLocationVerticalAccuracyMeters   float64 `json:"lostModeLocationVerticalAccuracyMeters,omitempty"`
	LostModeLocationAltitudeMeters           float64 `json:"lostModeLocationAltitudeMeters,omitempty"`
	LostModeLocationSpeedMetersPerSecond     float64 `json:"lostModeLocationSpeedMetersPerSecond,omitempty"`
	LostModeLocationCourseDegrees            float64 `json:"lostModeLocationCourseDegrees,omitempty"`
	LostModeLocationTimestamp                string  `json:"lostModeLocationTimestamp,omitempty"`
}

// MobileDeviceSubsetNetwork is the NETWORK section (iOS only).
type MobileDeviceSubsetNetwork struct {
	CellularTechnology       string `json:"cellularTechnology,omitempty"`
	VoiceRoamingEnabled      bool   `json:"voiceRoamingEnabled,omitempty"`
	Imei                     string `json:"imei,omitempty"`
	Iccid                    string `json:"iccid,omitempty"`
	Meid                     string `json:"meid,omitempty"`
	Eid                      string `json:"eid,omitempty"`
	CarrierSettingsVersion   string `json:"carrierSettingsVersion,omitempty"`
	CurrentCarrierNetwork    string `json:"currentCarrierNetwork,omitempty"`
	CurrentMobileCountryCode string `json:"currentMobileCountryCode,omitempty"`
	CurrentMobileNetworkCode string `json:"currentMobileNetworkCode,omitempty"`
	HomeCarrierNetwork       string `json:"homeCarrierNetwork,omitempty"`
	HomeMobileCountryCode    string `json:"homeMobileCountryCode,omitempty"`
	HomeMobileNetworkCode    string `json:"homeMobileNetworkCode,omitempty"`
	DataRoamingEnabled       bool   `json:"dataRoamingEnabled,omitempty"`
	Roaming                  bool   `json:"roaming,omitempty"`
	PersonalHotspotEnabled   bool   `json:"personalHotspotEnabled,omitempty"`
	PhoneNumber              string `json:"phoneNumber,omitempty"`
	PreferredVoiceNumber     string `json:"preferredVoiceNumber,omitempty"`
}

// MobileDeviceSubsetServiceSubscription is a single SERVICE_SUBSCRIPTIONS entry.
type MobileDeviceSubsetServiceSubscription struct {
	CarrierSettingsVersion   string `json:"carrierSettingsVersion,omitempty"`
	CurrentCarrierNetwork    string `json:"currentCarrierNetwork,omitempty"`
	CurrentMobileCountryCode string `json:"currentMobileCountryCode,omitempty"`
	CurrentMobileNetworkCode string `json:"currentMobileNetworkCode,omitempty"`
	SubscriberCarrierNetwork string `json:"subscriberCarrierNetwork,omitempty"`
	Eid                      string `json:"eid,omitempty"`
	Iccid                    string `json:"iccid,omitempty"`
	Imei                     string `json:"imei,omitempty"`
	DataPreferred            bool   `json:"dataPreferred,omitempty"`
	Roaming                  bool   `json:"roaming,omitempty"`
	VoicePreferred           bool   `json:"voicePreferred,omitempty"`
	Label                    string `json:"label,omitempty"`
	LabelID                  string `json:"labelId,omitempty"`
	Meid                     string `json:"meid,omitempty"`
	PhoneNumber              string `json:"phoneNumber,omitempty"`
	Slot                     string `json:"slot,omitempty"`
}

// MobileDeviceSubsetApplication is a single APPLICATIONS entry.
type MobileDeviceSubsetApplication struct {
	Identifier       string `json:"identifier,omitempty"`
	Name             string `json:"name,omitempty"`
	Version          string `json:"version,omitempty"`
	ShortVersion     string `json:"shortVersion,omitempty"`
	ManagementStatus string `json:"managementStatus,omitempty"`
	ValidationStatus bool   `json:"validationStatus,omitempty"`
	BundleSize       string `json:"bundleSize,omitempty"`
	DynamicSize      string `json:"dynamicSize,omitempty"`
}

// MobileDeviceSubsetEbook is a single EBOOKS entry.
type MobileDeviceSubsetEbook struct {
	Author          string `json:"author,omitempty"`
	Title           string `json:"title,omitempty"`
	Version         string `json:"version,omitempty"`
	Kind            string `json:"kind,omitempty"`
	ManagementState string `json:"managementState,omitempty"`
}

// MobileDeviceSubsetCertificate is a single CERTIFICATES entry.
type MobileDeviceSubsetCertificate struct {
	CommonName     string `json:"commonName,omitempty"`
	Identity       bool   `json:"identity,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
}

// MobileDeviceSubsetProfile is a single PROFILES entry.
type MobileDeviceSubsetProfile struct {
	DisplayName   string `json:"displayName,omitempty"`
	Version       string `json:"version,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	Identifier    string `json:"identifier,omitempty"`
	Removable     bool   `json:"removable,omitempty"`
	LastInstalled string `json:"lastInstalled,omitempty"`
}

// MobileDeviceSubsetUserProfile is a single USER_PROFILES entry.
type MobileDeviceSubsetUserProfile struct {
	DisplayName   string `json:"displayName,omitempty"`
	Version       string `json:"version,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	Identifier    string `json:"identifier,omitempty"`
	Removable     bool   `json:"removable,omitempty"`
	LastInstalled string `json:"lastInstalled,omitempty"`
	Username      string `json:"username,omitempty"`
}

// MobileDeviceSubsetProvisioningProfile is a single PROVISIONING_PROFILES entry.
type MobileDeviceSubsetProvisioningProfile struct {
	DisplayName    string `json:"displayName,omitempty"`
	UUID           string `json:"uuid,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
}

// MobileDeviceSubsetSharedUser is a single SHARED_USERS entry.
type MobileDeviceSubsetSharedUser struct {
	ManagedAppleID string `json:"managedAppleId,omitempty"`
	LoggedIn       bool   `json:"loggedIn,omitempty"`
	DataToSync     bool   `json:"dataToSync,omitempty"`
}

// MobileDeviceSubsetGroup is a single GROUPS (smart group membership) entry.
type MobileDeviceSubsetGroup struct {
	GroupID          string `json:"groupId,omitempty"`
	GroupName        string `json:"groupName,omitempty"`
	GroupDescription string `json:"groupDescription,omitempty"`
	Smart            bool   `json:"smart,omitempty"`
}

// MobileDeviceSubsetPurchasing is the PURCHASING section.
type MobileDeviceSubsetPurchasing struct {
	Purchased           bool                                   `json:"purchased,omitempty"`
	Leased              bool                                   `json:"leased,omitempty"`
	PoNumber            string                                 `json:"poNumber,omitempty"`
	Vendor              string                                 `json:"vendor,omitempty"`
	AppleCareID         string                                 `json:"appleCareId,omitempty"`
	PurchasePrice       string                                 `json:"purchasePrice,omitempty"`
	PurchasingAccount   string                                 `json:"purchasingAccount,omitempty"`
	PoDate              string                                 `json:"poDate,omitempty"`
	WarrantyExpiresDate string                                 `json:"warrantyExpiresDate,omitempty"`
	LeaseExpiresDate    string                                 `json:"leaseExpiresDate,omitempty"`
	LifeExpectancy      int                                    `json:"lifeExpectancy,omitempty"`
	PurchasingContact   string                                 `json:"purchasingContact,omitempty"`
	ExtensionAttributes []MobileDeviceSubsetExtensionAttribute `json:"extensionAttributes,omitempty"`
}

// MobileDeviceSubsetSite is a site reference (id, name). DivisionID is an 11.29
// additive field.
type MobileDeviceSubsetSite struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	DivisionID string `json:"divisionId,omitempty"`
}

// -----------------------------------------------------------------------------
// Mobile Device Details (V2) — GET /v2/mobile-devices/{id}/detail
//
// This endpoint uses a different, newer schema family (MobileDeviceDetailsGetV2)
// than the paginated /detail list. Device-type-specific data is carried in the
// ios / tvos / watchos / visionos objects, selected by the type discriminator.
// -----------------------------------------------------------------------------

// ResourceMobileDeviceDetailsV2 is the full mobile device record returned by
// GET /v2/mobile-devices/{id}/detail.
type ResourceMobileDeviceDetailsV2 struct {
	ID                                 string                                   `json:"id,omitempty"`
	Name                               string                                   `json:"name,omitempty"`
	EnforceName                        bool                                     `json:"enforceName,omitempty"`
	AssetTag                           string                                   `json:"assetTag,omitempty"`
	LastInventoryUpdateTimestamp       string                                   `json:"lastInventoryUpdateTimestamp,omitempty"`
	OsVersion                          string                                   `json:"osVersion,omitempty"`
	OsBuild                            string                                   `json:"osBuild,omitempty"`
	OsSupplementalBuildVersion         string                                   `json:"osSupplementalBuildVersion,omitempty"`
	OsRapidSecurityResponse            string                                   `json:"osRapidSecurityResponse,omitempty"`
	SoftwareUpdateDeviceID             string                                   `json:"softwareUpdateDeviceId,omitempty"`
	SerialNumber                       string                                   `json:"serialNumber,omitempty"`
	Udid                               string                                   `json:"udid,omitempty"`
	IPAddress                          string                                   `json:"ipAddress,omitempty"`
	WifiMacAddress                     string                                   `json:"wifiMacAddress,omitempty"`
	BluetoothMacAddress                string                                   `json:"bluetoothMacAddress,omitempty"`
	Managed                            bool                                     `json:"managed,omitempty"`
	TimeZone                           string                                   `json:"timeZone,omitempty"`
	InitialEntryTimestamp              string                                   `json:"initialEntryTimestamp,omitempty"`
	LastEnrollmentTimestamp            string                                   `json:"lastEnrollmentTimestamp,omitempty"`
	MdmProfileExpirationTimestamp      string                                   `json:"mdmProfileExpirationTimestamp,omitempty"`
	DeviceOwnershipLevel               string                                   `json:"deviceOwnershipLevel,omitempty"`
	EnrollmentMethod                   string                                   `json:"enrollmentMethod,omitempty"`
	EnrollmentSessionTokenValid        bool                                     `json:"enrollmentSessionTokenValid,omitempty"`
	DeclarativeDeviceManagementEnabled bool                                     `json:"declarativeDeviceManagementEnabled,omitempty"`
	Site                               *MobileDeviceSubsetSite                  `json:"site,omitempty"`
	ExtensionAttributes                []MobileDeviceSubsetExtensionAttributeV2 `json:"extensionAttributes,omitempty"`
	Location                           *MobileDeviceSubsetLocationV2            `json:"location,omitempty"`
	Type                               string                                   `json:"type,omitempty"`
	Ios                                *MobileDeviceSubsetDetailsV2             `json:"ios,omitempty"`
	Tvos                               *MobileDeviceSubsetTvOsDetails           `json:"tvos,omitempty"`
	Watchos                            *MobileDeviceSubsetWatchOsDetailsV2      `json:"watchos,omitempty"`
	Visionos                           *MobileDeviceSubsetDetailsV2             `json:"visionos,omitempty"`

	// MobileDeviceDetailsGetV2 additive (allOf inline) fields.
	ManagementID string                    `json:"managementId,omitempty"`
	Groups       []MobileDeviceSubsetGroup `json:"groups,omitempty"`
}

// MobileDeviceSubsetExtensionAttributeV2 is the extension attribute shape used by
// the /{id}/detail endpoint (no inventoryDisplay field).
type MobileDeviceSubsetExtensionAttributeV2 struct {
	ID                                  string   `json:"id,omitempty"`
	Name                                string   `json:"name,omitempty"`
	Type                                string   `json:"type,omitempty"`
	Value                               []string `json:"value,omitempty"`
	ExtensionAttributeCollectionAllowed bool     `json:"extensionAttributeCollectionAllowed,omitempty"`
}

// MobileDeviceSubsetLocationV2 is the location object for the /{id}/detail endpoint.
type MobileDeviceSubsetLocationV2 struct {
	Username     string `json:"username,omitempty"`
	RealName     string `json:"realName,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	Position     string `json:"position,omitempty"`
	PhoneNumber  string `json:"phoneNumber,omitempty"`
	DepartmentID string `json:"departmentId,omitempty"`
	BuildingID   string `json:"buildingId,omitempty"`
	Room         string `json:"room,omitempty"`
}

// MobileDeviceSubsetDetailsV2 is the ios / visionos device-type detail object.
type MobileDeviceSubsetDetailsV2 struct {
	Model                       string                                   `json:"model,omitempty"`
	ModelIdentifier             string                                   `json:"modelIdentifier,omitempty"`
	ModelNumber                 string                                   `json:"modelNumber,omitempty"`
	Supervised                  bool                                     `json:"supervised,omitempty"`
	BatteryLevel                int                                      `json:"batteryLevel,omitempty"`
	BatteryHealth               string                                   `json:"batteryHealth,omitempty"`
	LastBackupTimestamp         string                                   `json:"lastBackupTimestamp,omitempty"`
	CapacityMb                  int                                      `json:"capacityMb,omitempty"`
	AvailableMb                 int                                      `json:"availableMb,omitempty"`
	PercentageUsed              int                                      `json:"percentageUsed,omitempty"`
	Shared                      bool                                     `json:"shared,omitempty"`
	DeviceLocatorServiceEnabled bool                                     `json:"deviceLocatorServiceEnabled,omitempty"`
	DoNotDisturbEnabled         bool                                     `json:"doNotDisturbEnabled,omitempty"`
	CloudBackupEnabled          bool                                     `json:"cloudBackupEnabled,omitempty"`
	LastCloudBackupTimestamp    string                                   `json:"lastCloudBackupTimestamp,omitempty"`
	LocationServicesEnabled     bool                                     `json:"locationServicesEnabled,omitempty"`
	ItunesStoreAccountActive    bool                                     `json:"iTunesStoreAccountActive,omitempty"`
	BleCapable                  bool                                     `json:"bleCapable,omitempty"`
	UnlockToken                 string                                   `json:"unlockToken,omitempty"`
	Computer                    *MobileDeviceSubsetIdAndNameV2           `json:"computer,omitempty"`
	Purchasing                  *MobileDeviceSubsetPurchasingV2          `json:"purchasing,omitempty"`
	Security                    *MobileDeviceSubsetSecurityV2            `json:"security,omitempty"`
	Network                     *MobileDeviceSubsetNetwork               `json:"network,omitempty"`
	ServiceSubscriptions        []MobileDeviceSubsetServiceSubscription  `json:"serviceSubscriptions,omitempty"`
	Applications                []MobileDeviceSubsetApplicationV2        `json:"applications,omitempty"`
	Certificates                []MobileDeviceSubsetCertificateV2        `json:"certificates,omitempty"`
	Ebooks                      []MobileDeviceSubsetEbookV2              `json:"ebooks,omitempty"`
	MdmCapableUsers             []MobileDeviceSubsetMdmCapableUser       `json:"mdmCapableUsers,omitempty"`
	ConfigurationProfiles       []MobileDeviceSubsetConfigurationProfile `json:"configurationProfiles,omitempty"`
	ProvisioningProfiles        []MobileDeviceSubsetProvisioningProfile  `json:"provisioningProfiles,omitempty"`
	Attachments                 []MobileDeviceSubsetAttachmentV2         `json:"attachments,omitempty"`
}

// MobileDeviceSubsetTvOsDetails is the tvos device-type detail object.
type MobileDeviceSubsetTvOsDetails struct {
	Model                 string                                   `json:"model,omitempty"`
	ModelIdentifier       string                                   `json:"modelIdentifier,omitempty"`
	ModelNumber           string                                   `json:"modelNumber,omitempty"`
	Supervised            bool                                     `json:"supervised,omitempty"`
	AirplayPassword       string                                   `json:"airplayPassword,omitempty"`
	DeviceID              string                                   `json:"deviceId,omitempty"`
	Locales               string                                   `json:"locales,omitempty"`
	Purchasing            *MobileDeviceSubsetPurchasingV2          `json:"purchasing,omitempty"`
	ConfigurationProfiles []MobileDeviceSubsetConfigurationProfile `json:"configurationProfiles,omitempty"`
	Certificates          []MobileDeviceSubsetCertificateV2        `json:"certificates,omitempty"`
	Applications          []MobileDeviceSubsetApplicationV2        `json:"applications,omitempty"`
}

// MobileDeviceSubsetWatchOsDetailsV2 is the watchos device-type detail object.
type MobileDeviceSubsetWatchOsDetailsV2 struct {
	Model                       string                                   `json:"model,omitempty"`
	ModelIdentifier             string                                   `json:"modelIdentifier,omitempty"`
	ModelNumber                 string                                   `json:"modelNumber,omitempty"`
	Supervised                  bool                                     `json:"supervised,omitempty"`
	BatteryLevel                int                                      `json:"batteryLevel,omitempty"`
	CapacityMb                  int                                      `json:"capacityMb,omitempty"`
	AvailableMb                 int                                      `json:"availableMb,omitempty"`
	PercentageUsed              int                                      `json:"percentageUsed,omitempty"`
	DeviceLocatorServiceEnabled bool                                     `json:"deviceLocatorServiceEnabled,omitempty"`
	DoNotDisturbEnabled         bool                                     `json:"doNotDisturbEnabled,omitempty"`
	LastCloudBackupTimestamp    string                                   `json:"lastCloudBackupTimestamp,omitempty"`
	ItunesStoreAccountActive    bool                                     `json:"iTunesStoreAccountActive,omitempty"`
	BleCapable                  bool                                     `json:"bleCapable,omitempty"`
	UnlockToken                 string                                   `json:"unlockToken,omitempty"`
	Security                    *MobileDeviceSubsetSecurityV2            `json:"security,omitempty"`
	Applications                []MobileDeviceSubsetApplicationV2        `json:"applications,omitempty"`
	Certificates                []MobileDeviceSubsetCertificateV2        `json:"certificates,omitempty"`
	ConfigurationProfiles       []MobileDeviceSubsetConfigurationProfile `json:"configurationProfiles,omitempty"`
	ProvisioningProfiles        []MobileDeviceSubsetProvisioningProfile  `json:"provisioningProfiles,omitempty"`
	Attachments                 []MobileDeviceSubsetAttachmentV2         `json:"attachments,omitempty"`
}

// MobileDeviceSubsetIdAndNameV2 is a simple id/name reference.
type MobileDeviceSubsetIdAndNameV2 struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// MobileDeviceSubsetPurchasingV2 is the purchasing object for the /{id}/detail endpoint.
type MobileDeviceSubsetPurchasingV2 struct {
	Purchased           bool   `json:"purchased,omitempty"`
	Leased              bool   `json:"leased,omitempty"`
	PoNumber            string `json:"poNumber,omitempty"`
	Vendor              string `json:"vendor,omitempty"`
	AppleCareID         string `json:"appleCareId,omitempty"`
	PurchasePrice       string `json:"purchasePrice,omitempty"`
	PurchasingAccount   string `json:"purchasingAccount,omitempty"`
	PoDate              string `json:"poDate,omitempty"`
	WarrantyExpiresDate string `json:"warrantyExpiresDate,omitempty"`
	LeaseExpiresDate    string `json:"leaseExpiresDate,omitempty"`
	LifeExpectancy      int    `json:"lifeExpectancy,omitempty"`
	PurchasingContact   string `json:"purchasingContact,omitempty"`
}

// MobileDeviceSubsetSecurityV2 is the security object for the /{id}/detail endpoint.
type MobileDeviceSubsetSecurityV2 struct {
	DataProtected                 bool   `json:"dataProtected,omitempty"`
	BlockLevelEncryptionCapable   bool   `json:"blockLevelEncryptionCapable,omitempty"`
	FileLevelEncryptionCapable    bool   `json:"fileLevelEncryptionCapable,omitempty"`
	PasscodePresent               bool   `json:"passcodePresent,omitempty"`
	PasscodeCompliant             bool   `json:"passcodeCompliant,omitempty"`
	PasscodeCompliantWithProfile  bool   `json:"passcodeCompliantWithProfile,omitempty"`
	HardwareEncryption            int    `json:"hardwareEncryption,omitempty"`
	ActivationLockEnabled         bool   `json:"activationLockEnabled,omitempty"`
	JailBreakDetected             bool   `json:"jailBreakDetected,omitempty"`
	AttestationStatus             string `json:"attestationStatus,omitempty"`
	LastAttestationAttemptDate    string `json:"lastAttestationAttemptDate,omitempty"`
	LastSuccessfulAttestationDate string `json:"lastSuccessfulAttestationDate,omitempty"`
	BootstrapToken                string `json:"bootstrapToken,omitempty"`
	BootstrapTokenEscrowed        string `json:"bootstrapTokenEscrowed,omitempty"`
}

// MobileDeviceSubsetApplicationV2 is the application object for the /{id}/detail endpoint.
type MobileDeviceSubsetApplicationV2 struct {
	Identifier   string `json:"identifier,omitempty"`
	Name         string `json:"name,omitempty"`
	Version      string `json:"version,omitempty"`
	ShortVersion string `json:"shortVersion,omitempty"`
}

// MobileDeviceSubsetCertificateV2 is the certificate object for the /{id}/detail endpoint.
type MobileDeviceSubsetCertificateV2 struct {
	CommonName          string `json:"commonName,omitempty"`
	Identity            bool   `json:"identity,omitempty"`
	ExpirationDateEpoch string `json:"expirationDateEpoch,omitempty"`
	SubjectName         string `json:"subjectName,omitempty"`
	SerialNumber        string `json:"serialNumber,omitempty"`
	Sha1Fingerprint     string `json:"sha1Fingerprint,omitempty"`
	IssuedDateEpoch     string `json:"issuedDateEpoch,omitempty"`
	CertificateStatus   string `json:"certificateStatus,omitempty"`
	LifecycleStatus     string `json:"lifecycleStatus,omitempty"`
}

// MobileDeviceSubsetEbookV2 is the ebook object for the /{id}/detail endpoint.
type MobileDeviceSubsetEbookV2 struct {
	Author  string `json:"author,omitempty"`
	Title   string `json:"title,omitempty"`
	Version string `json:"version,omitempty"`
}

// MobileDeviceSubsetMdmCapableUser is a single mdmCapableUsers entry.
type MobileDeviceSubsetMdmCapableUser struct {
	UserShortName string `json:"userShortName,omitempty"`
	ManagementID  string `json:"managementId,omitempty"`
}

// MobileDeviceSubsetConfigurationProfile is a single configurationProfiles entry.
type MobileDeviceSubsetConfigurationProfile struct {
	DisplayName string `json:"displayName,omitempty"`
	Version     string `json:"version,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
}

// MobileDeviceSubsetAttachmentV2 is a single attachments entry.
type MobileDeviceSubsetAttachmentV2 struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
