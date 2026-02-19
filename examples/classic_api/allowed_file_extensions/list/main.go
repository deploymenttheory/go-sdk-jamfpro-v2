// Package main demonstrates ListAllowedFileExtensions — returns all allowed file extensions from the Classic API.
//
// Run with: go run ./examples/classic_api/allowed_file_extensions/list
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	list, resp, err := client.AllowedFileExtensions.ListAllowedFileExtensions(context.Background())
	if err != nil {
		log.Fatalf("ListAllowedFileExtensions failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total allowed file extensions: %d\n", list.Size)
	for _, e := range list.Results {
		fmt.Printf("  ID=%-5d  Extension=%s\n", e.ID, e.Extension)
	}
}
