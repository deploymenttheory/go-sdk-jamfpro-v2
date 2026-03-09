package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	venafiID := "1" // Replace with actual Venafi configuration ID
	_, data, err := jamfClient.JamfProAPI.Venafi.GetJamfPublicKey(context.Background(), venafiID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if err := os.WriteFile("jamf-public-key.pem", data, 0600); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	fmt.Printf("Downloaded %d bytes to jamf-public-key.pem\n", len(data))
}
