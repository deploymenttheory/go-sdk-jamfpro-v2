package volume_purchasing_locations

// ResourceVolumePurchasingLocation represents a volume purchasing location resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id
type ResourceVolumePurchasingLocation struct {
	ID                                    string                           `json:"id,omitempty"`
	Name                                  string                           `json:"name,omitempty"`
	AppleID                               string                           `json:"appleId,omitempty"`
	OrganizationName                      string                           `json:"organizationName,omitempty"`
	TokenExpiration                       string                           `json:"tokenExpiration,omitempty"`
	CountryCode                           string                           `json:"countryCode,omitempty"`
	LocationName                          string                           `json:"locationName,omitempty"`
	ClientContextMismatch                 bool                             `json:"clientContextMismatch,omitempty"`
	AutomaticallyPopulatePurchasedContent bool                             `json:"automaticallyPopulatePurchasedContent,omitempty"`
	SendNotificationWhenNoLongerAssigned  bool                             `json:"sendNotificationWhenNoLongerAssigned,omitempty"`
	AutoRegisterManagedUsers              bool                             `json:"autoRegisterManagedUsers,omitempty"`
	SiteID                                string                           `json:"siteId,omitempty"`
	LastSyncTime                          string                           `json:"lastSyncTime,omitempty"`
	TotalPurchasedLicenses                int                              `json:"totalPurchasedLicenses,omitempty"`
	TotalUsedLicenses                     int                              `json:"totalUsedLicenses,omitempty"`
	ServiceToken                          string                           `json:"serviceToken,omitempty"`
	Content                               []VolumePurchasingSubsetContent `json:"content,omitempty"`
}

// VolumePurchasingSubsetContent represents content associated with a volume purchasing location.
type VolumePurchasingSubsetContent struct {
	Name                 string   `json:"name"`
	LicenseCountTotal    int      `json:"licenseCountTotal"`
	LicenseCountInUse    int      `json:"licenseCountInUse"`
	LicenseCountReported int      `json:"licenseCountReported"`
	IconURL              string   `json:"iconUrl"`
	DeviceTypes          []string `json:"deviceTypes"`
	ContentType          string   `json:"contentType"`
	PricingParam         string   `json:"pricingParam"`
	AdamId               string   `json:"adamId"`
}

// ListResponse is the response for ListVolumePurchasingLocationsV1.
type ListResponse struct {
	TotalCount int                                `json:"totalCount"`
	Results    []ResourceVolumePurchasingLocation `json:"results"`
}

// RequestVolumePurchasingLocation is the body for creating and updating volume purchasing locations.
type RequestVolumePurchasingLocation struct {
	Name                                  string `json:"name,omitempty"`
	AutomaticallyPopulatePurchasedContent bool   `json:"automaticallyPopulatePurchasedContent"`
	SendNotificationWhenNoLongerAssigned  bool   `json:"sendNotificationWhenNoLongerAssigned"`
	AutoRegisterManagedUsers              bool   `json:"autoRegisterManagedUsers"`
	SiteID                                string `json:"siteId,omitempty"`
	ServiceToken                          string `json:"serviceToken"`
}

// CreateResponse is the response for CreateVolumePurchasingLocationV1.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ContentListResponse is the response for GetContentV1.
type ContentListResponse struct {
	TotalCount int                           `json:"totalCount"`
	Results    []VolumePurchasingSubsetContent `json:"results"`
}

// HistoryListResponse is the response for GetHistoryV1.
type HistoryListResponse struct {
	TotalCount int            `json:"totalCount"`
	Results    []HistoryEntry `json:"results"`
}

// HistoryEntry represents an individual history record.
type HistoryEntry struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// AddHistoryNotesRequest is the request for AddHistoryNotesV1.
type AddHistoryNotesRequest struct {
	ObjectHistoryNote string `json:"objectHistoryNote"`
}
