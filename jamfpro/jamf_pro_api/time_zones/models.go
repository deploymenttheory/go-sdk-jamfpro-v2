package time_zones

// ResourceTimeZone represents a time zone resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-time-zones
type ResourceTimeZone struct {
	ZoneId      string `json:"zoneId"`
	Region      string `json:"region"`
	DisplayName string `json:"displayName"`
}
