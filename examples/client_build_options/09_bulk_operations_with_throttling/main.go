package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	authConfig := jamfpro.AuthConfigFromEnv()

	jamfClient, err := jamfpro.NewClient(
		authConfig,
		jamfpro.WithMaxConcurrentRequests(3),
		jamfpro.WithMandatoryRequestDelay(1*time.Second),
		jamfpro.WithTotalRetryDuration(5*time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	result, _, err := jamfClient.
		JamfProAPI.
		Buildings.
		ListV1(context.Background(), map[string]string{"page": "0", "page-size": "100"})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Buildings (with throttling):\n%s\n", string(out))
}
