// Package constants defines shared constants used across the Jamf Pro SDK,
// including media types for HTTP content negotiation. It has no dependencies
// and may be imported by any layer — transport, services, or tests.
package constants

const (
	ApplicationJSON                = "application/json"
	ApplicationMergePatchJSON      = "application/merge-patch+json"
	TextCSV                        = "text/csv"
	TextCSVApplicationJSON         = "text/csv,application/json"
	ApplicationXML                 = "application/xml"
	TextXML                        = "text/xml"
	ApplicationPKIXCert            = "application/pkix-cert"
	ApplicationPEMCertificateChain = "application/pem-certificate-chain"
	ApplicationXPEMFile            = "application/x-pem-file"
	ApplicationOctetStream         = "application/octet-stream"
	ApplicationXAppleAspenConfig   = "application/x-apple-aspen-config"
	MultipartFormData              = "multipart/form-data"
	TextPlain                      = "text/plain"
	ImageAny                       = "image/*"
	AcceptAny                      = "*/*"
)
