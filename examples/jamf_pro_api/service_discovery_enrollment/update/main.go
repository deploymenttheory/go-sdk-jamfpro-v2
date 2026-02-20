package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/service_discovery_enrollment"
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

	request := &service_discovery_enrollment.WellKnownSettingsResponseV1{
		WellKnownSettings: []service_discovery_enrollment.ResourceWellKnownSettingV1{
			{OrgName: "Example Org", ServerUUID: "A1B2C3D4-E5F6-7890-ABCD-EF1234567890", EnrollmentType: "USER_ENROLLMENT"},
		},
	}

	_, resp, err := jamfClient.ServiceDiscoveryEnrollment.UpdateV1(context.Background(), request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated well-known settings (HTTP %d)\n", resp.StatusCode)
}
