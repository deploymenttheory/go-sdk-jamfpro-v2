package environment_type

// ResourceEnvironmentType is the response for GetV2.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-environment-type
type ResourceEnvironmentType struct {
	Environment string `json:"environment"`
}
