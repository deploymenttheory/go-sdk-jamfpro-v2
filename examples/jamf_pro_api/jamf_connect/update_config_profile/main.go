package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_connect"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uuid := "YOUR-PROFILE-UUID-HERE"
	request := &jamf_connect.ResourceJamfConnectConfigProfileUpdate{
		JamfConnectVersion: "2.1.0",
		AutoDeploymentType: "INSTALL_AUTOMATICALLY",
	}

	result, resp, err := client.JamfConnect.UpdateConfigProfileByUUIDV1(context.Background(), uuid, request)
	if err != nil {
		log.Fatalf("Failed to update Jamf Connect config profile: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Updated Profile UUID: %s\n", result.UUID)
	fmt.Printf("Updated Version: %s\n", result.Version)
	fmt.Printf("Updated Auto Deployment Type: %s\n", result.AutoDeploymentType)
}
