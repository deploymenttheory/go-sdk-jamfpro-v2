package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	groupID := "8"

	result, _, err := jamfClient.JamfProAPI.MobileDeviceGroups.GetStaticGroupMembershipV2(context.Background(), groupID, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Members: %d\n\n", result.TotalCount)
	for _, member := range result.Results {
		fmt.Printf("ID: %s\n", member.MobileDeviceID)
		fmt.Printf("  Name: %s\n", member.DisplayName)
		fmt.Printf("  Serial: %s\n", member.SerialNumber)
		// lastContactDate was added to the membership response in Jamf Pro 11.30.
		fmt.Printf("  Last Contact: %s\n", member.LastContactDate)
		fmt.Println()
	}
}
