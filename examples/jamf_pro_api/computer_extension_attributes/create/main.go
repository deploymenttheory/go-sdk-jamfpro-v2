// Package main demonstrates CreateComputerExtensionAttributeV1 - creates a new computer extension attribute.
//
// Run with: go run ./examples/jamf_pro_api/computer_extension_attributes/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	enabled := true
	req := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 fmt.Sprintf("example-ea-%d", time.Now().UnixMilli()),
		Description:          "Example computer extension attribute",
		DataType:             "String",
		Enabled:              &enabled,
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}

	result, resp, err := client.ComputerExtensionAttributes.CreateComputerExtensionAttributeV1(ctx, req)
	if err != nil {
		log.Fatalf("CreateComputerExtensionAttributeV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created computer extension attribute ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	if _, err := client.ComputerExtensionAttributes.DeleteComputerExtensionAttributeByIDV1(ctx, result.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: computer extension attribute deleted")
	}
}
