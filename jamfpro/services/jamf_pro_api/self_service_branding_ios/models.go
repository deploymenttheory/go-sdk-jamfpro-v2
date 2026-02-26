package self_service_branding_ios

// ResourceSelfServiceBrandingMobile represents the details of a self-service branding
// configuration for mobile (iOS) devices.
type ResourceSelfServiceBrandingMobile struct {
	ID                    string `json:"id"`
	BrandingName          string `json:"brandingName"`
	IconId                    *int   `json:"iconId,omitempty"`
	HeaderBackgroundColorCode string `json:"headerBackgroundColorCode"`
	MenuIconColorCode     string `json:"menuIconColorCode"`
	BrandingNameColorCode string `json:"brandingNameColorCode"`
	StatusBarTextColor    string `json:"statusBarTextColor"`
}

// ListResponse is the response for ListV1.
type ListResponse struct {
	TotalCount int                             `json:"totalCount"`
	Results    []ResourceSelfServiceBrandingMobile `json:"results"`
}

// CreateResponse is the minimal response returned by the API when creating a branding
// configuration. It mirrors the lightweight server response { "id": "...", "href": "..." }.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
