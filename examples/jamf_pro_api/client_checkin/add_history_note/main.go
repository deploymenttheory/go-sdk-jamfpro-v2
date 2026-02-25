package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/client_checkin"
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

	noteRequest := &client_checkin.RequestClientCheckinHistoryNote{
		Note: "Manual note added via API - client check-in configuration change",
	}

	result, resp, err := jamfClient.ClientCheckin.AddHistoryNoteV3(context.Background(), noteRequest)
	if err != nil {
		fmt.Printf("Error adding history note: %v\n", err)
		return
	}
	fmt.Printf("History note added successfully (ID: %s, Status: %d)\n", result.ID, resp.StatusCode)
}
