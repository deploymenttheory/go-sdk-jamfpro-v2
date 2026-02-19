// Package main demonstrates UpdateScriptByIDV1 — updates an existing script.
//
// Run with: go run ./examples/jamf_pro_api/scripts/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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

	// Create a script to update
	createReq := &scripts.RequestScript{
		Name:           fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
		Priority:       scripts.ScriptPriorityAfter,
		ScriptContents: "#!/bin/bash\necho 'original'",
	}
	created, _, err := client.Scripts.CreateScriptV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateScriptV1 failed: %v", err)
	}
	id := created.ID

	// Update the script
	updateReq := &scripts.RequestScript{
		Name:           fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
		Priority:       scripts.ScriptPriorityBefore,
		Info:           "Updated by SDK example",
		ScriptContents: "#!/bin/bash\necho 'updated'",
	}
	result, resp, err := client.Scripts.UpdateScriptByIDV1(ctx, id, updateReq)
	if err != nil {
		_, _ = client.Scripts.DeleteScriptByIDV1(ctx, id)
		log.Fatalf("UpdateScriptByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated script ID: %s\n", result.ID)
	fmt.Printf("Name: %s\n", result.Name)
	fmt.Printf("Priority: %s\n", result.Priority)

	_, _ = client.Scripts.DeleteScriptByIDV1(ctx, id)
	fmt.Println("Cleanup: script deleted")
}
