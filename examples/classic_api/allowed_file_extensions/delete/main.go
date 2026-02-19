// Package main demonstrates DeleteAllowedFileExtensionByID — removes an allowed file extension via the Classic API.
//
// Run with: go run ./examples/classic_api/allowed_file_extensions/delete
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

	// Create an allowed file extension to delete
	createReq := &allowed_file_extensions.RequestAllowedFileExtension{
		Extension: fmt.Sprintf("del%d", time.Now().UnixMilli()%100000),
	}
	created, _, err := client.AllowedFileExtensions.CreateAllowedFileExtension(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateAllowedFileExtension failed: %v", err)
	}
	fmt.Printf("Created allowed file extension ID: %d\n", created.ID)

	resp, err := client.AllowedFileExtensions.DeleteAllowedFileExtensionByID(ctx, created.ID)
	if err != nil {
		log.Fatalf("DeleteAllowedFileExtensionByID failed: %v", err)
	}

	fmt.Printf("Status: %d (200 = success)\n", resp.StatusCode)
	fmt.Println("Allowed file extension deleted successfully")
}
