package api_roles

// ResourceAPIRole represents an API role resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-roles-id
type ResourceAPIRole struct {
	ID          string   `json:"id"`
	DisplayName string   `json:"displayName"`
	Privileges  []string `json:"privileges"`
}

// ListResponse is the response for ListAPIRolesV1.
type ListResponse struct {
	TotalCount int               `json:"totalCount"`
	Results    []ResourceAPIRole `json:"results"`
}

// RequestAPIRole is the body for creating and updating API roles.
type RequestAPIRole struct {
	DisplayName string   `json:"displayName"`
	Privileges  []string `json:"privileges"`
}
