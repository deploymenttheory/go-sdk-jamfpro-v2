package venafi

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"

// ResourceVenafi represents the writable Venafi PKI configuration payload.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-venafi
type ResourceVenafi struct {
	ID                   *int   `json:"id,omitempty"`
	Name                 string `json:"name"`
	ProxyAddress         string `json:"proxyAddress,omitempty"`
	RevocationEnabled   *bool  `json:"revocationEnabled,omitempty"`
	ClientID            string `json:"clientId,omitempty"`
	RefreshToken        string `json:"refreshToken,omitempty"`
	RefreshTokenConfigured *bool `json:"refreshTokenConfigured,omitempty"`
}

// ResponseVenafiCreated captures the identifier returned after creating a configuration.
type ResponseVenafiCreated struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResponseVenafi models the Venafi PKI configuration returned by the API.
type ResponseVenafi struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	ProxyAddress           string `json:"proxyAddress"`
	RevocationEnabled      bool   `json:"revocationEnabled"`
	ClientID               string `json:"clientId"`
	RefreshTokenConfigured bool   `json:"refreshTokenConfigured"`
}

// ResponseConnectionStatus represents the Venafi connection status.
type ResponseConnectionStatus struct {
	Status string `json:"status"`
}

// DependentProfile represents a configuration profile using the Venafi CA.
type DependentProfile struct {
	URLPath string `json:"urlPath"`
	Name    string `json:"name"`
}

// ResponseDependentProfiles represents the list of dependent profiles.
type ResponseDependentProfiles struct {
	TotalCount int                `json:"totalCount"`
	Results    []DependentProfile `json:"results"`
}

// HistoryItem is an alias to the shared history item struct (uses int ID).
type HistoryItem = shared.SharedHistoryItem

// ResponseHistory is an alias to the shared history response struct.
type ResponseHistory = shared.SharedHistoryResponse

// HistoryNoteRequest is an alias to the shared history note request struct.
type HistoryNoteRequest = shared.SharedHistoryNoteRequest
