// Package main demonstrates ListPrinters — returns all printers from the Classic API.
//
// Run with: go run ./examples/classic_api/printers/list
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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
		log.Fatalf("failed to create client: %v", err)
	}

	list, resp, err := client.Printers.ListPrinters(context.Background())
	if err != nil {
		log.Fatalf("ListPrinters failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total printers: %d\n", list.Size)
	for _, p := range list.Results {
		fmt.Printf("  ID=%-5d  Name=%s\n", p.ID, p.Name)
	}
}
