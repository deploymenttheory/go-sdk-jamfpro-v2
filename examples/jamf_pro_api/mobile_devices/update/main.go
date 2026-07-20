package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_devices"
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

	mobileDeviceID := "1"

	enforceName := true
	updateReq := &mobile_devices.RequestMobileDeviceUpdateV2{
		Name:        "Loaner iPad 01",
		EnforceName: &enforceName,
		AssetTag:    "8675309",
		TimeZone:    "Europe/London",
		Location: &mobile_devices.MobileDeviceSubsetLocationV2{
			Username:     "jappleseed",
			RealName:     "John Appleseed",
			EmailAddress: "jappleseed@example.com",
			Position:     "Field Engineer",
		},
	}

	result, _, err := jamfClient.JamfProAPI.MobileDevices.UpdateByIDV2(context.Background(), mobileDeviceID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Successfully updated mobile device: %s\n", result.ID)
	fmt.Printf("  Name: %s\n", result.Name)
	fmt.Printf("  Asset Tag: %s\n", result.AssetTag)
}
