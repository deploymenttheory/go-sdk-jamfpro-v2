// Package main demonstrates CreateV3 - creates a new mobile device prestage.
//
// Run with: go run ./examples/jamf_pro_api/mobile_device_prestages/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_prestages"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Helper function to create bool pointers
	boolPtr := func(b bool) *bool { return &b }
	intPtr := func(i int) *int { return &i }

	req := &mobile_device_prestages.ResourceMobileDevicePrestage{
		DisplayName:                        "go-sdk-v2-MobileDevicePrestage",
		Mandatory:                          boolPtr(true),
		MdmRemovable:                       boolPtr(false),
		SupportPhoneNumber:                 "555-0100",
		SupportEmailAddress:                "support@example.com",
		Department:                         "IT Department",
		DefaultPrestage:                    boolPtr(false),
		EnrollmentSiteID:                   "-1",
		KeepExistingSiteMembership:         boolPtr(false),
		KeepExistingLocationInformation:    boolPtr(false),
		RequireAuthentication:              boolPtr(true),
		AuthenticationPrompt:               "Please authenticate to enroll your device",
		PreventActivationLock:              boolPtr(true),
		EnableDeviceBasedActivationLock:    boolPtr(false),
		DeviceEnrollmentProgramInstanceID:  "1",
		AutoAdvanceSetup:                   boolPtr(true),
		AllowPairing:                       boolPtr(true),
		MultiUser:                          boolPtr(false),
		Supervised:                         boolPtr(true),
		MaximumSharedAccounts:              1,
		ConfigureDeviceBeforeSetupAssistant: boolPtr(false),
		SendTimezone:                       boolPtr(true),
		Timezone:                           "America/Chicago",
		StorageQuotaSizeMegabytes:          4096,
		UseStorageQuotaSize:                boolPtr(false),
		TemporarySessionOnly:               boolPtr(false),
		EnforceTemporarySessionTimeout:     boolPtr(false),
		TemporarySessionTimeout:            intPtr(30),
		EnforceUserSessionTimeout:          boolPtr(false),
		UserSessionTimeout:                 intPtr(60),
		Language:                           "en",
		Region:                             "US",
		SkipSetupItems: mobile_device_prestages.SubsetSkipSetupItems{
			Location:            boolPtr(true),
			Privacy:             boolPtr(true),
			Biometric:           boolPtr(false),
			SoftwareUpdate:      boolPtr(false),
			Diagnostics:         boolPtr(true),
			IMessageAndFaceTime: boolPtr(false),
			Intelligence:        boolPtr(true),
			Passcode:            boolPtr(false),
			SIMSetup:            boolPtr(false),
			ScreenTime:          boolPtr(false),
			Siri:                boolPtr(false),
			Restore:             boolPtr(false),
			AppleID:             boolPtr(false),
			Payment:             boolPtr(true),
			TOS:                 boolPtr(true),
			Welcome:             boolPtr(false),
		},
		LocationInformation: mobile_device_prestages.SubsetLocationInformation{
			Username:     "jdoe",
			Realname:     "John Doe",
			Phone:        "555-0101",
			Email:        "jdoe@example.com",
			Room:         "Room 101",
			Position:     "IT Manager",
			DepartmentId: "-1",
			BuildingId:   "-1",
		},
		PurchasingInformation: mobile_device_prestages.SubsetPurchasingInformation{
			Leased:            boolPtr(false),
			Purchased:         boolPtr(true),
			AppleCareId:       "AC123456789",
			PoNumber:          "PO-2024-001",
			Vendor:            "Apple Inc.",
			PurchasePrice:     "699.00",
			LifeExpectancy:    3,
			PurchasingAccount: "IT-Budget",
			PurchasingContact: "procurement@example.com",
			LeaseDate:         "2024-01-01",
			PoDate:            "2024-01-15",
			WarrantyDate:      "2027-01-15",
		},
		Names: mobile_device_prestages.SubsetNames{
			AssignNamesUsing:       "PATTERN",
			DeviceNamePrefix:       "iPad",
			DeviceNameSuffix:       "DEV",
			ManageNames:            boolPtr(true),
			DeviceNamingConfigured: boolPtr(true),
		},
	}

	result, _, err := jamfClient.MobileDevicePrestages.CreateV3(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Created Mobile Device Prestage:\n" + string(out))
}
