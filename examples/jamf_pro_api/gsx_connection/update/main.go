package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/gsx_connection"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example update request
	request := &gsx_connection.ResourceGSXConnection{
		Enabled:          true,
		Username:         "gsx_user@example.com",
		ServiceAccountNo: "12345",
		ShipToNo:         "67890",
		GsxKeystore: gsx_connection.GsxKeystore{
			Name:            "certificate.p12",
			ExpirationEpoch: 1691954900000,
			ErrorMessage:    "",
		},
	}

	result, resp, err := client.JamfProAPI.GsxConnection.UpdateV1(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to update GSX connection settings: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode())
	fmt.Printf("Updated GSX Connection Enabled: %t\n", result.Enabled)
	fmt.Printf("Updated Username: %s\n", result.Username)
	fmt.Printf("Updated Service Account No: %s\n", result.ServiceAccountNo)
}
