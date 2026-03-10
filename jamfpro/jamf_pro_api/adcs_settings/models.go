package adcs_settings

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ResourceAdcsSettings represents the writable AD CS configuration payload.
type ResourceAdcsSettings struct {
	ID                string            `json:"id,omitempty"`
	DisplayName       string            `json:"displayName,omitempty"`
	CAName            string            `json:"caName,omitempty"`
	FQDN              string            `json:"fqdn,omitempty"`
	AdcsURL           string            `json:"adcsUrl,omitempty"`
	ServerCert        *ResourceAdcsCert `json:"serverCert,omitempty"`
	ClientCert        *ResourceAdcsCert `json:"clientCert,omitempty"`
	RevocationEnabled *bool             `json:"revocationEnabled,omitempty"`
	APIClientID       string            `json:"apiClientId,omitempty"`
	Outbound          *bool             `json:"outbound,omitempty"`
}

// ResourceAdcsCert bundles the file metadata and certificate data.
type ResourceAdcsCert struct {
	Filename string   `json:"filename"`
	Data     []string `json:"data"`
	Password string   `json:"password,omitempty"`
}

// ResponseAdcsSettingsCreated captures the identifier returned after creating a configuration.
type ResponseAdcsSettingsCreated struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResponseAdcsSettings models the read-only fields returned for an AD CS configuration.
type ResponseAdcsSettings struct {
	ID                            string            `json:"id"`
	DisplayName                   string            `json:"displayName"`
	CAName                        string            `json:"caName"`
	FQDN                          string            `json:"fqdn"`
	AdcsURL                       string            `json:"adcsUrl"`
	ServerCert                    *ResponseAdcsCert `json:"serverCert,omitempty"`
	ClientCert                    *ResponseAdcsCert `json:"clientCert,omitempty"`
	RevocationEnabled             bool              `json:"revocationEnabled"`
	APIClientID                   string            `json:"apiClientId"`
	Outbound                      bool              `json:"outbound"`
	ConnectorLastCheckInTimestamp string            `json:"connectorLastCheckInTimestamp"`
}

// ResponseAdcsCert surfaces certificate details that Jamf Pro stores for AD CS.
type ResponseAdcsCert struct {
	Filename       string `json:"filename"`
	SerialNumber   string `json:"serialNumber"`
	Subject        string `json:"subject"`
	Issuer         string `json:"issuer"`
	ExpirationDate string `json:"expirationDate"`
}

// ValidateCertificateRequest represents the request for validating server or client certificates.
type ValidateCertificateRequest struct {
	Filename string   `json:"filename"`
	Data     []string `json:"data"`
	Password *string  `json:"password,omitempty"`
}

// DependenciesResponse represents the list of dependencies for an AD CS Settings configuration.
type DependenciesResponse struct {
	TotalCount int              `json:"totalCount"`
	Results    []DependencyItem `json:"results"`
}

// DependencyItem represents a single dependency (configuration profile).
type DependencyItem struct {
	ConfigProfileId   int    `json:"configProfileId"`
	ConfigProfileName string `json:"configProfileName"`
	ConfigProfileType string `json:"configProfileType"`
}

// HistoryResponse is an alias to the shared history response struct.
type HistoryResponse = models.SharedHistoryResponse

// HistoryItem is an alias to the shared history item struct (uses int ID).
type HistoryItem = models.SharedHistoryItem

// HistoryNoteRequest is an alias to the shared history note request struct.
type HistoryNoteRequest = models.SharedHistoryNoteRequest
