package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/policies"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"github.com/stretchr/testify/require"
)

// createMinimalPolicy creates a minimal valid policy for testing.
// This is the minimum configuration required by the Jamf Pro API.
func createMinimalPolicy(t *testing.T, name string) *policies.ResourcePolicy {
	return &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:                       name,
			Enabled:                    false,
			TriggerCheckin:             false,
			TriggerEnrollmentComplete:  false,
			TriggerLogin:               false,
			TriggerLogout:              false,
			TriggerNetworkStateChanged: false,
			TriggerStartup:             false,
			TriggerOther:               "EVENT",
			Frequency:                  "Once per computer",
			RetryEvent:                 "none",
			RetryAttempts:              -1,
			NotifyOnEachFailedRetry:    false,
			LocationUserOnly:           false,
			TargetDrive:                "/",
			Offline:                    false,
			Category: &shared.SharedResourceCategory{
				ID:   -1,
				Name: "No category assigned",
			},
			NetworkLimitations: &policies.PolicySubsetGeneralNetworkLimitations{
				MinimumNetworkConnection: "No Minimum",
				AnyIPAddress:             false,
				NetworkSegments:          "",
			},
			NetworkRequirements: "Any",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: policies.PolicySubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
		},
		SelfService: policies.PolicySubsetSelfService{
			UseForSelfService:           true,
			InstallButtonText:           "Install",
			ReinstallButtonText:         "Reinstall",
			ForceUsersToViewDescription: false,
			FeatureOnMainPage:           false,
			Notification:                false,
		},
		PackageConfiguration: policies.PolicySubsetPackageConfiguration{
			Packages:              []policies.PolicySubsetPackageConfigurationPackage{},
			DistributionPoint:     "",
		},
		Scripts: []policies.PolicySubsetScript{},
		Printers: policies.PolicySubsetPrinters{
			LeaveExistingDefault: false,
		},
		DockItems:              []policies.PolicySubsetDockItem{},
		AccountMaintenance: policies.PolicySubsetAccountMaintenance{},
		Maintenance: policies.PolicySubsetMaintenance{
			Recon:                    false,
			ResetName:                false,
			InstallAllCachedPackages: false,
			Heal:                     false,
			Prebindings:              false,
			Permissions:              false,
			Byhost:                   false,
			SystemCache:              false,
			UserCache:                false,
			Verify:                   false,
		},
		FilesProcesses: policies.PolicySubsetFilesProcesses{
			SearchByPath:         "",
			DeleteFile:           false,
			LocateFile:           "",
			UpdateLocateDatabase: false,
			SpotlightSearch:      "",
			SearchForProcess:     "",
			KillProcess:          false,
			RunCommand:           "",
		},
		UserInteraction: policies.PolicySubsetUserInteraction{
			MessageStart:          "",
			AllowUsersToDefer:     false,
			AllowDeferralUntilUtc: "",
			AllowDeferralMinutes:  0,
			MessageFinish:         "",
		},
		DiskEncryption: policies.PolicySubsetDiskEncryption{
			Action:                        "",
			DiskEncryptionConfigurationID: 0,
			AuthRestart:                   false,
			RemediateKeyType:              "Individual",
			RemediateDiskEncryptionConfigurationID: 0,
		},
		Reboot: policies.PolicySubsetReboot{
			Message:                      "",
			StartupDisk:                  "Current Startup Disk",
			SpecifyStartup:               "",
			NoUserLoggedIn:               "Do not restart",
			UserLoggedIn:                 "Do not restart",
			MinutesUntilReboot:           5,
			StartRebootTimerImmediately:  false,
			FileVault2Reboot:             false,
		},
	}
}

// createPolicyWithCleanup creates a policy and registers cleanup.
// Returns the created policy response and the policy ID.
func createPolicyWithCleanup(t *testing.T, ctx context.Context, svc *policies.Policies, req *policies.ResourcePolicy) (*policies.CreateUpdateResponse, int) {
	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, req)
	require.NoError(t, err, "Create policy should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	require.Positive(t, created.ID, "created policy ID should be positive")

	policyID := created.ID
	acc.LogTestSuccess(t, "Policy created with ID=%d name=%q", policyID, req.General.Name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, policyID)
		acc.LogCleanupDeleteError(t, "policy", fmt.Sprintf("%d", policyID), delErr)
	})

	return created, policyID
}
