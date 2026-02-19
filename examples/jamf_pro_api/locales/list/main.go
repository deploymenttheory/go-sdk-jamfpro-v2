// Package main demonstrates ListLocalesV1 - lists locales.
//
// Run with: go run ./examples/jamf_pro_api/locales/list
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	ctx := context.Background()

	result, resp, err := client.Locales.ListLocalesV1(ctx)
	if err != nil {
		log.Fatalf("ListLocalesV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Locales: %d\n", resp.StatusCode, len(result))
	for i, l := range result {
		if i >= 10 {
			break
		}
		fmt.Printf("  %s %s\n", l.Identifier, l.Description)
	}
}
