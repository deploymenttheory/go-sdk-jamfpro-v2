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

	dpID := 1 // Replace with the desired distribution point ID to delete
	_, err = jamfClient.ClassicFileShareDistributionPoints.DeleteByID(context.Background(), dpID)
	if err != nil {
		fmt.Printf("Error deleting file share distribution point by ID: %v\n", err)
		return
	}
	fmt.Println("File share distribution point by ID deleted successfully")
}
