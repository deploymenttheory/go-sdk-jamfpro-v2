package main

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	authConfig := &jamfpro.AuthConfig{
		InstanceDomain:           "https://your-instance.jamfcloud.com",
		AuthMethod:               constants.AuthMethodOAuth2,
		ClientID:                 "your-client-id",
		ClientSecret:             "your-client-secret",
		TokenRefreshBufferPeriod: 5 * time.Minute,
		HideSensitiveData:        true,
	}

	jamfClient, err := jamfpro.NewClient(authConfig)
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
