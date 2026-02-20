package icons

// Endpoints for the icons API (Jamf Pro API v1).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
const (
	EndpointIconsV1 = "/api/v1/icon"
	// EndpointIconsDownloadV1 is the path prefix for downloading icon image bytes.
	// Use with fmt.Sprintf("%s/%d", EndpointIconsDownloadV1, id).
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-icon-download-id
	EndpointIconsDownloadV1 = "/api/v1/icon/download"
)
