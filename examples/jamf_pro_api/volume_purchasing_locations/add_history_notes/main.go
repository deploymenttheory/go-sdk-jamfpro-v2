package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_locations"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Replace with actual volume purchasing location ID
	vppLocationID := "1"

	// Create a history note request
	noteRequest := &volume_purchasing_locations.AddHistoryNotesRequest{
		ObjectHistoryNote: "Example history note added via API",
	}

	resp, err := client.VolumePurchasingLocations.AddHistoryNotesV1(ctx, vppLocationID, noteRequest)
	if err != nil {
		log.Fatalf("Error adding history note: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("History note added successfully\n")

	// Optionally, verify by getting the history
	history, _, err := client.VolumePurchasingLocations.GetHistoryV1(ctx, vppLocationID, nil)
	if err != nil {
		log.Printf("Warning: Could not retrieve history: %v", err)
	} else {
		fmt.Printf("\nTotal History Entries: %d\n", history.TotalCount)
		if len(history.Results) > 0 {
			fmt.Printf("Latest Entry:\n")
			fmt.Printf("  Username: %s\n", history.Results[0].Username)
			fmt.Printf("  Date: %s\n", history.Results[0].Date)
			fmt.Printf("  Note: %s\n", history.Results[0].Note)
		}
	}
}
