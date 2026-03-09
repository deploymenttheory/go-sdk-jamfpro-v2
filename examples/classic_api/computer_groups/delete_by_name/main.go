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

	groupName := "go-sdk-v2-smart-group" // Replace with the desired computer group name to delete
	_, err = jamfClient.ClassicAPI.ComputerGroups.DeleteByName(context.Background(), groupName)
	if err != nil {
		fmt.Printf("Error deleting computer group by name: %v\n", err)
		return
	}
	fmt.Println("Computer group by name deleted successfully")
}
