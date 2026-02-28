package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Branding
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • DownloadBrandingImageV1(ctx, id) - Downloads a self-service branding image
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Reason: Service is download-only; requires an existing branding image
//     -- Tests: TestAcceptance_Branding_download_image_v1
//     -- Flow: List iOS branding configs → find one with iconId → download image
//
// Notes
// -----------------------------------------------------------------------------
//   • Branding images are associated with self-service branding configurations
//   • Image IDs come from the iconId field of iOS or macOS branding configs
//   • Test skips gracefully if no branding configs with icon IDs exist
//
// =============================================================================

func TestAcceptance_Branding_download_image_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Branding
	ctx := context.Background()

	// Find an existing iOS branding config that has an icon image
	list, _, err := acc.Client.SelfServiceBrandingIOS.ListV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, list)

	var imageID string
	for _, b := range list.Results {
		if b.IconId != nil && *b.IconId > 0 {
			imageID = fmt.Sprintf("%d", *b.IconId)
			break
		}
	}

	// Also check macOS branding configs if iOS had none
	if imageID == "" {
		macList, _, err := acc.Client.SelfServiceBrandingMacOS.List(ctx, nil)
		if err == nil && macList != nil {
			for _, b := range macList.Results {
				if b.IconId != nil && *b.IconId > 0 {
					imageID = fmt.Sprintf("%d", *b.IconId)
					break
				}
			}
		}
	}

	if imageID == "" {
		t.Skip("No self-service branding configurations with icon images found; skipping DownloadBrandingImage")
	}

	acc.LogTestStage(t, "Download", "Downloading branding image ID=%s", imageID)

	data, resp, err := svc.DownloadBrandingImageV1(ctx, imageID)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, data, "downloaded image should not be empty")

	acc.LogTestSuccess(t, "DownloadBrandingImageV1: ID=%s bytes=%d", imageID, len(data))
}
