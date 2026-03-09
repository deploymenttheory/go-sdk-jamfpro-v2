// Package main demonstrates GetByName — retrieves computer history by computer name.
//
// Run with: go run ./examples/classic_api/computer_history/get_by_name
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
// Optional: COMPUTER_NAME — uses first computer from inventory if not set.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyname
package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	var computerName string
	if raw := os.Getenv("COMPUTER_NAME"); raw != "" {
		computerName = raw
	} else {
		list, _, err := client.JamfProAPI.ComputerInventory.ListV3(ctx, nil)
		if err != nil || list == nil || len(list.Results) == 0 {
			log.Fatal("Set COMPUTER_NAME or ensure at least one computer exists in inventory")
		}
		computerName = list.Results[0].General.Name
		fmt.Printf("Using first computer name: %s\n", computerName)
	}

	history, resp, err := client.ClassicAPI.ComputerHistory.GetByName(ctx, computerName)
	if err != nil {
		log.Fatalf("GetByName failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())
	xmlOut, err := xml.MarshalIndent(history, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling computer history: %v", err)
	}
	fmt.Printf("\nComputer History:\n%s\n", string(xmlOut))
}
