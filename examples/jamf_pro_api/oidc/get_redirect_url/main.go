package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/oidc"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	req := &oidc.RequestOIDCRedirectURL{
		OriginalURL:  "https://yourserver.jamfcloud.com/",
		EmailAddress: "user@example.com",
	}

	result, _, err := jamfClient.JamfProAPI.Oidc.GetRedirectURLV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Redirect URL: %s\n", result.RedirectURL)

	// idpRedirects was added in Jamf Pro 11.30. logoUrl is nullable — it is null
	// when no custom icon is configured (Jamf ID connections always return null).
	fmt.Printf("\nIdentity Providers: %d\n", len(result.IdpRedirects))
	for _, idp := range result.IdpRedirects {
		fmt.Printf("  %s (%s)\n", idp.IdpName, idp.IdpType)
		fmt.Printf("    Redirect: %s\n", idp.RedirectURL)
		if idp.LogoURL != nil {
			fmt.Printf("    Logo: %s\n", *idp.LogoURL)
		} else {
			fmt.Printf("    Logo: (none configured)\n")
		}
	}
}
