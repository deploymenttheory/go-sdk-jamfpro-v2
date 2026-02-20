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

	// Example usage of UpdateWebhookByID — replace with the desired webhook ID and updated request
	webhookID := 1
	updateReq := &webhooks.RequestWebhook{
		Name:               "go-sdk-v2-Webhook-Updated",
		Enabled:            true,
		URL:                "https://server.com/updated",
		ContentType:        "application/json",
		Event:              "ComputerAdded",
		ConnectionTimeout:  30,
		ReadTimeout:        30,
		AuthenticationType: "NONE",
	}

	updatedWebhook, _, err := jamfClient.Webhooks.UpdateWebhookByID(context.Background(), webhookID, updateReq)
	if err != nil {
		fmt.Printf("Error updating webhook by ID: %v\n", err)
		return
	}
	fmt.Printf("Updated Webhook: %+v\n", updatedWebhook)
}
