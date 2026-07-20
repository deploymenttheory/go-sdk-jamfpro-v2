package computer_inventory

// ComputerSection represents a section of computer inventory data.
const (
	ComputerSectionGeneral                = "GENERAL"
	ComputerSectionDiskEncryption         = "DISK_ENCRYPTION"
	ComputerSectionPurchasing             = "PURCHASING"
	ComputerSectionApplications           = "APPLICATIONS"
	ComputerSectionStorage                = "STORAGE"
	ComputerSectionUserAndLocation        = "USER_AND_LOCATION"
	ComputerSectionConfigurationProfiles  = "CONFIGURATION_PROFILES"
	ComputerSectionPrinters               = "PRINTERS"
	ComputerSectionServices               = "SERVICES"
	ComputerSectionHardware               = "HARDWARE"
	ComputerSectionLocalUserAccounts      = "LOCAL_USER_ACCOUNTS"
	ComputerSectionCertificates           = "CERTIFICATES"
	ComputerSectionAttachments            = "ATTACHMENTS"
	ComputerSectionPlugins                = "PLUGINS"
	ComputerSectionPackageReceipts        = "PACKAGE_RECEIPTS"
	ComputerSectionFonts                  = "FONTS"
	ComputerSectionSecurity               = "SECURITY"
	ComputerSectionOperatingSystem        = "OPERATING_SYSTEM"
	ComputerSectionLicensedSoftware       = "LICENSED_SOFTWARE"
	ComputerSectionIbeacons               = "IBEACONS"
	ComputerSectionSoftwareUpdates        = "SOFTWARE_UPDATES"
	ComputerSectionExtensionAttributes    = "EXTENSION_ATTRIBUTES"
	ComputerSectionContentCaching         = "CONTENT_CACHING"
	ComputerSectionGroupMemberships       = "GROUP_MEMBERSHIPS"
)

// ComputerSectionV4* constants represent a section of v4 computer inventory data.
// Jamf Pro 11.30 removed the PLUGINS and FONTS sections from the v4 surface; every
// other section carries over from v3 unchanged.
const (
	ComputerSectionV4General               = "GENERAL"
	ComputerSectionV4DiskEncryption        = "DISK_ENCRYPTION"
	ComputerSectionV4Purchasing            = "PURCHASING"
	ComputerSectionV4Applications          = "APPLICATIONS"
	ComputerSectionV4Storage               = "STORAGE"
	ComputerSectionV4UserAndLocation       = "USER_AND_LOCATION"
	ComputerSectionV4ConfigurationProfiles = "CONFIGURATION_PROFILES"
	ComputerSectionV4Printers              = "PRINTERS"
	ComputerSectionV4Services              = "SERVICES"
	ComputerSectionV4Hardware              = "HARDWARE"
	ComputerSectionV4LocalUserAccounts     = "LOCAL_USER_ACCOUNTS"
	ComputerSectionV4Certificates          = "CERTIFICATES"
	ComputerSectionV4Attachments           = "ATTACHMENTS"
	ComputerSectionV4PackageReceipts       = "PACKAGE_RECEIPTS"
	ComputerSectionV4Security              = "SECURITY"
	ComputerSectionV4OperatingSystem       = "OPERATING_SYSTEM"
	ComputerSectionV4LicensedSoftware      = "LICENSED_SOFTWARE"
	ComputerSectionV4Ibeacons              = "IBEACONS"
	ComputerSectionV4SoftwareUpdates       = "SOFTWARE_UPDATES"
	ComputerSectionV4ExtensionAttributes   = "EXTENSION_ATTRIBUTES"
	ComputerSectionV4ContentCaching        = "CONTENT_CACHING"
	ComputerSectionV4GroupMemberships      = "GROUP_MEMBERSHIPS"
)

// ComputerSecurityAttestationStatus represents the attestation status of a computer.
const (
	AttestationStatusPending                        = "PENDING"
	AttestationStatusSuccess                        = "SUCCESS"
	AttestationStatusCertificateInvalid             = "CERTIFICATE_INVALID"
	AttestationStatusDevicePropertiesMismatch       = "DEVICE_PROPERTIES_MISMATCH"
	AttestationStatusMdaUnsupportedDueToHardware    = "MDA_UNSUPPORTED_DUE_TO_HARDWARE"
	AttestationStatusMdaUnsupportedDueToSoftware    = "MDA_UNSUPPORTED_DUE_TO_SOFTWARE"
)

// ComputerSecurityBootstrapTokenEscrowedStatus represents the bootstrap token escrowed status.
const (
	BootstrapTokenEscrowedStatusEscrowed    = "ESCROWED"
	BootstrapTokenEscrowedStatusNotEscrowed = "NOT_ESCROWED"
	BootstrapTokenEscrowedStatusNotSupported = "NOT_SUPPORTED"
)

// ComputerSecurityExternalBootLevel represents the external boot level setting.
const (
	ExternalBootLevelAllowBootingFromExternalMedia    = "ALLOW_BOOTING_FROM_EXTERNAL_MEDIA"
	ExternalBootLevelDisallowBootingFromExternalMedia = "DISALLOW_BOOTING_FROM_EXTERNAL_MEDIA"
	ExternalBootLevelNotSupported                     = "NOT_SUPPORTED"
	ExternalBootLevelUnknown                          = "UNKNOWN"
)

// ComputerSecurityGatekeeperStatus represents the Gatekeeper status.
const (
	GatekeeperStatusNotCollected                   = "NOT_COLLECTED"
	GatekeeperStatusDisabled                       = "DISABLED"
	GatekeeperStatusAppStoreAndIdentifiedDevelopers = "APP_STORE_AND_IDENTIFIED_DEVELOPERS"
	GatekeeperStatusAppStore                       = "APP_STORE"
)

// ComputerSecuritySecureBootLevel represents the secure boot level setting.
const (
	SecureBootLevelNoSecurity   = "NO_SECURITY"
	SecureBootLevelMediumSecurity = "MEDIUM_SECURITY"
	SecureBootLevelFullSecurity  = "FULL_SECURITY"
	SecureBootLevelNotSupported = "NOT_SUPPORTED"
	SecureBootLevelUnknown      = "UNKNOWN"
)

// ComputerSecuritySipStatus represents the System Integrity Protection status.
const (
	SipStatusNotCollected  = "NOT_COLLECTED"
	SipStatusNotAvailable  = "NOT_AVAILABLE"
	SipStatusDisabled      = "DISABLED"
	SipStatusEnabled       = "ENABLED"
)

// ComputerDiskEncryptionIndividualRecoveryKeyValidityStatus represents the validity status of an individual recovery key.
const (
	IndividualRecoveryKeyValidityStatusValid         = "VALID"
	IndividualRecoveryKeyValidityStatusInvalid       = "INVALID"
	IndividualRecoveryKeyValidityStatusUnknown       = "UNKNOWN"
	IndividualRecoveryKeyValidityStatusNotApplicable = "NOT_APPLICABLE"
)

// ComputerLocalUserAccountUserAccountType represents the type of a local user account.
const (
	UserAccountTypeLocal   = "LOCAL"
	UserAccountTypeMobile  = "MOBILE"
	UserAccountTypeUnknown = "UNKNOWN"
)

// ComputerLocalUserAccountAzureActiveDirectoryId represents the Azure Active Directory ID status.
const (
	AzureActiveDirectoryIdActivated   = "ACTIVATED"
	AzureActiveDirectoryIdDeactivated = "DEACTIVATED"
	AzureActiveDirectoryIdUnresponsive = "UNRESPONSIVE"
	AzureActiveDirectoryIdUnknown     = "UNKNOWN"
)

// ComputerPartitionPartitionType represents the type of a disk partition.
const (
	PartitionTypeBoot     = "BOOT"
	PartitionTypeRecovery = "RECOVERY"
	PartitionTypeOther    = "OTHER"
)

// ComputerHardwareBatteryHealth represents the battery health status.
const (
	BatteryHealthNonGenuine         = "NON_GENUINE"
	BatteryHealthNormal             = "NORMAL"
	BatteryHealthServiceRecommended = "SERVICE_RECOMMENDED"
	BatteryHealthUnknown            = "UNKNOWN"
	BatteryHealthUnsupported        = "UNSUPPORTED"
)

// ComputerOperatingSystemFileVault2Status represents the FileVault 2 encryption status.
const (
	FileVault2StatusNotApplicable = "NOT_APPLICABLE"
	FileVault2StatusNotEncrypted  = "NOT_ENCRYPTED"
	FileVault2StatusBootEncrypted = "BOOT_ENCRYPTED"
	FileVault2StatusSomeEncrypted = "SOME_ENCRYPTED"
	FileVault2StatusAllEncrypted  = "ALL_ENCRYPTED"
)

// ComputerContentCachingRegistrationStatus represents the content caching registration status.
const (
	ContentCachingRegistrationStatusFailed    = "CONTENT_CACHING_FAILED"
	ContentCachingRegistrationStatusPending   = "CONTENT_CACHING_PENDING"
	ContentCachingRegistrationStatusSucceeded = "CONTENT_CACHING_SUCCEEDED"
)

// ComputerContentCachingTetheratorStatus represents the content caching tetherator status.
const (
	TetheratorStatusUnknown  = "CONTENT_CACHING_UNKNOWN"
	TetheratorStatusDisabled = "CONTENT_CACHING_DISABLED"
	TetheratorStatusEnabled  = "CONTENT_CACHING_ENABLED"
)

// ComputerCertificateCertificateStatus represents the status of a certificate.
const (
	CertificateStatusExpiring      = "EXPIRING"
	CertificateStatusExpired       = "EXPIRED"
	CertificateStatusRevoked       = "REVOKED"
	CertificateStatusPendingRevoke = "PENDING_REVOKE"
	CertificateStatusIssued        = "ISSUED"
)

// ComputerCertificateLifecycleStatus represents the lifecycle status of a certificate.
const (
	CertificateLifecycleStatusActive   = "ACTIVE"
	CertificateLifecycleStatusInactive = "INACTIVE"
)

// ComputerPartitionFileVault2State represents the FileVault 2 state of a partition.
const (
	FileVault2StateUnknown          = "UNKNOWN"
	FileVault2StateUnencrypted      = "UNENCRYPTED"
	FileVault2StateIneligible       = "INELIGIBLE"
	FileVault2StateDecrypted        = "DECRYPTED"
	FileVault2StateDecrypting       = "DECRYPTING"
	FileVault2StateEncrypted        = "ENCRYPTED"
	FileVault2StateEncrypting       = "ENCRYPTING"
	FileVault2StateRestartNeeded    = "RESTART_NEEDED"
	FileVault2StateOptimizing       = "OPTIMIZING"
	FileVault2StateDecryptingPaused = "DECRYPTING_PAUSED"
	FileVault2StateEncryptingPaused = "ENCRYPTING_PAUSED"
)

// ComputerGeneralPlatform represents the platform of a computer.
const (
	PlatformWindows = "WINDOWS"
	PlatformMac     = "MAC"
	PlatformNone    = "NONE"
)
