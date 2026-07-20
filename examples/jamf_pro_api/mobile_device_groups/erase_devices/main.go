package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_groups"
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

	groupID := "8"

	// WARNING: this erases every mobile device in the group. All fields are
	// optional; omit the body entirely by passing an empty request.
	preserveDataPlan := true
	clearActivationLock := false
	req := &mobile_device_groups.RequestEraseDevices{
		PreserveDataPlan:    &preserveDataPlan,
		ClearActivationLock: &clearActivationLock,
	}

	_, err = jamfClient.JamfProAPI.MobileDeviceGroups.EraseDevicesByGroupIDV2(context.Background(), groupID, req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Erase command sent to all devices in group %s\n", groupID)
}
