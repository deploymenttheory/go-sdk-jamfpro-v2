// Package main demonstrates DeleteVolumePurchasingLocationByIDV1 - deletes a volume purchasing location by ID.
//
// Run with: go run ./examples/jamf_pro_api/volume_purchasing_locations/delete <id>
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/volume_purchasing_locations/delete <id>")
	}
	id := os.Args[1]

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	resp, err := client.VolumePurchasingLocations.DeleteVolumePurchasingLocationByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("DeleteVolumePurchasingLocationByIDV1 failed: %v", err)
	}

	log.Printf("Deleted volume purchasing location ID=%s (status %d)", id, resp.StatusCode)
}
