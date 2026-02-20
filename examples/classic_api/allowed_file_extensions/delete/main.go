package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

	// Example usage of DeleteAllowedFileExtensionByID
	extensionID := 1 // Replace with the desired allowed file extension ID to delete
	_, err = jamfClient.AllowedFileExtensions.DeleteAllowedFileExtensionByID(context.Background(), extensionID)
	if err != nil {
		fmt.Printf("Error deleting allowed file extension by ID: %v\n", err)
		return
	}
	fmt.Println("Allowed file extension by ID deleted successfully")
}
