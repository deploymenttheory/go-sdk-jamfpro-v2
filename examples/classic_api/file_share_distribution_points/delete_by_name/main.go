package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
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

	dpName := "Main File Share DP" // Replace with the desired distribution point name to delete
	_, err = jamfClient.ClassicAPI.FileShareDistributionPoints.DeleteByName(context.Background(), dpName)
	if err != nil {
		fmt.Printf("Error deleting file share distribution point by name: %v\n", err)
		return
	}
	fmt.Println("File share distribution point by name deleted successfully")
}
