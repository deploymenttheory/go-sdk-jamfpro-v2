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

	current, _, err := jamfClient.JamfProAPI.SelfServiceSettings.Get(context.Background())
	if err != nil {
		fmt.Printf("Error getting current: %v\n", err)
		return
	}

	current.ConfigurationSettings.NotificationsEnabled = !current.ConfigurationSettings.NotificationsEnabled
	updated, _, err := jamfClient.JamfProAPI.SelfServiceSettings.Update(context.Background(), current)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated self-service settings (notificationsEnabled=%v)\n", updated.ConfigurationSettings.NotificationsEnabled)
}
