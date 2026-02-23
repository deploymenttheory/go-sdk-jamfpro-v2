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
		log.Fatalf("Failed to create client: %v", err)
	}

	// Optional: add sorting or pagination parameters
	rsqlQuery := map[string]string{
		"sort": "date:desc",
	}

	result, resp, err := client.Enrollment.GetHistoryV2(context.Background(), rsqlQuery)
	if err != nil {
		log.Fatalf("Failed to get enrollment history: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Total History Entries: %d\n", result.TotalCount)

	for i, entry := range result.Results {
		fmt.Printf("\nHistory Entry #%d:\n", i+1)
		fmt.Printf("  ID: %d\n", entry.ID)
		fmt.Printf("  Username: %s\n", entry.Username)
		fmt.Printf("  Date: %s\n", entry.Date)
		fmt.Printf("  Note: %s\n", entry.Note)
		fmt.Printf("  Details: %s\n", entry.Details)
	}
}
