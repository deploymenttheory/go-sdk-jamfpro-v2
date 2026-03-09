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

	// Example 1: Get all history
	result, _, err := jamfClient.JamfProAPI.CloudDistributionPoint.GetHistoryV1(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting cloud distribution point history: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("All cloud distribution point history:\n%s\n\n", string(out))

	// Example 2: Get history with RSQL filter
	rsqlQuery := map[string]string{
		"filter": `username==admin`,
		"sort":   "date:desc",
	}
	result, _, err = jamfClient.JamfProAPI.CloudDistributionPoint.GetHistoryV1(context.Background(), rsqlQuery)
	if err != nil {
		fmt.Printf("Error getting filtered history: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Cloud distribution point history for admin:\n%s\n", string(out))

	// Example 3: Pagination with page and page-size
	rsqlQuery = map[string]string{
		"page":      "0",
		"page-size": "10",
	}
	result, _, err = jamfClient.JamfProAPI.CloudDistributionPoint.GetHistoryV1(context.Background(), rsqlQuery)
	if err != nil {
		fmt.Printf("Error getting paginated history: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Paginated history (page 0, size 10):\n%s\n", string(out))
}
