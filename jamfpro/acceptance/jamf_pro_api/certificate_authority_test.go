package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Certificate Authority
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Retrieves active certificate authority X.509 details
//   • GetActiveCertificateAuthorityDERV1(ctx) - Retrieves active CA in DER format
//   • GetActiveCertificateAuthorityPEMV1(ctx) - Retrieves active CA in PEM format
//   • GetCertificateAuthorityByIDV1(ctx, id) - Retrieves CA by ID in X.509 format
//   • GetCertificateAuthorityByIDDERV1(ctx, id) - Retrieves CA by ID in DER format
//   • GetCertificateAuthorityByIDPEMV1(ctx, id) - Retrieves CA by ID in PEM format
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service only provides read access to certificate authority data
//     -- Tests: TestAcceptance_CertificateAuthority_GetV1
//     -- Flow: Get active CA → Verify response structure and required fields
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get active certificate authority (X.509 format)
//   ✓ Verify response structure (200 status)
//   ✓ Validate required fields (SubjectX500Principal is present and non-empty)
//   ✗ Get active CA in DER format (not yet tested - should be added)
//   ✗ Get active CA in PEM format (not yet tested - should be added)
//   ✗ Get CA by ID in X.509 format (not yet tested - should be added)
//   ✗ Get CA by ID in DER format (not yet tested - should be added)
//   ✗ Get CA by ID in PEM format (not yet tested - should be added)
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a read-only endpoint - no mutations possible
//   • Multiple format options available (X.509, DER, PEM) for certificate export
//   • GetV1 retrieves the active/current certificate authority
//   • GetCertificateAuthorityByID operations require a valid CA ID
//   • No cleanup needed as this is a read-only operation
//   • TODO: Add tests for DER and PEM format retrieval operations
//   • TODO: Add tests for GetByID operations (requires existing CA ID)
//
// =============================================================================

func TestAcceptance_CertificateAuthority_GetV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.CertificateAuthority
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.SubjectX500Principal)
}
