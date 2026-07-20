package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	mcp "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/macos_configuration_profiles"
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
//   • UpdateByPayloadUUID(ctx, id, profile) - Replaces an existing profile
//   • DeleteByPayloadUUID(ctx, id) - Deletes a profile by payload UUID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Tests: TestAcceptance_MacOSConfigurationProfiles_lifecycle
//     -- Flow: Create → Get → Update → Get (verify) → Delete
//   ✓ Pattern 3: Read-Only Information
//     -- GetSchemaList is read-only and safe to run repeatedly
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_MacOSConfigurationProfiles_validation_errors
//
// Notes
// -----------------------------------------------------------------------------
//   • There is NO List operation. /api/config-profiles/macos allows POST only;
//     GET on the collection returns 405. Created profiles cannot be verified
//     against a collection -- verification is via GetByPayloadUUID only.
//   • DELETE is a known server-side defect: it returns 204 but does not remove
//     the profile. Confirmed against Jamf Pro 11.30.0 by polling for 10 minutes
//     after a successful delete. Tests MUST NOT assert read-after-delete 404,
//     and the lifecycle test leaks one profile per run as a result.
//   • `level` is write-only and never returned by GET -- do not assert on it.
//   • Profiles created here do NOT appear in the Classic API
//     /JSSResource/osxconfigurationprofiles list. This endpoint manages a
//     payload-authoring object, not a fully scoped configuration profile.
//   • Supported payloadTypes are an undocumented allowlist wider than Custom
//     Settings alone -- com.apple.notificationsettings and
//     com.apple.mobiledevice.passwordpolicy also create successfully, and
//     several more are recognised with their own field rules. Certificates,
//     TCC/PPPC, VPN and FileVault payloads are not supported. The SDK does not
//     validate payloadType; see the service package documentation.
//   • preferenceDomain is required only by com.apple.ManagedClient.preferences.
//
// =============================================================================

const testProfilePlist = `<?xml version="1.0" encoding="UTF-8"?><plist version="1.0"><dict><key>TestKey</key><string>TestValue</string></dict></plist>`

const updatedProfilePlist = `<?xml version="1.0" encoding="UTF-8"?><plist version="1.0"><dict><key>TestKey</key><string>UpdatedValue</string></dict></plist>`

func TestAcceptance_MacOSConfigurationProfiles_get_schema_list(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.MacosConfigurationProfiles
	ctx := context.Background()

	result, resp, err := svc.GetSchemaList(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	// Schema list can be empty or contain buckets
	assert.NotNil(t, result)
}

func TestAcceptance_MacOSConfigurationProfiles_lifecycle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.MacosConfigurationProfiles
	ctx := context.Background()

	domain := acc.UniqueName("com.example.acctest")

	acc.LogTestStage(t, "CREATE", "creating profile with preferenceDomain %s", domain)
	profile := &mcp.ResourceConfigProfile{
		Level: mcp.ConfigProfileLevelSystem,
		PayloadContent: []mcp.PayloadContentItem{
			{
				PayloadType:        mcp.PayloadTypeManagedClientPreferences,
				PayloadVersion:     1,
				PayloadIdentifier:  domain,
				PayloadDisplayName: acc.UniqueName("acctest-profile"),
				PreferenceDomain:   domain,
				Forced:             &mcp.ForcedSettings{Plist: testProfilePlist},
			},
		},
	}

	created, resp, err := svc.Create(ctx, profile)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 201, resp.StatusCode())
	require.NotEmpty(t, created.UUID)
	profileUUID := created.UUID
	acc.LogTestSuccess(t, "created profile %s", profileUUID)

	acc.Cleanup(t, func() {
		// Best-effort only: DELETE returns 204 but does not remove the
		// profile on Jamf Pro 11.30.0, so this run leaks one record.
		if _, err := svc.DeleteByPayloadUUID(ctx, profileUUID); err != nil {
			acc.LogCleanupDeleteError(t, "macOS config profile custom settings", profileUUID, err)
		}
	})

	acc.LogTestStage(t, "READ", "reading profile %s back", profileUUID)
	got, resp, err := svc.GetByPayloadUUID(ctx, profileUUID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, profileUUID, got.PayloadUUID)
	require.Len(t, got.PayloadContent, 1)
	assert.Equal(t, domain, got.PayloadContent[0].PreferenceDomain)
	// Do not assert on got.Level -- the API never returns it.

	acc.LogTestStage(t, "UPDATE", "updating profile %s", profileUUID)
	profile.PayloadContent[0].Forced.Plist = updatedProfilePlist
	updated, resp, err := svc.UpdateByPayloadUUID(ctx, profileUUID, profile)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, resp.StatusCode())
	// Update returns the create envelope, not the resource.
	assert.Equal(t, profileUUID, updated.UUID)

	acc.LogTestStage(t, "VERIFY", "confirming the update landed")
	got, _, err = svc.GetByPayloadUUID(ctx, profileUUID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Len(t, got.PayloadContent, 1)
	require.NotNil(t, got.PayloadContent[0].Forced)
	assert.Equal(t, updatedProfilePlist, got.PayloadContent[0].Forced.Plist)
	acc.LogTestSuccess(t, "update verified")

	acc.LogTestStage(t, "DELETE", "deleting profile %s", profileUUID)
	resp, err = svc.DeleteByPayloadUUID(ctx, profileUUID)
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
	// Stop here. A follow-up GET still returns 200 -- see the notes above.
	acc.LogTestSuccess(t, "delete accepted")
}

func TestAcceptance_MacOSConfigurationProfiles_get_by_payload_uuid_not_found(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.MacosConfigurationProfiles
	ctx := context.Background()

	result, resp, err := svc.GetByPayloadUUID(ctx, "00000000-0000-0000-0000-000000000000")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// TestAcceptance_MacOSConfigurationProfiles_validation_errors exercises
// the client-side guards, so it issues no HTTP requests.
func TestAcceptance_MacOSConfigurationProfiles_validation_errors(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.MacosConfigurationProfiles
	ctx := context.Background()

	valid := &mcp.ResourceConfigProfile{
		PayloadContent: []mcp.PayloadContentItem{
			{
				PayloadType:       mcp.PayloadTypeManagedClientPreferences,
				PayloadVersion:    1,
				PayloadIdentifier: "com.example.app",
				PreferenceDomain:  "com.example.app",
			},
		},
	}

	t.Run("create nil profile", func(t *testing.T) {
		_, _, err := svc.Create(ctx, nil)
		assert.ErrorContains(t, err, "profile is required")
	})

	t.Run("get empty id", func(t *testing.T) {
		_, _, err := svc.GetByPayloadUUID(ctx, "")
		assert.ErrorContains(t, err, "payload UUID is required")
	})

	t.Run("update empty id", func(t *testing.T) {
		_, _, err := svc.UpdateByPayloadUUID(ctx, "", valid)
		assert.ErrorContains(t, err, "payload UUID is required")
	})

	t.Run("update nil profile", func(t *testing.T) {
		_, _, err := svc.UpdateByPayloadUUID(ctx, "some-uuid", nil)
		assert.ErrorContains(t, err, "profile is required")
	})

	t.Run("delete empty id", func(t *testing.T) {
		_, err := svc.DeleteByPayloadUUID(ctx, "")
		assert.ErrorContains(t, err, "payload UUID is required")
	})

	t.Run("title case level", func(t *testing.T) {
		// The form real .mobileconfig files use; the API rejects it.
		profile := &mcp.ResourceConfigProfile{Level: "System"}
		_, resp, err := svc.Create(ctx, profile)
		assert.ErrorContains(t, err, "invalid level")
		assert.Nil(t, resp)
	})
}
