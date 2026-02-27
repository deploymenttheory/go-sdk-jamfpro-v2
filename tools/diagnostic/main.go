package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/policies"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

func main() {
	authConfig := client.AuthConfigFromEnv()

	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Test: Create Policy with minimum config
	fmt.Println("=== Test: Policy Create with Minimum Config ===")
	policyReq := &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:                       fmt.Sprintf("diagnostic-policy-%d", time.Now().Unix()),
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

	created, resp, err := jamfClient.ClassicPolicies.Create(ctx, policyReq)
	if err != nil {
		fmt.Printf("Policy Create Error: %v\n", err)
	}
	if resp != nil {
		fmt.Printf("Policy Create Status: %d\n", resp.StatusCode)
		fmt.Printf("Policy Create Response Body:\n%s\n\n", string(resp.Body))
	}
	
	if created != nil && created.ID > 0 {
		policyID := created.ID
		fmt.Printf("Created Policy ID: %d\n\n", policyID)
		
		// Test GET
		fmt.Println("=== Test: Policy GET ===")
		time.Sleep(1 * time.Second)
		fetched, resp2, err := jamfClient.ClassicPolicies.GetByID(ctx, policyID)
		if err != nil {
			fmt.Printf("Policy GET Error: %v\n", err)
		}
		if resp2 != nil {
			fmt.Printf("Policy GET Status: %d\n", resp2.StatusCode)
			fmt.Printf("Policy GET Response Body (first 2000 chars):\n%s\n\n", string(resp2.Body[:min(2000, len(resp2.Body))]))
		}
		
		// Test UPDATE
		if fetched != nil {
			fmt.Println("=== Test: Policy UPDATE ===")
			fmt.Printf("DiskEncryption.RemediateKeyType before update: %q\n", fetched.DiskEncryption.RemediateKeyType)
			fmt.Printf("DiskEncryption.Action before update: %q\n", fetched.DiskEncryption.Action)
			
			fetched.General.Enabled = true
			fetched.SelfService.SelfServiceIcon = nil
			fetched.DiskEncryption = policies.PolicySubsetDiskEncryption{
				Action:                        "",
				DiskEncryptionConfigurationID: 0,
				AuthRestart:                   false,
				RemediateKeyType:              "Individual",
				RemediateDiskEncryptionConfigurationID: 0,
			}
			
			_, resp3, err := jamfClient.ClassicPolicies.UpdateByID(ctx, policyID, fetched)
			if err != nil {
				fmt.Printf("Policy UPDATE Error: %v\n", err)
			}
			if resp3 != nil {
				fmt.Printf("Policy UPDATE Status: %d\n", resp3.StatusCode)
				fmt.Printf("Policy UPDATE Response Body:\n%s\n\n", string(resp3.Body))
			}
		}
		
		// Clean up
		jamfClient.ClassicPolicies.DeleteByID(ctx, policyID)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
