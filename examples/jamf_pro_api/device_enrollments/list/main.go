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

	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "id:asc",
	}

	enrollments, resp, err := jamfClient.DeviceEnrollments.ListV1(ctx, rsqlQuery)
	if err != nil {
		log.Fatalf("Failed to list device enrollments: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Total Count: %d\n", enrollments.TotalCount)
	fmt.Printf("Results:\n")
	for _, enrollment := range enrollments.Results {
		fmt.Printf("  - ID: %s, Name: %s, Server: %s\n", enrollment.ID, enrollment.Name, enrollment.ServerName)
	}
}
