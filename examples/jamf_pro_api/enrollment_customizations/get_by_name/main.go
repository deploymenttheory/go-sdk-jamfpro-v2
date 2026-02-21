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

	// Replace with actual enrollment customization display name
	displayName := "Default Customization"

	customization, resp, err := client.EnrollmentCustomizations.GetByNameV2(ctx, displayName)
	if err != nil {
		log.Fatalf("Error getting enrollment customization by name: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", customization.ID)
	fmt.Printf("Display Name: %s\n", customization.DisplayName)
	fmt.Printf("Description: %s\n", customization.Description)
}
