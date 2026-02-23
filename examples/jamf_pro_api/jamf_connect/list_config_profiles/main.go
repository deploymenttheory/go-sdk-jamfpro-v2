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

	result, resp, err := client.JamfConnect.ListConfigProfilesV1(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to list Jamf Connect config profiles: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Total Count: %d\n", result.TotalCount)

	for i, profile := range result.Results {
		fmt.Printf("\nProfile #%d:\n", i+1)
		fmt.Printf("  UUID: %s\n", profile.UUID)
		fmt.Printf("  Profile ID: %d\n", profile.ProfileID)
		fmt.Printf("  Profile Name: %s\n", profile.ProfileName)
		fmt.Printf("  Version: %s\n", profile.Version)
		fmt.Printf("  Auto Deployment Type: %s\n", profile.AutoDeploymentType)
	}
}
