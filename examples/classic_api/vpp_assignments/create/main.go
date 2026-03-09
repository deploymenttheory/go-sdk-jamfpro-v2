package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/vpp_assignments"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Replace vppAdminAccountID and vppAdminAccountName with an existing VPP account
	vppAdminAccountID := 1
	vppAdminAccountName := "My VPP Account"

	createReq := &vpp_assignments.Resource{
		General: vpp_assignments.SubsetGeneral{
			Name:                "go-sdk-v2-vpp-assignment",
			VPPAdminAccountID:   vppAdminAccountID,
			VPPAdminAccountName: vppAdminAccountName,
		},
		Scope: vpp_assignments.SubsetScope{
			AllJSSUsers: false,
		},
	}

	created, _, err := jamfClient.ClassicAPI.VppAssignments.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("VPP Assignment Created: ID=%d\n", created.ID)
}
