package main

import (
	"context"
	"encoding/json"
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

	// Example 1: Get all history
	result, _, err := jamfClient.JamfProAPI.ActivationCode.GetHistoryV1(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting activation code history: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("All activation code history:\n%s\n\n", string(out))

	// Example 2: Get history with RSQL filter
	rsqlQuery := map[string]string{
		"filter": `username==admin`,
		"sort":   "date:desc",
	}
	result, _, err = jamfClient.JamfProAPI.ActivationCode.GetHistoryV1(context.Background(), rsqlQuery)
	if err != nil {
		fmt.Printf("Error getting filtered history: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Activation code history for admin:\n%s\n", string(out))
}
