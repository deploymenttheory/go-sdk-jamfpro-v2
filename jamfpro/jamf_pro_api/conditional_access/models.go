package conditional_access

// ResourceDeviceComplianceStatus represents the device compliance feature enablement status for Conditional Access.
type ResourceDeviceComplianceStatus struct {
	SharedDeviceFeatureEnabled bool `json:"sharedDeviceFeatureEnabled"`
}

// ResourceDeviceComplianceInfo represents a single device compliance information record.
// Returned by GetDeviceComplianceInformationComputerV1 and GetDeviceComplianceInformationMobileV1.
type ResourceDeviceComplianceInfo struct {
	DeviceId                          string                              `json:"deviceId"`
	Applicable                        bool                                `json:"applicable"`
	ComplianceState                   string                              `json:"complianceState"` // UNKNOWN, NON_COMPLIANT, COMPLIANT
	ComplianceVendor                  string                              `json:"complianceVendor"`
	ComplianceVendorDeviceInformation []ComplianceVendorDeviceInformation `json:"complianceVendorDeviceInformation"`
}

// ComplianceVendorDeviceInformation holds vendor-specific device details (e.g. Intune device IDs).
type ComplianceVendorDeviceInformation struct {
	Registrations            []ComplianceVendorRegistration `json:"registrations,omitempty"`
	GracePeriodSettingsState *GracePeriodSettingsState      `json:"gracePeriodSettingsState,omitempty"`
	DeviceIds                []string                       `json:"deviceIds"`
}

type ComplianceVendorRegistration struct {
	EntraDeviceID          *string `json:"entraDeviceId"`
	EntraUserID            *string `json:"entraUserId"`
	ResourceName           *string `json:"resourceName"`
	OriginID               *string `json:"originId"`
	ComplianceState        string  `json:"complianceState"`
	ComplianceStateChanged *string `json:"complianceStateChanged"`
	Applicable             bool    `json:"applicable"`
	ApplicableChanged      *string `json:"applicableChanged"`
	LastSynced             *string `json:"lastSynced"`
	GracePeriodUntil       *string `json:"gracePeriodUntil"`
	SharedDevice           *bool   `json:"sharedDevice"`
	SharedDeviceRegistered *string `json:"sharedDeviceRegistered"`
}

type GracePeriodSettingsState struct {
	Enabled       bool `json:"enabled"`
	DurationHours *int `json:"durationHours"`
}
