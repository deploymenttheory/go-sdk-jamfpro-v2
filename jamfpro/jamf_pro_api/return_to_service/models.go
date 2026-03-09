package return_to_service

// ListResponse is the response for List (GET /api/v1/return-to-service).
type ListResponse struct {
	TotalCount int                         `json:"totalCount"`
	Results    []ResponseConfigurationItem `json:"results"`
}

// ResponseConfigurationItem is a single Return to Service configuration in list/get.
type ResponseConfigurationItem struct {
	ID            string `json:"id"`
	DisplayName   string `json:"displayName"`
	WifiProfileID string `json:"wifiProfileId"`
}

// CreateResponse is the response for Create (POST).
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResourceReturnToServiceConfiguration is the request/response body for create and update.
type ResourceReturnToServiceConfiguration struct {
	DisplayName                                    string `json:"displayName"`
	SsoForEnrollmentEnabled                        bool   `json:"ssoForEnrollmentEnabled"`
	SsoBypassAllowed                               bool   `json:"ssoBypassAllowed"`
	SsoEnabled                                     bool   `json:"ssoEnabled"`
	SsoForMacOsSelfServiceEnabled                  bool   `json:"ssoForMacOsSelfServiceEnabled"`
	TokenExpirationDisabled                        bool   `json:"tokenExpirationDisabled"`
	UserAttributeEnabled                           bool   `json:"userAttributeEnabled"`
	UserAttributeName                              string `json:"userAttributeName"`
	UserMapping                                    string `json:"userMapping"`
	EnrollmentSsoForAccountDrivenEnrollmentEnabled bool   `json:"enrollmentSsoForAccountDrivenEnrollmentEnabled"`
	GroupEnrollmentAccessEnabled                   bool   `json:"groupEnrollmentAccessEnabled"`
	GroupAttributeName                             string `json:"groupAttributeName"`
	GroupRdnKey                                    string `json:"groupRdnKey"`
	GroupEnrollmentAccessName                      string `json:"groupEnrollmentAccessName"`
	IdpProviderType                                string `json:"idpProviderType"`
	OtherProviderTypeName                          string `json:"otherProviderTypeName"`
	MetadataSource                                 string `json:"metadataSource"`
	SessionTimeout                                 int    `json:"sessionTimeout"`
	Title                                          string `json:"title"`
	Description                                    string `json:"description"`
	Priority                                       int    `json:"priority"`
	WifiProfileID                                  string `json:"wifiProfileId"`
}
