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

	// Replace with actual image file path
	imagePath := "/path/to/your/icon.png"

	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalf("Error opening image file: %v", err)
	}
	defer file.Close()

	// Get file info for size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file info: %v", err)
	}

	result, resp, err := client.EnrollmentCustomizations.UploadImageV2(ctx, file, fileInfo.Size(), fileInfo.Name())
	if err != nil {
		log.Fatalf("Error uploading image: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Uploaded Image ID: %s\n", result.ID)
	fmt.Printf("Image URL: %s\n", result.URL)
}
