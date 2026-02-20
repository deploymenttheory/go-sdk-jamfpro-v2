package api_integrations

// ListResponse is the response for ListV1.
type ListResponse struct {
	TotalCount int                      `json:"totalCount"`
	Results    []ResourceApiIntegration `json:"results"`
}

// ResourceApiIntegration represents an API integration resource.
type ResourceApiIntegration struct {
	ID                         int      `json:"id,omitempty"`
	AuthorizationScopes        []string `json:"authorizationScopes,omitempty"`
	DisplayName                string   `json:"displayName,omitempty"`
	Enabled                    bool     `json:"enabled,omitempty"`
	AccessTokenLifetimeSeconds int      `json:"accessTokenLifetimeSeconds,omitempty"`
	AppType                    string   `json:"appType,omitempty"`
	ClientID                   string   `json:"clientId,omitempty"`
}

// ResourceClientCredentials represents API client credentials returned by RefreshClientCredentials.
type ResourceClientCredentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
