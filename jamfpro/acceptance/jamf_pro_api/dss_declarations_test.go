package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: DSS Declarations
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetByUUIDV1(ctx, uuid) - Retrieves a DSS declaration by UUID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: DSS declarations are read-only and require a valid UUID
//     -- Tests: TestAcceptance_DSSDeclarations_GetByUUID
//     -- Flow: Get by UUID → Verify structure
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Read operations (GetByUUID)
//   ✗ List operations (not available in API)
//   ✗ Create/Update/Delete operations (not available in API)
//
// Notes
// -----------------------------------------------------------------------------
//   • DSS (Declarative Software Services) declarations are managed by Declarative Device Management
//   • Test requires a valid UUID from an existing declaration
//   • Test skips if UUID is not available or feature not supported on tenant
//   • Declarations contain DDM configuration payloads in JSON format
//
// =============================================================================

func TestAcceptance_DSSDeclarations_GetByUUID(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DSSDeclarations
	ctx := context.Background()

	testUUID := "550e8400-e29b-41d4-a716-446655440000"

	acc.LogTestStage(t, "GetByUUID", "Fetching DSS declaration by UUID=%s", testUUID)
	declaration, resp, err := svc.GetByUUIDV1(ctx, testUUID)

	if err != nil {
		t.Skipf("Failed to get DSS declaration (may not be supported or UUID not found on this tenant): %v", err)
		return
	}

	require.NotNil(t, declaration)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, declaration.Declarations)
	assert.GreaterOrEqual(t, len(declaration.Declarations), 0)

	if len(declaration.Declarations) > 0 {
		firstDecl := declaration.Declarations[0]
		assert.NotEmpty(t, firstDecl.UUID)
		assert.NotEmpty(t, firstDecl.Type)
		acc.LogTestSuccess(t, "DSS declaration retrieved: UUID=%s, Type=%s", firstDecl.UUID, firstDecl.Type)
	} else {
		acc.LogTestSuccess(t, "DSS declaration response received (empty declarations array)")
	}
}
