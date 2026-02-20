// Package main demonstrates CreateRemoveableMacAddress — creates a new removeable MAC address via the Classic API.
//
// Run with: go run ./examples/classic_api/removeable_mac_addresses/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a removeable MAC address then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/removeable_mac_addresses"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &removeable_mac_addresses.RequestRemoveableMacAddress{
		Name: fmt.Sprintf("AA:BB:CC:DD:EE:%d", time.Now().UnixMilli()%1000),
	}

	created, resp, err := client.RemoveableMacAddresses.CreateRemoveableMacAddress(ctx, req)
	if err != nil {
		log.Fatalf("CreateRemoveableMacAddress failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created removeable MAC address ID: %d\n", created.ID)
	fmt.Printf("Name: %s\n", created.Name)

	// Cleanup: delete the created removeable MAC address
	if _, err := client.RemoveableMacAddresses.DeleteRemoveableMacAddressByID(ctx, created.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: removeable MAC address deleted")
	}
}
