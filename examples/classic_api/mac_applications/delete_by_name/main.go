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

	appName := "Sample Mac App" // Replace with the desired Mac application name to delete
	_, err = jamfClient.ClassicAPI.MacApplications.DeleteByName(context.Background(), appName)
	if err != nil {
		fmt.Printf("Error deleting Mac application by name: %v\n", err)
		return
	}
	fmt.Println("Mac application deleted successfully")
}
