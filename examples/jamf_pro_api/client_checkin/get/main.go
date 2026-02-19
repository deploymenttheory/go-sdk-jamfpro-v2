// Package main demonstrates GetV3 - gets client check-in settings.
//
// Run with: go run ./examples/jamf_pro_api/client_checkin/get
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

	result, resp, err := client.ClientCheckin.GetV3(ctx)
	if err != nil {
		log.Fatalf("GetV3 failed: %v", err)
	}
	fmt.Printf("Status: %d CheckInFrequency: %d CreateHooks: %v\n", resp.StatusCode, result.CheckInFrequency, result.CreateHooks)
}
