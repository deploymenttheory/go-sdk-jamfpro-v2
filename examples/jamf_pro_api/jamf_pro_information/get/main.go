// Package main demonstrates GetV2 - gets Jamf Pro information (feature flags).
//
// Run with: go run ./examples/jamf_pro_api/jamf_pro_information/get
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

	result, resp, err := client.JamfProInformation.GetV2(ctx)
	if err != nil {
		log.Fatalf("GetV2 failed: %v", err)
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	if result.VppTokenEnabled != nil {
		fmt.Printf("VPP enabled: %v\n", *result.VppTokenEnabled)
	}
	if result.PatchEnabled != nil {
		fmt.Printf("Patch enabled: %v\n", *result.PatchEnabled)
	}
}
