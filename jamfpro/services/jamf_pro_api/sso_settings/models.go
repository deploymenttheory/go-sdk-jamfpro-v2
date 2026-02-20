package sso_settings

// ResourceSsoSettings is the SSO settings resource (get/update).
type ResourceSsoSettings struct {
	SsoEnabled                                      bool                 `json:"ssoEnabled"`
	ConfigurationType                               string               `json:"configurationType"`
	OidcSettings                                    *OidcSettings        `json:"oidcSettings"`
	SamlSettings                                    *SamlSettings        `json:"samlSettings"`
	SsoBypassAllowed                                bool                 `json:"ssoBypassAllowed"`
	SsoForEnrollmentEnabled                         bool                 `json:"ssoForEnrollmentEnabled"`
	SsoForMacOsSelfServiceEnabled                   bool                 `json:"ssoForMacOsSelfServiceEnabled"`
	EnrollmentSsoForAccountDrivenEnrollmentEnabled bool                 `json:"enrollmentSsoForAccountDrivenEnrollmentEnabled"`
	GroupEnrollmentAccessEnabled                    bool                 `json:"groupEnrollmentAccessEnabled"`
	GroupEnrollmentAccessName                        string               `json:"groupEnrollmentAccessName"`
	EnrollmentSsoConfig                             *EnrollmentSsoConfig `json:"enrollmentSsoConfig,omitempty"`
}

// OidcSettings holds OIDC provider settings.
type OidcSettings struct {
	UserMapping                   string `json:"userMapping"`
	JamfIdAuthenticationEnabled   *bool  `json:"jamfIdAuthenticationEnabled,omitempty"`
	UsernameAttributeClaimMapping string `json:"usernameAttributeClaimMapping,omitempty"`
}

// SamlSettings holds SAML provider settings.
type SamlSettings struct {
	IdpUrl                  string `json:"idpUrl,omitempty"`
	EntityId                string `json:"entityId,omitempty"`
	MetadataSource          string `json:"metadataSource,omitempty"`
	UserMapping             string `json:"userMapping,omitempty"`
	IdpProviderType         string `json:"idpProviderType,omitempty"`
	GroupRdnKey             string `json:"groupRdnKey"`
	UserAttributeName       string `json:"userAttributeName"`
	GroupAttributeName       string `json:"groupAttributeName,omitempty"`
	UserAttributeEnabled     bool   `json:"userAttributeEnabled"`
	MetadataFileName        string `json:"metadataFileName,omitempty"`
	OtherProviderTypeName   string `json:"otherProviderTypeName"`
	FederationMetadataFile  string `json:"federationMetadataFile,omitempty"`
	TokenExpirationDisabled bool   `json:"tokenExpirationDisabled"`
	SessionTimeout          int    `json:"sessionTimeout,omitempty"`
}

// EnrollmentSsoConfig holds enrollment SSO host and hint.
type EnrollmentSsoConfig struct {
	Hosts          []string `json:"hosts,omitempty"`
	ManagementHint string   `json:"managementHint,omitempty"`
}

// ResponseSsoEnrollmentCustomizationDependencies is the response for GetEnrollmentCustomizationDependencies.
type ResponseSsoEnrollmentCustomizationDependencies struct {
	Dependencies []EnrollmentCustomizationDependency `json:"dependencies"`
}

// EnrollmentCustomizationDependency represents one dependency (name, link, human name).
type EnrollmentCustomizationDependency struct {
	Name              string `json:"name"`
	Hyperlink         string `json:"hyperlink"`
	HumanReadableName string `json:"humanReadableName"`
}
