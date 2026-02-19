// Package main demonstrates GetScriptByID — retrieves a single script by ID.
//
// Run with: go run ./examples/jamf_pro_api/scripts/get
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
		list, _, err := client.Scripts.ListScripts(ctx, map[string]string{"page": "0", "page-size": "1"})
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set SCRIPT_ID or ensure at least one script exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first script ID: %s\n", id)
	}

	script, resp, err := client.Scripts.GetScriptByID(ctx, id)
	if err != nil {
		log.Fatalf("GetScriptByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID:             %s\n", script.ID)
	fmt.Printf("Name:           %s\n", script.Name)
	fmt.Printf("Priority:       %s\n", script.Priority)
	fmt.Printf("CategoryName:   %s\n", script.CategoryName)
	fmt.Printf("OSRequirements: %s\n", script.OSRequirements)
	fmt.Printf("Info:           %s\n", script.Info)
	if script.ScriptContents != "" {
		preview := script.ScriptContents
		if len(preview) > 100 {
			preview = preview[:100] + "..."
		}
		fmt.Printf("ScriptContents: %s\n", preview)
	}
}
