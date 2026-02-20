package cloud_ldap_keystore

// ValidateKeystoreRequest represents the request for validating a Cloud LDAP keystore.
type ValidateKeystoreRequest struct {
	Password  string `json:"password"`
	FileBytes string `json:"fileBytes"`
	FileName  string `json:"fileName"`
}

// ResponseValidateKeystore represents the response from validating a Cloud LDAP keystore.
type ResponseValidateKeystore struct {
	Type           string `json:"type"`
	ExpirationDate string `json:"expirationDate"`
	Subject        string `json:"subject"`
	FileName       string `json:"fileName"`
}
