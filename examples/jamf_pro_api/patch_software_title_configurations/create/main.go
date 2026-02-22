// Package main demonstrates CreateV2 - creates a new patch software title configuration.
//
// Run with: go run ./examples/jamf_pro_api/patch_software_title_configurations/create
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

	// Replace with an actual software title ID from your Jamf Pro environment
	req := &patch_software_title_configurations.ResourcePatchSoftwareTitleConfiguration{
		DisplayName:        "go-sdk-v2-PatchConfig",
		SoftwareTitleID:    "1", // Replace with valid software title ID
		CategoryID:         "1",
		UINotifications:    true,
		EmailNotifications: false,
		ExtensionAttributes: []patch_software_title_configurations.SubsetExtensionAttribute{
			{
				Accepted: true,
				EAID:     "1",
			},
		},
		Packages: []patch_software_title_configurations.SubsetPackage{
			{
				PackageID:   "1",
				Version:     "1.0.0",
				DisplayName: "MyApp-1.0.0.pkg",
			},
		},
	}

	result, _, err := jamfClient.PatchSoftwareTitleConfigurations.CreateV2(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Created Patch Software Title Configuration:\n" + string(out))
}
