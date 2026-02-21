package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	authConfig := client.AuthConfigFromEnv()
	if err := authConfig.Validate(); err != nil {
		log.Fatalf("Invalid auth config: %v", err)
	}

	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()
	declarationUUID := "550e8400-e29b-41d4-a716-446655440000"

	declaration, resp, err := jamfClient.DSSDeclarations.GetByUUIDV1(ctx, declarationUUID)
	if err != nil {
		log.Fatalf("Failed to get DSS declaration: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Declarations:\n")
	for _, decl := range declaration.Declarations {
		fmt.Printf("  - UUID: %s\n", decl.UUID)
		fmt.Printf("    Type: %s\n", decl.Type)
		fmt.Printf("    Group: %s\n", decl.Group)
		fmt.Printf("    Payload: %s\n", decl.PayloadJson)
	}
}
