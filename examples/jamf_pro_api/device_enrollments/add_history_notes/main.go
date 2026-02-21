package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/device_enrollments"
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

	request := &device_enrollments.RequestAddHistoryNotes{
		Note: "Renewed MDM server token from Apple Business Manager",
	}

	result, resp, err := jamfClient.DeviceEnrollments.AddHistoryNotesV1(ctx, enrollmentID, request)
	if err != nil {
		log.Fatalf("Failed to add history notes: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("History Note Added:\n")
	fmt.Printf("  ID: %s\n", result.ID)
	fmt.Printf("  Href: %s\n", result.Href)
}
