// Package main demonstrates CreateAttachment with ResourceIDTypeName — uploads a file
// to a policy by name via the Classic API.
//
// Run with: go run ./examples/classic_api/file_uploads/create_attachment_by_name
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
//
// Usage: pass policy name and file path as arguments:
//
//	go run . "My Policy" /path/to/file.pdf
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/file_uploads"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <policy_name> <file_path>", os.Args[0])
	}

	policyName := os.Args[1]
	filePath := os.Args[2]

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	resp, err := client.ClassicFileUploads.CreateAttachment(ctx, "policies", file_uploads.ResourceIDTypeName, policyName, filePath, false)
	if err != nil {
		log.Fatalf("CreateAttachment failed: %v", err)
	}

	fmt.Printf("File uploaded to policy %q successfully. Status: %d\n", policyName, resp.StatusCode)
}
