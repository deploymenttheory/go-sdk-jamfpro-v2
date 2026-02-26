package mobile_device_extension_attributes

// ResourceMobileDeviceExtensionAttribute represents a mobile device extension attribute resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes-id
type ResourceMobileDeviceExtensionAttribute struct {
	ID                            string   `json:"id"`
	Name                          string   `json:"name"`
	Description                   string   `json:"description,omitempty"`
	DataType                      string   `json:"dataType"`
	InventoryDisplayType          string   `json:"inventoryDisplayType"`
	InputType                     string   `json:"inputType"`
	PopupMenuChoices              []string `json:"popupMenuChoices,omitempty"`
	LDAPAttributeMapping          string   `json:"ldapAttributeMapping,omitempty"`
	LDAPExtensionAttributeAllowed *bool    `json:"ldapExtensionAttributeAllowed,omitempty"`
}

// ListResponse is the response for ListMobileDeviceExtensionAttributesV1.
type ListResponse struct {
	TotalCount int                                    `json:"totalCount"`
	Results    []ResourceMobileDeviceExtensionAttribute `json:"results"`
}

// RequestMobileDeviceExtensionAttribute is the body for creating and updating mobile device extension attributes.
type RequestMobileDeviceExtensionAttribute struct {
	Name                            string   `json:"name"`
	Description                     string   `json:"description,omitempty"`
	DataType                        string   `json:"dataType"`
	Enabled                         *bool    `json:"enabled,omitempty"`
	InventoryDisplayType             string   `json:"inventoryDisplayType"`
	InputType                       string   `json:"inputType"`
	PopupMenuChoices                []string `json:"popupMenuChoices,omitempty"`
	LDAPAttributeMapping            string   `json:"ldapAttributeMapping,omitempty"`
	LDAPExtensionAttributeAllowed   *bool    `json:"ldapExtensionAttributeAllowed,omitempty"`
}

// CreateResponse is the response for CreateMobileDeviceExtensionAttributeV1.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// DeleteMobileDeviceExtensionAttributesByIDRequest is the body for DeleteMobileDeviceExtensionAttributesByIDV1 (delete multiple).
type DeleteMobileDeviceExtensionAttributesByIDRequest struct {
	IDs []string `json:"ids"`
}

// HistoryItem represents a single mobile device extension attribute history entry.
type HistoryItem struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// HistoryResponse is the response for GetHistoryByIDV1.
type HistoryResponse struct {
	TotalCount int           `json:"totalCount"`
	Results    []HistoryItem `json:"results"`
}

// AddHistoryNoteRequest is the body for AddHistoryNoteByIDV1.
type AddHistoryNoteRequest struct {
	Note string `json:"note"`
}

// DataDependencyItem represents a single dependent object (e.g., smart group).
type DataDependencyItem struct {
	ID                int    `json:"id"`
	ObjectID          int    `json:"objectId"`
	NameLocalization  string `json:"nameLocalization"`
	IdentifiableName  string `json:"identifiableName"`
	Hyperlink         string `json:"hyperlink"`
}

// DataDependencyResponse is the response for GetDataDependencyByIDV1.
type DataDependencyResponse struct {
	TotalCount int                  `json:"totalCount"`
	Results    []DataDependencyItem `json:"results"`
}
