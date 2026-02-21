package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment_customizations"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Replace with actual enrollment customization ID
	customizationID := "1"

	noteRequest := &enrollment_customizations.RequestAddHistoryNotes{
		Note: "Added via Go SDK example",
	}

	result, resp, err := client.EnrollmentCustomizations.AddHistoryNotesV2(ctx, customizationID, noteRequest)
	if err != nil {
		log.Fatalf("Error adding history note: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Note added - ID: %d\n", result.ID)
	fmt.Printf("Username: %s\n", result.Username)
	fmt.Printf("Date: %s\n", result.Date)
	fmt.Printf("Note: %s\n", result.Note)
}
