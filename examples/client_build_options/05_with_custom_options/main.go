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
		jamfpro.WithTimeout(60*time.Second),
		jamfpro.WithRetryCount(5),
		jamfpro.WithMaxConcurrentRequests(3),
		jamfpro.WithMandatoryRequestDelay(500*time.Millisecond),
		jamfpro.WithDebug(),
	)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
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
}
