// Package main demonstrates DeleteByIDV2 - deletes a patch software title configuration.
//
// Run with: go run ./examples/jamf_pro_api/patch_software_title_configurations/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

	// Replace "1" with an actual patch software title configuration ID to delete
	configID := "1"

	resp, err := jamfClient.PatchSoftwareTitleConfigurations.DeleteByIDV2(context.Background(), configID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Successfully deleted Patch Software Title Configuration (ID: %s). Status code: %d\n", configID, resp.StatusCode)
}
