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

	customization := &enrollment_customizations.ResourceEnrollmentCustomization{
		DisplayName: "Example Enrollment Customization",
		Description: "Created via Go SDK example",
		SiteID:      "-1", // -1 for None
		BrandingSettings: enrollment_customizations.SubsetBrandingSettings{
			TextColor:       "#000000",
			ButtonColor:     "#0066CC",
			ButtonTextColor: "#FFFFFF",
			BackgroundColor: "#F5F5F7",
			IconUrl:         "",
		},
	}

	created, resp, err := client.EnrollmentCustomizations.CreateV2(ctx, customization)
	if err != nil {
		log.Fatalf("Error creating enrollment customization: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created Enrollment Customization ID: %s\n", created.ID)
	fmt.Printf("Href: %s\n", created.Href)
}
