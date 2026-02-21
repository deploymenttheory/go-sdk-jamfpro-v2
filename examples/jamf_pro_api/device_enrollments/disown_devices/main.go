package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/device_enrollments"
)

func main() {
	authConfig := client.AuthConfigFromEnv()
	if err := authConfig.Validate(); err != nil {
		log.Fatalf("Invalid auth config: %v", err)
	}

	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()
	enrollmentID := "1"

	request := &device_enrollments.RequestDisown{
		Devices: []string{
			"C02ABCDEFGH1",
			"C02ABCDEFGH2",
			"C02ABCDEFGH3",
		},
	}

	result, resp, err := jamfClient.DeviceEnrollments.DisownDevicesByIDV1(ctx, enrollmentID, request)
	if err != nil {
		log.Fatalf("Failed to disown devices: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Disown Results:\n")
	for serial, status := range result.Devices {
		fmt.Printf("  Device %s: %s\n", serial, status)
	}
}
