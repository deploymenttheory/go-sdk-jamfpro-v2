package mdm

// Endpoints for the MDM commands API (Jamf Pro API v1/v2).
//
// Jamf Pro API docs:
//   - https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-commands
//   - https://developer.jamf.com/jamf-pro/reference/post_v2-mdm-blank-push
const (
	EndpointBlankPush     = "/api/v2/mdm/blank-push"
	EndpointCommands      = "/api/v2/mdm/commands"
	EndpointDeployPackage = "/api/v1/deploy-package"
	EndpointProfileRenewal = "/api/v1/mdm/renew-profile"
)
