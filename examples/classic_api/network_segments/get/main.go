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

	// Example usage of GetNetworkSegmentByID
	segmentID := 1 // Replace with the desired network segment ID
	segment, _, err := jamfClient.NetworkSegments.GetNetworkSegmentByID(context.Background(), segmentID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Pretty print the network segment details in XML
	segmentXML, err := xml.MarshalIndent(segment, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling network segment data: %v", err)
	}
	fmt.Println("Network Segment Details (ID " + strconv.Itoa(segmentID) + "):\n" + string(segmentXML))
}
