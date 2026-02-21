package jamf_pro_api

import (
	"context"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_prestages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Computer Prestages
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   Prestage CRUD (V3 API):
//   • ListV3(ctx, query) - Lists computer prestages with pagination/sorting (page, page-size, sort)
//   • GetByIDV3(ctx, id) - Retrieves a computer prestage by ID
//   • GetByNameV3(ctx, name) - Retrieves a computer prestage by display name (helper)
//   • CreateV3(ctx, request) - Creates a new computer prestage
//   • UpdateByIDV3(ctx, id, request) - Updates an existing computer prestage
//   • UpdateByNameV3(ctx, name, request) - Updates a computer prestage by display name (helper)
//   • DeleteByIDV3(ctx, id) - Deletes a computer prestage by ID
//   • DeleteByNameV3(ctx, name) - Deletes a computer prestage by display name (helper)
//
//   Device Scope Management (V2 API):
//   • GetDeviceScopeByIDV2(ctx, id) - Retrieves device scope for a prestage
//   • ReplaceDeviceScopeByIDV2(ctx, id, request) - Replaces device scope for a prestage
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_ComputerPrestages_LifecycleReplaceScope
//     -- Flow: Create → GetByID → GetByName → GetDeviceScope → ReplaceScope → Update → Delete
//
//   ✓ List Operations
//     -- Tests: TestAcceptance_ComputerPrestages_ListV3
//     -- Flow: List all prestages → Verify response structure
//
//   Note: RSQL Filter Testing NOT applicable
//     -- ListV3 supports pagination (page, page-size) and sorting (sort), not RSQL filtering
//     -- Query params are for pagination/sorting only, not RSQL filter expressions
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (prestage creation with comprehensive settings)
//   ✓ Read operations (GetByID, GetByName, List with pagination)
//   ✓ Update operations (full resource update with version locking)
//   ✓ Delete operations (single delete by ID)
//   ✓ Device scope operations (get scope, replace scope with version locking)
//   ✓ Cleanup and resource management
//   ✗ Input validation and error handling (not yet tested)
//   ✗ Update by name operations (not yet tested)
//   ✗ Delete by name operations (not yet tested)
//
// Notes
// -----------------------------------------------------------------------------
//   • Computer prestages define automated enrollment settings for DEP/ABM devices
//   • Requires enrollment site and DEP instance - test handles gracefully if unavailable
//   • Uses optimistic locking with VersionLock field (critical for updates)
//   • CRUD operations use V3 API, device scope operations use V2 API
//   • GetByName, UpdateByName, DeleteByName are helper methods (use ListV3 for lookup)
//   • Device scope defines which serial numbers are assigned to this prestage
//   • ReplaceDeviceScope requires current VersionLock to prevent conflicts
//   • All tests register cleanup handlers to remove test prestages
//   • Tests use acc.UniqueName() to avoid conflicts in shared test environments
//   • SkipSetupItems controls which setup assistant screens are skipped
//   • AccountSettings, LocationInformation, PurchasingInformation are complex nested objects
//   • TODO: Add validation error tests for empty IDs, nil requests, etc.
//   • TODO: Add tests for UpdateByNameV3 and DeleteByNameV3 operations
//
// =============================================================================

// Helper to create bool pointers
func boolPtr(b bool) *bool { return &b }

func TestAcceptance_ComputerPrestages_ListV3(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ComputerPrestages
	ctx := context.Background()

	result, resp, err := svc.ListV3(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
	assert.NotNil(t, result.Results)
}

func TestAcceptance_ComputerPrestages_LifecycleReplaceScope(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ComputerPrestages
	ctx := context.Background()
	name := acc.UniqueName("acc-computer-prestage")

	// Helper to create bool pointers
	boolPtr := func(b bool) *bool { return &b }

	created, resp, err := svc.CreateV3(ctx, &computer_prestages.ResourceComputerPrestage{
		DisplayName:                       name,
		Mandatory:                         boolPtr(true),
		MDMRemovable:                      boolPtr(true),
		SupportPhoneNumber:                "111-222-3333",
		SupportEmailAddress:               "email@company.com",
		Department:                        "IT Department",
		DefaultPrestage:                   boolPtr(false),
		EnrollmentSiteId:                  "-1",
		KeepExistingSiteMembership:        boolPtr(false),
		KeepExistingLocationInformation:   boolPtr(false),
		RequireAuthentication:             boolPtr(false),
		AuthenticationPrompt:              "Welcome to your enterprise managed macOS device",
		PreventActivationLock:             boolPtr(false),
		EnableDeviceBasedActivationLock:   boolPtr(false),
		DeviceEnrollmentProgramInstanceId: "1",
		SkipSetupItems: &computer_prestages.SkipSetupItems{
			Biometric:          boolPtr(false),
			TermsOfAddress:     boolPtr(false),
			FileVault:          boolPtr(false),
			ICloudDiagnostics:  boolPtr(false),
			Diagnostics:        boolPtr(false),
			Accessibility:      boolPtr(false),
			AppleID:            boolPtr(false),
			ScreenTime:         boolPtr(false),
			Siri:               boolPtr(false),
			DisplayTone:        boolPtr(false),
			Restore:            boolPtr(false),
			Appearance:         boolPtr(false),
			Privacy:            boolPtr(false),
			Payment:            boolPtr(false),
			Registration:       boolPtr(false),
			TOS:                boolPtr(false),
			ICloudStorage:      boolPtr(false),
			Location:           boolPtr(false),
			Intelligence:       boolPtr(false),
			EnableLockdownMode: boolPtr(false),
			Welcome:            boolPtr(false),
			Wallpaper:          boolPtr(false),
		},
		LocationInformation: computer_prestages.LocationInformation{
			ID:           "-1",
			Username:     "testuser",
			Realname:     "Test User",
			Phone:        "555-0000",
			Email:        "test@example.com",
			Room:         "Unassigned",
			Position:     "Employee",
			DepartmentId: "-1",
			BuildingId:   "-1",
			VersionLock:  1,
		},
		PurchasingInformation: computer_prestages.PurchasingInformation{
			ID:                "-1",
			Leased:            boolPtr(false),
			Purchased:         boolPtr(true),
			AppleCareId:       "Not Applicable",
			PONumber:          "N/A",
			Vendor:            "Apple",
			PurchasePrice:     "0.00",
			LifeExpectancy:    5,
			PurchasingAccount: "IT Budget",
			PurchasingContact: "IT Department",
			LeaseDate:         "1970-01-01",
			PODate:            "1970-01-01",
			WarrantyDate:      "1970-01-01",
			VersionLock:       1,
		},
		EnrollmentCustomizationId:          "0",
		AutoAdvanceSetup:                   boolPtr(false),
		InstallProfilesDuringSetup:         boolPtr(true),
		PrestageInstalledProfileIds:        []string{},
		CustomPackageIds:                   []string{},
		CustomPackageDistributionPointId:   "-1",
		EnableRecoveryLock:                 boolPtr(false),
		RecoveryLockPasswordType:           "",
		RecoveryLockPassword:               "",
		RotateRecoveryLockPassword:         boolPtr(false),
		PrestageMinimumOsTargetVersionType: "NO_ENFORCEMENT",
		MinimumOsSpecificVersion:           "",
		SiteId:                             "-1",
		VersionLock:                        0,
		AccountSettings: &computer_prestages.AccountSettings{
			PayloadConfigured:                       boolPtr(true),
			LocalAdminAccountEnabled:                boolPtr(false),
			AdminUsername:                           "",
			AdminPassword:                           "",
			HiddenAdminAccount:                      boolPtr(false),
			LocalUserManaged:                        boolPtr(false),
			UserAccountType:                         "ADMINISTRATOR",
			VersionLock:                             0,
			PrefillPrimaryAccountInfoFeatureEnabled: boolPtr(false),
			PrefillType:                             "UNKNOWN",
			PrefillAccountFullName:                  "",
			PrefillAccountUserName:                  "",
			PreventPrefillInfoFromModification:      boolPtr(false),
		},
	})
	if err != nil {
		t.Skipf("CreateV3 failed (may require enrollment site / DEP): %v", err)
		return
	}
	require.NotNil(t, created)
	assert.Contains(t, []int{200, 201}, resp.StatusCode)

	id := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV3(cleanupCtx, id)
	})

	getByID, resp, err := svc.GetByIDV3(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, getByID)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, name, getByID.DisplayName)

	byName, resp, err := svc.GetByNameV3(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, byName)
	assert.Equal(t, name, byName.DisplayName)

	scope, resp, err := svc.GetDeviceScopeByIDV2(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, scope)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, id, scope.PrestageId)

	replaceReq := &computer_prestages.ReplaceDeviceScopeRequest{
		SerialNumbers: []string{"ACC-TEST-SERIAL"},
		VersionLock:   scope.VersionLock,
	}
	replaced, resp, err := svc.ReplaceDeviceScopeByIDV2(ctx, id, replaceReq)
	require.NoError(t, err)
	require.NotNil(t, replaced)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, id, replaced.PrestageId)
	assert.GreaterOrEqual(t, replaced.VersionLock, scope.VersionLock)

	scope2, resp, err := svc.GetDeviceScopeByIDV2(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, scope2)
	assert.Equal(t, 200, resp.StatusCode)

	updated, resp, err := svc.UpdateByIDV3(ctx, id, &computer_prestages.ResourceComputerPrestage{
		DisplayName: name,
		VersionLock: getByID.VersionLock,
	})
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, resp.StatusCode)

	delResp, err := svc.DeleteByIDV3(ctx, id)
	require.NoError(t, err)
	assert.Contains(t, []int{200, 204}, delResp.StatusCode)
}
