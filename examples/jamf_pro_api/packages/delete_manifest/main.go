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

	_, err = jamfClient.JamfProAPI.Packages.DeletePackageManifestV1(context.Background(), packageID)
	if err != nil {
		fmt.Printf("Error deleting manifest: %v\n", err)
		return
	}
	fmt.Printf("Manifest deleted successfully from package %s\n", packageID)
}
