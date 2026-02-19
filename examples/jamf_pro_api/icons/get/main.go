// Package main demonstrates GetByIDV1 - gets icon metadata by ID.
//
// Run with: go run ./examples/jamf_pro_api/icons/get <id>
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/icons/get <id>")
	}
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("invalid id: %v", err)
	}

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	ctx := context.Background()

	result, resp, err := client.Icons.GetByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("GetByIDV1 failed: %v", err)
	}
	fmt.Printf("Status: %d ID: %d Name: %s URL: %s\n", resp.StatusCode, result.ID, result.Name, result.URL)
}
