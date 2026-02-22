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

	// First, get the existing customization
	existing, _, err := client.EnrollmentCustomizations.GetByIDV2(ctx, customizationID)
	if err != nil {
		log.Fatalf("Error getting existing customization: %v", err)
	}

	// Update the customization
	existing.DisplayName = existing.DisplayName + " - Updated"
	existing.Description = "Updated via Go SDK example"
	existing.BrandingSettings.ButtonColor = "#FF0000" // Change button color to red

	updated, resp, err := client.EnrollmentCustomizations.UpdateByIDV2(ctx, customizationID, existing)
	if err != nil {
		log.Fatalf("Error updating enrollment customization: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated Enrollment Customization ID: %s\n", updated.ID)
	fmt.Printf("New Display Name: %s\n", updated.DisplayName)
	fmt.Printf("New Button Color: %s\n", updated.BrandingSettings.ButtonColor)
}
