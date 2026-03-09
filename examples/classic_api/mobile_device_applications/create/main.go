package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_applications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
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

	internalApp := true
	createReq := &mobile_device_applications.Resource{
		General: mobile_device_applications.SubsetGeneral{
			Name:        "go-sdk-v2-mobile-app",
			DisplayName: "go-sdk-v2-mobile-app",
			BundleID:    "com.apple.mobilesafari",
			Version:     "1.0",
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
			SelfServiceDescription: "Safari web browser for iOS",
		},
	}

	created, _, err := jamfClient.ClassicAPI.MobileDeviceApplications.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Application Created: ID=%d\n", created.ID)
}

func boolPtr(b bool) *bool {
	return &b
}
