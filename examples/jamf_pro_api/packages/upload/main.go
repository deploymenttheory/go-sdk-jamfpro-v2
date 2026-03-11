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
	filePath := "/path/to/your/package.pkg"

	result, _, err := jamfClient.JamfProAPI.Packages.UploadV1(context.Background(), packageID, filePath)
	if err != nil {
		fmt.Printf("Error uploading package: %v\n", err)
		return
	}
	fmt.Printf("Package uploaded successfully: %+v\n", result)
}
