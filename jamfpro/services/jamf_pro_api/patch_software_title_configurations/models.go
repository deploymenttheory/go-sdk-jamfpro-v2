package patch_software_title_configurations

// ListResponse represents a list of patch software title configurations.
type ListResponse []ResourcePatchSoftwareTitleConfiguration

// ResourcePatchSoftwareTitleConfiguration represents a patch software title configuration resource.
type ResourcePatchSoftwareTitleConfiguration struct {
	ID                     string                                                    `json:"id,omitempty"`
	DisplayName            string                                                    `json:"displayName"`
	SoftwareTitleID        string                                                    `json:"softwareTitleId"`
	CategoryID             string                                                    `json:"categoryId,omitempty"`
	SiteID                 string                                                    `json:"siteId,omitempty"`
	UINotifications        bool                                                      `json:"uiNotifications,omitempty"`
	EmailNotifications     bool                                                      `json:"emailNotifications,omitempty"`
	ExtensionAttributes    []SubsetExtensionAttribute                                `json:"extensionAttributes,omitempty"`
	SoftwareTitleName      string                                                    `json:"softwareTitleName,omitempty"`
	SoftwareTitleNameID    string                                                    `json:"softwareTitleNameId,omitempty"`
	SoftwareTitlePublisher string                                                    `json:"softwareTitlePublisher,omitempty"`
	JamfOfficial           bool                                                      `json:"jamfOfficial,omitempty"`
	PatchSourceName        string                                                    `json:"patchSourceName,omitempty"`
	PatchSourceEnabled     bool                                                      `json:"patchSourceEnabled,omitempty"`
	Packages               []SubsetPackage                                           `json:"packages,omitempty"`
}

// SubsetExtensionAttribute represents an extension attribute in a patch software title configuration.
type SubsetExtensionAttribute struct {
	Accepted bool   `json:"accepted,omitempty"`
	EAID     string `json:"eaId,omitempty"`
}

// SubsetPackage represents a package in a patch software title configuration.
type SubsetPackage struct {
	PackageID   string `json:"packageId,omitempty"`
	Version     string `json:"version,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// CreateResponse represents the response when creating a patch software title configuration.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
