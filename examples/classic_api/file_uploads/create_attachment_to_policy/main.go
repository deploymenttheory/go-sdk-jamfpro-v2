// Package main demonstrates CreateAttachment — uploads a file to a policy via the Classic API.
//
// Run with: go run ./examples/classic_api/file_uploads/create_attachment_to_policy
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
//
// Usage: pass policy ID and file path as arguments, or use defaults:
//
//	go run . 123 /path/to/file.pdf
//	go run .  # uses policy ID 1 and creates a temp file
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/file_uploads"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	policyID := "1"
	filePath := ""

	if len(os.Args) >= 3 {
		policyID = os.Args[1]
		filePath = os.Args[2]
	} else if len(os.Args) == 2 {
		policyID = os.Args[1]
	}

	if filePath == "" {
		tmpDir := os.TempDir()
		filePath = filepath.Join(tmpDir, "jamf-file-upload-example.txt")
		if err := os.WriteFile(filePath, []byte("Example file attachment for Jamf Pro policy\n"), 0644); err != nil {
			log.Fatalf("failed to create temp file: %v", err)
		}
		defer os.Remove(filePath)
	}

	ctx := context.Background()

	resp, err := client.ClassicFileUploads.CreateAttachment(ctx, "policies", file_uploads.ResourceIDTypeID, policyID, filePath, false)
	if err != nil {
		log.Fatalf("CreateAttachment failed: %v", err)
	}

	fmt.Printf("File uploaded successfully. Status: %d\n", resp.StatusCode)
}
