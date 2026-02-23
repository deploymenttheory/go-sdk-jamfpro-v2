package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_settings"
)

func main() {
	// Initialize Jamf Pro client using environment variables
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Construct the history note request
	noteReq := &sso_settings.AddHistoryNoteRequest{
		Note: "Example history note added via SDK",
	}

	// Add the history note
	result, resp, err := client.SsoSettings.AddHistoryNoteV3(ctx, noteReq)
	if err != nil {
		log.Fatalf("Error adding SSO history note: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("History Note Added:\n")
	fmt.Printf("  ID: %s\n", result.ID)
	fmt.Printf("  Href: %s\n", result.Href)

	// Pretty print the response
	resultJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling result: %v", err)
	}
	fmt.Printf("\nFull Response:\n%s\n", string(resultJSON))

	// Optional: Verify by fetching the history
	fmt.Println("\nVerifying by fetching SSO history...")
	history, _, err := client.SsoSettings.GetHistoryV3(ctx, nil)
	if err != nil {
		log.Fatalf("Error getting SSO history: %v", err)
	}
	if len(history.Results) > 0 {
		fmt.Printf("Most recent history entry:\n")
		fmt.Printf("  Username: %s\n", history.Results[0].Username)
		fmt.Printf("  Date: %s\n", history.Results[0].Date)
		fmt.Printf("  Note: %s\n", history.Results[0].Note)
	}
}
