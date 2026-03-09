package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/command_flush"
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

	// Create XML request to clear pending commands for multiple mobile devices
	req := &command_flush.RequestCommandFlush{
		Status: "Pending",
		MobileDevices: &command_flush.MobileDevices{
			MobileDevice: []command_flush.DeviceID{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
		},
	}

	resp, err := jamfClient.ClassicAPI.CommandFlush.FlushWithXML(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Successfully cleared pending MDM commands for batch of mobile devices (Status: %d)\n", resp.StatusCode())
}
