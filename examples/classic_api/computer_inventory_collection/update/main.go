// Package main demonstrates Update — updates computer inventory collection settings.
//
// Run with: go run ./examples/classic_api/computer_inventory_collection/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
//
// Doc: https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_inventory_collection"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Fetch current settings
	current, _, err := client.ClassicAPI.ComputerInventoryCollection.Get(ctx)
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}
	fmt.Printf("Current settings: local_user_accounts=%v inclue_applications=%v\n",
		current.LocalUserAccounts, current.InclueApplications)

	// Update with modified settings (toggle one boolean for demo)
	settings := &computer_inventory_collection.ResourceComputerInventoryCollection{
		LocalUserAccounts:             current.LocalUserAccounts,
		HomeDirectorySizes:            current.HomeDirectorySizes,
		HiddenAccounts:                current.HiddenAccounts,
		Printers:                      current.Printers,
		ActiveServices:                current.ActiveServices,
		MobileDeviceAppPurchasingInfo:  current.MobileDeviceAppPurchasingInfo,
		ComputerLocationInformation:    current.ComputerLocationInformation,
		PackageReceipts:               current.PackageReceipts,
		AvailableSoftwareUpdates:      current.AvailableSoftwareUpdates,
		InclueApplications:            current.InclueApplications,
		InclueFonts:                   current.InclueFonts,
		IncluePlugins:                 current.IncluePlugins,
		Applications:                  current.Applications,
		Fonts:                         current.Fonts,
		Plugins:                       current.Plugins,
	}

	resp, err := client.ClassicAPI.ComputerInventoryCollection.Update(ctx, settings)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())
	fmt.Println("Computer inventory collection settings updated successfully")
}
