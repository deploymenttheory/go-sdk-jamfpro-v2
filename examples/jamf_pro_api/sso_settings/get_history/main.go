package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	// Initialize Jamf Pro client using environment variables
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Optional: Add RSQL query parameters for filtering, pagination, or sorting
	// Example: rsqlQuery := map[string]string{
	//     "page": "0",
	//     "page-size": "100",
	//     "sort": "date:desc",
	//     "filter": "username=='admin'",
	// }
	// For this example, we'll retrieve all history entries
	history, resp, err := client.SsoSettings.GetHistoryV3(ctx, nil)
	if err != nil {
		log.Fatalf("Error getting SSO settings history: %v", err)
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
