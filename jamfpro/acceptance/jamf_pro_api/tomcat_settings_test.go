package jamf_pro_api

import (
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
)

// =============================================================================
// Acceptance Tests: Tomcat Settings
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • IssueTomcatSslCertificate(ctx) - Generates an SSL certificate via Jamf CA
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   • IssueTomcatSslCertificate is intentionally not called in acceptance tests.
//     Calling it would replace the Jamf Pro server's SSL certificate, which would
//     disrupt the test environment and require server reconfiguration.
//
//   • There are no validation entry points for this service (no required params),
//     so no validation error tests are possible.
//
// Notes
// -----------------------------------------------------------------------------
//   • This test file exists to document the service coverage decision
//   • IssueTomcatSslCertificate should only be called in controlled environments
//     with full knowledge of the downstream impact on TLS configuration
//
// =============================================================================

// TestAcceptance_TomcatSettings_note documents that IssueTomcatSslCertificate
// is intentionally not called in automated acceptance tests.
func TestAcceptance_TomcatSettings_note(t *testing.T) {
	acc.RequireClient(t)

	t.Log("TomcatSettings service is accessible.")
	t.Log("IssueTomcatSslCertificate is not exercised in acceptance tests because")
	t.Log("it replaces the Jamf Pro SSL certificate, which would disrupt the test environment.")
}
