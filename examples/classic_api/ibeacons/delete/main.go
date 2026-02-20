package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

	// Example usage of DeleteIBeaconByID
	ibeaconID := 1 // Replace with the desired iBeacon ID to delete
	_, err = jamfClient.IBeacons.DeleteIBeaconByID(context.Background(), ibeaconID)
	if err != nil {
		fmt.Printf("Error deleting iBeacon by ID: %v\n", err)
		return
	}
	fmt.Println("iBeacon by ID deleted successfully")
}
