package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example usage of DeleteByID
	printerID := 1 // Replace with the desired printer ID to delete
	_, err = jamfClient.ClassicAPI.Printers.DeleteByID(context.Background(), printerID)
	if err != nil {
		fmt.Printf("Error deleting printer by ID: %v\n", err)
		return
	}
	fmt.Println("Printer by ID deleted successfully")
}
