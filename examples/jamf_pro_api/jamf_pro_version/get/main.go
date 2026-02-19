// Package main demonstrates GetV1 - gets Jamf Pro server version.
//
// Run with: go run ./examples/jamf_pro_api/jamf_pro_version/get
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

	result, resp, err := client.JamfProVersion.GetV1(ctx)
	if err != nil {
		log.Fatalf("GetV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Version: %s\n", resp.StatusCode, stringPtrToStr(result.Version))
}

func stringPtrToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
