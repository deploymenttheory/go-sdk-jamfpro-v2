// Package main demonstrates UpdateComputerExtensionAttributeByIDV1 - updates a computer extension attribute by ID.
//
// Run with: go run ./examples/jamf_pro_api/computer_extension_attributes/update <id>
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/computer_extension_attributes/update <id>")
	}
	id := os.Args[1]

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	enabled := true
	req := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 "Updated EA Name",
		Description:          "Updated description",
		DataType:             "String",
		Enabled:              &enabled,
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}

	result, resp, err := client.ComputerExtensionAttributes.UpdateComputerExtensionAttributeByIDV1(ctx, id, req)
	if err != nil {
		log.Fatalf("UpdateComputerExtensionAttributeByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated ID: %s Name: %s\n", result.ID, result.Name)
}
