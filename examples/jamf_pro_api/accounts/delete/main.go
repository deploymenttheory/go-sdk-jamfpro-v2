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

	// Replace "123" with the actual account ID you want to delete
	accountID := "123"

	resp, err := jamfClient.Accounts.DeleteByIDV1(context.Background(), accountID)
	if err != nil {
		fmt.Printf("Error deleting account: %v\n", err)
		return
	}
	fmt.Printf("Successfully deleted account (Status: %d)\n", resp.StatusCode)
}
