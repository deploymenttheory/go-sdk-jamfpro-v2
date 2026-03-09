package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_enrollment_profiles"
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

	profileName := "My Enrollment Profile" // Replace with the desired profile name
	updateReq := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{
			Name:        "Updated Enrollment Profile",
			Description: "Updated description via SDK v2",
		},
	}

	updated, _, err := jamfClient.ClassicAPI.MobileDeviceEnrollmentProfiles.UpdateByName(context.Background(), profileName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Enrollment Profile Updated: ID=%d Name=%q\n", updated.General.ID, updated.General.Name)
}
