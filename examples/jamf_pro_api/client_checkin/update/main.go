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

	// Option 1: Get current settings, modify, then update
	current, _, err := jamfClient.JamfProAPI.ClientCheckin.GetV3(context.Background())
	if err != nil {
		fmt.Printf("Error getting current settings: %v\n", err)
		return
	}
	settings := *current
	settings.CheckInFrequency = 15
	settings.CreateHooks = true

	result, _, err := jamfClient.JamfProAPI.ClientCheckin.UpdateV3(context.Background(), &settings)
	if err != nil {
		fmt.Printf("Error updating client check-in settings: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Client check-in settings updated:\n" + string(out))
}
