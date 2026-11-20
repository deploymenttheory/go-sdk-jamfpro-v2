package main

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("provide clientconfig.json path")
	}
	auth, err := jamfpro.LoadAuthConfigFromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	sdk, err := jamfpro.NewClient(auth)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	result, err := sdk.JamfProAPI.DirectoryCache.SyncGroupV1(ctx, os.Getenv("DIRECTORY_SERVICE_ID"), os.Getenv("GROUP_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}
