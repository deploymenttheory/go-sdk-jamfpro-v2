// Package main demonstrates DeleteRestrictedSoftwareByID — removes restricted software via the Classic API.
//
// Run with: go run ./examples/classic_api/restricted_software/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates restricted software then deletes it.
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

	// Create restricted software to delete
	createReq := &restricted_software.RequestRestrictedSoftware{
		General: restricted_software.RequestGeneral{
			Name:        fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
			ProcessName: fmt.Sprintf("process_delete_%d.exe", time.Now().UnixMilli()),
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

	resp, err := client.RestrictedSoftware.DeleteRestrictedSoftwareByID(ctx, created.ID)
	if err != nil {
		log.Fatalf("DeleteRestrictedSoftwareByID failed: %v", err)
	}

	fmt.Printf("Status: %d (200 = success)\n", resp.StatusCode)
	fmt.Println("Restricted software deleted successfully")
}
