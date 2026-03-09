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

	itemID := 1 // Replace with the desired dock item ID to delete
	_, err = jamfClient.ClassicAPI.DockItems.DeleteByID(context.Background(), itemID)
	if err != nil {
		fmt.Printf("Error deleting dock item by ID: %v\n", err)
		return
	}
	fmt.Println("Dock item by ID deleted successfully")
}
