package main

// SetAllAvailableApiPriviledgesToRole creates a new API role with all available privileges
// in the Jamf Pro instance, named with the current Jamf Pro version.
//
// Required environment variables:
//   - INSTANCE_DOMAIN: Jamf Pro instance URL (e.g., https://example.jamfcloud.com)
//   - AUTH_METHOD: Authentication method ("oauth2" or "basic")
//
// For OAuth2 authentication:
//   - CLIENT_ID: OAuth2 client ID
//   - CLIENT_SECRET: OAuth2 client secret
//
// For Basic authentication:
//   - BASIC_AUTH_USERNAME: Jamf Pro username
//   - BASIC_AUTH_PASSWORD: Jamf Pro password
//
// Optional environment variables:
//   - TOKEN_REFRESH_BUFFER_SECONDS: Token refresh buffer in seconds (default: 300)
//   - HIDE_SENSITIVE_DATA: Hide sensitive data in logs (default: false)

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles"
)

func main() {
	jamfClient, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	versionInfo, _, err := jamfClient.JamfProVersion.GetV1(ctx)
	if err != nil {
		log.Fatalf("Failed to get Jamf Pro version: %v", err)
	}

	if versionInfo.Version == nil {
		log.Fatal("Received empty version information")
	}

	version := strings.TrimSpace(*versionInfo.Version)
	fmt.Printf("Found Jamf Pro version: %s\n", version)

	privileges, _, err := jamfClient.APIRolePrivileges.ListV1(ctx)
	if err != nil {
		log.Fatalf("Failed to get API privileges: %v", err)
	}

	if len(privileges.Privileges) == 0 {
		log.Fatal("No privileges found")
	}

	fmt.Printf("Found %d privileges\n", len(privileges.Privileges))

	roleName := fmt.Sprintf("all-jamfpro-privileges-%s", version)

	newRole := &api_roles.RequestAPIRole{
		DisplayName: roleName,
		Privileges:  privileges.Privileges,
	}

	createdRole, _, err := jamfClient.APIRoles.CreateV1(ctx, newRole)
	if err != nil {
		log.Fatalf("Failed to create API role: %v", err)
	}

	fmt.Printf("Successfully created role '%s' with ID: %s\n", createdRole.DisplayName, createdRole.ID)
	fmt.Printf("Total privileges assigned: %d\n", len(createdRole.Privileges))
}
