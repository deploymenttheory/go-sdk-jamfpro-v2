// Package main demonstrates UpdateByID — updates an existing policy via the Classic API.
//
// Run with: go run ./examples/classic_api/policies/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/policies"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Create a policy to update
	createReq := &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:      fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
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
	created, _, err := client.ClassicPolicies.Create(ctx, createReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("Created policy ID: %d\n", created.ID)

	// Update the policy - change enabled status and frequency
	updateReq := &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:      fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
			Enabled:   false,
			Frequency: "Ongoing",
		},
		Scope: policies.PolicySubsetScope{
			AllComputers: true,
		},
		Maintenance: policies.PolicySubsetMaintenance{
			Recon:       true,
			Permissions: true,
		},
	}
	updated, resp, err := client.ClassicPolicies.UpdateByID(ctx, created.ID, updateReq)
	if err != nil {
		_, _ = client.ClassicPolicies.DeleteByID(ctx, created.ID)
		log.Fatalf("UpdateByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated policy ID: %d\n", updated.ID)

	_, _ = client.ClassicPolicies.DeleteByID(ctx, created.ID)
	fmt.Println("Cleanup: policy deleted")
}
