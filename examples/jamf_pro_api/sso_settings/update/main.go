package main

import (
	"context"
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

	current, _, err := jamfClient.SsoSettings.GetV3(context.Background())
	if err != nil {
		fmt.Printf("Error getting current: %v\n", err)
		return
	}

	current.SsoBypassAllowed = !current.SsoBypassAllowed
	updated, _, err := jamfClient.SsoSettings.UpdateV3(context.Background(), current)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated SSO settings (ssoBypassAllowed=%v)\n", updated.SsoBypassAllowed)
}
