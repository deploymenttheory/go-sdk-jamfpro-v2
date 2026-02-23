package smart_computer_groups

// Endpoints for the smart computer groups API (Jamf Pro API v2).
// API reference: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
const (
	// EndpointBaseV2 is the base path for computer groups v2 API.
	EndpointBaseV2 = "/api/v2/computer-groups"
	// EndpointSmartGroupsV2 is the path for smart groups list/create.
	EndpointSmartGroupsV2 = EndpointBaseV2 + "/smart-groups"
	// EndpointSmartGroupMembershipV2 is the path for smart group membership.
	EndpointSmartGroupMembershipV2 = EndpointBaseV2 + "/smart-group-membership"
)
