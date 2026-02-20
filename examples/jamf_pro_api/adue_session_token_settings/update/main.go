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

	current, _, err := jamfClient.AdueSessionTokenSettings.GetV1(context.Background())
	if err != nil {
		fmt.Printf("Error getting current: %v\n", err)
		return
	}

	current.Enabled = !current.Enabled
	updated, _, err := jamfClient.AdueSessionTokenSettings.UpdateV1(context.Background(), current)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated ADUE session token settings (enabled=%v)\n", updated.Enabled)
}
