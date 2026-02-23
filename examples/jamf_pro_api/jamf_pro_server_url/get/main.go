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

	result, resp, err := client.JamfProServerURL.GetV1(context.Background())
	if err != nil {
		log.Fatalf("Failed to get Jamf Pro server URL settings: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("URL: %s\n", result.URL)
	fmt.Printf("Unsecured Enrollment URL: %s\n", result.UnsecuredEnrollmentUrl)
}
