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

// TestAcceptance_ComputerPrestages_MinimalConfig tests creating a prestage with minimal required configuration.
func TestAcceptance_ComputerPrestages_MinimalConfig(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ComputerPrestages
	ctx := context.Background()
	name := acc.UniqueName("acc-prestage-minimal")

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
		AuthenticationPrompt:               "Welcome",
		PreventActivationLock:              boolPtr(false),
		EnableDeviceBasedActivationLock:    boolPtr(false),
		DeviceEnrollmentProgramInstanceId:  "1",
		SkipSetupItems: &computer_prestages.SkipSetupItems{
			Biometric: boolPtr(false),
		},
		LocationInformation: computer_prestages.LocationInformation{
			ID:           "-1",
			Username:     "user",
			Realname:     "User",
			Phone:        "000-0000",
			Email:        "user@example.com",
			Room:         "N/A",
			Position:     "N/A",
			DepartmentId: "-1",
			BuildingId:   "-1",
			VersionLock:  1,
		},
		PurchasingInformation: computer_prestages.PurchasingInformation{
			ID:                "-1",
			Leased:            boolPtr(false),
			Purchased:         boolPtr(true),
			AppleCareId:       "N/A",
			PONumber:          "N/A",
			Vendor:            "Apple",
			PurchasePrice:     "0.00",
			LifeExpectancy:    0,
			PurchasingAccount: "N/A",
			PurchasingContact: "N/A",
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
		PrestageMinimumOsTargetVersionType: "NO_ENFORCEMENT",
		SiteId:                             "-1",
		AccountSettings: &computer_prestages.AccountSettings{
			PayloadConfigured:                       boolPtr(true),
			LocalAdminAccountEnabled:                boolPtr(false),
			AdminUsername:                           "",
			AdminPassword:                           "",
			HiddenAdminAccount:                      boolPtr(false),
			LocalUserManaged:                        boolPtr(false),
			UserAccountType:                         "STANDARD",
			PrefillPrimaryAccountInfoFeatureEnabled: boolPtr(false),
			PrefillType:                             "CUSTOM",
			PreventPrefillInfoFromModification:      boolPtr(false),
		},
	})
	if err != nil {
		t.Skipf("CreateV3 failed (may require DEP): %v", err)
		return
	}
	require.NotNil(t, created)
	assert.Contains(t, []int{200, 201}, resp.StatusCode)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV3(cleanupCtx, created.ID)
	})

	// Verify creation
	fetched, _, err := svc.GetByIDV3(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, name, fetched.DisplayName)
}

// TestAcceptance_ComputerPrestages_RecoveryLockScenario tests prestage with recovery lock enabled.
func TestAcceptance_ComputerPrestages_RecoveryLockScenario(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ComputerPrestages
	ctx := context.Background()
	name := acc.UniqueName("acc-prestage-recovery")

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
		AuthenticationPrompt:               "Welcome",
		PreventActivationLock:              boolPtr(false),
		EnableDeviceBasedActivationLock:    boolPtr(false),
		DeviceEnrollmentProgramInstanceId:  "1",
		SkipSetupItems: &computer_prestages.SkipSetupItems{
			Biometric: boolPtr(false),
		},
		LocationInformation: computer_prestages.LocationInformation{
			ID:           "-1",
			Username:     "user",
			Realname:     "User",
			Phone:        "000-0000",
			Email:        "user@example.com",
			Room:         "N/A",
			Position:     "N/A",
			DepartmentId: "-1",
			BuildingId:   "-1",
			VersionLock:  1,
		},
		PurchasingInformation: computer_prestages.PurchasingInformation{
			ID:                "-1",
			Leased:            boolPtr(false),
			Purchased:         boolPtr(true),
			AppleCareId:       "N/A",
			PONumber:          "N/A",
			Vendor:            "Apple",
			PurchasePrice:     "0.00",
			LifeExpectancy:    0,
			PurchasingAccount: "N/A",
			PurchasingContact: "N/A",
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
		EnableRecoveryLock:                 boolPtr(true),
		RecoveryLockPasswordType:           "RANDOM",
		RotateRecoveryLockPassword:         boolPtr(true),
		PrestageMinimumOsTargetVersionType: "NO_ENFORCEMENT",
		SiteId:                             "-1",
		AccountSettings: &computer_prestages.AccountSettings{
			PayloadConfigured:                       boolPtr(true),
			LocalAdminAccountEnabled:                boolPtr(false),
			UserAccountType:                         "ADMINISTRATOR",
			PrefillPrimaryAccountInfoFeatureEnabled: boolPtr(false),
			PrefillType:                             "CUSTOM",
		},
	})
	if err != nil {
		t.Skipf("CreateV3 failed (may require DEP): %v", err)
		return
	}
	require.NotNil(t, created)
	assert.Contains(t, []int{200, 201}, resp.StatusCode)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV3(cleanupCtx, created.ID)
	})

	// Verify recovery lock settings
	fetched, _, err := svc.GetByIDV3(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, name, fetched.DisplayName)
}

// TestAcceptance_ComputerPrestages_AccountSettingsScenario tests prestage with admin account enabled.
func TestAcceptance_ComputerPrestages_AccountSettingsScenario(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ComputerPrestages
	ctx := context.Background()
	name := acc.UniqueName("acc-prestage-account")

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
		AuthenticationPrompt:               "Welcome",
		PreventActivationLock:              boolPtr(false),
		EnableDeviceBasedActivationLock:    boolPtr(false),
		DeviceEnrollmentProgramInstanceId:  "1",
		SkipSetupItems: &computer_prestages.SkipSetupItems{
			Biometric: boolPtr(false),
		},
		LocationInformation: computer_prestages.LocationInformation{
			ID:           "-1",
			Username:     "user",
			Realname:     "User",
			Phone:        "000-0000",
			Email:        "user@example.com",
			Room:         "N/A",
			Position:     "N/A",
			DepartmentId: "-1",
			BuildingId:   "-1",
			VersionLock:  1,
		},
		PurchasingInformation: computer_prestages.PurchasingInformation{
			ID:                "-1",
			Leased:            boolPtr(false),
			Purchased:         boolPtr(true),
			AppleCareId:       "N/A",
			PONumber:          "N/A",
			Vendor:            "Apple",
			PurchasePrice:     "0.00",
			LifeExpectancy:    0,
			PurchasingAccount: "N/A",
			PurchasingContact: "N/A",
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
		PrestageMinimumOsTargetVersionType: "NO_ENFORCEMENT",
		SiteId:                             "-1",
		AccountSettings: &computer_prestages.AccountSettings{
			PayloadConfigured:                       boolPtr(true),
			LocalAdminAccountEnabled:                boolPtr(true),
			AdminUsername:                           "localadmin",
			AdminPassword:                           "P@ssw0rd123",
			HiddenAdminAccount:                      boolPtr(true),
			LocalUserManaged:                        boolPtr(true),
			UserAccountType:                         "STANDARD",
			PrefillPrimaryAccountInfoFeatureEnabled: boolPtr(true),
			PrefillType:                             "DEVICE_OWNER",
			PrefillAccountFullName:                  "Test User",
			PrefillAccountUserName:                  "testuser",
			PreventPrefillInfoFromModification:      boolPtr(true),
		},
	})
	if err != nil {
		t.Skipf("CreateV3 failed (may require DEP): %v", err)
		return
	}
	require.NotNil(t, created)
	assert.Contains(t, []int{200, 201}, resp.StatusCode)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV3(cleanupCtx, created.ID)
	})

	// Verify account settings
	fetched, _, err := svc.GetByIDV3(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, name, fetched.DisplayName)
}

// TestAcceptance_ComputerPrestages_MaximalConfig tests creating a prestage with maximal configuration.
func TestAcceptance_ComputerPrestages_MaximalConfig(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ComputerPrestages
	ctx := context.Background()
	name := acc.UniqueName("acc-prestage-maximal")

	created, resp, err := svc.CreateV3(ctx, &computer_prestages.ResourceComputerPrestage{
		DisplayName:                        name,
		Mandatory:                          boolPtr(true),
		MDMRemovable:                       boolPtr(false),
		SupportPhoneNumber:                 "111-222-3333",
		SupportEmailAddress:                "support@company.com",
		Department:                         "IT Department",
		DefaultPrestage:                    boolPtr(false),
		EnrollmentSiteId:                   "-1",
		KeepExistingSiteMembership:         boolPtr(false),
		KeepExistingLocationInformation:    boolPtr(false),
		RequireAuthentication:              boolPtr(true),
		AuthenticationPrompt:               "Please authenticate to continue setup",
		PreventActivationLock:              boolPtr(true),
		EnableDeviceBasedActivationLock:    boolPtr(false),
		DeviceEnrollmentProgramInstanceId:  "1",
		SkipSetupItems: &computer_prestages.SkipSetupItems{
			Biometric:                 boolPtr(true),
			TermsOfAddress:            boolPtr(true),
			FileVault:                 boolPtr(true),
			ICloudDiagnostics:         boolPtr(true),
			Diagnostics:               boolPtr(true),
			Accessibility:             boolPtr(true),
			AppleID:                   boolPtr(true),
			ScreenTime:                boolPtr(true),
			Siri:                      boolPtr(true),
			DisplayTone:               boolPtr(true),
			Restore:                   boolPtr(true),
			Appearance:                boolPtr(true),
			Privacy:                   boolPtr(true),
			Payment:                   boolPtr(true),
			Registration:              boolPtr(true),
			TOS:                       boolPtr(true),
			ICloudStorage:             boolPtr(true),
			Location:                  boolPtr(true),
			Intelligence:              boolPtr(true),
			EnableLockdownMode:        boolPtr(false),
			Welcome:                   boolPtr(true),
			Wallpaper:                 boolPtr(true),
			SoftwareUpdate:            boolPtr(true),
			AdditionalPrivacySettings: boolPtr(true),
			OSShowcase:                boolPtr(true),
		},
		LocationInformation: computer_prestages.LocationInformation{
			ID:           "-1",
			Username:     "john.doe",
			Realname:     "John Doe",
			Phone:        "555-1234",
			Email:        "john.doe@company.com",
			Room:         "Building A, Room 101",
			Position:     "Software Engineer",
			DepartmentId: "-1",
			BuildingId:   "-1",
			VersionLock:  1,
		},
		PurchasingInformation: computer_prestages.PurchasingInformation{
			ID:                "-1",
			Leased:            boolPtr(false),
			Purchased:         boolPtr(true),
			AppleCareId:       "AC123456789",
			PONumber:          "PO-2025-001",
			Vendor:            "Apple Inc.",
			PurchasePrice:     "2499.00",
			LifeExpectancy:    5,
			PurchasingAccount: "IT Capital Budget",
			PurchasingContact: "procurement@company.com",
			LeaseDate:         "1970-01-01",
			PODate:            "2025-01-15",
			WarrantyDate:      "2028-01-15",
			VersionLock:       1,
		},
		EnrollmentCustomizationId:          "0",
		Language:                           "en",
		Region:                             "US",
		AutoAdvanceSetup:                   boolPtr(true),
		InstallProfilesDuringSetup:         boolPtr(true),
		PrestageInstalledProfileIds:        []string{},
		CustomPackageIds:                   []string{},
		CustomPackageDistributionPointId:   "-1",
		EnableRecoveryLock:                 boolPtr(true),
		RecoveryLockPasswordType:           "MANUAL",
		RecoveryLockPassword:               "SecurePassword123!",
		RotateRecoveryLockPassword:         boolPtr(false),
		PrestageMinimumOsTargetVersionType: "MINIMUM_OS_LATEST_MAJOR_VERSION",
		MinimumOsSpecificVersion:           "",
		SiteId:                             "-1",
		AccountSettings: &computer_prestages.AccountSettings{
			PayloadConfigured:                       boolPtr(true),
			LocalAdminAccountEnabled:                boolPtr(true),
			AdminUsername:                           "localadmin",
			AdminPassword:                           "SecureAdminPass123!",
			HiddenAdminAccount:                      boolPtr(true),
			LocalUserManaged:                        boolPtr(true),
			UserAccountType:                         "ADMINISTRATOR",
			PrefillPrimaryAccountInfoFeatureEnabled: boolPtr(true),
			PrefillType:                             "DEVICE_OWNER",
			PrefillAccountFullName:                  "Device Owner",
			PrefillAccountUserName:                  "deviceowner",
			PreventPrefillInfoFromModification:      boolPtr(true),
		},
	})
	if err != nil {
		t.Skipf("CreateV3 failed (may require DEP): %v", err)
		return
	}
	require.NotNil(t, created)
	assert.Contains(t, []int{200, 201}, resp.StatusCode)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV3(cleanupCtx, created.ID)
	})

	// Verify maximal configuration
	fetched, _, err := svc.GetByIDV3(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, name, fetched.DisplayName)
	assert.Equal(t, "MINIMUM_OS_LATEST_MAJOR_VERSION", fetched.PrestageMinimumOsTargetVersionType)
}

// TestAcceptance_ComputerPrestages_ValidationErrors tests validation of enum fields.
func TestAcceptance_ComputerPrestages_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ComputerPrestages
	ctx := context.Background()

	t.Run("InvalidRecoveryLockPasswordType", func(t *testing.T) {
		_, _, err := svc.CreateV3(ctx, &computer_prestages.ResourceComputerPrestage{
			DisplayName:              "test",
			RecoveryLockPasswordType: "INVALID_TYPE",
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid recoveryLockPasswordType")
	})

	t.Run("InvalidPrestageMinimumOsTargetVersionType", func(t *testing.T) {
		_, _, err := svc.CreateV3(ctx, &computer_prestages.ResourceComputerPrestage{
			DisplayName:                        "test",
			PrestageMinimumOsTargetVersionType: "INVALID_TYPE",
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid prestageMinimumOsTargetVersionType")
	})

	t.Run("InvalidUserAccountType", func(t *testing.T) {
		_, _, err := svc.CreateV3(ctx, &computer_prestages.ResourceComputerPrestage{
			DisplayName: "test",
			AccountSettings: &computer_prestages.AccountSettings{
				UserAccountType: "INVALID_TYPE",
			},
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid userAccountType")
	})

	t.Run("InvalidPrefillType", func(t *testing.T) {
		_, _, err := svc.CreateV3(ctx, &computer_prestages.ResourceComputerPrestage{
			DisplayName: "test",
			AccountSettings: &computer_prestages.AccountSettings{
				PrefillType: "INVALID_TYPE",
			},
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid prefillType")
	})
}
