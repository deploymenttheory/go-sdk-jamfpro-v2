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

	// Optional: add query parameters for pagination/sorting
	query := map[string]string{
		"sort": "id:desc",
	}

	result, _, err := jamfClient.JamfProAPI.SelfServiceBrandingIos.ListV1(context.Background(), query)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Self-Service Branding Mobile (iOS):")
	fmt.Println(string(out))
}
