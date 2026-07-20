package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
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

	mobileDeviceID := "1"

	result, _, err := jamfClient.JamfProAPI.MobileDevices.GetDetailByIDV2(context.Background(), mobileDeviceID, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("ID: %s\n", result.ID)
	fmt.Printf("Name: %s\n", result.Name)
	fmt.Printf("Serial: %s\n", result.SerialNumber)
	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Last Inventory Update: %s\n", result.LastInventoryUpdateTimestamp)
	// lastContactTimestamp was added in Jamf Pro 11.30.
	fmt.Printf("Last Contact: %s\n", result.LastContactTimestamp)
}
