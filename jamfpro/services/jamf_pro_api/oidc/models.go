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
}
