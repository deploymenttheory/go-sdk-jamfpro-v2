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
	result, _, err := sdk.JamfProAPI.UserSessions.GetActiveV2(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}
