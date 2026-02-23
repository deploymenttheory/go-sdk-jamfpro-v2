package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	// Initialize Jamf Pro client using environment variables
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// WARNING: This will disable SSO for the entire Jamf Pro instance.
	// Users will no longer be able to authenticate via SSO until it is re-enabled.
	// Use this operation with caution.
	resp, err := client.SsoSettings.DisableV3(ctx)
	if err != nil {
		log.Fatalf("Error disabling SSO: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Println("SSO has been disabled successfully")
}
