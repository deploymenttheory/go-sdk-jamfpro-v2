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

	current, _, err := jamfClient.Reenrollment.Get(context.Background())
	if err != nil {
		fmt.Printf("Error getting current: %v\n", err)
		return
	}

	current.FlushPolicyHistory = !current.FlushPolicyHistory
	updated, _, err := jamfClient.Reenrollment.Update(context.Background(), current)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated re-enrollment settings (flushPolicyHistory=%v)\n", updated.FlushPolicyHistory)
}
