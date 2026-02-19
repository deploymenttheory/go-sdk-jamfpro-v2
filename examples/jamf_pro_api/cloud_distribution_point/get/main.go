// Package main demonstrates GetV1 - retrieves cloud distribution point configuration.
//
// Run with: go run ./examples/jamf_pro_api/cloud_distribution_point/get
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

	result, resp, err := client.CloudDistributionPoint.GetV1(ctx)
	if err != nil {
		log.Fatalf("GetV1 failed: %v", err)
	}
	fmt.Printf("Status: %d CdnType: %s Master: %v HasConnectionSucceeded: %v\n",
		resp.StatusCode, result.CdnType, result.Master, result.HasConnectionSucceeded)
}
