package oidc

// ResourceOIDCDirectIdPLoginURL represents the direct IdP login URL for OIDC.
type ResourceOIDCDirectIdPLoginURL struct {
	URL string `json:"url"` // Direct IdP login URL
}

// ResourceOIDCPublicKey represents the OIDC public key set (JWKS).
type ResourceOIDCPublicKey struct {
	Keys []ResourceOIDCKey `json:"keys"` // Array of JSON Web Keys
}

// ResourceOIDCKey represents a single JSON Web Key (JWK) in OIDC public key set.
type ResourceOIDCKey struct {
	Kty string `json:"kty"` // Key type (e.g., "RSA")
	E   string `json:"e"`   // Exponent (RSA public key parameter)
	Use string `json:"use"` // Public key use (e.g., "sig" for signature)
	Kid string `json:"kid"` // Key ID
	Alg string `json:"alg"` // Algorithm (e.g., "RS256")
	Iat int64  `json:"iat"` // Issued at timestamp
	N   string `json:"n"`   // Modulus (RSA public key parameter)
}

// RequestOIDCRedirectURL represents a request for OIDC redirect URL.
type RequestOIDCRedirectURL struct {
	OriginalURL  string `json:"originalUrl"`  // Original URL before OIDC redirect
	EmailAddress string `json:"emailAddress"` // User email address for OIDC authentication
}

// ResourceOIDCRedirectURL represents the OIDC redirect URL response.
type ResourceOIDCRedirectURL struct {
	RedirectURL string `json:"redirectUrl"` // URL to redirect for OIDC authentication
	// IdpRedirects lists the configured identity providers the user may be
	// redirected to. Added in Jamf Pro 11.30.
	IdpRedirects []ResourceOIDCIdpRedirect `json:"idpRedirects,omitempty"`
}

// ResourceOIDCIdpRedirect represents a single identity provider redirect option
// returned by the OIDC dispatch endpoint.
type ResourceOIDCIdpRedirect struct {
	RedirectURL string `json:"redirectUrl"` // URL to redirect for this identity provider
	IdpName     string `json:"idpName"`     // Display name of the identity provider
	IdpType     string `json:"idpType"`     // Identity provider type (e.g. "GENERIC_OIDC")
	// LogoURL is a customer-provided icon served by the Jamf Account URL for this
	// connection. Null when no custom icon is configured (Jamf ID connections always
	// return null). Added in Jamf Pro 11.30.
	LogoURL *string `json:"logoUrl,omitempty"`
}

// ResourcePublicFeatures represents the public OIDC configuration features.
type ResourcePublicFeatures struct {
	JamfIdAuthenticationEnabled bool `json:"jamfIdAuthenticationEnabled"` // Whether Jamf ID authentication is enabled
}
