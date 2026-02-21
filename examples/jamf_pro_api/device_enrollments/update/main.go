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

	request := &device_enrollments.RequestUpdate{
		Name:                  "Updated Device Enrollment Name",
		SupervisionIdentityId: "2",
		SiteId:                "1",
	}

	updated, resp, err := jamfClient.DeviceEnrollments.UpdateByIDV1(ctx, enrollmentID, request)
	if err != nil {
		log.Fatalf("Failed to update device enrollment: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Updated Device Enrollment:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Server Name: %s\n", updated.ServerName)
	fmt.Printf("  Site ID: %s\n", updated.SiteId)
	fmt.Printf("  Supervision Identity ID: %s\n", updated.SupervisionIdentityId)
}
