// Package main demonstrates CreatePrinter — creates a new printer via the Classic API.
//
// Run with: go run ./examples/classic_api/printers/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a printer then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/printers"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &printers.RequestPrinter{
		Name:     fmt.Sprintf("example-printer-%d", time.Now().UnixMilli()),
		CUPSName: "example_printer",
		URI:      "ipp://printer.example.com/ipp",
		Location: "Example Lab",
		Model:    "Example Printer Model",
	}

	created, resp, err := client.Printers.CreatePrinter(ctx, req)
	if err != nil {
		log.Fatalf("CreatePrinter failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created printer ID: %d\n", created.ID)

	// Cleanup: delete the created printer
	if _, err := client.Printers.DeletePrinterByID(ctx, created.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: printer deleted")
	}
}
