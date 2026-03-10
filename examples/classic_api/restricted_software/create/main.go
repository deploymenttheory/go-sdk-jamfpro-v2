// Package main demonstrates Create — creates new restricted software via the Classic API.
//
// Run with: go run ./examples/classic_api/restricted_software/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates restricted software then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/restricted_software"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &restricted_software.RequestRestrictedSoftware{
		General: restricted_software.RequestGeneral{
			Name:        fmt.Sprintf("example-restricted-%d", time.Now().UnixMilli()),
			ProcessName: fmt.Sprintf("process_%d.exe", time.Now().UnixMilli()),
			Site: &models.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: restricted_software.Scope{
			AllComputers: true,
		},
	}

	created, resp, err := client.ClassicAPI.RestrictedSoftware.Create(ctx, req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())
	fmt.Printf("Created restricted software ID: %d\n", created.ID)

	// Cleanup: delete the created restricted software
	if _, err := client.ClassicAPI.RestrictedSoftware.DeleteByID(ctx, created.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: restricted software deleted")
	}
}
