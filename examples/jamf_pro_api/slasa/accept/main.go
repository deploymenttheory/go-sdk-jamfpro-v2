// Package main demonstrates how to accept the SLASA (Software License Agreement
// Service Acceptance) using the Jamf Pro SDK.
//
// This must be done before managed software updates can be used. It is typically
// a one-time operation per Jamf Pro instance.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	configFilePath := "/path/to/clientconfig.json"
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	resp, err := jamfClient.JamfProAPI.Slasa.AcceptV1(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("SLASA accepted successfully (status: %d)\n", resp.StatusCode())
}
