package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

	// Example usage of GetWebhookByID
	webhookID := 4 // Replace with the desired webhook ID
	webhookByID, _, err := jamfClient.Webhooks.GetWebhookByID(context.Background(), webhookID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Pretty print the webhook details in XML
	webhooksXML, err := xml.MarshalIndent(webhookByID, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling webhook data: %v", err)
	}
	fmt.Println("Webhook Details (ID " + strconv.Itoa(webhookID) + "):\n" + string(webhooksXML))
}
