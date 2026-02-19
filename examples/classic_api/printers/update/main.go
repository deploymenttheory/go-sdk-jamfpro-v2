// Package main demonstrates UpdatePrinterByID — updates an existing printer via the Classic API.
//
// Run with: go run ./examples/classic_api/printers/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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

	// Create a printer to update
	createReq := &printers.RequestPrinter{
		Name:     fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
		CUPSName: "example_update_printer",
		URI:      "ipp://printer.example.com/ipp",
		Location: "Example Lab",
		Model:    "Example Printer Model",
	}
	created, _, err := client.Printers.CreatePrinter(ctx, createReq)
	if err != nil {
		log.Fatalf("CreatePrinter failed: %v", err)
	}
	fmt.Printf("Created printer ID: %d\n", created.ID)

	// Update the printer
	updateReq := &printers.RequestPrinter{
		Name:     fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
		CUPSName: "example_updated_printer",
		URI:      "ipp://printer.example.com/ipp-updated",
		Location: "Updated Lab",
		Model:    "Updated Printer Model",
	}
	updated, resp, err := client.Printers.UpdatePrinterByID(ctx, created.ID, updateReq)
	if err != nil {
		_, _ = client.Printers.DeletePrinterByID(ctx, created.ID)
		log.Fatalf("UpdatePrinterByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated printer ID: %d\n", updated.ID)

	_, _ = client.Printers.DeletePrinterByID(ctx, created.ID)
	fmt.Println("Cleanup: printer deleted")
}
