package device_enrollments

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"

// ListResponse represents the paginated response for device enrollments list.
type ListResponse struct {
	TotalCount int                        `json:"totalCount"`
	Results    []ResourceDeviceEnrollment `json:"results"`
}

// HistoryResponse represents the paginated response for device enrollment history.
type HistoryResponse = shared.SharedHistoryResponse

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
type ResourceHistoryEntry = shared.SharedHistoryItem

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
type RequestAddHistoryNotes = shared.SharedHistoryNoteRequest

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

// ResourceEnrolledDevice represents a device assigned to a device enrollment instance.
type ResourceEnrolledDevice struct {
	ID                                string                        `json:"id"`
	DeviceEnrollmentProgramInstanceId string                        `json:"deviceEnrollmentProgramInstanceId"`
	PrestageId                        string                        `json:"prestageId,omitempty"`
	SerialNumber                      string                        `json:"serialNumber"`
	Description                       string                        `json:"description,omitempty"`
	Model                             string                        `json:"model,omitempty"`
	Color                             string                        `json:"color,omitempty"`
	AssetTag                          string                        `json:"assetTag,omitempty"`
	ProfileStatus                     string                        `json:"profileStatus"`
	ProfileUuid                       string                        `json:"profileUuid,omitempty"`
	ProfileAssignTime                 string                        `json:"profileAssignTime,omitempty"`
	ProfilePushTime                   string                        `json:"profilePushTime,omitempty"`
	DeviceAssignedDate                string                        `json:"deviceAssignedDate"`
	DeviceAssignedBy                  string                        `json:"deviceAssignedBy,omitempty"`
	Os                                string                        `json:"os,omitempty"`
	DeviceFamily                      string                        `json:"deviceFamily,omitempty"`
	SyncState                         *EnrolledDeviceSyncState      `json:"syncState,omitempty"`
}

// EnrolledDeviceSyncState represents the sync state subset for an enrolled device.
type EnrolledDeviceSyncState struct {
	ProfileAssignTime   string `json:"profileAssignTime,omitempty"`
	ProfilePushTime     string `json:"profilePushTime,omitempty"`
	DeviceAssignedDate  string `json:"deviceAssignedDate,omitempty"`
}

// DevicesResponse represents the paginated response for devices assigned to a device enrollment instance.
type DevicesResponse struct {
	TotalCount int                      `json:"totalCount"`
	Results    []ResourceEnrolledDevice `json:"results"`
}
