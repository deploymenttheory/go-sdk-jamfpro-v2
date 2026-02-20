package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/allowed_file_extensions"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example usage of CreateAllowedFileExtension
	createReq := &allowed_file_extensions.RequestAllowedFileExtension{
		Extension: "example", // Replace with the desired file extension
	}

	created, _, err := jamfClient.AllowedFileExtensions.CreateAllowedFileExtension(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error creating allowed file extension: %v\n", err)
		return
	}
	fmt.Printf("Created Allowed File Extension: %+v\n", created)
}
