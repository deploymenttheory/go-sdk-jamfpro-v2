package device_enrollments

// ListResponse represents the paginated response for device enrollments list.
type ListResponse struct {
	TotalCount int                        `json:"totalCount"`
	Results    []ResourceDeviceEnrollment `json:"results"`
}

// HistoryResponse represents the paginated response for device enrollment history.
type HistoryResponse struct {
	TotalCount int                       `json:"totalCount"`
	Results    []ResourceHistoryEntry    `json:"results"`
}

// ResourceDeviceEnrollment represents a single device enrollment instance.
type ResourceDeviceEnrollment struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	SupervisionIdentityId string `json:"supervisionIdentityId"`
	SiteId                string `json:"siteId"`
	ServerName            string `json:"serverName"`
	ServerUuid            string `json:"serverUuid"`
	AdminId               string `json:"adminId"`
	OrgName               string `json:"orgName"`
	OrgEmail              string `json:"orgEmail"`
	OrgPhone              string `json:"orgPhone"`
	OrgAddress            string `json:"orgAddress"`
	TokenExpirationDate   string `json:"tokenExpirationDate"`
}

// ResourceHistoryEntry represents a single device enrollment history entry.
type ResourceHistoryEntry struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// ResourceSyncState represents a single device enrollment sync state instance.
type ResourceSyncState struct {
	SyncState  string `json:"syncState"`
	InstanceID string `json:"instanceId"`
	Timestamp  string `json:"timestamp"`
}

// CreateResponse represents the response after creating a device enrollment with token upload.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// RequestTokenUpload represents the request body for token upload (create or update).
type RequestTokenUpload struct {
	TokenFileName string `json:"tokenFileName,omitempty"`
	EncodedToken  string `json:"encodedToken"`
}

// RequestUpdate represents the request body for updating device enrollment metadata.
type RequestUpdate struct {
	Name                  string `json:"name"`
	SupervisionIdentityId string `json:"supervisionIdentityId,omitempty"`
	SiteId                string `json:"siteId,omitempty"`
}

// RequestDisown represents the request body for disowning devices.
type RequestDisown struct {
	Devices []string `json:"devices"`
}

// ResponseDisown represents the response after disowning devices.
type ResponseDisown struct {
	Devices map[string]string `json:"devices"`
}

// RequestAddHistoryNotes represents the request body for adding history notes.
type RequestAddHistoryNotes struct {
	Note string `json:"note"`
}

// ResponseAddHistoryNotes represents the response after adding history notes.
type ResponseAddHistoryNotes struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResourceLatestSyncState represents the latest sync state for a device enrollment instance.
type ResourceLatestSyncState struct {
	SyncState  string `json:"syncState"`
	InstanceID string `json:"instanceId"`
	Timestamp  string `json:"timestamp"`
}
