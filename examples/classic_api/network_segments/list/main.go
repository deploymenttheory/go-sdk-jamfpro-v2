package main

import (
	"context"
	"encoding/xml"
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

	// Call the ListNetworkSegments function to retrieve the list of network segments
	list, _, err := jamfClient.NetworkSegments.ListNetworkSegments(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Pretty print the network segments details in XML
	listXML, err := xml.MarshalIndent(list, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling network segments data: %v", err)
	}
	fmt.Println("Network Segments Details:\n" + string(listXML))
}
