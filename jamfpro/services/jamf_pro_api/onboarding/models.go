package onboarding

// ResponseOnboardingSettings is the response for Get (onboarding configuration).
type ResponseOnboardingSettings struct {
	ID              string                 `json:"id"`
	Enabled         bool                   `json:"enabled"`
	OnboardingItems []OnboardingItemResponse `json:"onboardingItems"`
}

// OnboardingItemResponse is an item in the onboarding configuration response.
type OnboardingItemResponse struct {
	ID                    string `json:"id,omitempty"`
	EntityID              string `json:"entityId"`
	EntityName             string `json:"entityName,omitempty"`
	ScopeDescription       string `json:"scopeDescription,omitempty"`
	SiteDescription        string `json:"siteDescription,omitempty"`
	SelfServiceEntityType  string `json:"selfServiceEntityType"`
	Priority               int    `json:"priority"`
}

// ResourceUpdateOnboardingSettings is the request body for Update.
type ResourceUpdateOnboardingSettings struct {
	Enabled         bool                          `json:"enabled"`
	OnboardingItems []SubsetOnboardingItemRequest `json:"onboardingItems"`
}

// SubsetOnboardingItemRequest is an item in the onboarding update request.
type SubsetOnboardingItemRequest struct {
	ID                    string `json:"id,omitempty"`
	EntityID              string `json:"entityId"`
	SelfServiceEntityType string `json:"selfServiceEntityType"`
	Priority              int    `json:"priority"`
}

// ResponseEligibilityList is the response for GetEligibleApps, GetEligibleConfigurationProfiles, GetEligiblePolicies.
type ResponseEligibilityList struct {
	TotalCount int                           `json:"totalCount"`
	Results    []ResourceEligibilityListItem `json:"results"`
}

// ResourceEligibilityListItem is a single eligible app/profile/policy item.
type ResourceEligibilityListItem struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ScopeDescription string `json:"scopeDescription"`
	SiteDescription  string `json:"siteDescription"`
}

type HistoryResponse struct {
	TotalCount int                    `json:"totalCount"`
	Results    []ResourceHistoryEntry `json:"results"`
}

type ResourceHistoryEntry struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Date     string  `json:"date"`
	Note     string  `json:"note"`
	Details  *string `json:"details"`
}

type RequestAddHistoryNotes struct {
	Note string `json:"note"`
}

type ResponseAddHistoryNotes struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

type RequestExportHistory struct {
	Page     *int     `json:"page,omitempty"`
	PageSize *int     `json:"pageSize,omitempty"`
	Sort     []string `json:"sort,omitempty"`
	Filter   *string  `json:"filter,omitempty"`
	Fields   []string `json:"fields,omitempty"`
}
