package jamf_remote_assist

// SessionHistory represents a single Jamf Remote Assist session history item.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session
type SessionHistory struct {
	TenantID                   string `json:"tenantId"`
	SessionID                  string `json:"sessionId"`
	DeviceID                   string `json:"deviceId"`
	SessionStartedTimestamp    string `json:"sessionStartedTimestamp"`
	SessionEndedTimestamp      string `json:"sessionEndedTimestamp"`
	SessionType                string `json:"sessionType"`                // ATTENDED, UNATTENDED
	StatusType                 string `json:"statusType"`                 // STARTED, FINISHED, ERROR
	SessionAdminID             string `json:"sessionAdminId"`
	Comment                    string `json:"comment"`
	Details                    *SessionDetails `json:"details,omitempty"` // Only returned in detail endpoints
}

// SessionDetails represents the details of a session including file transfers.
type SessionDetails struct {
	FileTransferItemList []FileTransferItem `json:"fileTransferItemList"`
}

// FileTransferItem represents a single file transfer within a session.
type FileTransferItem struct {
	FilePath           string `json:"filePath"`
	TransferTimestamp  string `json:"transferTimestamp"`
	FileTransferType   string `json:"fileTransferType"` // DOWNLOAD, UPLOAD
}

// ListSessionsResponse is the response for ListSessionsV2.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session
type ListSessionsResponse struct {
	TotalCount int              `json:"totalCount"`
	Results    []SessionHistory `json:"results"`
}

// ExportSessionsRequest represents the request body for exporting sessions.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-jamf-remote-assist-session-export
type ExportSessionsRequest struct {
	Page     *int               `json:"page,omitempty"`
	PageSize *int               `json:"pageSize,omitempty"`
	Sort     []string           `json:"sort,omitempty"`
	Filter   string             `json:"filter,omitempty"`
	Fields   []ExportFieldOrder `json:"fields,omitempty"`
}

// ExportFieldOrder represents field ordering for export.
type ExportFieldOrder struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}
