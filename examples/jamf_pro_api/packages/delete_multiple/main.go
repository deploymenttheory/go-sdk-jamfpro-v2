// Package main demonstrates DeletePackagesByIDV1 - deletes multiple packages by their IDs.
//
// Run with: go run ./examples/jamf_pro_api/packages/delete_multiple
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates two packages then bulk deletes them.
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

	// Create two packages
	ids := make([]string, 0, 2)
	for i := 0; i < 2; i++ {
		req := &packages.RequestPackage{
			PackageName:           fmt.Sprintf("example-bulk-%d-%d", i, time.Now().UnixMilli()),
			FileName:              "example.pkg",
			CategoryID:            "-1",
			Info:                  "Bulk delete test",
			Priority:              5,
			FillUserTemplate:      packages.BoolPtr(false),
			RebootRequired:        packages.BoolPtr(false),
			OSInstall:            packages.BoolPtr(false),
			SuppressUpdates:       packages.BoolPtr(false),
			SuppressFromDock:      packages.BoolPtr(false),
			SuppressEula:          packages.BoolPtr(false),
			SuppressRegistration:  packages.BoolPtr(false),
		}
		created, _, err := client.Packages.CreatePackageV1(ctx, req)
		if err != nil {
			log.Fatalf("CreatePackageV1 %d failed: %v", i, err)
		}
		ids = append(ids, created.ID)
		fmt.Printf("Created package ID: %s\n", created.ID)
	}

	// Bulk delete
	bulkReq := &packages.DeletePackagesByIDRequest{IDs: ids}
	resp, err := client.Packages.DeletePackagesByIDV1(ctx, bulkReq)
	if err != nil {
		log.Fatalf("DeletePackagesByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d (204 = success)\n", resp.StatusCode)
	fmt.Printf("Deleted %d packages: %v\n", len(ids), ids)
}
