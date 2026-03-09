package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_configuration_profiles"
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

	createReq := &mobile_device_configuration_profiles.RequestResource{
		General: mobile_device_configuration_profiles.SubsetGeneral{
			Name:             "go-sdk-v2-md-profile",
			Description:      "Created via go-sdk-jamfpro-v2",
			DeploymentMethod: "Install Automatically",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: &mobile_device_configuration_profiles.SubsetScope{
			AllMobileDevices: true,
			AllJSSUsers:     false,
		},
	}

	created, _, err := jamfClient.ClassicAPI.MobileDeviceConfigurationProfiles.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Configuration Profile Created: ID=%d\n", created.ID)
}
