package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mdm"
)

func main() {
	// Initialize the Jamf Pro client from environment variables
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Get device UDIDs from environment (comma-separated) or use example values
	udidsEnv := os.Getenv("JAMF_DEVICE_UDIDS")
	var udids []string
	if udidsEnv != "" {
		udids = strings.Split(udidsEnv, ",")
		for i, u := range udids {
			udids[i] = strings.TrimSpace(u)
		}
	} else {
		udids = []string{"udid-001", "udid-002"}
	}

	req := &mdm.RenewProfileRequest{
		UDIDs: udids,
	}

	result, resp, err := client.MDM.RenewProfile(ctx, req)
	if err != nil {
		log.Fatalf("Failed to renew MDM profile: %v (HTTP %d)", err, resp.StatusCode)
	}

	fmt.Printf("MDM profile renewal completed (HTTP %d)\n", resp.StatusCode)
	if len(result.UDIDsNotProcessed.UDIDs) > 0 {
		fmt.Printf("UDIDs not processed: %v\n", result.UDIDsNotProcessed.UDIDs)
	} else {
		fmt.Println("All devices processed successfully.")
	}
}
