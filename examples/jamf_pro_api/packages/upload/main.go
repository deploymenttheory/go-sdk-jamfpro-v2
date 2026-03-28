package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/packages"
)

const (
	firefoxURL = "https://ftp.mozilla.org/pub/firefox/releases/147.0/mac/en-GB/Firefox%20147.0.pkg"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro-sdkv2/clientconfig.json"
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig, jamfpro.WithTimeout(15*time.Minute))
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Download Firefox 147.0 to a temp file.
	fmt.Println("Downloading Firefox 147.0...")
	tmpDir, err := os.MkdirTemp("", "jamf-pkg-upload-*")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	pkgPath := filepath.Join(tmpDir, "Firefox_147.0.pkg")
	dlCtx, dlCancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer dlCancel()

	req, err := http.NewRequestWithContext(dlCtx, http.MethodGet, firefoxURL, nil)
	if err != nil {
		log.Fatalf("Failed to create download request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to download package: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Download returned status %d", resp.StatusCode)
	}

	out, err := os.Create(pkgPath)
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
	}
	if _, err = io.Copy(out, resp.Body); err != nil {
		out.Close()
		log.Fatalf("Failed to write package: %v", err)
	}
	out.Close()
	fmt.Printf("Downloaded to: %s\n", pkgPath)

	// Create metadata + upload + verify SHA-512.
	ctx := context.Background()
	packageName := fmt.Sprintf("Firefox 147.0 (%s)", time.Now().Format("20060102-150405"))
	createReq := &packages.RequestPackage{
		PackageName:          packageName,
		FileName:             "Firefox_147.0.pkg",
		CategoryID:           "-1",
		Priority:             10,
		FillUserTemplate:     packages.BoolPtr(false),
		FillExistingUsers:    packages.BoolPtr(false),
		RebootRequired:       packages.BoolPtr(false),
		OSInstall:            packages.BoolPtr(false),
		SuppressUpdates:      packages.BoolPtr(false),
		SuppressFromDock:     packages.BoolPtr(false),
		SuppressEula:         packages.BoolPtr(false),
		SuppressRegistration: packages.BoolPtr(false),
	}

	fmt.Println("Creating package metadata, uploading file, and verifying SHA-512...")
	created, _, err := jamfClient.JamfProAPI.Packages.CreateAndUpload(ctx, pkgPath, createReq)
	if err != nil {
		fmt.Printf("CreateAndUpload error: %v\n", err)
		if created != nil {
			fmt.Printf("Package was created with ID=%s before error — you may need to clean it up\n", created.ID)
		}
		return
	}

	fmt.Printf("Package created and uploaded successfully: ID=%s\n", created.ID)
}
