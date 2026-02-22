package jamf_pro_api

import (
	"context"
	"net/url"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/managed_software_updates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Managed Software Updates
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   Managed Software Updates Operations (V1 API):
//   • GetAvailableUpdates(ctx) - Retrieves available software updates
//   • GetPlans(ctx, params) - Lists all managed software update plans with pagination
//   • GetPlanByUUID(ctx, uuid) - Retrieves a specific plan by UUID
//   • GetDeclarationsByPlanUUID(ctx, uuid) - Gets declarations for a plan
//   • CreatePlanByDeviceID(ctx, plan) - Creates a plan for specific devices
//   • CreatePlanByGroupID(ctx, plan) - Creates a plan for a device group
//   • GetPlansByGroupID(ctx, groupID, groupType) - Gets plans for a group
//   • GetFeatureToggle(ctx) - Retrieves feature toggle settings
//   • UpdateFeatureToggle(ctx, toggle) - Updates feature toggle
//   • GetFeatureToggleStatus(ctx) - Gets feature toggle status
//   • ForceStopFeatureToggleProcess(ctx) - Emergency stop for toggle process
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Operations
//     -- Reason: Read operations for available updates, plans, and feature toggles
//     -- Tests: TestAcceptance_ManagedSoftwareUpdates_GetAvailableUpdates,
//               TestAcceptance_ManagedSoftwareUpdates_GetPlans,
//               TestAcceptance_ManagedSoftwareUpdates_FeatureToggle
//     -- Flow: Get Available Updates → List Plans → Get Plan Details
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get available updates (macOS and iOS versions)
//   ✓ List operations (all plans with pagination)
//   ✓ Read operations (GetPlanByUUID, GetDeclarationsByPlanUUID if plans exist)
//   ✓ Feature toggle operations (get, status)
//   ✗ Create/Update operations (require specific test environment setup)
//   ✗ Input validation and error handling (not yet tested)
//
// Notes
// -----------------------------------------------------------------------------
//   • Managed Software Updates enable automated OS updates for macOS and iOS devices
//   • Plans define update schedules, deferral limits, and installation behavior
//   • Feature toggle controls whether the managed updates feature is enabled
//   • Declarations are DDM (Declarative Device Management) payloads
//   • Tests may skip if no plans exist in the environment
//   • Create operations require devices and should be tested in controlled environments
//   • ForceStopFeatureToggleProcess is a "break glass" operation and not tested here
//   • TODO: Add validation error tests for empty IDs/UUIDs
//
// =============================================================================

func TestAcceptance_ManagedSoftwareUpdates_GetAvailableUpdates(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ManagedSoftwareUpdates
	ctx := context.Background()

	result, resp, err := svc.GetAvailableUpdates(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result.AvailableUpdates)

	if len(result.AvailableUpdates.MacOS) > 0 {
		acc.LogTestSuccess(t, "Found %d macOS updates available", len(result.AvailableUpdates.MacOS))
		acc.LogTestSuccess(t, "Latest macOS version: %s", result.AvailableUpdates.MacOS[0])
	} else {
		acc.LogTestSuccess(t, "No macOS updates available")
	}

	if len(result.AvailableUpdates.IOS) > 0 {
		acc.LogTestSuccess(t, "Found %d iOS updates available", len(result.AvailableUpdates.IOS))
		acc.LogTestSuccess(t, "Latest iOS version: %s", result.AvailableUpdates.IOS[0])
	} else {
		acc.LogTestSuccess(t, "No iOS updates available")
	}
}

func TestAcceptance_ManagedSoftwareUpdates_GetPlans(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ManagedSoftwareUpdates
	ctx := context.Background()

	params := url.Values{}
	result, resp, err := svc.GetPlans(ctx, params)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
	assert.NotNil(t, result.Results)

	if result.TotalCount > 0 {
		acc.LogTestSuccess(t, "Found %d managed software update plans", result.TotalCount)

		// Verify first plan has expected fields
		plan := result.Results[0]
		assert.NotEmpty(t, plan.PlanUuid)
		assert.NotEmpty(t, plan.Device.DeviceId)
		assert.NotEmpty(t, plan.Device.ObjectType)
		assert.NotEmpty(t, plan.UpdateAction)
		assert.NotEmpty(t, plan.VersionType)
		acc.LogTestSuccess(t, "Sample plan: UUID=%s, Device=%s, Action=%s, Version=%s",
			plan.PlanUuid, plan.Device.DeviceId, plan.UpdateAction, plan.VersionType)

		// Test GetPlanByUUID
		acc.LogTestStage(t, "Read", "Getting plan by UUID")
		planDetail, resp, err := svc.GetPlanByUUID(ctx, plan.PlanUuid)
		require.NoError(t, err)
		require.NotNil(t, planDetail)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, plan.PlanUuid, planDetail.PlanUuid)
		acc.LogTestSuccess(t, "Retrieved plan by UUID: %s", planDetail.PlanUuid)

		// Test GetDeclarationsByPlanUUID
		acc.LogTestStage(t, "Read", "Getting declarations for plan")
		declarations, resp, err := svc.GetDeclarationsByPlanUUID(ctx, plan.PlanUuid)
		require.NoError(t, err)
		require.NotNil(t, declarations)
		assert.Equal(t, 200, resp.StatusCode)
		acc.LogTestSuccess(t, "Retrieved %d declarations for plan", len(declarations.Declarations))

		if len(declarations.Declarations) > 0 {
			decl := declarations.Declarations[0]
			assert.NotEmpty(t, decl.UUID)
			assert.NotEmpty(t, decl.Type)
			acc.LogTestSuccess(t, "Sample declaration: UUID=%s, Type=%s", decl.UUID, decl.Type)
		}
	} else {
		acc.LogTestSuccess(t, "No managed software update plans found (empty list OK)")
	}
}

func TestAcceptance_ManagedSoftwareUpdates_FeatureToggle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ManagedSoftwareUpdates
	ctx := context.Background()

	// Get current feature toggle state
	acc.LogTestStage(t, "Read", "Getting current feature toggle state")
	toggle, resp, err := svc.GetFeatureToggle(ctx)
	require.NoError(t, err)
	require.NotNil(t, toggle)
	assert.Equal(t, 200, resp.StatusCode)
	acc.LogTestSuccess(t, "Current feature toggle state: %v", toggle.Toggle)

	// Get feature toggle status
	acc.LogTestStage(t, "Read", "Getting feature toggle background status")
	status, resp, err := svc.GetFeatureToggleStatus(ctx)
	require.NoError(t, err)
	require.NotNil(t, status)
	assert.Equal(t, 200, resp.StatusCode)

	if status.ToggleOn != nil {
		acc.LogTestSuccess(t, "Toggle ON status: State=%s, Progress=%s",
			status.ToggleOn.State, status.ToggleOn.FormattedPercentComplete)
		assert.NotEmpty(t, status.ToggleOn.State)
	}

	if status.ToggleOff != nil {
		acc.LogTestSuccess(t, "Toggle OFF status: State=%s, Progress=%s",
			status.ToggleOff.State, status.ToggleOff.FormattedPercentComplete)
		assert.NotEmpty(t, status.ToggleOff.State)
	}

	if status.ToggleOn == nil && status.ToggleOff == nil {
		acc.LogTestSuccess(t, "No active toggle operations in progress")
	}
}

func TestAcceptance_ManagedSoftwareUpdates_CreatePlanByDeviceID(t *testing.T) {
	t.Skip("Skipping create plan test - requires specific device setup")

	acc.RequireClient(t)
	svc := acc.Client.ManagedSoftwareUpdates
	ctx := context.Background()

	// This is an example of how to create a plan - adapt to your environment
	plan := &managed_software_updates.RequestPlanCreate{
		Devices: []managed_software_updates.PlanObject{
			{
				ObjectType: "COMPUTER",
				DeviceId:   "TEST_DEVICE_ID", // Replace with actual device ID
			},
		},
		Config: managed_software_updates.PlanConfig{
			UpdateAction: "DOWNLOAD_INSTALL_ALLOW_DEFERRAL",
			VersionType:  "LATEST_MAJOR",
			MaxDeferrals: 3,
		},
	}

	acc.LogTestStage(t, "Create", "Creating managed software update plan")
	result, resp, err := svc.CreatePlanByDeviceID(ctx, plan)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Len(t, result.Plans, 1)
	acc.LogTestSuccess(t, "Created plan with UUID: %s", result.Plans[0].PlanID)
}

func TestAcceptance_ManagedSoftwareUpdates_GetPlansByGroupID(t *testing.T) {
	t.Skip("Skipping get plans by group ID test - requires specific group setup")

	acc.RequireClient(t)
	svc := acc.Client.ManagedSoftwareUpdates
	ctx := context.Background()

	groupID := "TEST_GROUP_ID" // Replace with actual group ID
	groupType := "COMPUTER"

	acc.LogTestStage(t, "Read", "Getting plans for group")
	result, resp, err := svc.GetPlansByGroupID(ctx, groupID, groupType)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	acc.LogTestSuccess(t, "Found %d plans for group %s", result.TotalCount, groupID)
}
