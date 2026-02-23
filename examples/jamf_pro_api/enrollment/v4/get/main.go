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

	result, resp, err := client.Enrollment.GetV4(context.Background())
	if err != nil {
		log.Fatalf("Failed to get enrollment settings: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Install Single Profile: %t\n", result.InstallSingleProfile)
	fmt.Printf("Signing MDM Profile Enabled: %t\n", result.SigningMdmProfileEnabled)
	fmt.Printf("Restrict Re-enrollment: %t\n", result.RestrictReenrollment)
	fmt.Printf("macOS Enterprise Enrollment Enabled: %t\n", result.MacOsEnterpriseEnrollmentEnabled)
	fmt.Printf("Management Username: %s\n", result.ManagementUsername)
	fmt.Printf("Create Management Account: %t\n", result.CreateManagementAccount)
	fmt.Printf("iOS Enterprise Enrollment Enabled: %t\n", result.IosEnterpriseEnrollmentEnabled)
	fmt.Printf("iOS Personal Enrollment Enabled: %t\n", result.IosPersonalEnrollmentEnabled)
}
