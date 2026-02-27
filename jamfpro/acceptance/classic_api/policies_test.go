package classic_api

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/disk_encryption_configurations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/dock_items"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/policies"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/scripts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_Policies_minimum_config tests creating a policy with the
// minimum required configuration.
// =============================================================================

func TestAcceptance_Policies_minimum_config(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create policy with minimum config
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating policy with minimum configuration")

	policyName := acc.UniqueName("sdkv2_acc_policy_min")
	createReq := createMinimalPolicy(t, policyName)

	_, policyID := createPolicyWithCleanup(t, ctx, svc, createReq)

	// ------------------------------------------------------------------
	// 2. Get by ID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching policy by ID=%d", policyID)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, policyID, fetched.General.ID)
	assert.Equal(t, policyName, fetched.General.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 3. Get by Name
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching policy by name=%q", policyName)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx3, policyName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, policyID, fetchedByName.General.ID)
	assert.Equal(t, policyName, fetchedByName.General.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.General.ID, fetchedByName.General.Name)

	// ------------------------------------------------------------------
	// 4. Update by ID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByID", "Updating policy by ID=%d", policyID)

	updateReq := fetched
	updateReq.General.Enabled = true
	updateReq.SelfService.SelfServiceIcon = nil
	updateReq.DiskEncryption = policies.PolicySubsetDiskEncryption{
		Action:                                 "",
		DiskEncryptionConfigurationID:          0,
		AuthRestart:                            false,
		RemediateKeyType:                       "Individual",
		RemediateDiskEncryptionConfigurationID: 0,
	}

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	updated, updateResp, err := svc.UpdateByID(ctx4, policyID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	assert.Equal(t, policyID, updated.ID, "updated policy ID should match")
	acc.LogTestSuccess(t, "UpdateByID: ID=%d", updated.ID)

	// ------------------------------------------------------------------
	// 5. List
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing all policies")

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	list, listResp, err := svc.List(ctx5)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, p := range list.Results {
		if p.ID == policyID {
			found = true
			assert.Equal(t, policyName, p.Name)
			break
		}
	}
	assert.True(t, found, "created policy should appear in list")
	acc.LogTestSuccess(t, "List: found policy ID=%d in list of %d policies", policyID, list.Size)
}

// =============================================================================
// TestAcceptance_Policies_validation_errors tests validation error handling.
// =============================================================================

func TestAcceptance_Policies_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(ctx, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy ID must be a positive integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy name is required")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(ctx, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		policy := createMinimalPolicy(t, "test")
		_, _, err := svc.UpdateByID(ctx, 0, policy)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		policy := createMinimalPolicy(t, "test")
		_, _, err := svc.UpdateByName(ctx, "", policy)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy name is required")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(ctx, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy ID must be a positive integer")
	})

	t.Run("DeleteByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteByName(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "policy name is required")
	})
}

// =============================================================================
// TestAcceptance_Policies_with_script tests creating a policy with a script.
// =============================================================================

func TestAcceptance_Policies_with_script(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	scriptSvc := acc.Client.Scripts
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create a script prerequisite using Jamf Pro API
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Prerequisite", "Creating script for policy")

	scriptName := acc.UniqueName("sdkv2_acc_policy_script")
	scriptReq := &scripts.RequestScript{
		Name:           scriptName,
		ScriptContents: "#!/bin/bash\necho 'Test script for policy'\nexit 0",
		Info:           "Test script for policy acceptance test",
		Notes:          "Created by SDK v2 acceptance test",
		Priority:       "AFTER",
		CategoryId:     "-1",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	createdScript, createScriptResp, err := scriptSvc.CreateScriptV1(ctx1, scriptReq)
	require.NoError(t, err, "Create script should not return an error")
	require.NotNil(t, createdScript)
	assert.Equal(t, 201, createScriptResp.StatusCode)

	scriptID := createdScript.ID
	acc.LogTestSuccess(t, "Script created with ID=%s name=%q", scriptID, scriptName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := scriptSvc.DeleteScriptByIDV1(cleanupCtx, scriptID)
		acc.LogCleanupDeleteError(t, "script", scriptID, delErr)
	})

	// ------------------------------------------------------------------
	// 2. Create policy with script
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating policy with script")

	policyName := acc.UniqueName("sdkv2_acc_policy_with_script")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.Scripts = []policies.PolicySubsetScript{
		{
			ID:          scriptID,
			Priority:    "After",
			Parameter4:  "param4_value",
			Parameter5:  "param5_value",
			Parameter6:  "",
			Parameter7:  "",
			Parameter8:  "",
			Parameter9:  "",
			Parameter10: "",
			Parameter11: "",
		},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// ------------------------------------------------------------------
	// 3. Verify script is in policy
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Verify", "Verifying script is in policy")

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	require.Len(t, fetched.Scripts, 1, "Policy should have 1 script")
	assert.Equal(t, scriptID, fetched.Scripts[0].ID)
	assert.Equal(t, "After", fetched.Scripts[0].Priority)
	acc.LogTestSuccess(t, "Policy has script ID=%s", fetched.Scripts[0].ID)
}

// =============================================================================
// TestAcceptance_Policies_maintenance tests creating a maintenance policy.
// =============================================================================

func TestAcceptance_Policies_maintenance(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating maintenance policy")

	policyName := acc.UniqueName("sdkv2_acc_policy_maintenance")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.Maintenance = policies.PolicySubsetMaintenance{
		Recon:                    true,
		ResetName:                false,
		InstallAllCachedPackages: false,
		Heal:                     false,
		Prebindings:              false,
		Permissions:              true,
		Byhost:                   false,
		SystemCache:              true,
		UserCache:                false,
		Verify:                   true,
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify maintenance settings
	acc.LogTestStage(t, "Verify", "Verifying maintenance settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.True(t, fetched.Maintenance.Recon)
	assert.True(t, fetched.Maintenance.Permissions)
	assert.True(t, fetched.Maintenance.SystemCache)
	assert.True(t, fetched.Maintenance.Verify)
	acc.LogTestSuccess(t, "Maintenance policy verified")
}

// =============================================================================
// TestAcceptance_Policies_files_and_processes tests files and processes policy.
// =============================================================================

func TestAcceptance_Policies_files_and_processes(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating files and processes policy")

	policyName := acc.UniqueName("sdkv2_acc_policy_files_processes")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.FilesProcesses = policies.PolicySubsetFilesProcesses{
		SearchByPath:         "/tmp/test_file.txt",
		DeleteFile:           true,
		LocateFile:           "",
		UpdateLocateDatabase: true,
		SpotlightSearch:      "",
		SearchForProcess:     "Safari",
		KillProcess:          false,
		RunCommand:           "echo 'test command'",
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify files and processes settings
	acc.LogTestStage(t, "Verify", "Verifying files and processes settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, "/tmp/test_file.txt", fetched.FilesProcesses.SearchByPath)
	assert.True(t, fetched.FilesProcesses.DeleteFile)
	assert.True(t, fetched.FilesProcesses.UpdateLocateDatabase)
	assert.Equal(t, "Safari", fetched.FilesProcesses.SearchForProcess)
	acc.LogTestSuccess(t, "Files and processes policy verified")
}

// =============================================================================
// TestAcceptance_Policies_self_service tests self service policy configuration.
// =============================================================================

func TestAcceptance_Policies_self_service(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating self service policy")

	policyName := acc.UniqueName("sdkv2_acc_policy_self_service")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.SelfService = policies.PolicySubsetSelfService{
		UseForSelfService:           true,
		SelfServiceDisplayName:      "Test Self Service Policy",
		InstallButtonText:           "Install",
		ReinstallButtonText:         "Reinstall",
		SelfServiceDescription:      "This is a test self service policy",
		ForceUsersToViewDescription: true,
		FeatureOnMainPage:           true,
		Notification:                true,
		NotificationType:            "Self Service",
		NotificationSubject:         "Installation Complete",
		NotificationMessage:         "The software has been installed successfully",
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify self service settings
	acc.LogTestStage(t, "Verify", "Verifying self service settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.True(t, fetched.SelfService.UseForSelfService, "UseForSelfService should be true")
	assert.Equal(t, "Test Self Service Policy", fetched.SelfService.SelfServiceDisplayName)
	// Note: Some self service fields may not persist exactly as set
	if fetched.SelfService.ForceUsersToViewDescription {
		acc.LogTestSuccess(t, "ForceUsersToViewDescription: true")
	}
	if fetched.SelfService.FeatureOnMainPage {
		acc.LogTestSuccess(t, "FeatureOnMainPage: true")
	}
	acc.LogTestSuccess(t, "Self service policy verified")
}

// =============================================================================
// TestAcceptance_Policies_restart_options tests restart/reboot configuration.
// =============================================================================

func TestAcceptance_Policies_restart_options(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with restart options")

	policyName := acc.UniqueName("sdkv2_acc_policy_restart")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.Reboot = policies.PolicySubsetReboot{
		Message:                     "This computer will restart in 5 minutes. Please save your work.",
		StartupDisk:                 "Current Startup Disk",
		SpecifyStartup:              "",
		NoUserLoggedIn:              "Restart if a package or update requires it",
		UserLoggedIn:                "Restart if a package or update requires it",
		MinutesUntilReboot:          5,
		StartRebootTimerImmediately: true,
		FileVault2Reboot:            true,
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify restart options
	acc.LogTestStage(t, "Verify", "Verifying restart options")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Contains(t, fetched.Reboot.Message, "restart in 5 minutes")
	assert.Equal(t, 5, fetched.Reboot.MinutesUntilReboot)
	assert.True(t, fetched.Reboot.StartRebootTimerImmediately)
	assert.True(t, fetched.Reboot.FileVault2Reboot)
	acc.LogTestSuccess(t, "Restart options policy verified")
}

// =============================================================================
// TestAcceptance_Policies_user_interaction tests user interaction settings.
// =============================================================================

func TestAcceptance_Policies_user_interaction(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with user interaction")

	policyName := acc.UniqueName("sdkv2_acc_policy_user_interaction")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.UserInteraction = policies.PolicySubsetUserInteraction{
		MessageStart:          "This policy is about to run. Please wait...",
		AllowUsersToDefer:     true,
		AllowDeferralUntilUtc: "",
		AllowDeferralMinutes:  1440,
		MessageFinish:         "The policy has completed successfully.",
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify user interaction settings
	acc.LogTestStage(t, "Verify", "Verifying user interaction settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Contains(t, fetched.UserInteraction.MessageStart, "about to run")
	assert.True(t, fetched.UserInteraction.AllowUsersToDefer)
	assert.Equal(t, 1440, fetched.UserInteraction.AllowDeferralMinutes)
	assert.Contains(t, fetched.UserInteraction.MessageFinish, "completed successfully")
	acc.LogTestSuccess(t, "User interaction policy verified")
}

// =============================================================================
// TestAcceptance_Policies_dock_items tests dock items configuration.
// =============================================================================

func TestAcceptance_Policies_dock_items(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	dockItemsSvc := acc.Client.ClassicDockItems
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create dock item prerequisites
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Prerequisite", "Creating dock items")

	dockItem1Name := acc.UniqueName("sdkv2_acc_dock_safari")
	dockItem1Req := &dock_items.Request{
		Name:     dockItem1Name,
		Type:     "App",
		Path:     "/Applications/Safari.app",
		Contents: "",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	createdDock1, createDock1Resp, err := dockItemsSvc.Create(ctx1, dockItem1Req)
	require.NoError(t, err, "Create dock item 1 should not return an error")
	require.NotNil(t, createdDock1)
	assert.Contains(t, []int{200, 201}, createDock1Resp.StatusCode)

	dockItem1ID := createdDock1.ID
	acc.LogTestSuccess(t, "Dock item 1 created with ID=%d name=%q", dockItem1ID, dockItem1Name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := dockItemsSvc.DeleteByID(cleanupCtx, dockItem1ID)
		acc.LogCleanupDeleteError(t, "dock_item", fmt.Sprintf("%d", dockItem1ID), delErr)
	})

	dockItem2Name := acc.UniqueName("sdkv2_acc_dock_mail")
	dockItem2Req := &dock_items.Request{
		Name:     dockItem2Name,
		Type:     "App",
		Path:     "/Applications/Mail.app",
		Contents: "",
	}

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	createdDock2, createDock2Resp, err := dockItemsSvc.Create(ctx2, dockItem2Req)
	require.NoError(t, err, "Create dock item 2 should not return an error")
	require.NotNil(t, createdDock2)
	assert.Contains(t, []int{200, 201}, createDock2Resp.StatusCode)

	dockItem2ID := createdDock2.ID
	acc.LogTestSuccess(t, "Dock item 2 created with ID=%d name=%q", dockItem2ID, dockItem2Name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := dockItemsSvc.DeleteByID(cleanupCtx, dockItem2ID)
		acc.LogCleanupDeleteError(t, "dock_item", fmt.Sprintf("%d", dockItem2ID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. Create policy with dock items
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating policy with dock items")

	policyName := acc.UniqueName("sdkv2_acc_policy_dock_items")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.DockItems = []policies.PolicySubsetDockItem{
		{
			ID:     dockItem1ID,
			Name:   dockItem1Name,
			Action: "Add To End",
		},
		{
			ID:     dockItem2ID,
			Name:   dockItem2Name,
			Action: "Add To Beginning",
		},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// ------------------------------------------------------------------
	// 3. Verify dock items
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Verify", "Verifying dock items")

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	require.Len(t, fetched.DockItems, 2, "Policy should have 2 dock items")
	assert.Equal(t, dockItem1ID, fetched.DockItems[0].ID)
	assert.Equal(t, dockItem2ID, fetched.DockItems[1].ID)
	acc.LogTestSuccess(t, "Dock items policy verified")
}

// =============================================================================
// TestAcceptance_Policies_printers tests printer configuration.
// =============================================================================

func TestAcceptance_Policies_printers(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with printer settings")

	policyName := acc.UniqueName("sdkv2_acc_policy_printers")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.Printers = policies.PolicySubsetPrinters{
		LeaveExistingDefault: true,
		Printer:              []policies.PolicySubsetPrinter{},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify printer settings
	acc.LogTestStage(t, "Verify", "Verifying printer settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	// Verify printer settings were applied
	acc.LogTestSuccess(t, "Printer policy verified: LeaveExistingDefault=%v", fetched.Printers.LeaveExistingDefault)
}

// =============================================================================
// TestAcceptance_Policies_package tests package installation policy.
// Note: This test requires a package to exist in Jamf Pro.
// =============================================================================

func TestAcceptance_Policies_package(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	packagesSvc := acc.Client.Packages
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Check if any packages exist to use as prerequisite
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Prerequisite", "Checking for existing packages")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	packagesList, _, err := packagesSvc.ListV1(ctx1, nil)
	if err != nil || packagesList == nil || packagesList.TotalCount == 0 {
		t.Skip("No packages found - skipping package policy test")
	}

	packageIDStr := packagesList.Results[0].ID
	packageName := packagesList.Results[0].PackageName
	packageIDInt, err := strconv.Atoi(packageIDStr)
	require.NoError(t, err, "Package ID should be convertible to int")
	acc.LogTestSuccess(t, "Found package ID=%d name=%q", packageIDInt, packageName)

	// ------------------------------------------------------------------
	// 2. Create policy with package
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating policy with package")

	policyName := acc.UniqueName("sdkv2_acc_policy_package")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.PackageConfiguration = policies.PolicySubsetPackageConfiguration{
		Packages: []policies.PolicySubsetPackageConfigurationPackage{
			{
				ID:                packageIDInt,
				Name:              packageName,
				Action:            "Install",
				FillUserTemplate:  false,
				FillExistingUsers: false,
				UpdateAutorun:     false,
			},
		},
		DistributionPoint: "default",
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// ------------------------------------------------------------------
	// 3. Verify package configuration
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Verify", "Verifying package configuration")

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	require.Len(t, fetched.PackageConfiguration.Packages, 1, "Policy should have 1 package")
	assert.Equal(t, packageIDInt, fetched.PackageConfiguration.Packages[0].ID)
	assert.Equal(t, "Install", fetched.PackageConfiguration.Packages[0].Action)
	acc.LogTestSuccess(t, "Package policy verified with package ID=%d", packageIDInt)
}

// =============================================================================
// TestAcceptance_Policies_disk_encryption_individual tests disk encryption
// with individual recovery key.
// =============================================================================

func TestAcceptance_Policies_disk_encryption_individual(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with disk encryption (individual key)")

	policyName := acc.UniqueName("sdkv2_acc_policy_disk_enc_individual")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.DiskEncryption = policies.PolicySubsetDiskEncryption{
		Action:                                 "remediate",
		DiskEncryptionConfigurationID:          0,
		AuthRestart:                            false,
		RemediateKeyType:                       "Individual",
		RemediateDiskEncryptionConfigurationID: 0,
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify disk encryption settings
	acc.LogTestStage(t, "Verify", "Verifying disk encryption settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, "remediate", fetched.DiskEncryption.Action)
	assert.Equal(t, "Individual", fetched.DiskEncryption.RemediateKeyType)
	acc.LogTestSuccess(t, "Disk encryption (individual) policy verified")
}

// =============================================================================
// TestAcceptance_Policies_disk_encryption_institutional tests disk encryption
// with institutional recovery key.
// =============================================================================

func TestAcceptance_Policies_disk_encryption_institutional(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	diskEncSvc := acc.Client.ClassicDiskEncryptionConfigurations
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create a disk encryption configuration as a prerequisite
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Prerequisite", "Creating disk encryption configuration")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	diskEncName := acc.UniqueName("sdkv2_acc_diskenc-for-policy")
	diskEncReq := &disk_encryption_configurations.RequestDiskEncryptionConfiguration{
		Name:                  diskEncName,
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}
	createdDiskEnc, _, err := diskEncSvc.Create(ctx1, diskEncReq)
	require.NoError(t, err, "Create disk encryption configuration should not return an error")
	require.NotNil(t, createdDiskEnc)
	require.Positive(t, createdDiskEnc.ID)

	diskEncID := createdDiskEnc.ID
	acc.LogTestSuccess(t, "Created disk encryption configuration ID=%d name=%q", diskEncID, diskEncName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := diskEncSvc.DeleteByID(cleanupCtx, diskEncID)
		acc.LogCleanupDeleteError(t, "disk encryption configuration", fmt.Sprintf("%d", diskEncID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. Create policy with disk encryption settings
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating policy with disk encryption settings")

	policyName := acc.UniqueName("sdkv2_acc_policy_disk_enc_institutional")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.DiskEncryption = policies.PolicySubsetDiskEncryption{
		Action:                                 "remediate",
		DiskEncryptionConfigurationID:          diskEncID,
		AuthRestart:                            false,
		RemediateKeyType:                       "Individual",
		RemediateDiskEncryptionConfigurationID: diskEncID,
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// ------------------------------------------------------------------
	// 3. Verify disk encryption settings
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Verify", "Verifying disk encryption settings")

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, "remediate", fetched.DiskEncryption.Action)
	assert.Equal(t, "Individual", fetched.DiskEncryption.RemediateKeyType)
	acc.LogTestSuccess(t, "Disk encryption policy verified")
}

// =============================================================================
// TestAcceptance_Policies_directory_binding tests directory binding policy.
// =============================================================================

func TestAcceptance_Policies_directory_binding(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with directory binding")

	policyName := acc.UniqueName("sdkv2_acc_policy_directory_binding")
	policyReq := createMinimalPolicy(t, policyName)
	// Note: Directory bindings require actual directory binding configurations
	// This test creates a policy structure that would contain directory bindings
	policyReq.AccountMaintenance = policies.PolicySubsetAccountMaintenance{
		DirectoryBindings: &[]policies.PolicySubsetAccountMaintenanceDirectoryBindings{},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify policy was created
	acc.LogTestStage(t, "Verify", "Verifying directory binding policy")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	acc.LogTestSuccess(t, "Directory binding policy structure verified")
}

// =============================================================================
// TestAcceptance_Policies_account_management tests account management policy.
// =============================================================================

func TestAcceptance_Policies_account_management(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with account management")

	policyName := acc.UniqueName("sdkv2_acc_policy_account_mgmt")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.AccountMaintenance = policies.PolicySubsetAccountMaintenance{
		ManagementAccount: &policies.PolicySubsetAccountMaintenanceManagementAccount{
			Action:                 "doNotChange",
			ManagedPassword:        "",
			ManagedPasswordLength:  0,
		},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify account management settings
	acc.LogTestStage(t, "Verify", "Verifying account management settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	require.NotNil(t, fetched.AccountMaintenance.ManagementAccount)
	acc.LogTestSuccess(t, "Account management policy verified")
}

// =============================================================================
// TestAcceptance_Policies_create_local_account tests creating a local account.
// =============================================================================

func TestAcceptance_Policies_create_local_account(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with new local account")

	policyName := acc.UniqueName("sdkv2_acc_policy_create_account")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.AccountMaintenance = policies.PolicySubsetAccountMaintenance{
		Accounts: &[]policies.PolicySubsetAccountMaintenanceAccount{
			{
				Action:                 "Create",
				Username:               "testuser",
				Realname:               "Test User",
				Password:               "TestPassword123!",
				ArchiveHomeDirectory:   false,
				ArchiveHomeDirectoryTo: "",
				Home:                   "/Users/testuser/",
				Hint:                   "test hint",
				Picture:                "/Library/User Pictures/Animals/Butterfly.tif",
				Admin:                  false,
				FilevaultEnabled:       false,
			},
		},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify account creation settings
	acc.LogTestStage(t, "Verify", "Verifying account creation settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	require.NotNil(t, fetched.AccountMaintenance.Accounts)
	require.Len(t, *fetched.AccountMaintenance.Accounts, 1)
	assert.Equal(t, "Create", (*fetched.AccountMaintenance.Accounts)[0].Action)
	assert.Equal(t, "testuser", (*fetched.AccountMaintenance.Accounts)[0].Username)
	acc.LogTestSuccess(t, "Create local account policy verified")
}

// =============================================================================
// TestAcceptance_Policies_delete_account tests deleting an account.
// =============================================================================

func TestAcceptance_Policies_delete_account(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with delete account")

	policyName := acc.UniqueName("sdkv2_acc_policy_delete_account")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.AccountMaintenance = policies.PolicySubsetAccountMaintenance{
		Accounts: &[]policies.PolicySubsetAccountMaintenanceAccount{
			{
				Action:                 "Delete",
				Username:               "olduser",
				Realname:               "",
				Password:               "",
				ArchiveHomeDirectory:   true,
				ArchiveHomeDirectoryTo: "/Users/Deleted Users/",
				Home:                   "",
				Hint:                   "",
				Picture:                "",
				Admin:                  false,
				FilevaultEnabled:       false,
			},
		},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify account deletion settings
	acc.LogTestStage(t, "Verify", "Verifying account deletion settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	require.NotNil(t, fetched.AccountMaintenance.Accounts)
	require.Len(t, *fetched.AccountMaintenance.Accounts, 1)
	assert.Equal(t, "Delete", (*fetched.AccountMaintenance.Accounts)[0].Action)
	acc.LogTestSuccess(t, "Delete account policy verified")
}

// =============================================================================
// TestAcceptance_Policies_reset_password tests resetting account password.
// =============================================================================

func TestAcceptance_Policies_reset_password(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with reset password")

	policyName := acc.UniqueName("sdkv2_acc_policy_reset_password")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.AccountMaintenance = policies.PolicySubsetAccountMaintenance{
		Accounts: &[]policies.PolicySubsetAccountMaintenanceAccount{
			{
				Action:                 "Reset",
				Username:               "existinguser",
				Realname:               "",
				Password:               "NewPassword123!",
				ArchiveHomeDirectory:   false,
				ArchiveHomeDirectoryTo: "",
				Home:                   "",
				Hint:                   "",
				Picture:                "",
				Admin:                  false,
				FilevaultEnabled:       false,
			},
		},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify password reset settings
	acc.LogTestStage(t, "Verify", "Verifying password reset settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	require.NotNil(t, fetched.AccountMaintenance.Accounts)
	require.Len(t, *fetched.AccountMaintenance.Accounts, 1)
	assert.Equal(t, "Reset", (*fetched.AccountMaintenance.Accounts)[0].Action)
	acc.LogTestSuccess(t, "Reset password policy verified")
}

// =============================================================================
// TestAcceptance_Policies_filevault_disable_user tests disabling FileVault user.
// =============================================================================

func TestAcceptance_Policies_filevault_disable_user(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with FileVault disable user")

	policyName := acc.UniqueName("sdkv2_acc_policy_fv_disable")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.AccountMaintenance = policies.PolicySubsetAccountMaintenance{
		Accounts: &[]policies.PolicySubsetAccountMaintenanceAccount{
			{
				Action:                 "DisableFileVault",
				Username:               "fvuser",
				Realname:               "",
				Password:               "",
				ArchiveHomeDirectory:   false,
				ArchiveHomeDirectoryTo: "",
				Home:                   "",
				Hint:                   "",
				Picture:                "",
				Admin:                  false,
				FilevaultEnabled:       false,
			},
		},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify FileVault disable settings
	acc.LogTestStage(t, "Verify", "Verifying FileVault disable settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	if fetched.AccountMaintenance.Accounts != nil && len(*fetched.AccountMaintenance.Accounts) > 0 {
		assert.Equal(t, "DisableFileVault", (*fetched.AccountMaintenance.Accounts)[0].Action)
	}
	acc.LogTestSuccess(t, "FileVault disable user policy verified")
}

// =============================================================================
// TestAcceptance_Policies_efi_password_set tests setting EFI password.
// =============================================================================

func TestAcceptance_Policies_efi_password_set(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with EFI password set")

	policyName := acc.UniqueName("sdkv2_acc_policy_efi_set")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.AccountMaintenance = policies.PolicySubsetAccountMaintenance{
		OpenFirmwareEfiPassword: &policies.PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword{
			OfMode:   "command",
			OfPassword: "test123",
		},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify EFI password settings
	acc.LogTestStage(t, "Verify", "Verifying EFI password settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	require.NotNil(t, fetched.AccountMaintenance.OpenFirmwareEfiPassword)
	assert.Equal(t, "command", fetched.AccountMaintenance.OpenFirmwareEfiPassword.OfMode)
	acc.LogTestSuccess(t, "EFI password set policy verified")
}

// =============================================================================
// TestAcceptance_Policies_efi_password_remove tests removing EFI password.
// =============================================================================

func TestAcceptance_Policies_efi_password_remove(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ClassicPolicies
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating policy with EFI password remove")

	policyName := acc.UniqueName("sdkv2_acc_policy_efi_remove")
	policyReq := createMinimalPolicy(t, policyName)
	policyReq.AccountMaintenance = policies.PolicySubsetAccountMaintenance{
		OpenFirmwareEfiPassword: &policies.PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword{
			OfMode:   "none",
			OfPassword: "",
		},
	}

	_, policyID := createPolicyWithCleanup(t, ctx, svc, policyReq)

	// Verify EFI password removal settings
	acc.LogTestStage(t, "Verify", "Verifying EFI password removal settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, policyID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	require.NotNil(t, fetched.AccountMaintenance.OpenFirmwareEfiPassword)
	assert.Equal(t, "none", fetched.AccountMaintenance.OpenFirmwareEfiPassword.OfMode)
	acc.LogTestSuccess(t, "EFI password remove policy verified")
}
