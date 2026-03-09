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
		log.Fatalf("Failed to create client: %v", err)
	}

	result, resp, err := client.JamfProAPI.Enrollment.ListAccessGroupsV3(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to list ADUE access groups: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode())
	fmt.Printf("Total Count: %d\n", result.TotalCount)

	for i, group := range result.Results {
		fmt.Printf("\nAccess Group #%d:\n", i+1)
		fmt.Printf("  ID: %s\n", group.ID)
		fmt.Printf("  Name: %s\n", group.Name)
		fmt.Printf("  Group ID: %s\n", group.GroupID)
		fmt.Printf("  Enterprise Enrollment: %t\n", group.EnterpriseEnrollmentEnabled)
		fmt.Printf("  Personal Enrollment: %t\n", group.PersonalEnrollmentEnabled)
		fmt.Printf("  ADUE Enabled: %t\n", group.AccountDrivenUserEnrollmentEnabled)
	}
}
