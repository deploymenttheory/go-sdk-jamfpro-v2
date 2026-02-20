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

	// Example usage of CreateIBeacon
	newIBeacon := &ibeacons.RequestIBeacon{
		Name:  "go-sdk-v2-iBeacon",
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 1,
	}

	createdIBeacon, _, err := jamfClient.IBeacons.CreateIBeacon(context.Background(), newIBeacon)
	if err != nil {
		fmt.Printf("Error creating iBeacon: %v\n", err)
		return
	}
	fmt.Printf("Created iBeacon: %+v\n", createdIBeacon)
}
