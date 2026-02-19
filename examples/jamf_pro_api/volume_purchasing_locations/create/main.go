// Package main demonstrates CreateVolumePurchasingLocationV1 - creates a new volume purchasing location.
// Requires a valid Apple VPP service token (ServiceToken). Reclaim is not called here; call
// ReclaimVolumePurchasingLocationByIDV1 after create if needed.
//
// Run with: go run ./examples/jamf_pro_api/volume_purchasing_locations/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, auth env vars, and a valid VPP service token.
package main

import (
	"context"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_locations"
)

func main() {
	token := os.Getenv("VPP_SERVICE_TOKEN")
	if token == "" {
		log.Fatal("VPP_SERVICE_TOKEN is required for this example")
	}

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &volume_purchasing_locations.RequestVolumePurchasingLocation{
		Name:                                  "Example VPL",
		ServiceToken:                          token,
		AutomaticallyPopulatePurchasedContent: true,
		SendNotificationWhenNoLongerAssigned:  false,
		AutoRegisterManagedUsers:              false,
	}

	result, resp, err := client.VolumePurchasingLocations.CreateVolumePurchasingLocationV1(ctx, req)
	if err != nil {
		log.Fatalf("CreateVolumePurchasingLocationV1 failed: %v", err)
	}

	log.Printf("Status: %d Created ID: %s Href: %s", resp.StatusCode, result.ID, result.Href)
}
