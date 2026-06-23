// Package mobile_devices provides types and constants for the Jamf Pro mobile devices API.
//
// Implementation note (11.27.0): PATCH /v2/mobile-devices/{id} response uses V1SiteBase
// (fields: id, name only) for the site property, not the full V1Site schema.
// When implementing CRUD for this endpoint, use a SiteBase struct with only id and name.
package mobile_devices

// MobileDeviceType represents the type of a mobile device.
const (
	MobileDeviceTypeIos      = "ios"
	MobileDeviceTypeTvos     = "tvos"
	MobileDeviceTypeWatchos  = "watchos"
	MobileDeviceTypeVisionos = "visionos"
	MobileDeviceTypeUnknown  = "unknown"
)

// MobileDeviceGeneralDeviceOwnershipType represents the ownership type of a mobile device.
const (
	DeviceOwnershipTypeInstitutional                 = "Institutional"
	DeviceOwnershipTypeUserEnrollment                = "UserEnrollment"
	DeviceOwnershipTypeAccountDrivenUserEnrollment   = "AccountDrivenUserEnrollment"
	DeviceOwnershipTypeAccountDrivenDeviceEnrollment = "AccountDrivenDeviceEnrollment"
)

// MobileDeviceSecurityAttestationStatus represents the attestation status of a mobile device.
const (
	MobileDeviceAttestationStatusPending                     = "PENDING"
	MobileDeviceAttestationStatusSuccess                     = "SUCCESS"
	MobileDeviceAttestationStatusCertificateInvalid          = "CERTIFICATE_INVALID"
	MobileDeviceAttestationStatusDevicePropertiesMismatch    = "DEVICE_PROPERTIES_MISMATCH"
	MobileDeviceAttestationStatusMdaUnsupportedDueToHardware = "MDA_UNSUPPORTED_DUE_TO_HARDWARE"
	MobileDeviceAttestationStatusMdaUnsupportedDueToSoftware = "MDA_UNSUPPORTED_DUE_TO_SOFTWARE"
)

// MobileDeviceSecurityBootstrapTokenEscrowed represents the bootstrap token escrowed status.
const (
	MobileDeviceBootstrapTokenEscrowedEscrowed     = "ESCROWED"
	MobileDeviceBootstrapTokenEscrowedNotEscrowed  = "NOT_ESCROWED"
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

// MobileDeviceExceptionHandling represents the values for the `exception-handling`
// query parameter on GET /v2/mobile-devices/detail (added in Jamf Pro 11.29).
// STRICT (the default) fails the request if any device record cannot be returned;
// LENIENT skips records that cannot be returned and returns the remainder.
const (
	MobileDeviceExceptionHandlingStrict  = "STRICT"
	MobileDeviceExceptionHandlingLenient = "LENIENT"
)

// MobileDeviceSection represents a section of mobile device inventory data.
const (
	MobileDeviceSectionGeneral              = "GENERAL"
	MobileDeviceSectionHardware             = "HARDWARE"
	MobileDeviceSectionUserAndLocation      = "USER_AND_LOCATION"
	MobileDeviceSectionPurchasing           = "PURCHASING"
	MobileDeviceSectionSecurity             = "SECURITY"
	MobileDeviceSectionApplications         = "APPLICATIONS"
	MobileDeviceSectionEbooks               = "EBOOKS"
	MobileDeviceSectionNetwork              = "NETWORK"
	MobileDeviceSectionServiceSubscriptions = "SERVICE_SUBSCRIPTIONS"
	MobileDeviceSectionCertificates         = "CERTIFICATES"
	MobileDeviceSectionProfiles             = "PROFILES"
	MobileDeviceSectionUserProfiles         = "USER_PROFILES"
	MobileDeviceSectionProvisioningProfiles = "PROVISIONING_PROFILES"
	MobileDeviceSectionSharedUsers          = "SHARED_USERS"
	MobileDeviceSectionGroups               = "GROUPS"
	MobileDeviceSectionExtensionAttributes  = "EXTENSION_ATTRIBUTES"
)
