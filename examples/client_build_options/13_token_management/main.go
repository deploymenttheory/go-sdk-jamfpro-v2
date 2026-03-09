package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	authConfig := jamfpro.AuthConfigFromEnv()

	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	if err := jamfClient.GetTransport().KeepAliveToken(); err != nil {
		log.Printf("Warning: Failed to keep token alive: %v", err)
	}

	result, _, err := jamfClient.
		JamfProAPI.
		Buildings.
		ListV1(context.Background(), map[string]string{"page": "0", "page-size": "50"})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Buildings:\n" + string(out))

	if err := jamfClient.GetTransport().InvalidateToken(); err != nil {
		log.Printf("Warning: Failed to invalidate token: %v", err)
	}
}
