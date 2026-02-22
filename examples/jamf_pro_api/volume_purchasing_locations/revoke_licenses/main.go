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

	// Replace with actual volume purchasing location ID
	vppLocationID := "1"

	// WARNING: This operation will revoke all licenses for the specified VPP location
	// Use with caution in production environments

	resp, err := client.VolumePurchasingLocations.RevokeVolumePurchasingLocationLicensesByIDV1(ctx, vppLocationID)
	if err != nil {
		log.Fatalf("Error revoking licenses: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Licenses revoked successfully for VPP location ID: %s\n", vppLocationID)

	// Optionally, verify by getting the location details
	location, _, err := client.VolumePurchasingLocations.GetByIDV1(ctx, vppLocationID)
	if err != nil {
		log.Printf("Warning: Could not retrieve location details: %v", err)
	} else {
		fmt.Printf("\nLocation Details:\n")
		fmt.Printf("  Name: %s\n", location.Name)
		fmt.Printf("  Organization: %s\n", location.OrganizationName)
		fmt.Printf("  Total Purchased Licenses: %d\n", location.TotalPurchasedLicenses)
		fmt.Printf("  Total Used Licenses: %d\n", location.TotalUsedLicenses)
	}
}
