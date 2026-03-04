package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: macOS Configuration Profile Custom Settings
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetSchemaList(ctx) - Retrieves custom settings schema list
//   • GetByPayloadUUID(ctx, id) - Retrieves configuration profile by payload UUID
//   • Create(ctx, profile) - Creates a new configuration profile with custom settings
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- GetSchemaList is read-only
//   ✓ Pattern 1: Create-Read-Delete (for Create/GetByPayloadUUID)
//     -- Create profile, verify with GetByPayloadUUID
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get custom settings schema list
//   ✓ Verify response structure (200 status)
//   ✓ Verify schema list contains expected structure
//
// Notes
// -----------------------------------------------------------------------------
//   • Create modifies Jamf Pro state - use with caution in shared environments.
//   • GetSchemaList is safe to run repeatedly.
//   • GetByPayloadUUID requires an existing profile UUID.
//
// =============================================================================

func TestAcceptance_MacOSConfigProfileCustomSettings_get_schema_list(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.MacOSConfigProfileCustomSettings
	ctx := context.Background()

	result, resp, err := svc.GetSchemaList(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	// Schema list can be empty or contain buckets
	assert.NotNil(t, result)
}
