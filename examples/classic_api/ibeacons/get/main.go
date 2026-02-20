package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

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

	// Example usage of GetIBeaconByID
	ibeaconID := 1 // Replace with the desired iBeacon ID
	ibeaconByID, _, err := jamfClient.IBeacons.GetIBeaconByID(context.Background(), ibeaconID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Pretty print the iBeacon details in XML
	ibeaconXML, err := xml.MarshalIndent(ibeaconByID, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling iBeacon data: %v", err)
	}
	fmt.Println("iBeacon Details (ID " + strconv.Itoa(ibeaconID) + "):\n" + string(ibeaconXML))
}
