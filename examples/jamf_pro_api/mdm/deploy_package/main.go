package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mdm"
)

func main() {
	// Initialize the Jamf Pro client from environment variables
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Configure deploy package request
	// Replace with your package URL, hash, device IDs, and group ID
	req := &mdm.DeployPackageRequest{
		Manifest: mdm.PackageManifest{
			HashType:         "SHA256",
			URL:              os.Getenv("JAMF_PACKAGE_URL"), // e.g. https://example.com/pkg.dmg
			Hash:             os.Getenv("JAMF_PACKAGE_HASH"), // e.g. abc123...
			DisplayImageURL:  "",
			FullSizeImageURL: "",
			BundleID:         "com.example.app",
			BundleVersion:    "1.0.0",
			Subtitle:         "Example App",
			Title:            "Example Application",
			SizeInBytes:      1024000,
		},
		InstallAsManaged: true,
		Devices:          []int{1001, 1002}, // Replace with actual device IDs
		GroupID:          "1",               // Replace with actual group ID
	}

	if req.Manifest.URL == "" {
		req.Manifest.URL = "https://example.com/sample.dmg"
	}
	if req.Manifest.Hash == "" {
		req.Manifest.Hash = "sha256-hash-of-package"
	}

	result, resp, err := client.MDM.DeployPackage(ctx, req)
	if err != nil {
		log.Fatalf("Failed to deploy package: %v (HTTP %d)", err, resp.StatusCode)
	}

	fmt.Printf("Package deployment initiated (HTTP %d)\n", resp.StatusCode)
	fmt.Printf("Queued commands: %d\n", len(result.QueuedCommands))
	for _, qc := range result.QueuedCommands {
		fmt.Printf("  Device %d: %s\n", qc.Device, qc.CommandUUID)
	}
	if len(result.Errors) > 0 {
		fmt.Printf("Errors: %+v\n", result.Errors)
	}
}
