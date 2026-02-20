// Package main demonstrates UpdateRestrictedSoftwareByID — updates an existing restricted software via the Classic API.
//
// Run with: go run ./examples/classic_api/restricted_software/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/restricted_software"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Create restricted software to update
	createReq := &restricted_software.RequestRestrictedSoftware{
		General: restricted_software.RequestGeneral{
			Name:        fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
			ProcessName: fmt.Sprintf("process_create_%d.exe", time.Now().UnixMilli()),
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: restricted_software.Scope{
			AllComputers: true,
		},
	}
	created, _, err := client.RestrictedSoftware.CreateRestrictedSoftware(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateRestrictedSoftware failed: %v", err)
	}
	fmt.Printf("Created restricted software ID: %d\n", created.ID)

	// Update the restricted software
	updateReq := &restricted_software.RequestRestrictedSoftware{
		General: restricted_software.RequestGeneral{
			Name:        fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
			ProcessName: fmt.Sprintf("process_updated_%d.exe", time.Now().UnixMilli()),
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: restricted_software.Scope{
			AllComputers: true,
		},
	}
	updated, resp, err := client.RestrictedSoftware.UpdateRestrictedSoftwareByID(ctx, created.ID, updateReq)
	if err != nil {
		_, _ = client.RestrictedSoftware.DeleteRestrictedSoftwareByID(ctx, created.ID)
		log.Fatalf("UpdateRestrictedSoftwareByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated restricted software ID: %d\n", updated.ID)

	_, _ = client.RestrictedSoftware.DeleteRestrictedSoftwareByID(ctx, created.ID)
	fmt.Println("Cleanup: restricted software deleted")
}
