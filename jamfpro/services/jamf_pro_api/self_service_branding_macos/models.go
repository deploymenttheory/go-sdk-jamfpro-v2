package self_service_branding_macos

// ResourceSelfServiceBrandingMacOS represents the details of a self-service branding configuration for macOS.
type ResourceSelfServiceBrandingMacOS struct {
	ID                    string `json:"id"`
	ApplicationName       string `json:"applicationName"`
	BrandingName          string `json:"brandingName"`
	BrandingNameSecondary string `json:"brandingNameSecondary"`
	IconId                *int   `json:"iconId,omitempty"`
	BrandingHeaderImageId *int   `json:"brandingHeaderImageId,omitempty"`
	HomeHeading           string `json:"homeHeading"`
	HomeSubheading        string `json:"homeSubheading"`
}

// ListResponse is the response for ListV1.
type ListResponse struct {
	TotalCount int                              `json:"totalCount"`
	Results    []ResourceSelfServiceBrandingMacOS `json:"results"`
}
