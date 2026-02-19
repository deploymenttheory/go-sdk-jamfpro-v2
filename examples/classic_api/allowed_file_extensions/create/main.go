// Package main demonstrates CreateAllowedFileExtension — creates a new allowed file extension via the Classic API.
//
// Run with: go run ./examples/classic_api/allowed_file_extensions/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates an allowed file extension then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/allowed_file_extensions"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &allowed_file_extensions.RequestAllowedFileExtension{
		Extension: fmt.Sprintf("ex%d", time.Now().UnixMilli()%100000),
	}

	created, resp, err := client.AllowedFileExtensions.CreateAllowedFileExtension(ctx, req)
	if err != nil {
		log.Fatalf("CreateAllowedFileExtension failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created allowed file extension ID: %d\n", created.ID)
	fmt.Printf("Extension: %s\n", created.Extension)

	// Cleanup: delete the created allowed file extension
	if _, err := client.AllowedFileExtensions.DeleteAllowedFileExtensionByID(ctx, created.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: allowed file extension deleted")
	}
}
