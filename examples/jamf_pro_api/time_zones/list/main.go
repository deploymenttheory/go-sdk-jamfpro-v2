// Package main demonstrates ListTimeZonesV1 - lists time zones.
//
// Run with: go run ./examples/jamf_pro_api/time_zones/list
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

	result, resp, err := client.TimeZones.ListTimeZonesV1(ctx)
	if err != nil {
		log.Fatalf("ListTimeZonesV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Time zones: %d\n", resp.StatusCode, len(result))
	for i, z := range result {
		if i >= 10 {
			break
		}
		fmt.Printf("  %s %s\n", z.ZoneId, z.DisplayName)
	}
}
