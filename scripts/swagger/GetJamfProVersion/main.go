package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	authConfig := client.AuthConfigFromEnv()
	if err := authConfig.Validate(); err != nil {
		log.Fatalf("Failed to validate auth config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	result, _, err := jamfClient.JamfProVersion.GetV1(context.Background())
	if err != nil {
		log.Fatalf("Get Jamf Pro version: %v", err)
	}
	if result == nil || result.Version == nil || *result.Version == "" {
		log.Fatal("Get Jamf Pro version: empty version")
	}
	fmt.Println(*result.Version)
}
