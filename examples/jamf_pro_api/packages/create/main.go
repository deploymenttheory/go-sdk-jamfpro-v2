// Package main demonstrates CreatePackageV1 - creates a new package (metadata only, no file upload).
//
// Run with: go run ./examples/jamf_pro_api/packages/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/packages"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &packages.RequestPackage{
		PackageName:           fmt.Sprintf("example-package-%d", time.Now().UnixMilli()),
		FileName:              "example.pkg",
		CategoryID:            "-1",
		Info:                  "Example package created via SDK",
		Notes:                 "Metadata only - use UploadPackageV1 to add file",
		Priority:              10,
		FillUserTemplate:      packages.BoolPtr(true),
		FillExistingUsers:     packages.BoolPtr(false),
		RebootRequired:        packages.BoolPtr(false),
		OSInstall:             packages.BoolPtr(false),
		SuppressUpdates:       packages.BoolPtr(false),
		SuppressFromDock:      packages.BoolPtr(false),
		SuppressEula:          packages.BoolPtr(false),
		SuppressRegistration: packages.BoolPtr(false),
	}

	result, resp, err := client.Packages.CreatePackageV1(ctx, req)
	if err != nil {
		log.Fatalf("CreatePackageV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created package ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	if _, err := client.Packages.DeletePackageByIDV1(ctx, result.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: package deleted")
	}
}
