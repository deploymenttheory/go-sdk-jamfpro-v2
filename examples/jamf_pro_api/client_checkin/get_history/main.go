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

	// Example 1: Get all history (nil query)
	result, _, err := jamfClient.JamfProAPI.ClientCheckin.GetHistoryV3(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting client check-in history: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("All client check-in history (totalCount=%d):\n%s\n\n", result.TotalCount, string(out))

	// Example 2: Get history with RSQL filter and sort
	rsqlQuery := map[string]string{
		"filter": `username=="admin"`,
		"sort":   "date:desc",
		"page":   "0",
		"page-size": "10",
	}
	result, _, err = jamfClient.JamfProAPI.ClientCheckin.GetHistoryV3(context.Background(), rsqlQuery)
	if err != nil {
		fmt.Printf("Error getting filtered history: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Filtered client check-in history (username==admin):\n%s\n", string(out))
}
