package computer_extension_attributes

// ResourceComputerExtensionAttribute represents a computer extension attribute resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-extension-attributes-id
type ResourceComputerExtensionAttribute struct {
	ID                            string   `json:"id"`
	Name                          string   `json:"name"`
	Description                   string   `json:"description,omitempty"`
	DataType                      string   `json:"dataType"`
	Enabled                       *bool    `json:"enabled"`
	InventoryDisplayType          string   `json:"inventoryDisplayType"`
	InputType                     string   `json:"inputType"`
	ScriptContents                string   `json:"scriptContents,omitempty"`
	PopupMenuChoices              []string `json:"popupMenuChoices,omitempty"`
	LDAPAttributeMapping          string   `json:"ldapAttributeMapping,omitempty"`
	LDAPExtensionAttributeAllowed *bool    `json:"ldapExtensionAttributeAllowed,omitempty"`
}

// ListResponse is the response for ListComputerExtensionAttributesV1.
type ListResponse struct {
	TotalCount int                                `json:"totalCount"`
	Results    []ResourceComputerExtensionAttribute `json:"results"`
}

// RequestComputerExtensionAttribute is the body for creating and updating computer extension attributes.
type RequestComputerExtensionAttribute struct {
	Name                            string   `json:"name"`
	Description                     string   `json:"description,omitempty"`
	DataType                        string   `json:"dataType"`
	Enabled                         *bool    `json:"enabled,omitempty"`
	InventoryDisplayType            string   `json:"inventoryDisplayType"`
	InputType                       string   `json:"inputType"`
	ScriptContents                  string   `json:"scriptContents,omitempty"`
	PopupMenuChoices                []string `json:"popupMenuChoices,omitempty"`
	LDAPAttributeMapping            string   `json:"ldapAttributeMapping,omitempty"`
	LDAPExtensionAttributeAllowed   *bool    `json:"ldapExtensionAttributeAllowed,omitempty"`
}

// CreateResponse is the response for CreateComputerExtensionAttributeV1.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// DeleteComputerExtensionAttributesByIDRequest is the body for DeleteComputerExtensionAttributesByIDV1 (delete multiple).
type DeleteComputerExtensionAttributesByIDRequest struct {
	IDs []string `json:"ids"`
}
