package jamf_pro_api

import (
	"context"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_prestages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Mobile Device Prestages
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   Prestage CRUD (V3 API):
//   • ListV3(ctx) - Lists mobile device prestages with automatic pagination
//   • GetByIDV3(ctx, id) - Retrieves a mobile device prestage by ID
//   • GetByNameV3(ctx, name) - Retrieves a mobile device prestage by display name (helper)
//   • CreateV3(ctx, request) - Creates a new mobile device prestage
//   • UpdateByIDV3(ctx, id, request) - Updates an existing mobile device prestage
//   • UpdateByNameV3(ctx, name, request) - Updates a mobile device prestage by display name (helper)
//   • DeleteByIDV3(ctx, id) - Deletes a mobile device prestage by ID
//   • DeleteByNameV3(ctx, name) - Deletes a mobile device prestage by display name (helper)
//
//   Device Scope Management (V2 API):
//   • GetScopeByIDV2(ctx, id) - Retrieves device scope for a prestage
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_MobileDevicePrestages_LifecycleWithScope
//     -- Flow: Create → GetByID → GetByName → GetScope → Update → Delete
//
//   ✓ List Operations
//     -- Tests: TestAcceptance_MobileDevicePrestages_ListV3
//     -- Flow: List all prestages → Verify response structure
//
//   Note: RSQL Filter Testing NOT applicable
//     -- ListV3 uses automatic pagination, not RSQL filtering
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (prestage creation with comprehensive settings)
//   ✓ Read operations (GetByID, GetByName, List with pagination)
//   ✓ Update operations (full resource update with version locking)
//   ✓ Delete operations (single delete by ID)
//   ✓ Device scope operations (get scope)
//   ✓ Cleanup and resource management
//   ✗ Input validation and error handling (not yet tested)
//   ✗ Update by name operations (not yet tested)
//   ✗ Delete by name operations (not yet tested)
//
// Notes
// -----------------------------------------------------------------------------
//   • Mobile device prestages define automated enrollment settings for iOS/iPadOS devices
//   • Requires enrollment site and DEP instance - test handles gracefully if unavailable
//   • Uses optimistic locking with VersionLock field (critical for updates)
//   • CRUD operations use V3 API, device scope operations use V2 API
//   • GetByName, UpdateByName, DeleteByName are helper methods (use ListV3 for lookup)
//   • Device scope defines which serial numbers are assigned to this prestage
//   • All tests register cleanup handlers to remove test prestages
//   • Tests use acc.UniqueName() to avoid conflicts in shared test environments
//   • SkipSetupItems controls which setup assistant screens are skipped
//   • LocationInformation, PurchasingInformation, Names are complex nested objects
//   • Supports device naming configuration and shared iPad settings
//   • TODO: Add validation error tests for empty IDs, nil requests, etc.
//   • TODO: Add tests for UpdateByNameV3 and DeleteByNameV3 operations
//
// =============================================================================

func TestAcceptance_MobileDevicePrestages_ListV3(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.MobileDevicePrestages
	ctx := context.Background()

	result, resp, err := svc.ListV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
	assert.NotNil(t, result.Results)
}

func TestAcceptance_MobileDevicePrestages_LifecycleWithScope(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.MobileDevicePrestages
	ctx := context.Background()
	name := acc.UniqueName("acc-mobile-prestage")

	// Create
	acc.LogTestStage(t, "Create", "Creating mobile device prestage")
	created, resp, err := svc.CreateV3(ctx, &mobile_device_prestages.ResourceMobileDevicePrestage{
		DisplayName:                       name,
		Mandatory:                         boolPtr(true),
		MdmRemovable:                      boolPtr(true),
		SupportPhoneNumber:                "111-222-3333",
		SupportEmailAddress:               "email@company.com",
		Department:                        "IT Department",
		DefaultPrestage:                   boolPtr(false),
		EnrollmentSiteID:                  "-1",
		KeepExistingSiteMembership:        boolPtr(false),
		KeepExistingLocationInformation:   boolPtr(false),
		RequireAuthentication:             boolPtr(false),
		AuthenticationPrompt:              "Welcome to your enterprise managed iOS device",
		PreventActivationLock:             boolPtr(false),
		EnableDeviceBasedActivationLock:   boolPtr(false),
		DeviceEnrollmentProgramInstanceID: "1",
		SkipSetupItems: mobile_device_prestages.SubsetSkipSetupItems{
			Location:              boolPtr(false),
			Privacy:               boolPtr(false),
			Biometric:             boolPtr(false),
			SoftwareUpdate:        boolPtr(false),
			Diagnostics:           boolPtr(false),
			IMessageAndFaceTime:   boolPtr(false),
			Intelligence:          boolPtr(false),
			TVRoom:                boolPtr(false),
			Passcode:              boolPtr(false),
			SIMSetup:              boolPtr(false),
			ScreenTime:            boolPtr(false),
			RestoreCompleted:      boolPtr(false),
			TVProviderSignIn:      boolPtr(false),
			Siri:                  boolPtr(false),
			Restore:               boolPtr(false),
			ScreenSaver:           boolPtr(false),
			HomeButtonSensitivity: boolPtr(false),
			CloudStorage:          boolPtr(false),
			ActionButton:          boolPtr(false),
			TransferData:          boolPtr(false),
			EnableLockdownMode:    boolPtr(false),
			Zoom:                  boolPtr(false),
			PreferredLanguage:     boolPtr(false),
			VoiceSelection:        boolPtr(false),
			TVHomeScreenSync:      boolPtr(false),
			Safety:                boolPtr(false),
			TermsOfAddress:        boolPtr(false),
			ExpressLanguage:       boolPtr(false),
			CameraButton:          boolPtr(false),
			AppleID:               boolPtr(false),
			DisplayTone:           boolPtr(false),
			WatchMigration:        boolPtr(false),
			UpdateCompleted:       boolPtr(false),
			Appearance:            boolPtr(false),
			Android:               boolPtr(false),
			Payment:               boolPtr(false),
			OnBoarding:            boolPtr(false),
			TOS:                   boolPtr(false),
			Welcome:               boolPtr(false),
			SafetyAndHandling:     boolPtr(false),
			TapToSetup:            boolPtr(false),
		},
		LocationInformation: mobile_device_prestages.SubsetLocationInformation{
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
		PurchasingInformation: mobile_device_prestages.SubsetPurchasingInformation{
			ID:                "-1",
			Leased:            boolPtr(false),
			Purchased:         boolPtr(true),
			AppleCareId:       "Not Applicable",
			PoNumber:          "N/A",
			Vendor:            "Apple",
			PurchasePrice:     "0.00",
			LifeExpectancy:    5,
			PurchasingAccount: "IT Budget",
			PurchasingContact: "IT Department",
			LeaseDate:         "1970-01-01",
			PoDate:            "1970-01-01",
			WarrantyDate:      "1970-01-01",
			VersionLock:       1,
		},
		EnrollmentCustomizationID:       "0",
		AutoAdvanceSetup:                boolPtr(false),
		AllowPairing:                    boolPtr(true),
		MultiUser:                       boolPtr(false),
		Supervised:                      boolPtr(true),
		MaximumSharedAccounts:           10,
		ConfigureDeviceBeforeSetupAssistant: boolPtr(false),
		Names: mobile_device_prestages.SubsetNames{
			AssignNamesUsing:       "STATIC",
			PrestageDeviceNames:    []mobile_device_prestages.SubsetNamesName{},
			DeviceNamePrefix:       "",
			DeviceNameSuffix:       "",
			SingleDeviceName:       "",
			ManageNames:            boolPtr(false),
			DeviceNamingConfigured: boolPtr(false),
		},
		SendTimezone:                   boolPtr(false),
		Timezone:                       "America/New_York",
		StorageQuotaSizeMegabytes:      0,
		UseStorageQuotaSize:            boolPtr(false),
		TemporarySessionOnly:           boolPtr(false),
		EnforceTemporarySessionTimeout: boolPtr(false),
		EnforceUserSessionTimeout:      boolPtr(false),
		SiteId:                         "-1",
		VersionLock:                    0,
	})
	if err != nil {
		t.Skipf("CreateV3 failed (may require enrollment site / DEP): %v", err)
		return
	}
	require.NotNil(t, created)
	assert.Contains(t, []int{200, 201}, resp.StatusCode)
	acc.LogTestSuccess(t, "Created mobile device prestage with ID: %s", created.ID)

	id := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV3(cleanupCtx, id)
	})

	// Read by ID
	acc.LogTestStage(t, "Read", "Getting mobile device prestage by ID")
	getByID, resp, err := svc.GetByIDV3(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, getByID)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, name, getByID.DisplayName)
	acc.LogTestSuccess(t, "Retrieved prestage by ID: %s", getByID.DisplayName)

	// Read by Name
	acc.LogTestStage(t, "Read", "Getting mobile device prestage by name")
	byName, resp, err := svc.GetByNameV3(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, byName)
	assert.Equal(t, name, byName.DisplayName)
	assert.Equal(t, id, byName.ID)
	acc.LogTestSuccess(t, "Retrieved prestage by name: %s", byName.DisplayName)

	// Get Device Scope
	acc.LogTestStage(t, "Scope", "Getting device scope")
	scope, resp, err := svc.GetScopeByIDV2(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, scope)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, id, scope.PrestageId)
	acc.LogTestSuccess(t, "Retrieved scope - VersionLock: %d, Assignments: %d", scope.VersionLock, len(scope.Assignments))

	// Update
	acc.LogTestStage(t, "Update", "Updating mobile device prestage")
	updated, resp, err := svc.UpdateByIDV3(ctx, id, &mobile_device_prestages.ResourceMobileDevicePrestage{
		DisplayName:         name,
		SupportPhoneNumber:  "999-888-7777",
		SupportEmailAddress: "updated@company.com",
		VersionLock:         getByID.VersionLock,
	})
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, resp.StatusCode)
	acc.LogTestSuccess(t, "Updated mobile device prestage")

	// Verify Update
	acc.LogTestStage(t, "Verify", "Verifying updated prestage")
	verifyUpdated, resp, err := svc.GetByIDV3(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, verifyUpdated)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "999-888-7777", verifyUpdated.SupportPhoneNumber)
	assert.Equal(t, "updated@company.com", verifyUpdated.SupportEmailAddress)
	acc.LogTestSuccess(t, "Verified update - Phone: %s", verifyUpdated.SupportPhoneNumber)

	// Delete
	acc.LogTestStage(t, "Delete", "Deleting mobile device prestage")
	delResp, err := svc.DeleteByIDV3(ctx, id)
	require.NoError(t, err)
	assert.Contains(t, []int{200, 204}, delResp.StatusCode)
	acc.LogTestSuccess(t, "Deleted mobile device prestage: %s", id)

	// Verify Deletion
	acc.LogTestStage(t, "Verify Deletion", "Verifying prestage is deleted")
	_, getAfterDeleteResp, err := svc.GetByIDV3(ctx, id)
	assert.Error(t, err)
	if getAfterDeleteResp != nil {
		assert.Equal(t, 404, getAfterDeleteResp.StatusCode)
	}
	acc.LogTestSuccess(t, "Verified deletion - prestage no longer exists")
}
