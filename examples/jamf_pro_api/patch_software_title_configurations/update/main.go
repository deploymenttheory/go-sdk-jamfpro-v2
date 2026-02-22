// Package main demonstrates UpdateByIDV2 - updates an existing patch software title configuration.
//
// Run with: go run ./examples/jamf_pro_api/patch_software_title_configurations/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/patch_software_title_configurations"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Replace "1" with an actual patch software title configuration ID
	configID := "1"

	// First, get the existing configuration
	existing, _, err := jamfClient.PatchSoftwareTitleConfigurations.GetByIDV2(context.Background(), configID)
	if err != nil {
		fmt.Printf("Error getting existing config: %v\n", err)
		return
	}

	// Update fields as needed
	existing.UINotifications = false
	existing.EmailNotifications = true

	result, _, err := jamfClient.PatchSoftwareTitleConfigurations.UpdateByIDV2(context.Background(), configID, existing)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Updated Patch Software Title Configuration (ID: %s):\n%s\n", configID, string(out))
}
