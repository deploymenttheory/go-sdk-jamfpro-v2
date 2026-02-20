// Package main demonstrates UpdateRemoveableMacAddressByID — updates a removeable MAC address by ID via the Classic API.
//
// Run with: go run ./examples/classic_api/removeable_mac_addresses/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
// Set REMOVEABLE_MAC_ADDRESS_ID environment variable to specify which ID to update.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/removeable_mac_addresses"
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

	req := &removeable_mac_addresses.RequestRemoveableMacAddress{
		Name: fmt.Sprintf("Updated-MAC-%d", time.Now().UnixMilli()%1000),
	}

	updated, resp, err := client.RemoveableMacAddresses.UpdateRemoveableMacAddressByID(ctx, id, req)
	if err != nil {
		log.Fatalf("UpdateRemoveableMacAddressByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated removeable MAC address ID: %d\n", updated.ID)
	fmt.Printf("New name: %s\n", updated.Name)
}
