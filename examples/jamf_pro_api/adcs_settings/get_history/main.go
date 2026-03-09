package main

import (
	"context"
	"encoding/json"
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

	// Replace "1" with the actual AD CS Settings ID
	adcsID := "1"

	// Example 1: Get all history
	result, _, err := jamfClient.JamfProAPI.AdcsSettings.GetHistoryByIDV1(context.Background(), adcsID, nil)
	if err != nil {
		fmt.Printf("Error getting AD CS Settings history: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("AD CS Settings history:\n%s\n\n", string(out))

	// Example 2: Get history with RSQL filter
	rsqlQuery := map[string]string{
		"filter": `username==admin`,
		"sort":   "date:desc",
	}
	result, _, err = jamfClient.JamfProAPI.AdcsSettings.GetHistoryByIDV1(context.Background(), adcsID, rsqlQuery)
	if err != nil {
		fmt.Printf("Error getting filtered history: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Filtered AD CS Settings history:\n%s\n", string(out))
}
