// Package main demonstrates how to create a macOS configuration profile
// with custom settings schema using the Jamf Pro SDK.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/macos_configuration_profile_custom_settings"
)

func main() {
	configFilePath := "/path/to/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	profile := &macos_configuration_profile_custom_settings.ResourceConfigProfile{
		PayloadUUID: "profile-payload-uuid",
		PayloadContent: []macos_configuration_profile_custom_settings.PayloadContentItem{
			{
				PayloadType:       "com.apple.ManagedClient.preferences",
				PayloadVersion:    1,
				PayloadIdentifier: "com.example.app",
				PayloadUUID:       "payload-uuid-1",
				PayloadDisplayName: "Example App Settings",
			},
		},
	}

	result, resp, err := jamfClient.MacOSConfigProfileCustomSettings.Create(context.Background(), profile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Configuration profile created successfully (status: %d)\n", resp.StatusCode)
	fmt.Printf("New profile UUID: %s\n", result.UUID)
}
