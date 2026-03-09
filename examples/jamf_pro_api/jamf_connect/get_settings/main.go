package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	result, resp, err := client.JamfProAPI.JamfConnect.GetSettingsV1(context.Background())
	if err != nil {
		log.Fatalf("Failed to get Jamf Connect settings: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode())
	fmt.Printf("ID: %s\n", result.ID)
	fmt.Printf("Display Name: %s\n", result.DisplayName)
	fmt.Printf("Enabled: %t\n", result.Enabled)
	fmt.Printf("Version: %s\n", result.Version)
}
