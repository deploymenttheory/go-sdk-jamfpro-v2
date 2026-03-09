package self_service_branding_upload

// ResourceBrandingImage represents the response from the branding upload endpoint.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_self-service-branding-images
type ResourceBrandingImage struct {
	URL string `json:"url"`
}
