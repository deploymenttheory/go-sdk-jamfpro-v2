package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_devices"
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

	// Create a new mobile device record.
	// Note: Creating via Classic API typically requires valid UDID/serial from MDM enrollment.
	// This example shows the structure; adjust fields for your environment.
	createReq := &mobile_devices.ResponseMobileDevice{
		General: mobile_devices.MobileDeviceSubsetGeneral{
			DisplayName:     "test-device-01",
			DeviceName:      "Test iPhone",
			Name:            "test-device-01",
			SerialNumber:    "TEST1234567",
			UDID:            "00008030-001234567890001E",
			Model:           "iPhone",
			ModelIdentifier: "iPhone14,2",
			ModelDisplay:    "iPhone 13 Pro",
			OSType:          "iOS",
			OSVersion:       "17.0",
		},
	}

	created, _, err := jamfClient.ClassicAPI.MobileDevices.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile device created with ID=%d name=%q\n", created.General.ID, created.General.Name)
}
