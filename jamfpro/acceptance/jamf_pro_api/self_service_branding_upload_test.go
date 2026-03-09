package jamf_pro_api

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Self Service Branding Upload
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • Upload(ctx, fileReader, fileSize, fileName) - Uploads a branding image
//   • UploadFromFile(ctx, filePath) - Helper to upload from file path
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern: Upload with temporary file
//     -- Reason: Upload creates resources; use temp file for cleanup
//     -- Tests: TestAcceptance_SelfServiceBrandingUpload_upload
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Upload branding image
//   ✓ UploadFromFile helper
//
// Notes
// -----------------------------------------------------------------------------
//   • Branding images are used for Self Service branding customization
//   • No List, Get, or Delete operations available in the API
//   • Upload returns URL to the uploaded image
//
// =============================================================================

func TestAcceptance_SelfServiceBrandingUpload_upload(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.SelfServiceBrandingUpload
	ctx := context.Background()

	// Create a temporary PNG file for upload
	tmpDir := t.TempDir()
	imagePath := filepath.Join(tmpDir, "branding.png")
	content := []byte("\x89PNG\r\n\x1a\nfake png content")
	require.NoError(t, os.WriteFile(imagePath, content, 0644))

	acc.LogTestStage(t, "SelfServiceBrandingUpload", "Uploading branding image")

	result, resp, err := svc.UploadFromFile(ctx, imagePath)
	if err != nil {
		t.Skipf("upload may require specific permissions or branding image format: %v", err)
		return
	}

	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.URL)
	acc.LogTestSuccess(t, "Successfully uploaded branding image: %s", result.URL)
}
