// Package main demonstrates GetByName — retrieves a single policy by name.
//
// Run with: go run ./examples/classic_api/policies/get_by_name
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set POLICY_NAME or uses first from list.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	var name string
	if name = os.Getenv("POLICY_NAME"); name == "" {
		list, _, err := client.ClassicAPI.Policies.List(ctx)
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set POLICY_NAME or ensure at least one policy exists")
		}
		name = list.Results[0].Name
		fmt.Printf("Using first policy name: %s\n", name)
	}

	policy, resp, err := client.ClassicAPI.Policies.GetByName(ctx, name)
	if err != nil {
		log.Fatalf("GetByName failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())

	// Display policy details as JSON for readability
	policyJSON, err := json.MarshalIndent(policy, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling policy: %v", err)
	}
	fmt.Printf("\nPolicy Details:\n%s\n", string(policyJSON))
}
