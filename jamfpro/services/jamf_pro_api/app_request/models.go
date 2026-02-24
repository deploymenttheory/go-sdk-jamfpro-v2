package app_request

// ResourceFormInputField represents a form input field resource.
// Priority: 1 is highest, 255 is lowest.
type ResourceFormInputField struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Priority    int     `json:"priority"`
}

// FormInputFieldListResponse is the response for ListFormInputFieldsV1.
type FormInputFieldListResponse struct {
	TotalCount int                      `json:"totalCount"`
	Results    []ResourceFormInputField  `json:"results"`
}

// RequestFormInputField is the body for creating and updating form input fields.
// IDs in the request body are ignored for create and update.
type RequestFormInputField struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Priority    int     `json:"priority"`
}

// ResourceAppRequestSettings represents the app request settings resource.
type ResourceAppRequestSettings struct {
	IsEnabled             bool     `json:"isEnabled"`
	AppStoreLocale        string   `json:"appStoreLocale"`
	RequesterUserGroupID  int      `json:"requesterUserGroupId"`
	ApproverEmails        []string `json:"approverEmails"`
}
