package enrollment_customization_preview

// RequestParseMarkdown is the request body for ParseMarkdown.
type RequestParseMarkdown struct {
	Markdown string `json:"markdown"`
}

// ResponseParseMarkdown is the response for ParseMarkdown.
// The API returns parsed HTML in the markdown field.
type ResponseParseMarkdown struct {
	Markdown string `json:"markdown"`
}

// PanelListResponse is the response for GetAllPanels.
type PanelListResponse struct {
	Panels []ResourcePanel `json:"panels"`
}

// ResourcePanel is the base panel structure returned by GetAllPanels and GetPanelByID.
type ResourcePanel struct {
	DisplayName string `json:"displayName"`
	Rank        int    `json:"rank"`
	ID          int    `json:"id"`
	Type        string `json:"type"`
}

// ResourceLdapGroupAccess represents LDAP group access for an LDAP panel.
type ResourceLdapGroupAccess struct {
	LdapServerID int    `json:"ldapServerId,omitempty"`
	GroupName    string `json:"groupName,omitempty"`
	ID           int    `json:"id,omitempty"`
}

// ResourceLdapPanel is the LDAP panel model for create/update/get.
type ResourceLdapPanel struct {
	DisplayName       string                     `json:"displayName"`
	Rank              int                        `json:"rank"`
	UsernameLabel     string                     `json:"usernameLabel"`
	PasswordLabel     string                     `json:"passwordLabel"`
	Title             string                     `json:"title"`
	BackButtonText    string                     `json:"backButtonText"`
	ContinueButtonText string                    `json:"continueButtonText"`
	LdapGroupAccess   []ResourceLdapGroupAccess  `json:"ldapGroupAccess,omitempty"`
	ID                int                        `json:"id,omitempty"`
	Type              string                     `json:"type,omitempty"`
}

// ResourceSsoPanel is the SSO panel model for create/update/get.
type ResourceSsoPanel struct {
	DisplayName                 string `json:"displayName"`
	Rank                        int    `json:"rank"`
	IsUseJamfConnect            bool   `json:"isUseJamfConnect"`
	LongNameAttribute           string `json:"longNameAttribute"`
	ShortNameAttribute          string `json:"shortNameAttribute"`
	IsGroupEnrollmentAccessEnabled bool `json:"isGroupEnrollmentAccessEnabled"`
	GroupEnrollmentAccessName   string `json:"groupEnrollmentAccessName"`
	ID                          int    `json:"id,omitempty"`
	Type                        string `json:"type,omitempty"`
}

// ResourceTextPanel is the text panel model for create/update/get.
type ResourceTextPanel struct {
	DisplayName       string `json:"displayName"`
	Rank              int    `json:"rank"`
	Body              string `json:"body"`
	Subtext           string `json:"subtext,omitempty"`
	Title             string `json:"title"`
	BackButtonText    string `json:"backButtonText"`
	ContinueButtonText string `json:"continueButtonText"`
	ID                int    `json:"id,omitempty"`
	Type              string `json:"type,omitempty"`
}

// ResponseTextPanelMarkdown is the response for GetTextPanelMarkdown.
type ResponseTextPanelMarkdown struct {
	Markdown string `json:"markdown"`
}
