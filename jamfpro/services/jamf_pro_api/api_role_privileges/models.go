package api_role_privileges

// ListResponse is the response for ListPrivilegesV1 and SearchPrivilegesByNameV1.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/api-role-privileges
type ListResponse struct {
	Privileges []string `json:"privileges"`
}
