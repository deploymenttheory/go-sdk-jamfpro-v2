package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/file_share_distribution_points"
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

	dpName := "Main File Share DP" // Replace with the desired distribution point name
	updateReq := &file_share_distribution_points.RequestFileShareDistributionPoint{
		Name:                     "go-sdk-v2-file-share-dp-updated",
		IsMaster:                 false,
		LocalPath:                "/path/to/share",
		ConnectionType:           "SMB",
		ShareName:                "JamfShare",
		SharePort:                445,
		HTTPDownloadsEnabled:     true,
		HTTPURL:                  "http://192.168.1.100:8080",
		NoAuthenticationRequired: false,
		UsernamePasswordRequired: true,
	}

	updated, _, err := jamfClient.ClassicFileShareDistributionPoints.UpdateByName(context.Background(), dpName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("File Share Distribution Point Updated: ID=%d\n", updated.ID)
}
