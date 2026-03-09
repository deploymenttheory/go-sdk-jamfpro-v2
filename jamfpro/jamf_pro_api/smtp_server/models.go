package smtp_server

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"

// ResourceSMTPServer represents the SMTP server configuration (singleton).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
type ResourceSMTPServer struct {
	Enabled               bool                                    `json:"enabled"`
	AuthenticationType    string                                  `json:"authenticationType"`
	ConnectionSettings    *ResourceSMTPServerConnectionSettings    `json:"connectionSettings,omitempty"`
	SenderSettings        *ResourceSMTPServerSenderSettings        `json:"senderSettings,omitempty"`
	BasicAuthCredentials  *ResourceSMTPServerBasicAuthCredentials  `json:"basicAuthCredentials,omitempty"`
	GraphApiCredentials   *ResourceSMTPServerGraphApiCredentials   `json:"graphApiCredentials,omitempty"`
	GoogleMailCredentials *ResourceSMTPServerGoogleMailCredentials `json:"googleMailCredentials,omitempty"`
}

// ResourceSMTPServerConnectionSettings holds SMTP connection settings.
type ResourceSMTPServerConnectionSettings struct {
	Host              string `json:"host"`
	Port              int    `json:"port"`
	EncryptionType    string `json:"encryptionType"`
	ConnectionTimeout int    `json:"connectionTimeout"`
}

// ResourceSMTPServerSenderSettings holds sender display name and email.
type ResourceSMTPServerSenderSettings struct {
	DisplayName  string `json:"displayName,omitempty"`
	EmailAddress string `json:"emailAddress"`
}

// ResourceSMTPServerBasicAuthCredentials holds username/password for SMTP auth.
type ResourceSMTPServerBasicAuthCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ResourceSMTPServerGraphApiCredentials holds Microsoft Graph API credentials for SMTP.
type ResourceSMTPServerGraphApiCredentials struct {
	TenantId     string `json:"tenantId"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// ResourceSMTPServerAuthentication represents an authentication entry (e.g. for Google Mail).
type ResourceSMTPServerAuthentication struct {
	EmailAddress string `json:"emailAddress"`
	Status       string `json:"status"`
}

// ResourceSMTPServerGoogleMailCredentials holds Google Mail API credentials.
type ResourceSMTPServerGoogleMailCredentials struct {
	ClientId        string                               `json:"clientId"`
	ClientSecret    string                               `json:"clientSecret"`
	Authentications []ResourceSMTPServerAuthentication `json:"authentications,omitempty"`
}

// HistoryObject is an alias to the shared history item struct with string IDs.
type HistoryObject = shared.SharedHistoryItemString

// HistoryResponse is an alias to the shared history response struct with string IDs.
type HistoryResponse = shared.SharedHistoryResponseString

// AddHistoryNoteRequest is an alias to the shared history note request struct.
type AddHistoryNoteRequest = shared.SharedHistoryNoteRequest

// AddHistoryNoteResponse is the response for AddHistoryNoteV1 (201 Created).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-smtp-server-history
type AddHistoryNoteResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// TestRequest is the request body for TestV1.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-smtp-server-test
type TestRequest struct {
	RecipientEmail string `json:"recipientEmail"`
}
