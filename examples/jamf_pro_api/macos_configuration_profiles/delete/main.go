// Package main demonstrates how to delete a macOS configuration profile
// with custom settings schema using the Jamf Pro SDK.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
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

	resp, err := jamfClient.JamfProAPI.MacosConfigurationProfiles.DeleteByPayloadUUID(context.Background(), profilePayloadUUID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// A 204 does not guarantee removal. Against Jamf Pro 11.30.0 the profile
	// remained readable via GetByPayloadUUID indefinitely after a successful
	// delete -- see the package documentation.
	fmt.Printf("Delete request accepted (status: %d)\n", resp.StatusCode())
}
