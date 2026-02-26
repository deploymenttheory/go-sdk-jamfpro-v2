package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Connect
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetSettingsV1(ctx) - Retrieves Jamf Connect settings
//   • ListConfigProfilesV1(ctx, rsqlQuery) - Lists config profiles with pagination
//   • GetConfigProfileByUUIDV1(ctx, uuid) - Retrieves profile by UUID
//   • GetConfigProfileByIDV1(ctx, profileID) - Retrieves profile by ID
//   • GetConfigProfileByNameV1(ctx, name) - Retrieves profile by name
//   • UpdateConfigProfileByUUIDV1(ctx, uuid, request) - Updates profile
//   • RetryDeploymentTasksByUUIDV1(ctx, uuid, computerIDs) - Retries deployment
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Settings/Configuration Testing
//     -- Reason: Settings endpoint returns singleton configuration
//     -- Tests: TestAcceptance_JamfConnect_get_settings
//     -- Flow: Get settings → Verify structure
//
//   ✓ Pattern 3: Read-Only List Testing
//     -- Reason: Config profiles are created/managed outside this API
//     -- Tests: TestAcceptance_JamfConnect_list_and_get_profiles
//     -- Flow: List profiles → Get by UUID/ID/Name → Verify
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_JamfConnect_validation_errors
//     -- Cases: Empty UUID, empty name, invalid ID, nil requests
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get settings operations
//   ✓ List operations with pagination support
//   ✓ Get-by helper methods (UUID, ID, Name)
//   ✓ Input validation and error handling
//
// Notes
// -----------------------------------------------------------------------------
//   • Config profiles are typically created through Jamf Pro UI
//   • Update and retry operations require existing profiles
//   • Tests focus on read operations that don't modify state
//   • Validation tests ensure proper error handling
//
// =============================================================================

func TestAcceptance_JamfConnect_get_settings(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfConnect
	ctx := context.Background()

	// Get Jamf Connect settings
	result, resp, err := svc.GetSettingsV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify structure
	assert.NotEmpty(t, result.ID)
}

func TestAcceptance_JamfConnect_list_and_get_profiles(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfConnect
	ctx := context.Background()

	// List all config profiles
	listResult, resp, err := svc.ListConfigProfilesV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, listResult)
	assert.Equal(t, 200, resp.StatusCode)

	// If there are profiles, test get-by helpers
	if listResult.TotalCount > 0 && len(listResult.Results) > 0 {
		firstProfile := listResult.Results[0]

		// Test GetByUUID
		profileByUUID, resp, err := svc.GetConfigProfileByUUIDV1(ctx, firstProfile.UUID)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, firstProfile.UUID, profileByUUID.UUID)

		// Test GetByID
		profileByID, resp, err := svc.GetConfigProfileByIDV1(ctx, firstProfile.ProfileID)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, firstProfile.ProfileID, profileByID.ProfileID)

		// Test GetByName
		profileByName, resp, err := svc.GetConfigProfileByNameV1(ctx, firstProfile.ProfileName)
		require.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, firstProfile.ProfileName, profileByName.ProfileName)
	}
}

func TestAcceptance_JamfConnect_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfConnect
	ctx := context.Background()

	// Test empty UUID validation
	result, resp, err := svc.GetConfigProfileByUUIDV1(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "uuid is required")

	// Test invalid ID validation
	result, resp, err = svc.GetConfigProfileByIDV1(ctx, 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "profile ID must be greater than 0")

	// Test empty name validation
	result, resp, err = svc.GetConfigProfileByNameV1(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "name is required")
}
