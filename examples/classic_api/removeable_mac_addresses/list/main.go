// Package main demonstrates ListRemoveableMacAddresses — lists all removeable MAC addresses via the Classic API.
//
// Run with: go run ./examples/classic_api/removeable_mac_addresses/list
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

	list, resp, err := client.RemoveableMacAddresses.ListRemoveableMacAddresses(ctx)
	if err != nil {
		log.Fatalf("ListRemoveableMacAddresses failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total count: %d\n", list.Size)
	fmt.Println("\nRemoveable MAC Addresses:")
	for _, mac := range list.Results {
		fmt.Printf("  ID: %d, Name: %s\n", mac.ID, mac.Name)
	}
}
