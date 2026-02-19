// Package main demonstrates GetComputerExtensionAttributeByIDV1 - retrieves a single computer extension attribute by ID.
//
// Run with: go run ./examples/jamf_pro_api/computer_extension_attributes/get <id>
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/computer_extension_attributes/get <id>")
	}
	id := os.Args[1]

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	result, resp, err := client.ComputerExtensionAttributes.GetComputerExtensionAttributeByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("GetComputerExtensionAttributeByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", result.ID)
	fmt.Printf("Name: %s\n", result.Name)
	fmt.Printf("DataType: %s\n", result.DataType)
	fmt.Printf("InputType: %s\n", result.InputType)
	fmt.Printf("InventoryDisplayType: %s\n", result.InventoryDisplayType)
	if result.Description != "" {
		fmt.Printf("Description: %s\n", result.Description)
	}
}
