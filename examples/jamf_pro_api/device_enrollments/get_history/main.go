package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	authConfig := client.AuthConfigFromEnv()
	if err := authConfig.Validate(); err != nil {
		log.Fatalf("Invalid auth config: %v", err)
	}

	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()
	enrollmentID := "1"

	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "date:desc",
	}

	history, resp, err := jamfClient.DeviceEnrollments.GetHistoryV1(ctx, enrollmentID, rsqlQuery)
	if err != nil {
		log.Fatalf("Failed to get device enrollment history: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Total Count: %d\n", history.TotalCount)
	fmt.Printf("History Entries:\n")
	for _, entry := range history.Results {
		fmt.Printf("  - ID: %d, User: %s, Date: %s, Note: %s\n", entry.ID, entry.Username, entry.Date, entry.Note)
	}
}
