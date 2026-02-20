package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ibeacons"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example usage of UpdateIBeaconByID — replace with the desired iBeacon ID and updated request
	ibeaconID := 1
	updateReq := &ibeacons.RequestIBeacon{
		Name:  "go-sdk-v2-iBeacon-Updated",
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 2,
	}

	updatedIBeacon, _, err := jamfClient.IBeacons.UpdateIBeaconByID(context.Background(), ibeaconID, updateReq)
	if err != nil {
		fmt.Printf("Error updating iBeacon by ID: %v\n", err)
		return
	}
	fmt.Printf("Updated iBeacon: %+v\n", updatedIBeacon)
}
