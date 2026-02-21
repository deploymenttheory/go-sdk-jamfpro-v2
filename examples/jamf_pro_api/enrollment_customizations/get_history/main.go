package main

import (
	"context"
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

	// Replace with actual enrollment customization ID
	customizationID := "1"

	// Get history with pagination
	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "date:desc",
	}

	history, resp, err := client.EnrollmentCustomizations.GetHistoryV2(ctx, customizationID, rsqlQuery)
	if err != nil {
		log.Fatalf("Error getting enrollment customization history: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total history entries: %d\n", history.TotalCount)

	for _, entry := range history.Results {
		fmt.Printf("\nHistory Entry %d:\n", entry.ID)
		fmt.Printf("  Username: %s\n", entry.Username)
		fmt.Printf("  Date: %s\n", entry.Date)
		fmt.Printf("  Note: %s\n", entry.Note)
		if entry.Details != nil {
			fmt.Printf("  Details: %s\n", *entry.Details)
		}
	}
}
