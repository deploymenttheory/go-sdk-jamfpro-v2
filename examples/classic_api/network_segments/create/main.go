package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/network_segments"
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

	// Example usage of CreateNetworkSegment
	newSegment := &network_segments.RequestNetworkSegment{
		Name:            "go-sdk-v2-NetworkSegment",
		StartingAddress: "10.10.10.0",
		EndingAddress:   "10.10.10.255",
	}

	createdSegment, _, err := jamfClient.NetworkSegments.CreateNetworkSegment(context.Background(), newSegment)
	if err != nil {
		fmt.Printf("Error creating network segment: %v\n", err)
		return
	}
	fmt.Printf("Created Network Segment: %+v\n", createdSegment)
}
