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
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Replace with actual enrollment customization ID
	customizationID := "1"

	prestages, resp, err := client.EnrollmentCustomizations.GetPrestagesV2(ctx, customizationID)
	if err != nil {
		log.Fatalf("Error getting prestages: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total prestages using this customization: %d\n", len(prestages.Dependencies))

	for _, prestage := range prestages.Dependencies {
		fmt.Printf("\nPrestage:\n")
		fmt.Printf("  Name: %s\n", prestage.Name)
		fmt.Printf("  Human Readable Name: %s\n", prestage.HumanReadableName)
		fmt.Printf("  Hyperlink: %s\n", prestage.Hyperlink)
	}
}
