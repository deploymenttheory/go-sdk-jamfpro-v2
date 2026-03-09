package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_applications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
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

	appName := "go-sdk-v2-mobile-app" // Replace with the desired mobile device application name
	internalApp := true
	updateReq := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:        "go-sdk-v2-mobile-app-updated",
			DisplayName: "go-sdk-v2-mobile-app-updated",
			BundleID:    "com.apple.mobilesafari",
			Version:     "1.1",
			InternalApp: &internalApp,
			OsType:      "iOS",
			Site: &shared.SharedResourceSite{
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

	_, _, err = jamfClient.ClassicAPI.MobileDeviceApplications.UpdateByName(context.Background(), appName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Application %q updated successfully\n", appName)
}

func boolPtr(b bool) *bool {
	return &b
}
