package conditional_access

// ResourceDeviceComplianceStatus represents the device compliance feature enablement status for Conditional Access.
type ResourceDeviceComplianceStatus struct {
	SharedDeviceFeatureEnabled bool `json:"sharedDeviceFeatureEnabled"`
}

// ResourceDeviceComplianceInfo represents a single device compliance information record.
// Returned by GetDeviceComplianceInformationComputerV1 and GetDeviceComplianceInformationMobileV1.
type ResourceDeviceComplianceInfo struct {
	DeviceId                          string                                `json:"deviceId"`
	Applicable                        bool                                  `json:"applicable"`
	ComplianceState                   string                                `json:"complianceState"` // UNKNOWN, NON_COMPLIANT, COMPLIANT
	ComplianceVendor                  string                                `json:"complianceVendor"`
	ComplianceVendorDeviceInformation []ComplianceVendorDeviceInformation   `json:"complianceVendorDeviceInformation"`
}

// ComplianceVendorDeviceInformation holds vendor-specific device details (e.g. Intune device IDs).
type ComplianceVendorDeviceInformation struct {
	DeviceIds []string `json:"deviceIds"`
}
