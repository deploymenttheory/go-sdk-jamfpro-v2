// Package main demonstrates GetPrinterByID — retrieves a single printer by ID.
//
// Run with: go run ./examples/classic_api/printers/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set PRINTER_ID or uses first from list.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	var id int
	if raw := os.Getenv("PRINTER_ID"); raw != "" {
		id, err = strconv.Atoi(raw)
		if err != nil {
			log.Fatalf("invalid PRINTER_ID %q: %v", raw, err)
		}
	} else {
		list, _, err := client.Printers.ListPrinters(ctx)
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set PRINTER_ID or ensure at least one printer exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first printer ID: %d\n", id)
	}

	printer, resp, err := client.Printers.GetPrinterByID(ctx, id)
	if err != nil {
		log.Fatalf("GetPrinterByID failed: %v", err)
	}

	fmt.Printf("Status:   %d\n", resp.StatusCode)
	fmt.Printf("ID:       %d\n", printer.ID)
	fmt.Printf("Name:     %s\n", printer.Name)
	fmt.Printf("CUPSName: %s\n", printer.CUPSName)
	fmt.Printf("URI:      %s\n", printer.URI)
	fmt.Printf("Location: %s\n", printer.Location)
	fmt.Printf("Model:    %s\n", printer.Model)
}
