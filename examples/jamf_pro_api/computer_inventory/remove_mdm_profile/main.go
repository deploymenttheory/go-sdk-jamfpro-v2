package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

	result, _, err := jamfClient.ComputerInventory.RemoveMDMProfileByIDV1(context.Background(), computerID)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Successfully initiated MDM profile removal\n")
	fmt.Printf("  Device ID: %s\n", result.DeviceID)
	fmt.Printf("  Command UUID: %s\n", result.CommandUUID)
}
