// Package main demonstrates GetSMTPServerV2 - gets SMTP server configuration.
//
// Run with: go run ./examples/jamf_pro_api/smtp_server/get
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

	result, resp, err := client.SMTPServer.GetSMTPServerV2(ctx)
	if err != nil {
		log.Fatalf("GetSMTPServerV2 failed: %v", err)
	}
	fmt.Printf("Status: %d Enabled: %v AuthType: %s\n", resp.StatusCode, result.Enabled, result.AuthenticationType)
	if result.ConnectionSettings != nil {
		fmt.Printf("  Host: %s Port: %d\n", result.ConnectionSettings.Host, result.ConnectionSettings.Port)
	}
}
