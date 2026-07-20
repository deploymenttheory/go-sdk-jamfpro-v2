// Package main demonstrates how to create a macOS configuration profile
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

	// Level must be one of the SDK constants. The title-case "System" that
	// real .mobileconfig files carry in PayloadScope is rejected by the API.
	//
	// forced.plist must be a complete plist document; a bare <dict> fragment
	// is rejected.
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
					Plist: `<?xml version="1.0" encoding="UTF-8"?><plist version="1.0"><dict><key>ExampleKey</key><string>ExampleValue</string></dict></plist>`,
				},
			},
		},
	}

	result, resp, err := jamfClient.JamfProAPI.MacosConfigurationProfiles.Create(context.Background(), profile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Configuration profile created successfully (status: %d)\n", resp.StatusCode())
	fmt.Printf("New profile UUID: %s\n", result.UUID)
}
