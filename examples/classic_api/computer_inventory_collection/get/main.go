// Package main demonstrates Get — retrieves computer inventory collection settings.
//
// Run with: go run ./examples/classic_api/computer_inventory_collection/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
//
// Doc: https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
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

	settings, resp, err := client.ClassicAPI.ComputerInventoryCollection.Get(ctx)
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())
	fmt.Printf("Local User Accounts: %v\n", settings.LocalUserAccounts)
	fmt.Printf("Home Directory Sizes: %v\n", settings.HomeDirectorySizes)
	fmt.Printf("Hidden Accounts: %v\n", settings.HiddenAccounts)
	fmt.Printf("Printers: %v\n", settings.Printers)
	fmt.Printf("Active Services: %v\n", settings.ActiveServices)
	fmt.Printf("Inclue Applications: %v\n", settings.InclueApplications)
	fmt.Printf("Inclue Fonts: %v\n", settings.InclueFonts)
	fmt.Printf("Inclue Plugins: %v\n", settings.IncluePlugins)
	fmt.Printf("Applications count: %d\n", len(settings.Applications))
	fmt.Printf("Fonts count: %d\n", len(settings.Fonts))
	fmt.Printf("Plugins count: %d\n", len(settings.Plugins))
}
