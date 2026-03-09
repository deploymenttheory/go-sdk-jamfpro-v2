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

	profileName := "YOUR-PROFILE-NAME-HERE"
	result, resp, err := client.JamfProAPI.JamfConnect.GetConfigProfileByNameV1(context.Background(), profileName)
	if err != nil {
		log.Fatalf("Failed to get Jamf Connect config profile by name: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode())
	fmt.Printf("UUID: %s\n", result.UUID)
	fmt.Printf("Profile ID: %d\n", result.ProfileID)
	fmt.Printf("Profile Name: %s\n", result.ProfileName)
}
