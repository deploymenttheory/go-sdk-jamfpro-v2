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

	prestageID := "1" // Replace with the desired prestage ID

	// First, get the current prestage to retrieve the versionLock
	current, _, err := jamfClient.MobileDevicePrestages.GetByIDV3(context.Background(), prestageID)
	if err != nil {
		fmt.Printf("Error getting prestage: %v\n", err)
		return
	}

	// Helper function to create bool pointers
	boolPtr := func(b bool) *bool { return &b }

	// Update the prestage with new settings
	updateReq := &mobile_device_prestages.ResourceMobileDevicePrestage{
		DisplayName:                        "go-sdk-v2-MobileDevicePrestage-Updated",
		Mandatory:                          boolPtr(true),
		MdmRemovable:                       boolPtr(false),
		SupportPhoneNumber:                 "555-0200",
		SupportEmailAddress:                "support-updated@example.com",
		Department:                         "IT Operations",
		DefaultPrestage:                    current.DefaultPrestage,
		EnrollmentSiteID:                   current.EnrollmentSiteID,
		KeepExistingSiteMembership:         current.KeepExistingSiteMembership,
		KeepExistingLocationInformation:    current.KeepExistingLocationInformation,
		RequireAuthentication:              boolPtr(true),
		AuthenticationPrompt:               "Updated: Please authenticate to enroll your device",
		PreventActivationLock:              boolPtr(true),
		EnableDeviceBasedActivationLock:    boolPtr(false),
		DeviceEnrollmentProgramInstanceID:  current.DeviceEnrollmentProgramInstanceID,
		AutoAdvanceSetup:                   boolPtr(true),
		AllowPairing:                       boolPtr(true),
		MultiUser:                          boolPtr(false),
		Supervised:                         boolPtr(true),
		MaximumSharedAccounts:              2,
		ConfigureDeviceBeforeSetupAssistant: boolPtr(false),
		SendTimezone:                       boolPtr(true),
		Timezone:                           "America/New_York",
		StorageQuotaSizeMegabytes:          8192,
		UseStorageQuotaSize:                boolPtr(false),
		TemporarySessionOnly:               current.TemporarySessionOnly,
		EnforceTemporarySessionTimeout:     current.EnforceTemporarySessionTimeout,
		TemporarySessionTimeout:            current.TemporarySessionTimeout,
		EnforceUserSessionTimeout:          current.EnforceUserSessionTimeout,
		UserSessionTimeout:                 current.UserSessionTimeout,
		Language:                           "en",
		Region:                             "US",
		SkipSetupItems:                     current.SkipSetupItems,
		LocationInformation: mobile_device_prestages.SubsetLocationInformation{
			Username:     "jadmin",
			Realname:     "Jane Admin",
			Phone:        "555-0202",
			Email:        "jadmin@example.com",
			Room:         "Room 202",
			Position:     "Senior IT Manager",
			DepartmentId: current.LocationInformation.DepartmentId,
			BuildingId:   current.LocationInformation.BuildingId,
			ID:           current.LocationInformation.ID,
			VersionLock:  current.LocationInformation.VersionLock,
		},
		PurchasingInformation: current.PurchasingInformation,
		Names:                 current.Names,
		VersionLock:           current.VersionLock, // Important: include versionLock for optimistic locking
	}

	result, _, err := jamfClient.MobileDevicePrestages.UpdateByIDV3(context.Background(), prestageID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Updated Mobile Device Prestage:\n" + string(out))
}
