// Package main demonstrates GetBySerialNumber — retrieves computer history by serial number.
//
// Run with: go run ./examples/classic_api/computer_history/get_by_serial
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
// Optional: SERIAL_NUMBER — uses first computer's serial from inventory if not set.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerhistorybyserialnumber
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

	var serialNumber string
	if raw := os.Getenv("SERIAL_NUMBER"); raw != "" {
		serialNumber = raw
	} else {
		list, _, err := client.ComputerInventory.ListV3(ctx, nil)
		if err != nil || list == nil || len(list.Results) == 0 {
			log.Fatal("Set SERIAL_NUMBER or ensure at least one computer exists in inventory")
		}
		serialNumber = list.Results[0].Hardware.SerialNumber
		if serialNumber == "" {
			log.Fatal("First computer has no serial number; set SERIAL_NUMBER env var")
		}
		fmt.Printf("Using first computer serial: %s\n", serialNumber)
	}

	history, resp, err := client.ClassicComputerHistory.GetBySerialNumber(ctx, serialNumber)
	if err != nil {
		log.Fatalf("GetBySerialNumber failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	xmlOut, err := xml.MarshalIndent(history, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling computer history: %v", err)
	}
	fmt.Printf("\nComputer History:\n%s\n", string(xmlOut))
}
