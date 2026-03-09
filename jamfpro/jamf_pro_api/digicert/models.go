package digicert

// ResourceDigicertTrustLifecycleManager represents the writable DigiCert Trust Lifecycle Manager configuration payload.
type ResourceDigicertTrustLifecycleManager struct {
	ID                string                      `json:"id,omitempty"`
	CAName            string                      `json:"caName,omitempty"`
	FQDN              string                      `json:"fqdn,omitempty"`
	RevocationEnabled *bool                       `json:"revocationEnabled,omitempty"`
	ClientCert        *ResourceDigicertClientCert `json:"clientCert,omitempty"`
}

// ResourceDigicertClientCert bundles the file metadata and certificate data for client authentication.
type ResourceDigicertClientCert struct {
	Filename string   `json:"filename"`
	Data     []string `json:"data"`
	Password string   `json:"password,omitempty"`
}

// ResponseDigicertTrustLifecycleManagerCreated captures the identifier returned after creating a configuration.
type ResponseDigicertTrustLifecycleManagerCreated struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResponseDigicertTrustLifecycleManager models the read-only fields returned for a DigiCert Trust Lifecycle Manager configuration.
type ResponseDigicertTrustLifecycleManager struct {
	ID                string                      `json:"id"`
	CAName            string                      `json:"caName"`
	FQDN              string                      `json:"fqdn"`
	RevocationEnabled bool                        `json:"revocationEnabled"`
	ClientCert        *ResponseDigicertClientCert `json:"clientCert,omitempty"`
}

// ResponseDigicertClientCert surfaces certificate details that Jamf Pro stores for DigiCert.
type ResponseDigicertClientCert struct {
	Filename       string `json:"filename"`
	SerialNumber   string `json:"serialNumber"`
	Subject        string `json:"subject"`
	Issuer         string `json:"issuer"`
	ExpirationDate string `json:"expirationDate"`
}

// ValidateClientCertificateRequest represents the request for validating the DigiCert client certificate.
type ValidateClientCertificateRequest struct {
	Filename string   `json:"filename"`
	Data     []string `json:"data"`
	Password *string  `json:"password,omitempty"`
}

// ConnectionStatusResponse represents the connection status of a DigiCert Trust Lifecycle Manager configuration.
type ConnectionStatusResponse struct {
	Status string `json:"status"`
}

// DependenciesResponse represents the list of dependencies for a DigiCert Trust Lifecycle Manager configuration.
type DependenciesResponse struct {
	TotalCount int                 `json:"totalCount"`
	Results    []DependencyItem    `json:"results"`
}

// DependencyItem represents a single dependency (configuration profile).
type DependencyItem struct {
	ConfigProfileId   int    `json:"configProfileId"`
	ConfigProfileName string `json:"configProfileName"`
	ConfigProfileType string `json:"configProfileType"`
}
