// Package main demonstrates GetByIDAndSubset — retrieves a subset of computer history by ID.
//
// Run with: go run ./examples/classic_api/computer_history/get_by_id_with_subset
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
// Optional: COMPUTER_ID — uses first computer from inventory if not set.
// Optional: SUBSET — defaults to "General". Other subsets: Audits, PolicyLogs, etc.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyid
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

	var computerID string
	if raw := os.Getenv("COMPUTER_ID"); raw != "" {
		computerID = raw
	} else {
		list, _, err := client.ComputerInventory.ListV3(ctx, nil)
		if err != nil || list == nil || len(list.Results) == 0 {
			log.Fatal("Set COMPUTER_ID or ensure at least one computer exists in inventory")
		}
		computerID = list.Results[0].ID
		fmt.Printf("Using first computer ID: %s\n", computerID)
	}

	subset := os.Getenv("SUBSET")
	if subset == "" {
		subset = "General"
	}

	history, resp, err := client.ClassicComputerHistory.GetByIDAndSubset(ctx, computerID, subset)
	if err != nil {
		log.Fatalf("GetByIDAndSubset failed: %v", err)
	}

	fmt.Printf("Status: %d (subset=%q)\n", resp.StatusCode, subset)
	xmlOut, err := xml.MarshalIndent(history, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling computer history: %v", err)
	}
	fmt.Printf("\nComputer History (subset %q):\n%s\n", subset, string(xmlOut))
}
