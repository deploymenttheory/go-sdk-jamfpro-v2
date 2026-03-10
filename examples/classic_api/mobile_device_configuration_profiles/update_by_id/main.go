package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_configuration_profiles"
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

	profileID := 1
	updateReq := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:             "Updated Wi-Fi Profile",
			Description:      "Updated via go-sdk-jamfpro-v2",
			DeploymentMethod: "Install Automatically",
			Site: &models.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: &mobile_device_configuration_profiles.SubsetScope{
			AllMobileDevices: true,
			AllJSSUsers:     false,
		},
	}

	updated, _, err := jamfClient.ClassicAPI.MobileDeviceConfigurationProfiles.UpdateByID(context.Background(), profileID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Configuration Profile Updated: ID=%d\n", updated.ID)
}
