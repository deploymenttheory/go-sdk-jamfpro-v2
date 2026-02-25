package apns_client_push_status

// PushStatusEntry represents an MDM client with push notifications disabled.
type PushStatusEntry struct {
	DeviceType   string `json:"deviceType"`   // MOBILE_DEVICE or COMPUTER
	ClientID     string `json:"clientId"`     // Id of the Computer or Device record in Jamf Pro
	DisabledAt   string `json:"disabledAt"`   // ISO 8601 timestamp when push was disabled
	ManagementID string `json:"managementId"` // Unique identifier for the device management record
}

// ListResponse is the response for listing clients with push notifications disabled.
type ListResponse struct {
	TotalCount int               `json:"totalCount"`
	Results    []PushStatusEntry `json:"results"`
}

// EnableAllClientsStatusResponse is the response for GetEnableAllClientsStatusV1.
// Status can be QUEUED, STARTED, or COMPLETED.
type EnableAllClientsStatusResponse struct {
	RequestedTime  string  `json:"requestedTime"`  // ISO-8601 timestamp when the request was created
	Status         string  `json:"status"`         // QUEUED, STARTED, or COMPLETED
	ProcessedTime  *string `json:"processedTime"`   // ISO-8601 timestamp when processed, null if not yet processed
}

// EnableClientRequest is the request body for EnableClientV1.
type EnableClientRequest struct {
	ManagementID string `json:"managementId"` // Unique identifier for the device management record to enable push for
}
