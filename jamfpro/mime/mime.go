// Package mime defines media type constants used for HTTP content negotiation
// across the Jamf Pro SDK. It has no dependencies and may be imported by any
// layer — transport, services, or tests.
package mime

const (
	ApplicationJSON                = "application/json"
	ApplicationMergePatchJSON      = "application/merge-patch+json"
	TextCSV                        = "text/csv"
	ApplicationXML                 = "application/xml"
	ApplicationPKIXCert            = "application/pkix-cert"
	ApplicationPEMCertificateChain = "application/pem-certificate-chain"
	ApplicationXPEMFile            = "application/x-pem-file"
	ApplicationOctetStream         = "application/octet-stream"
)
