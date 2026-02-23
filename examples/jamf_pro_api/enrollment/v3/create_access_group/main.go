package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	request := &enrollment.ResourceAccountDrivenUserEnrollmentAccessGroup{
		GroupID:                            "100",
		LdapServerID:                       "1",
		Name:                               "IT Department ADUE Group",
		SiteID:                             "1",
		EnterpriseEnrollmentEnabled:        true,
		PersonalEnrollmentEnabled:          false,
		AccountDrivenUserEnrollmentEnabled: true,
		RequireEula:                        true,
	}

	result, resp, err := client.Enrollment.CreateAccessGroupV3(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to create ADUE access group: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Created Access Group ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)
}
