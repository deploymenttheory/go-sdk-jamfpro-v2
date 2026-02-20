// Package main demonstrates DeleteRemoveableMacAddressByID — deletes a removeable MAC address by ID via the Classic API.
//
// Run with: go run ./examples/classic_api/removeable_mac_addresses/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
// Set REMOVEABLE_MAC_ADDRESS_ID environment variable to specify which ID to delete.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	idStr := os.Getenv("REMOVEABLE_MAC_ADDRESS_ID")
	if idStr == "" {
		log.Fatal("REMOVEABLE_MAC_ADDRESS_ID environment variable is required")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalf("invalid REMOVEABLE_MAC_ADDRESS_ID: %v", err)
	}

	ctx := context.Background()

	resp, err := client.RemoveableMacAddresses.DeleteRemoveableMacAddressByID(ctx, id)
	if err != nil {
		log.Fatalf("DeleteRemoveableMacAddressByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Successfully deleted removeable MAC address ID: %d\n", id)
}
