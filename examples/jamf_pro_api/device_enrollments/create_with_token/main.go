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

	// The MDM server token should be obtained from Apple Business Manager
	// This is typically a .p7m file that has been base64-encoded
	request := &device_enrollments.RequestTokenUpload{
		TokenFileName: "server-token.p7m",
		EncodedToken:  "BASE64_ENCODED_MDM_TOKEN_HERE",
	}

	created, resp, err := jamfClient.DeviceEnrollments.CreateWithTokenV1(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create device enrollment: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Created Device Enrollment:\n")
	fmt.Printf("  ID: %s\n", created.ID)
	fmt.Printf("  Href: %s\n", created.Href)
}
