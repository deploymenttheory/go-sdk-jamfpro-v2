// Package main demonstrates DeleteByID — removes a policy via the Classic API.
//
// Run with: go run ./examples/classic_api/policies/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a policy then deletes it.
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

	// Create a policy to delete
	createReq := &policies.ResourcePolicy{
		General: policies.PolicySubsetGeneral{
			Name:      fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
			Enabled:   true,
			Frequency: "Once per computer",
		},
		Scope: policies.PolicySubsetScope{
			AllComputers: true,
		},
	}
	created, _, err := client.ClassicPolicies.Create(ctx, createReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("Created policy ID: %d\n", created.ID)

	resp, err := client.ClassicPolicies.DeleteByID(ctx, created.ID)
	if err != nil {
		log.Fatalf("DeleteByID failed: %v", err)
	}

	fmt.Printf("Status: %d (200 = success)\n", resp.StatusCode)
	fmt.Println("Policy deleted successfully")
}
