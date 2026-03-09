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

	result, resp, err := client.JamfProAPI.GsxConnection.GetV1(context.Background())
	if err != nil {
		log.Fatalf("Failed to get GSX connection settings: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode())
	fmt.Printf("GSX Connection Enabled: %t\n", result.Enabled)
	fmt.Printf("Username: %s\n", result.Username)
	fmt.Printf("Service Account No: %s\n", result.ServiceAccountNo)
	fmt.Printf("Ship To No: %s\n", result.ShipToNo)
	fmt.Printf("Keystore Name: %s\n", result.GsxKeystore.Name)
	fmt.Printf("Keystore Expiration Epoch: %d\n", result.GsxKeystore.ExpirationEpoch)
}
