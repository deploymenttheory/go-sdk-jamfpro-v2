package api_integrations

// ListResponse is the response for ListV1.
type ListResponse struct {
	TotalCount int                      `json:"totalCount"`
	Results    []ResourceApiIntegration `json:"results"`
}

// ResourceApiIntegration represents an API integration resource.
type ResourceApiIntegration struct {
	ID                         int      `json:"id"`
	AuthorizationScopes        []string `json:"authorizationScopes"`
	DisplayName                string   `json:"displayName"`
	Enabled                    bool     `json:"enabled"`
	AccessTokenLifetimeSeconds int      `json:"accessTokenLifetimeSeconds,omitempty"` // optional
	AppType                    string   `json:"appType"`
	ClientID                   string   `json:"clientId"`
}

// ResourceClientCredentials represents API client credentials returned by RefreshClientCredentials.
type ResourceClientCredentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
