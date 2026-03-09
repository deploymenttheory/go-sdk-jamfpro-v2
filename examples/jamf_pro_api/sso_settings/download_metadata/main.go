package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	// Initialize Jamf Pro client using environment variables
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Download the SAML metadata file
	metadata, resp, err := client.JamfProAPI.SsoSettings.DownloadMetadataV3(ctx)
	if err != nil {
		log.Fatalf("Error downloading SAML metadata: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode())
	fmt.Printf("Downloaded metadata size: %d bytes\n", len(metadata))

	// Save the metadata to a file
	outputFile := "saml_metadata.xml"
	err = os.WriteFile(outputFile, metadata, 0644)
	if err != nil {
		log.Fatalf("Error writing metadata to file: %v", err)
	}

	// Get absolute path for display
	absPath, err := filepath.Abs(outputFile)
	if err != nil {
		absPath = outputFile
	}

	fmt.Printf("\nSAML metadata downloaded successfully:\n")
	fmt.Printf("  File: %s\n", absPath)
	fmt.Printf("  Size: %d bytes\n", len(metadata))
}
