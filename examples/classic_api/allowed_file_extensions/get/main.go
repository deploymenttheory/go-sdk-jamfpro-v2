// Package main demonstrates GetAllowedFileExtensionByID — retrieves a single allowed file extension by ID.
//
// Run with: go run ./examples/classic_api/allowed_file_extensions/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set ALLOWED_FILE_EXTENSION_ID or uses first from list.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	var id int
	if raw := os.Getenv("ALLOWED_FILE_EXTENSION_ID"); raw != "" {
		id, err = strconv.Atoi(raw)
		if err != nil {
			log.Fatalf("invalid ALLOWED_FILE_EXTENSION_ID %q: %v", raw, err)
		}
	} else {
		list, _, err := client.AllowedFileExtensions.ListAllowedFileExtensions(ctx)
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set ALLOWED_FILE_EXTENSION_ID or ensure at least one allowed file extension exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first allowed file extension ID: %d\n", id)
	}

	ext, resp, err := client.AllowedFileExtensions.GetAllowedFileExtensionByID(ctx, id)
	if err != nil {
		log.Fatalf("GetAllowedFileExtensionByID failed: %v", err)
	}

	fmt.Printf("Status:    %d\n", resp.StatusCode)
	fmt.Printf("ID:        %d\n", ext.ID)
	fmt.Printf("Extension: %s\n", ext.Extension)
}
