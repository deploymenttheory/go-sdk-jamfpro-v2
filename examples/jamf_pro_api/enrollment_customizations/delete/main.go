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

	resp, err := client.EnrollmentCustomizations.DeleteByIDV2(ctx, customizationID)
	if err != nil {
		log.Fatalf("Error deleting enrollment customization: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Enrollment Customization %s deleted successfully\n", customizationID)
}
