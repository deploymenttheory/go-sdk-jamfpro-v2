package branding

// Endpoints for the branding images API (Jamf Pro API v1).
//
// Jamf Pro API docs: GET /api/v1/branding-images/download/{id}
const (
	// EndpointBrandingImagesDownloadV1 is the path prefix for downloading self service branding images.
	// Use with fmt.Sprintf("%s/%s", EndpointBrandingImagesDownloadV1, id).
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-branding-images-download-id
	EndpointBrandingImagesDownloadV1 = "/api/v1/branding-images/download"
)
