package computer_extension_attributes

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"

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

// HistoryItem represents a single computer extension attribute history entry.
type HistoryItem = shared.SharedHistoryItem

// HistoryResponse is the response for GetHistoryByIDV1.
type HistoryResponse = shared.SharedHistoryResponse

// AddHistoryNoteRequest is the body for AddHistoryNoteByIDV1.
type AddHistoryNoteRequest = shared.SharedHistoryNoteRequest

// ResourceComputerExtensionAttributeTemplate represents a template for computer extension attributes.
type ResourceComputerExtensionAttributeTemplate struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Description         string   `json:"description,omitempty"`
	DataType            string   `json:"dataType"`
	Enabled             *bool    `json:"enabled,omitempty"`
	InventoryDisplayType string   `json:"inventoryDisplayType"`
	InputType           string   `json:"inputType"`
	ScriptContents      string   `json:"scriptContents,omitempty"`
	PopupMenuChoices    []string `json:"popupMenuChoices,omitempty"`
	TemplateCategory    string   `json:"templateCategory"`
	ManageExistingData  string   `json:"manageExistingData,omitempty"`
}

// TemplateListResponse is the response for ListTemplatesV1.
type TemplateListResponse struct {
	TotalCount int                                          `json:"totalCount"`
	Results    []ResourceComputerExtensionAttributeTemplate `json:"results"`
}

// DataDependencyItem represents a single dependent object (e.g., smart group or advanced search).
type DataDependencyItem struct {
	ID               int    `json:"id"`
	ObjectID         int    `json:"objectId"`
	NameLocalization string `json:"nameLocalization"`
	IdentifiableName string `json:"identifiableName"`
	Hyperlink        string `json:"hyperlink"`
}

// DataDependencyResponse is the response for GetDataDependencyByIDV1.
type DataDependencyResponse struct {
	TotalCount int                  `json:"totalCount"`
	Results    []DataDependencyItem `json:"results"`
}
