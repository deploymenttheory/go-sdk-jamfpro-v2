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

	// Get current failover URL
	before, _, err := jamfClient.JamfProAPI.SsoFailover.GetV1(context.Background())
	if err != nil {
		fmt.Printf("Error getting current failover URL: %v\n", err)
		return
	}
	fmt.Println("Current SSO failover URL:")
	out, _ := json.MarshalIndent(before, "", "    ")
	fmt.Println(string(out))

	// Regenerate failover URL
	result, _, err := jamfClient.JamfProAPI.SsoFailover.RegenerateV1(context.Background())
	if err != nil {
		fmt.Printf("Error regenerating failover URL: %v\n", err)
		return
	}
	fmt.Println("\nNew SSO failover URL:")
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Println(string(out))
}
