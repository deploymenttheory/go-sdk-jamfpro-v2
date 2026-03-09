package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_provisioning_profiles"
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

	createReq := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name:        "go-sdk-v2-provisioning-profile",
			DisplayName: "Created via go-sdk-jamfpro-v2",
			UUID:        "550e8400-e29b-41d4-a716-446655440000",
		},
	}

	created, _, err := jamfClient.ClassicAPI.MobileDeviceProvisioningProfiles.CreateByID(context.Background(), 0, createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Provisioning Profile Created: ID=%d\n", created.ID)
}
