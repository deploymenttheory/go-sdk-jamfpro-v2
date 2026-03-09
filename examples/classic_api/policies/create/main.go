// Package main demonstrates Create — creates a new policy via the Classic API.
//
// Run with: go run ./examples/classic_api/policies/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a policy then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/policies"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Define the policy to create with realistic settings
	req := &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:      fmt.Sprintf("example-policy-%d", time.Now().UnixMilli()),
			Enabled:   true,
			Frequency: "Once per computer",
		},
		Scope: policies.PolicySubsetScope{
			AllComputers: true,
		},
		Maintenance: policies.PolicySubsetMaintenance{
			Recon: true,
		},
	}

	created, resp, err := client.ClassicAPI.Policies.Create(ctx, req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())
	fmt.Printf("Created policy ID: %d\n", created.ID)

	// Cleanup: delete the created policy
	if _, err := client.ClassicAPI.Policies.DeleteByID(ctx, created.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: policy deleted")
	}
}
