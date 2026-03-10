package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_invitations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
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

	createReq := &computer_invitations.ResourceComputerInvitation{
		InvitationType:              "USER_INITIATED_ENROLLMENT",
		MultipleUsersAllowed:        false,
		CreateAccountIfDoesNotExist: true,
		KeepExistingSiteMembership:  true,
		Site: &models.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		EnrollIntoSite: &computer_invitations.ComputerInvitationSubsetEnrollIntoState{
			ID:   -1,
			Name: "None",
		},
	}

	created, _, err := jamfClient.ClassicAPI.ComputerInvitations.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Computer Invitation Created: ID=%d Invitation=%s\n", created.ID, created.Invitation)
}
