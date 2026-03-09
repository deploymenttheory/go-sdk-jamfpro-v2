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

	// Example 1: List all API integrations
	result, _, err := jamfClient.JamfProAPI.ApiIntegrations.ListV1(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error listing API integrations: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("All API integrations:\n%s\n\n", string(out))

	// Example 2: List with RSQL filter
	rsqlQuery := map[string]string{
		"filter": `displayName=="MyIntegration"`,
		"sort":   "displayName:asc",
	}
	result, _, err = jamfClient.JamfProAPI.ApiIntegrations.ListV1(context.Background(), rsqlQuery)
	if err != nil {
		fmt.Printf("Error listing filtered integrations: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Filtered API integrations:\n%s\n", string(out))
}
