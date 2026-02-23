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

	attrName := "Department" // Replace with the desired user extension attribute name to delete
	_, err = jamfClient.ClassicUserExtensionAttributes.DeleteByName(context.Background(), attrName)
	if err != nil {
		fmt.Printf("Error deleting user extension attribute by name: %v\n", err)
		return
	}
	fmt.Println("User extension attribute by name deleted successfully")
}
