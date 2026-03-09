package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/enrollment"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get current settings first
	current, _, err := client.JamfProAPI.Enrollment.GetV4(context.Background())
	if err != nil {
		log.Fatalf("Failed to get current enrollment settings: %v", err)
	}

	// Modify settings as needed
	request := &enrollment.ResourceEnrollment{
		InstallSingleProfile:             current.InstallSingleProfile,
		SigningMdmProfileEnabled:         current.SigningMdmProfileEnabled,
		MdmSigningCertificate:            current.MdmSigningCertificate,
		RestrictReenrollment:             true, // Enable re-enrollment restriction
		FlushLocationInformation:         true,
		FlushLocationHistoryInformation:  false,
		FlushPolicyHistory:               false,
		FlushExtensionAttributes:         false,
		FlushSoftwareUpdatePlans:         false,
		MacOsEnterpriseEnrollmentEnabled: true,
		ManagementUsername:               current.ManagementUsername,
		CreateManagementAccount:          current.CreateManagementAccount,
		HideManagementAccount:            current.HideManagementAccount,
		LaunchSelfService:                true,
		IosEnterpriseEnrollmentEnabled:   true,
		IosPersonalEnrollmentEnabled:     false,
	}

	result, resp, err := client.JamfProAPI.Enrollment.UpdateV4(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to update enrollment settings: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode())
	fmt.Printf("Updated - Restrict Re-enrollment: %t\n", result.RestrictReenrollment)
	fmt.Println("Enrollment settings updated successfully")
}
