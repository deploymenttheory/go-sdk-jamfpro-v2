package smtp_server

// Endpoints for the SMTP server API (Jamf Pro API v1 and v2).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
const (
	EndpointSMTPServerV2        = "/api/v2/smtp-server"
	EndpointSMTPServerHistoryV1 = "/api/v1/smtp-server/history"
	EndpointSMTPServerTestV1   = "/api/v1/smtp-server/test"
)
