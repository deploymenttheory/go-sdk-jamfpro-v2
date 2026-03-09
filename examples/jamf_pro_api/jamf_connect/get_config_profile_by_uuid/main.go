package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uuid := "YOUR-PROFILE-UUID-HERE"
	result, resp, err := client.JamfProAPI.JamfConnect.GetConfigProfileByUUIDV1(context.Background(), uuid)
	if err != nil {
		log.Fatalf("Failed to get Jamf Connect config profile by UUID: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode())
	fmt.Printf("UUID: %s\n", result.UUID)
	fmt.Printf("Profile ID: %d\n", result.ProfileID)
	fmt.Printf("Profile Name: %s\n", result.ProfileName)
	fmt.Printf("Version: %s\n", result.Version)
}
