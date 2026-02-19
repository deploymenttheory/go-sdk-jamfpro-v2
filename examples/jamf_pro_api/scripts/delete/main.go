// Package main demonstrates DeleteScriptByID — removes a script by ID.
//
// Run with: go run ./examples/jamf_pro_api/scripts/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a script then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/scripts"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Create a script to delete
	createReq := &scripts.RequestScript{
		Name:           fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
		Priority:       scripts.ScriptPriorityAfter,
		ScriptContents: "#!/bin/bash\necho 'to be deleted'",
	}
	created, _, err := client.Scripts.CreateScript(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateScript failed: %v", err)
	}
	id := created.ID
	fmt.Printf("Created script ID: %s\n", id)

	resp, err := client.Scripts.DeleteScriptByID(ctx, id)
	if err != nil {
		log.Fatalf("DeleteScriptByID failed: %v", err)
	}

	fmt.Printf("Status: %d (204 = success)\n", resp.StatusCode)
	fmt.Println("Script deleted successfully")
}
