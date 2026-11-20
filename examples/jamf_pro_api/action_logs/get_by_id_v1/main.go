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
	result, _, err := sdk.JamfProAPI.ActionLogs.GetByIDV1(ctx, os.Getenv("ACTION_LOG_ID"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}
