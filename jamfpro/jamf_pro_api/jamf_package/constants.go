package jamf_package

// Endpoints for the Jamf package API (Jamf Pro API v1 and v2).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-package
const (
	EndpointJamfPackageV1 = "/api/v1/jamf-package"
	EndpointJamfPackageV2 = "/api/v2/jamf-package"
)

// Application values for the application query parameter.
// Supported values: protect (Jamf Protect) or connect (Jamf Connect).
const (
	ApplicationProtect = "protect"
	ApplicationConnect = "connect"
)
