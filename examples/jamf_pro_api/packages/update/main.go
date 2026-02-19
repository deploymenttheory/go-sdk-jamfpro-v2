// Package main demonstrates UpdatePackageByIDV1 - updates an existing package.
//
// Run with: go run ./examples/jamf_pro_api/packages/update
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

	createReq := &packages.RequestPackage{
		PackageName:           fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
		FileName:              "example.pkg",
		CategoryID:            "-1",
		Info:                  "Original metadata",
		Notes:                 "Before update",
		Priority:              5,
		FillUserTemplate:      packages.BoolPtr(false),
		RebootRequired:        packages.BoolPtr(false),
		OSInstall:             packages.BoolPtr(false),
		SuppressUpdates:       packages.BoolPtr(false),
		SuppressFromDock:      packages.BoolPtr(false),
		SuppressEula:          packages.BoolPtr(false),
		SuppressRegistration: packages.BoolPtr(false),
	}
	created, _, err := client.Packages.CreatePackageV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreatePackageV1 failed: %v", err)
	}
	id := created.ID

	fetched, _, err := client.Packages.GetPackageByIDV1(ctx, id)
	if err != nil {
		_, _ = client.Packages.DeletePackageByIDV1(ctx, id)
		log.Fatalf("GetPackageByIDV1 failed: %v", err)
	}
	updateReq := fetched
	updateReq.PackageName = fmt.Sprintf("example-updated-%d", time.Now().UnixMilli())
	updateReq.Info = "Updated metadata via SDK"
	updateReq.Notes = "After update"
	updateReq.Priority = 15
	updateReq.FillUserTemplate = packages.BoolPtr(true)
	updateReq.FillExistingUsers = packages.BoolPtr(true)

	result, resp, err := client.Packages.UpdatePackageByIDV1(ctx, id, updateReq)
	if err != nil {
		_, _ = client.Packages.DeletePackageByIDV1(ctx, id)
		log.Fatalf("UpdatePackageByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated package ID: %s PackageName=%q\n", result.ID, result.PackageName)

	fetched2, _, _ := client.Packages.GetPackageByIDV1(ctx, id)
	if fetched2 != nil {
		fmt.Printf("Verified: packageName=%q info=%s\n", fetched2.PackageName, fetched2.Info)
	}

	_, _ = client.Packages.DeletePackageByIDV1(ctx, id)
	fmt.Println("Cleanup: package deleted")
}
