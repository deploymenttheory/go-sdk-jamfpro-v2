// Package main demonstrates DownloadScriptByIDV1 — downloads a script's contents as plain text.
//
// Run with: go run ./examples/jamf_pro_api/scripts/download
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set SCRIPT_ID or uses first from list.
package main

import (
	"context"
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
	id := os.Getenv("SCRIPT_ID")
	if id == "" {
		list, _, err := client.Scripts.ListScriptsV1(ctx, map[string]string{"page": "0", "page-size": "1"})
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set SCRIPT_ID or ensure at least one script exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first script ID: %s\n", id)
	}

	data, resp, err := client.Scripts.DownloadScriptByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("DownloadScriptByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Content length: %d bytes\n", len(data))
	fmt.Println("--- Script contents ---")
	fmt.Println(string(data))
}
