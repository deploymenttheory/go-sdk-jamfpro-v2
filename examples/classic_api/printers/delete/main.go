// Package main demonstrates DeletePrinterByID — removes a printer via the Classic API.
//
// Run with: go run ./examples/classic_api/printers/delete
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

	// Create a printer to delete
	createReq := &printers.RequestPrinter{
		Name:     fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
		CUPSName: "example_delete_printer",
		URI:      "ipp://printer.example.com/ipp",
	}
	created, _, err := client.Printers.CreatePrinter(ctx, createReq)
	if err != nil {
		log.Fatalf("CreatePrinter failed: %v", err)
	}
	fmt.Printf("Created printer ID: %d\n", created.ID)

	resp, err := client.Printers.DeletePrinterByID(ctx, created.ID)
	if err != nil {
		log.Fatalf("DeletePrinterByID failed: %v", err)
	}

	fmt.Printf("Status: %d (200 = success)\n", resp.StatusCode)
	fmt.Println("Printer deleted successfully")
}
