package sso_certificate

// ResourceSSOCertKeystore represents the SSO certificate keystore.
type ResourceSSOCertKeystore struct {
	Key               string            `json:"key,omitempty"`
	Keys              []ResourceCertKey `json:"keys,omitempty"`
	Type              string            `json:"type,omitempty"`
	KeystoreFileName  string            `json:"keystoreFileName,omitempty"`
	KeystoreSetupType string            `json:"keystoreSetupType,omitempty"`
}

// ResourceCertKey represents a certificate key entry.
type ResourceCertKey struct {
	ID    string `json:"id,omitempty"`
	Valid bool   `json:"valid"`
}

// ResourceSSOKeystoreDetails holds keystore details.
type ResourceSSOKeystoreDetails struct {
	Keys         []string `json:"keys,omitempty"`
	Issuer       string   `json:"issuer,omitempty"`
	Subject      string   `json:"subject,omitempty"`
	Expiration   string   `json:"expiration,omitempty"`
	SerialNumber int      `json:"serialNumber,omitempty"`
}

// ResourceSSOKeystoreResponse is the response for Get and Create.
type ResourceSSOKeystoreResponse struct {
	Keystore        ResourceSSOCertKeystore     `json:"keystore,omitempty"`
	KeystoreDetails *ResourceSSOKeystoreDetails `json:"keystoreDetails,omitempty"`
}

// UpdateKeystoreRequest is the request body for updating the SSO keystore.
type UpdateKeystoreRequest struct {
	KeystorePassword  string              `json:"keystorePassword,omitempty"`
	KeystoreFile      string              `json:"keystoreFile,omitempty"`
	KeystoreFileName  string              `json:"keystoreFileName,omitempty"`
	Keys              []ResourceCertKey   `json:"keys,omitempty"`
	Type              string              `json:"type,omitempty"`              // Allowed values: PKCS12, JKS, NONE
	KeystoreSetupType string              `json:"keystoreSetupType,omitempty"` // Allowed values: NONE, UPLOADED, GENERATED
}

// ParseKeystoreRequest is the request body for parsing a keystore file.
type ParseKeystoreRequest struct {
	KeystoreFile     string `json:"keystoreFile"`
	KeystorePassword string `json:"keystorePassword,omitempty"`
	KeystoreFileName string `json:"keystoreFileName,omitempty"`
}

// ParseKeystoreResponse is the response from parsing a keystore file.
type ParseKeystoreResponse struct {
	Keys              []ResourceCertKey `json:"keys,omitempty"`
	Type              string            `json:"type,omitempty"`              // Values: PKCS12, JKS, NONE
	KeystoreSetupType string            `json:"keystoreSetupType,omitempty"` // Values: NONE, UPLOADED, GENERATED
}
