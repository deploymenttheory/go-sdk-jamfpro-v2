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

	profileName := "My Provisioning Profile"
	updateReq := &mobile_device_provisioning_profiles.RequestResource{
		General: mobile_device_provisioning_profiles.SubsetGeneral{
			Name:        "Updated Provisioning Profile Name",
			DisplayName: "Updated Provisioning Profile Name",
			UUID:        "550e8400-e29b-41d4-a716-446655440000",
		},
	}

	updated, _, err := jamfClient.ClassicAPI.MobileDeviceProvisioningProfiles.UpdateByName(context.Background(), profileName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Provisioning Profile %q updated: ID=%d\n", profileName, updated.ID)
}
