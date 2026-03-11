package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	packageID := "1"
	manifestPath := "/path/to/your/manifest.plist"

	result, _, err := jamfClient.JamfProAPI.Packages.AssignManifestToPackageV1(context.Background(), packageID, manifestPath)
	if err != nil {
		fmt.Printf("Error assigning manifest: %v\n", err)
		return
	}
	fmt.Printf("Manifest assigned successfully: %+v\n", result)
}
