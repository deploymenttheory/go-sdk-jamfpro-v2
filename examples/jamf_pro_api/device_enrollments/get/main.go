package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

	enrollment, resp, err := jamfClient.DeviceEnrollments.GetByIDV1(ctx, enrollmentID)
	if err != nil {
		log.Fatalf("Failed to get device enrollment: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", enrollment.ID)
	fmt.Printf("Name: %s\n", enrollment.Name)
	fmt.Printf("Server Name: %s\n", enrollment.ServerName)
	fmt.Printf("Server UUID: %s\n", enrollment.ServerUuid)
	fmt.Printf("Organization: %s\n", enrollment.OrgName)
	fmt.Printf("Token Expiration: %s\n", enrollment.TokenExpirationDate)
}
