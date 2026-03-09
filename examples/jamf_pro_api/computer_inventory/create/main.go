package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/computer_inventory"
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

	computerInventoryRequest := &computer_inventory.ResourceComputerInventory{
		General: computer_inventory.ComputerInventorySubsetGeneral{
			Name:         "New-Test-Computer",
			AssetTag:     "ASSET-12345",
			Site:         computer_inventory.SharedResourceSiteProAPI{ID: "1", Name: "Default"},
			ManagementId: "test-mgmt-id",
		},
		Hardware: computer_inventory.ComputerInventorySubsetHardware{
			Make:                "Apple",
			Model:               "MacBook Pro",
			SerialNumber:        "C02TESTSERIAL",
			ProcessorSpeedMhz:   2600,
			TotalRamMegabytes:   16384,
		},
		UserAndLocation: computer_inventory.ComputerInventorySubsetUserAndLocation{
			Username:     "testuser",
			Realname:     "Test User",
			Email:        "testuser@example.com",
			Position:     "Developer",
			Phone:        "+1-555-0100",
			DepartmentId: "1",
			BuildingId:   "1",
			Room:         "101",
		},
	}

	result, _, err := jamfClient.JamfProAPI.ComputerInventory.CreateV3(context.Background(), computerInventoryRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Created computer inventory:\n%s\n", string(out))
}
