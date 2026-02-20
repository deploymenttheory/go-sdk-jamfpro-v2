package cloud_idp

// ListResponse represents the paginated response for Cloud Identity Providers.
type ListResponse struct {
	TotalCount int                       `json:"totalCount"`
	Results    []ResourceCloudIdProvider `json:"results"`
}

// ResourceCloudIdProvider represents a Cloud Identity Provider.
type ResourceCloudIdProvider struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName"`
	Enabled      bool   `json:"enabled"`
	ProviderName string `json:"providerName"`
}

// ResourceCloudIdProviderDetails represents detailed Cloud Identity Provider configuration.
type ResourceCloudIdProviderDetails struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName"`
	ProviderName string `json:"providerName"`
}

// ExportRequest represents the request body for exporting Cloud Identity Providers.
type ExportRequest struct {
	Page     *int      `json:"page,omitempty"`
	PageSize *int      `json:"pageSize,omitempty"`
	Sort     []string  `json:"sort,omitempty"`
	Filter   *string   `json:"filter,omitempty"`
	Fields   []ExportField `json:"fields,omitempty"`
}

// ExportField represents a field to export.
type ExportField struct {
	Name string `json:"name"`
}

// HistoryResponse represents the paginated history for a Cloud Identity Provider.
type HistoryResponse struct {
	TotalCount int           `json:"totalCount"`
	Results    []HistoryItem `json:"results"`
}

// HistoryItem represents a single history entry.
type HistoryItem struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Date     string  `json:"date"`
	Note     string  `json:"note"`
	Details  *string `json:"details,omitempty"`
}

// HistoryNoteRequest represents the request for adding a history note.
type HistoryNoteRequest struct {
	Note string `json:"note"`
}

// TestGroupSearchRequest represents the request for testing group search.
type TestGroupSearchRequest struct {
	GroupName string `json:"groupname"`
}

// TestGroupSearchResponse represents the response from testing group search.
type TestGroupSearchResponse struct {
	TotalCount int              `json:"totalCount"`
	Results    []TestGroupResult `json:"results"`
}

// TestGroupResult represents a single group search result.
type TestGroupResult struct {
	DistinguishedName string `json:"distinguishedName"`
	ID                string `json:"id"`
	UUID              string `json:"uuid"`
	ServerID          string `json:"serverId"`
	Name              string `json:"name"`
}

// TestUserSearchRequest represents the request for testing user search.
type TestUserSearchRequest struct {
	Username string `json:"username"`
}

// TestUserSearchResponse represents the response from testing user search.
type TestUserSearchResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []TestUserResult `json:"results"`
}

// TestUserResult represents a single user search result.
type TestUserResult struct {
	DistinguishedName string           `json:"distinguishedName"`
	ID                string           `json:"id"`
	UUID              string           `json:"uuid"`
	ServerID          string           `json:"serverId"`
	Name              string           `json:"name"`
	Attributes        TestUserAttributes `json:"attributes"`
}

// TestUserAttributes represents user attributes from search.
type TestUserAttributes struct {
	FullName       string `json:"fullName"`
	EmailAddress   string `json:"emailAddress"`
	PhoneNumber    string `json:"phoneNumber"`
	Position       string `json:"position"`
	Room           string `json:"room"`
	BuildingName   string `json:"buildingName"`
	DepartmentName string `json:"departmentName"`
}

// TestUserMembershipRequest represents the request for testing user membership.
type TestUserMembershipRequest struct {
	Username  string `json:"username"`
	GroupName string `json:"groupname"`
}

// TestUserMembershipResponse represents the response from testing user membership.
type TestUserMembershipResponse struct {
	TotalCount int              `json:"totalCount"`
	Results    []TestGroupResult `json:"results"`
}
