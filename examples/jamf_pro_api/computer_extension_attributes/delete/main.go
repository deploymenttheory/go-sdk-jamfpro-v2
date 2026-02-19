// Package main demonstrates DeleteComputerExtensionAttributeByIDV1 - deletes a computer extension attribute by ID.
//
// Run with: go run ./examples/jamf_pro_api/computer_extension_attributes/delete <id>
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/computer_extension_attributes/delete <id>")
	}
	id := os.Args[1]

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	resp, err := client.ComputerExtensionAttributes.DeleteComputerExtensionAttributeByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("DeleteComputerExtensionAttributeByIDV1 failed: %v", err)
	}

	log.Printf("Deleted computer extension attribute ID=%s (status %d)", id, resp.StatusCode)
}
