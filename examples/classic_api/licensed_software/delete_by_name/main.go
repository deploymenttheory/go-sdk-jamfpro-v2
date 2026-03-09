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

	licensedSoftwareName := "go-sdk-v2-licensed-software" // Replace with the desired licensed software name to delete
	_, err = jamfClient.ClassicAPI.LicensedSoftware.DeleteByName(context.Background(), licensedSoftwareName)
	if err != nil {
		fmt.Printf("Error deleting licensed software by name: %v\n", err)
		return
	}
	fmt.Println("Licensed software deleted successfully")
}
