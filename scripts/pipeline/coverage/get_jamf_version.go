package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

const defaultVersion = "unknown"

func main() {

	requiredEnvVars := []string{"INSTANCE_DOMAIN", "AUTH_METHOD", "CLIENT_ID", "CLIENT_SECRET"}
	missingVars := []string{}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			missingVars = append(missingVars, envVar)
		}
	}

	if len(missingVars) > 0 {
		fmt.Fprintf(os.Stderr, "Required environment variables not set: %v\n", missingVars)
		fmt.Println(defaultVersion)
		os.Exit(0)
	}

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create client: %v\n", err)
		fmt.Println(defaultVersion)
		os.Exit(0)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	version, _, err := client.JamfProVersion.GetV1(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get version: %v\n", err)
		fmt.Println(defaultVersion)
		os.Exit(0)
	}

	if version.Version != nil && *version.Version != "" {
		fmt.Println(*version.Version)
	} else {
		fmt.Fprintf(os.Stderr, "Version field is empty\n")
		fmt.Println(defaultVersion)
	}
}
