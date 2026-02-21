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

	// Replace with actual image ID
	imageID := "123"

	imageData, resp, err := client.EnrollmentCustomizations.GetImageByIdV2(ctx, imageID)
	if err != nil {
		log.Fatalf("Error downloading image: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Downloaded image size: %d bytes\n", len(imageData))

	// Save the image to a file
	outputPath := "downloaded_enrollment_icon.png"
	err = os.WriteFile(outputPath, imageData, 0644)
	if err != nil {
		log.Fatalf("Error saving image file: %v", err)
	}

	fmt.Printf("Image saved to: %s\n", outputPath)
}
