package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_devices"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	deviceID := "1" // Replace with the desired mobile device ID

	// Fetch existing device first
	existing, _, err := jamfClient.ClassicAPI.MobileDevices.GetByID(context.Background(), deviceID)
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

	updated, _, err := jamfClient.ClassicAPI.MobileDevices.UpdateByID(context.Background(), deviceID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile device ID %s updated. Department: %s\n", deviceID, updated.Location.Department)
}
