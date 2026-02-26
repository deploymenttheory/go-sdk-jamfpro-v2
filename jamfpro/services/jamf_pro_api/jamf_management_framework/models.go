package jamf_management_framework

// RedeployResponse is the response for POST /api/v1/jamf-management-framework/redeploy/{id}.
// Returned when the redeploy command is successfully queued (201 Created).
type RedeployResponse struct {
	DeviceID    string `json:"deviceId"`
	CommandUUID string `json:"commandUuid"`
}
