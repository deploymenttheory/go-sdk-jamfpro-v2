package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory_collection_settings"
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

	updateReq := &computer_inventory_collection_settings.ResourceComputerInventoryCollectionSettings{
		ComputerInventoryCollectionPreferences: computer_inventory_collection_settings.Preferences{
			MonitorApplicationUsage: true,
			IncludePackages:         true,
			IncludePrinters:         true,
		},
	}

	_, err = jamfClient.ComputerInventoryCollectionSettings.UpdateV2(context.Background(), updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Successfully updated computer inventory collection settings")
}
