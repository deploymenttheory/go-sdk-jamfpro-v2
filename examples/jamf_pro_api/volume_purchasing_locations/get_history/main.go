package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Replace with actual volume purchasing location ID
	vppLocationID := "1"

	history, resp, err := client.VolumePurchasingLocations.GetHistoryV1(ctx, vppLocationID, nil)
	if err != nil {
		log.Fatalf("Error getting volume purchasing location history: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total History Entries: %d\n", history.TotalCount)

	if len(history.Results) > 0 {
		fmt.Printf("\nHistory Entries:\n")
		for i, entry := range history.Results {
			fmt.Printf("  %d. ID: %s\n", i+1, entry.ID)
			fmt.Printf("     Username: %s\n", entry.Username)
			fmt.Printf("     Date: %s\n", entry.Date)
			fmt.Printf("     Note: %s\n", entry.Note)
			fmt.Printf("     Details: %s\n\n", entry.Details)
		}
	}

	// Pretty print full response
	historyJSON, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling history: %v", err)
	}
	fmt.Printf("\nFull History Response:\n%s\n", string(historyJSON))
}
