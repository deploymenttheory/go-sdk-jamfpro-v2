package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/digicert"
	"github.com/stretchr/testify/assert"
)

// =============================================================================
// Acceptance Tests: DigiCert Trust Lifecycle Manager
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • Create(ctx, request) - Creates a DigiCert TLM configuration
//   • GetByID(ctx, id) - Gets a configuration by ID
//   • UpdateByID(ctx, id, request) - Updates a configuration (PATCH/merge)
//   • DeleteByID(ctx, id) - Deletes a configuration
//   • ValidateClientCertificate(ctx, request) - Validates a client certificate
//   • GetConnectionStatusByID(ctx, id) - Gets connection status
//   • GetDependenciesByID(ctx, id) - Gets dependent configuration profiles
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Full lifecycle requires a real DigiCert account/certificate;
//                validation errors can be tested without external dependencies
//     -- Tests: TestAcceptance_Digicert_validation_errors
//
// Notes
// -----------------------------------------------------------------------------
//   • Full CRUD lifecycle requires valid DigiCert credentials and client cert
//   • Validation errors are the only tests that reliably run in all environments
//
// =============================================================================

func TestAcceptance_Digicert_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Digicert

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("GetByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("UpdateByID_EmptyID", func(t *testing.T) {
		_, err := svc.UpdateByID(context.Background(), "", &digicert.ResourceDigicertTrustLifecycleManager{
			CAName: "test",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("UpdateByID_NilRequest", func(t *testing.T) {
		_, err := svc.UpdateByID(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteByID_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("ValidateClientCertificate_NilRequest", func(t *testing.T) {
		_, err := svc.ValidateClientCertificate(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("GetConnectionStatusByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetConnectionStatusByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("GetDependenciesByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetDependenciesByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})
}
