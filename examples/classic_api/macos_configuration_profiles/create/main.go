package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/macos_configuration_profiles"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
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

	createReq := &macos_configuration_profiles.RequestResource{
		General: macos_configuration_profiles.SubsetGeneral{
			Name:          "go-sdk-v2-macos-profile",
			Description:   "Created via go-sdk-jamfpro-v2",
			UserRemovable: false,
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: &macos_configuration_profiles.SubsetScope{
			AllComputers: true,
			AllJSSUsers:  false,
		},
	}

	created, _, err := jamfClient.ClassicMacOSConfigurationProfiles.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("macOS Configuration Profile Created: ID=%d\n", created.ID)
}
