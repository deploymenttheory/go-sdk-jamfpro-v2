// Package main demonstrates GetVolumePurchasingLocationByIDV1 - retrieves a volume purchasing location by ID.
//
// Run with: go run ./examples/jamf_pro_api/volume_purchasing_locations/get <id>
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/volume_purchasing_locations/get <id>")
	}
	id := os.Args[1]

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	result, resp, err := client.VolumePurchasingLocations.GetVolumePurchasingLocationByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("GetVolumePurchasingLocationByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s Name: %s\n", result.ID, result.Name)
	fmt.Printf("AutomaticallyPopulatePurchasedContent: %v\n", result.AutomaticallyPopulatePurchasedContent)
}
