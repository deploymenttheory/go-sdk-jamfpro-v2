// Package main demonstrates CreateScript — creates a new script.
//
// Run with: go run ./examples/jamf_pro_api/scripts/create
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

	req := &scripts.RequestScript{
		Name:           fmt.Sprintf("example-script-%d", time.Now().UnixMilli()),
		Priority:       scripts.ScriptPriorityAfter,
		Info:           "Example script created by SDK",
		Notes:          "For testing purposes only",
		ScriptContents: "#!/bin/bash\necho 'Hello from Jamf Pro SDK example'",
	}

	result, resp, err := client.Scripts.CreateScript(ctx, req)
	if err != nil {
		log.Fatalf("CreateScript failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created script ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	// Cleanup: delete the created script
	if _, err := client.Scripts.DeleteScriptByID(ctx, result.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: script deleted")
	}
}
