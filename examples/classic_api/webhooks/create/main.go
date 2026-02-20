package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/webhooks"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example usage of CreateWebhook
	newWebhook := &webhooks.RequestWebhook{
		Name:               "go-sdk-v2-Webhook",
		Enabled:            true,
		URL:                "https://server.com",
		ContentType:        "application/json",
		Event:              "SmartGroupComputerMembershipChange",
		ConnectionTimeout:  30,
		ReadTimeout:        30,
		AuthenticationType: "NONE",
	}

	createdWebhook, _, err := jamfClient.Webhooks.CreateWebhook(context.Background(), newWebhook)
	if err != nil {
		fmt.Printf("Error creating webhook: %v\n", err)
		return
	}
	fmt.Printf("Created Webhook: %+v\n", createdWebhook)
}
