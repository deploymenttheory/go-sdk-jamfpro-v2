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

	// Returns the JWKS used to verify Jamf Pro OIDC tokens. Note this endpoint
	// returns HTTP 503 on instances where OIDC/SSO is not configured.
	result, _, err := jamfClient.JamfProAPI.Oidc.GetPublicKeyV1(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Keys: %d\n\n", len(result.Keys))
	for _, key := range result.Keys {
		fmt.Printf("Key ID: %s\n", key.Kid)
		fmt.Printf("  Type: %s\n", key.Kty)
		fmt.Printf("  Algorithm: %s\n", key.Alg)
		fmt.Printf("  Use: %s\n", key.Use)
		fmt.Println()
	}
}
