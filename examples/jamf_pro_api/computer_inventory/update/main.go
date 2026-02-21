package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory"
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

	computerID := "1"

	updateReq := &computer_inventory.ResourceComputerInventory{
		UserAndLocation: computer_inventory.ComputerInventorySubsetUserAndLocation{
			Username: "newuser",
			Realname: "New User",
			Email:    "newuser@example.com",
			Position: "Developer",
			Phone:    "555-0123",
		},
	}

	result, _, err := jamfClient.ComputerInventory.UpdateByIDV1(context.Background(), computerID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Successfully updated computer: %s\n", result.ID)
	fmt.Printf("  Username: %s\n", result.UserAndLocation.Username)
	fmt.Printf("  Email: %s\n", result.UserAndLocation.Email)
}
