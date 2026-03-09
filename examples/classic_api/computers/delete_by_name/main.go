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

	computerName := "MacBook-Pro-01" // Replace with the desired computer name to delete
	_, err = jamfClient.ClassicAPI.Computers.DeleteByName(context.Background(), computerName)
	if err != nil {
		fmt.Printf("Error deleting computer by name: %v\n", err)
		return
	}
	fmt.Println("Computer by name deleted successfully")
}
