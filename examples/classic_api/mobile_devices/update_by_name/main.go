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

	deviceName := "iPhone-01" // Replace with the desired mobile device name

	// Fetch existing device first
	existing, _, err := jamfClient.ClassicAPI.MobileDevices.GetByName(context.Background(), deviceName)
	if err != nil {
		fmt.Printf("Error fetching device: %v\n", err)
		return
	}

	// Update location info
	updateReq := &mobile_devices.ResponseMobileDevice{
		General:               existing.General,
		Location:              existing.Location,
		Purchasing:            existing.Purchasing,
		Applications:          existing.Applications,
		SecurityObject:        existing.SecurityObject,
		Network:               existing.Network,
		Certificates:          existing.Certificates,
		ConfigurationProfiles: existing.ConfigurationProfiles,
		ProvisioningProfiles:  existing.ProvisioningProfiles,
		MobileDeviceGroups:    existing.MobileDeviceGroups,
		ExtensionAttributes:   existing.ExtensionAttributes,
	}
	updateReq.Location.Department = "IT"
	updateReq.Location.Username = "jdoe"

	updated, _, err := jamfClient.ClassicAPI.MobileDevices.UpdateByName(context.Background(), deviceName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile device %q updated. Department: %s\n", deviceName, updated.Location.Department)
}
