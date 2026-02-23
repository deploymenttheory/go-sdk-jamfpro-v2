// Package main demonstrates GetByID — retrieves a single policy by ID.
//
// Run with: go run ./examples/classic_api/policies/get_by_id
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set POLICY_ID or uses first from list.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	var id int
	if raw := os.Getenv("POLICY_ID"); raw != "" {
		id, err = strconv.Atoi(raw)
		if err != nil {
			log.Fatalf("invalid POLICY_ID %q: %v", raw, err)
		}
	} else {
		list, _, err := client.ClassicPolicies.List(ctx)
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set POLICY_ID or ensure at least one policy exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first policy ID: %d\n", id)
	}

	policy, resp, err := client.ClassicPolicies.GetByID(ctx, id)
	if err != nil {
		log.Fatalf("GetByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)

	// Display policy details as JSON for readability
	policyJSON, err := json.MarshalIndent(policy, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling policy: %v", err)
	}
	fmt.Printf("\nPolicy Details:\n%s\n", string(policyJSON))
}
