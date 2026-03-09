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

	createReq := &mobile_device_enrollment_profiles.Resource{
		General: mobile_device_enrollment_profiles.SubsetGeneral{
			Name:        "go-sdk-v2-enrollment-profile",
			Description: "Mobile device enrollment profile created via SDK v2",
		},
		Location: &mobile_device_enrollment_profiles.SubsetLocation{
			Username:     "jdoe",
			Realname:     "John Doe",
			EmailAddress: "jdoe@example.com",
			Department:   "IT",
		},
		Purchasing: &mobile_device_enrollment_profiles.SubsetPurchasing{
			IsPurchased: true,
			IsLeased:    false,
		},
	}

	created, _, err := jamfClient.ClassicAPI.MobileDeviceEnrollmentProfiles.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Enrollment Profile Created: ID=%d Name=%q\n", created.General.ID, created.General.Name)
}
