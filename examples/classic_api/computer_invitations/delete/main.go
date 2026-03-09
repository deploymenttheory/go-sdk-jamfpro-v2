package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	invitationID := "1" // Replace with the desired computer invitation ID to delete
	_, err = jamfClient.ClassicAPI.ComputerInvitations.DeleteByID(context.Background(), invitationID)
	if err != nil {
		fmt.Printf("Error deleting computer invitation by ID: %v\n", err)
		return
	}
	fmt.Println("Computer invitation by ID deleted successfully")
}
