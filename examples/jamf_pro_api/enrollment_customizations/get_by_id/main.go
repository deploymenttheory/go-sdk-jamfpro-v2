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

	customization, resp, err := client.EnrollmentCustomizations.GetByIDV2(ctx, customizationID)
	if err != nil {
		log.Fatalf("Error getting enrollment customization: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", customization.ID)
	fmt.Printf("Display Name: %s\n", customization.DisplayName)
	fmt.Printf("Description: %s\n", customization.Description)
	fmt.Printf("Site ID: %s\n", customization.SiteID)
	fmt.Printf("Branding Settings:\n")
	fmt.Printf("  Text Color: %s\n", customization.BrandingSettings.TextColor)
	fmt.Printf("  Button Color: %s\n", customization.BrandingSettings.ButtonColor)
	fmt.Printf("  Button Text Color: %s\n", customization.BrandingSettings.ButtonTextColor)
	fmt.Printf("  Background Color: %s\n", customization.BrandingSettings.BackgroundColor)
	fmt.Printf("  Icon URL: %s\n", customization.BrandingSettings.IconUrl)
}
