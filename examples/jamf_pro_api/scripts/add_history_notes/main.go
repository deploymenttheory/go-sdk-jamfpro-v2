// Package main demonstrates AddScriptHistoryNotes — adds notes to a script's history.
//
// Run with: go run ./examples/jamf_pro_api/scripts/add_history_notes
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a script, adds a note, then deletes it.
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

	// Create a script
	createReq := &scripts.RequestScript{
		Name:           fmt.Sprintf("example-history-%d", time.Now().UnixMilli()),
		Priority:       scripts.ScriptPriorityAfter,
		ScriptContents: "#!/bin/bash\necho 'history note example'",
	}
	created, _, err := client.Scripts.CreateScript(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateScript failed: %v", err)
	}
	id := created.ID
	fmt.Printf("Created script ID: %s\n", id)

	// Add a history note
	noteReq := &scripts.AddScriptHistoryNotesRequest{
		Note: fmt.Sprintf("Example note added at %s", time.Now().Format(time.RFC3339)),
	}
	resp, err := client.Scripts.AddScriptHistoryNotes(ctx, id, noteReq)
	if err != nil {
		_, _ = client.Scripts.DeleteScriptByID(ctx, id)
		log.Fatalf("AddScriptHistoryNotes failed: %v", err)
	}

	fmt.Printf("Status: %d (201 = success)\n", resp.StatusCode)
	fmt.Println("History note added")

	// Fetch history to verify
	history, _, err := client.Scripts.GetScriptHistory(ctx, id, nil)
	if err == nil {
		fmt.Printf("History entries: %d\n", history.TotalCount)
		for _, e := range history.Results {
			if e.Note != "" {
				fmt.Printf("  Note: %s (by %s)\n", e.Note, e.Username)
			}
		}
	}

	_, _ = client.Scripts.DeleteScriptByID(ctx, id)
	fmt.Println("Cleanup: script deleted")
}
