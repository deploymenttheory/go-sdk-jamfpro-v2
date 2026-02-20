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

func TestAcceptance_ComputerPrestages_CreateGetUpdateDeleteReplaceScope(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ComputerPrestages
	ctx := context.Background()
	name := acc.UniqueName("acc-computer-prestage")

	// Helper to create bool pointers
	boolPtr := func(b bool) *bool { return &b }

	created, resp, err := svc.CreateV3(ctx, &computer_prestages.ResourceComputerPrestage{
		DisplayName:                        name,
		Mandatory:                          boolPtr(true),
		MDMRemovable:                       boolPtr(true),
		SupportPhoneNumber:                 "111-222-3333",
		SupportEmailAddress:                "email@company.com",
		Department:                         "IT Department",
		DefaultPrestage:                    boolPtr(false),
		EnrollmentSiteId:                   "-1",
		KeepExistingSiteMembership:         boolPtr(false),
		KeepExistingLocationInformation:    boolPtr(false),
		RequireAuthentication:              boolPtr(false),
		AuthenticationPrompt:               "Welcome to your enterprise managed macOS device",
		PreventActivationLock:              boolPtr(false),
		EnableDeviceBasedActivationLock:    boolPtr(false),
		DeviceEnrollmentProgramInstanceId:  "1",
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
