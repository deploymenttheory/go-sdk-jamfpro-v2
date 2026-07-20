// Package main demonstrates how to update a macOS configuration profile
// with custom settings schema using the Jamf Pro SDK.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/macos_configuration_profiles"
)

func main() {
	configFilePath := "/path/to/clientconfig.json"
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	profilePayloadUUID := "replace-with-an-existing-profile-payload-uuid"

	// Update replaces the profile wholesale. Level must be one of the SDK
	// constants -- the title-case "System" that real .mobileconfig files carry
	// in PayloadScope is rejected by the API. forced.plist must be a complete
	// plist document; a bare <dict> fragment is rejected.
	profile := &macos_configuration_profiles.ResourceConfigProfile{
		Level: macos_configuration_profiles.ConfigProfileLevelSystem,
		PayloadContent: []macos_configuration_profiles.PayloadContentItem{
			{
				PayloadType:        macos_configuration_profiles.PayloadTypeManagedClientPreferences,
				PayloadVersion:     1,
				PayloadIdentifier:  "com.example.app",
				PayloadDisplayName: "Example App Settings",
				PreferenceDomain:   "com.example.app",
				Forced: &macos_configuration_profiles.ForcedSettings{
					Plist: `<?xml version="1.0" encoding="UTF-8"?><plist version="1.0"><dict><key>ExampleKey</key><string>UpdatedValue</string></dict></plist>`,
				},
			},
		},
	}

	result, resp, err := jamfClient.JamfProAPI.MacosConfigurationProfiles.UpdateByPayloadUUID(context.Background(), profilePayloadUUID, profile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Update returns the UUID envelope, not the updated resource. Read the
	// profile back with GetByPayloadUUID to confirm the change landed.
	fmt.Printf("Configuration profile updated successfully (status: %d)\n", resp.StatusCode())
	fmt.Printf("Profile UUID: %s\n", result.UUID)
}
