// Package main demonstrates DeleteComputerExtensionAttributesByIDV1 - deletes multiple computer extension attributes by ID.
//
// Run with: go run ./examples/jamf_pro_api/computer_extension_attributes/delete_multiple <id1> [id2 ...]
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/computer_extension_attributes/delete_multiple <id1> [id2 ...]")
	}
	ids := os.Args[1:]

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	resp, err := client.ComputerExtensionAttributes.DeleteComputerExtensionAttributesByIDV1(ctx, &computer_extension_attributes.DeleteComputerExtensionAttributesByIDRequest{
		IDs: ids,
	})
	if err != nil {
		log.Fatalf("DeleteComputerExtensionAttributesByIDV1 failed: %v", err)
	}

	log.Printf("Deleted %d computer extension attribute(s) (status %d)", len(ids), resp.StatusCode)
}
