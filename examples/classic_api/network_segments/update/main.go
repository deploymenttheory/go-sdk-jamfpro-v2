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

	// Example usage of UpdateNetworkSegmentByID — replace with the desired network segment ID and updated request
	segmentID := 1
	updateReq := &network_segments.RequestNetworkSegment{
		Name:            "go-sdk-v2-NetworkSegment-Updated",
		StartingAddress: "10.20.20.0",
		EndingAddress:   "10.20.20.128",
	}

	updatedSegment, _, err := jamfClient.NetworkSegments.UpdateNetworkSegmentByID(context.Background(), segmentID, updateReq)
	if err != nil {
		fmt.Printf("Error updating network segment by ID: %v\n", err)
		return
	}
	fmt.Printf("Updated Network Segment: %+v\n", updatedSegment)
}
