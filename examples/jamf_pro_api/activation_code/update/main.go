package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/activation_code"
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

	// Replace with your actual activation code (32-39 characters, hyphens optional)
	req := &activation_code.ActivationCodeRequest{
		ActivationCode: "12345678-1234-1234-1234-123456789012",
	}

	resp, err := jamfClient.JamfProAPI.ActivationCode.UpdateV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error updating activation code: %v\n", err)
		return
	}
	fmt.Printf("Activation code updated successfully (Status: %d)\n", resp.StatusCode())
}
