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

	// The MDM server token should be obtained from Apple Business Manager
	// This is typically a .p7m file that has been base64-encoded
	request := &device_enrollments.RequestTokenUpload{
		TokenFileName: "updated-server-token.p7m",
		EncodedToken:  "BASE64_ENCODED_MDM_TOKEN_HERE",
	}

	updated, resp, err := jamfClient.DeviceEnrollments.UpdateTokenByIDV1(ctx, enrollmentID, request)
	if err != nil {
		log.Fatalf("Failed to update device enrollment token: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Updated Device Enrollment:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Server Name: %s\n", updated.ServerName)
	fmt.Printf("  Server UUID: %s\n", updated.ServerUuid)
	fmt.Printf("  Token Expiration: %s\n", updated.TokenExpirationDate)
}
