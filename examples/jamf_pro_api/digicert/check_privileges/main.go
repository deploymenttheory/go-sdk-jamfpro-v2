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
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// ID of the DigiCert Trust Lifecycle Manager settings to check.
	id := "12"

	// The endpoint answers 204 No Content when every required permission is
	// present, so there is no response body to print.
	resp, err := jamfClient.JamfProAPI.Digicert.CheckPrivilegesByID(context.Background(), id)
	if err != nil {
		if client.IsForbidden(err) {
			fmt.Printf("DigiCert account %s is missing required permissions: %v\n", id, err)
			return
		}
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("DigiCert account %s holds all required permissions (HTTP %d)\n", id, resp.StatusCode())
}
