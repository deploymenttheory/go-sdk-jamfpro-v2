package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Replace with actual branding image file path
	imagePath := "/path/to/your/branding.png"

	result, resp, err := client.SelfServiceBrandingUpload.UploadFromFile(ctx, imagePath)
	if err != nil {
		log.Fatalf("Error uploading branding image: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Uploaded Image URL: %s\n", result.URL)
}
