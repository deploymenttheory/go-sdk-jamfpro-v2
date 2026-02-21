package csa

// ResourceTokenExchangeDetails represents the CSA token exchange details.
type ResourceTokenExchangeDetails struct {
	TenantID                string   `json:"tenantId"`
	Subject                 string   `json:"subject"`
	RefreshExpiration       int      `json:"refreshExpiration"`
	Scopes                  []string `json:"scopes"`
	LegacyJamfSalesforceIds []string `json:"legacyJamfSalesforceIds"`
}

// ResourceTenantID represents the CSA tenant ID.
type ResourceTenantID struct {
	TenantID string `json:"tenantId"`
}
