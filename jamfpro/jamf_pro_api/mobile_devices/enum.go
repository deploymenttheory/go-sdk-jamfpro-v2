package mobile_devices

// MobileDeviceType represents the type of a mobile device.
const (
	MobileDeviceTypeIos     = "ios"
	MobileDeviceTypeTvos    = "tvos"
	MobileDeviceTypeWatchos = "watchos"
	MobileDeviceTypeVisionos = "visionos"
	MobileDeviceTypeUnknown = "unknown"
)

// MobileDeviceGeneralDeviceOwnershipType represents the ownership type of a mobile device.
const (
	DeviceOwnershipTypeInstitutional                  = "Institutional"
	DeviceOwnershipTypeUserEnrollment                 = "UserEnrollment"
	DeviceOwnershipTypeAccountDrivenUserEnrollment    = "AccountDrivenUserEnrollment"
	DeviceOwnershipTypeAccountDrivenDeviceEnrollment  = "AccountDrivenDeviceEnrollment"
)

// MobileDeviceSecurityAttestationStatus represents the attestation status of a mobile device.
const (
	MobileDeviceAttestationStatusPending                      = "PENDING"
	MobileDeviceAttestationStatusSuccess                      = "SUCCESS"
	MobileDeviceAttestationStatusCertificateInvalid           = "CERTIFICATE_INVALID"
	MobileDeviceAttestationStatusDevicePropertiesMismatch     = "DEVICE_PROPERTIES_MISMATCH"
	MobileDeviceAttestationStatusMdaUnsupportedDueToHardware  = "MDA_UNSUPPORTED_DUE_TO_HARDWARE"
	MobileDeviceAttestationStatusMdaUnsupportedDueToSoftware  = "MDA_UNSUPPORTED_DUE_TO_SOFTWARE"
)

// MobileDeviceSecurityBootstrapTokenEscrowed represents the bootstrap token escrowed status.
const (
	MobileDeviceBootstrapTokenEscrowedEscrowed    = "ESCROWED"
	MobileDeviceBootstrapTokenEscrowedNotEscrowed = "NOT_ESCROWED"
	MobileDeviceBootstrapTokenEscrowedNotSupported = "NOT_SUPPORTED"
)

// MobileDeviceCertificateCertificateStatus represents the status of a mobile device certificate.
const (
	MobileDeviceCertificateStatusExpiring      = "EXPIRING"
	MobileDeviceCertificateStatusExpired       = "EXPIRED"
	MobileDeviceCertificateStatusRevoked       = "REVOKED"
	MobileDeviceCertificateStatusPendingRevoke = "PENDING_REVOKE"
	MobileDeviceCertificateStatusIssued        = "ISSUED"
)

// MobileDeviceCertificateLifecycleStatus represents the lifecycle status of a mobile device certificate.
const (
	MobileDeviceCertificateLifecycleStatusActive   = "ACTIVE"
	MobileDeviceCertificateLifecycleStatusInactive = "INACTIVE"
)

// MobileDeviceHardwareBatteryHealth represents the battery health of a mobile device.
const (
	MobileDeviceBatteryHealthNonGenuine         = "NON_GENUINE"
	MobileDeviceBatteryHealthNormal             = "NORMAL"
	MobileDeviceBatteryHealthServiceRecommended = "SERVICE_RECOMMENDED"
	MobileDeviceBatteryHealthUnknown            = "UNKNOWN"
	MobileDeviceBatteryHealthUnsupported        = "UNSUPPORTED"
)

// MobileDeviceSection represents a section of mobile device inventory data.
const (
	MobileDeviceSectionGeneral               = "GENERAL"
	MobileDeviceSectionHardware              = "HARDWARE"
	MobileDeviceSectionUserAndLocation       = "USER_AND_LOCATION"
	MobileDeviceSectionPurchasing            = "PURCHASING"
	MobileDeviceSectionSecurity              = "SECURITY"
	MobileDeviceSectionApplications          = "APPLICATIONS"
	MobileDeviceSectionEbooks                = "EBOOKS"
	MobileDeviceSectionNetwork               = "NETWORK"
	MobileDeviceSectionServiceSubscriptions  = "SERVICE_SUBSCRIPTIONS"
	MobileDeviceSectionCertificates          = "CERTIFICATES"
	MobileDeviceSectionProfiles              = "PROFILES"
	MobileDeviceSectionUserProfiles          = "USER_PROFILES"
	MobileDeviceSectionProvisioningProfiles  = "PROVISIONING_PROFILES"
	MobileDeviceSectionSharedUsers           = "SHARED_USERS"
	MobileDeviceSectionGroups                = "GROUPS"
	MobileDeviceSectionExtensionAttributes   = "EXTENSION_ATTRIBUTES"
)
