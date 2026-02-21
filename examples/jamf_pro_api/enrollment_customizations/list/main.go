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

	// List all enrollment customizations
	customizations, resp, err := client.EnrollmentCustomizations.ListV2(ctx, nil)
	if err != nil {
		log.Fatalf("Error listing enrollment customizations: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total enrollment customizations: %d\n", customizations.TotalCount)
	for _, c := range customizations.Results {
		fmt.Printf("- ID: %s, Name: %s, Description: %s\n", c.ID, c.DisplayName, c.Description)
	}

	// List with pagination and filtering
	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "10",
		"sort":      "id:asc",
	}

	paginatedList, paginatedResp, err := client.EnrollmentCustomizations.ListV2(ctx, rsqlQuery)
	if err != nil {
		log.Fatalf("Error listing enrollment customizations with pagination: %v", err)
	}

	fmt.Printf("\nPaginated Status: %d\n", paginatedResp.StatusCode)
	fmt.Printf("Paginated Total: %d\n", paginatedList.TotalCount)
}
