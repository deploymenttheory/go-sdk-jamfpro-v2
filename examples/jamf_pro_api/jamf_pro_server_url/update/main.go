package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_server_url"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example update request
	request := &jamf_pro_server_url.ResourceJamfProServerURL{
		URL:                    "https://jamf.example.com",
		UnsecuredEnrollmentUrl: "http://jamf.example.com:8080",
	}

	result, resp, err := client.JamfProAPI.JamfProServerUrl.UpdateV1(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to update Jamf Pro server URL settings: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode())
	fmt.Printf("Updated URL: %s\n", result.URL)
	fmt.Printf("Updated Unsecured Enrollment URL: %s\n", result.UnsecuredEnrollmentUrl)
}
