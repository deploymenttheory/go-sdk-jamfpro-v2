// Package main demonstrates DownloadV1 - downloads icon image bytes by ID.
//
// Run with: go run ./examples/jamf_pro_api/icons/download <id> [res] [scale]
// res: original, 300, or 512 (default original). scale: 0 or 1 (default 0).
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/icons/download <id> [res] [scale]")
	}
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("invalid id: %v", err)
	}
	res, scale := "original", "0"
	if len(os.Args) > 2 {
		res = os.Args[2]
	}
	if len(os.Args) > 3 {
		scale = os.Args[3]
	}

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	ctx := context.Background()

	body, resp, err := client.Icons.DownloadV1(ctx, id, res, scale)
	if err != nil {
		log.Fatalf("DownloadV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Size: %d bytes\n", resp.StatusCode, len(body))
}
