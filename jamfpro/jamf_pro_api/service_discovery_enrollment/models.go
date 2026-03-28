package service_discovery_enrollment

// ServiceDiscoveryVersion constants define the valid enrollment types for well-known settings.
// Used in ResourceWellKnownSettingV1.EnrollmentType.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-service-discovery-enrollment-well-known-settings
const (
	ServiceDiscoveryVersionNone    = "none"
	ServiceDiscoveryVersionMDMBYOD = "mdm-byod"
	ServiceDiscoveryVersionMDMADDE = "mdm-adde"
)

// ResourceWellKnownSettingV1 represents a single well-known setting entry.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-service-discovery-enrollment-well-known-settings
type ResourceWellKnownSettingV1 struct {
	OrgName        string `json:"orgName,omitempty"`
	ServerUUID     string `json:"serverUuid"`
	EnrollmentType string `json:"enrollmentType"`
}

// WellKnownSettingsResponseV1 is the response for GetWellKnownSettingsV1 and the request body for UpdateWellKnownSettingsV1.
type WellKnownSettingsResponseV1 struct {
	WellKnownSettings []ResourceWellKnownSettingV1 `json:"wellKnownSettings"`
}
