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

	accountName := "testuser"
	_, err = jamfClient.ClassicAPI.Accounts.DeleteByName(context.Background(), accountName)
	if err != nil {
		fmt.Printf("Error deleting account by name: %v\n", err)
		return
	}
	fmt.Printf("Account %q deleted successfully\n", accountName)
}
