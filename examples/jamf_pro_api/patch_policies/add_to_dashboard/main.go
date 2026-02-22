package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Replace "1" with the actual patch policy ID you want to add to dashboard
	policyID := "1"

	resp, err := jamfClient.PatchPolicies.AddToDashboardV2(context.Background(), policyID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Successfully added patch policy (ID: %s) to dashboard. Status: %d\n", policyID, resp.StatusCode)

	// Verify it was added
	status, _, err := jamfClient.PatchPolicies.GetDashboardStatusV2(context.Background(), policyID)
	if err != nil {
		fmt.Printf("Error verifying status: %v\n", err)
		return
	}
	fmt.Printf("Dashboard Status - OnDashboard: %v\n", status.OnDashboard)
}
