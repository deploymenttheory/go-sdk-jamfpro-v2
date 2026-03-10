package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_applications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	appID := 1 // Replace with the desired mobile device application ID
	internalApp := true
	updateReq := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:        "go-sdk-v2-mobile-app-updated",
			DisplayName: "go-sdk-v2-mobile-app-updated",
			BundleID:    "com.apple.mobilesafari",
			Version:     "1.1",
			InternalApp: &internalApp,
			OsType:      "iOS",
			Site: &models.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: mobile_device_applications.SubsetScope{
			AllMobileDevices: boolPtr(true),
			AllJSSUsers:     boolPtr(false),
		},
		SelfService: mobile_device_applications.SubsetSelfService{
			SelfServiceDescription: "Safari web browser for iOS (updated)",
		},
	}

	_, _, err = jamfClient.ClassicAPI.MobileDeviceApplications.UpdateByID(context.Background(), appID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Application ID %d updated successfully\n", appID)
}

func boolPtr(b bool) *bool {
	return &b
}
