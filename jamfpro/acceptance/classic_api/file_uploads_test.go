package classic_api

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/file_uploads"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_FileUploads_create_attachment creates a policy, uploads a file
// attachment to it, then deletes the policy.
// =============================================================================

func TestAcceptance_FileUploads_create_attachment(t *testing.T) {
	acc.RequireClient(t)

	policySvc := acc.Client.ClassicPolicies
	fileSvc := acc.Client.ClassicFileUploads
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create a policy to attach a file to
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test policy for file attachment")

	policyName := acc.UniqueName("sdkv2_acc_acc-test-fileupload-policy")
	createReq := createMinimalPolicy(t, policyName)

	_, policyID := createPolicyWithCleanup(t, ctx, policySvc, createReq)

	// ------------------------------------------------------------------
	// 2. Create a temporary PNG file to upload
	// Policy attachments require a supported file format; PNG is accepted as an icon.
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Prepare", "Creating temporary PNG file for upload")

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "attachment.png")
	pngFile, err := os.Create(tmpFile)
	require.NoError(t, err)
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	require.NoError(t, png.Encode(pngFile, img))
	pngFile.Close()

	// ------------------------------------------------------------------
	// 3. Upload file to the policy
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Upload", "Uploading file to policy ID=%d", policyID)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	uploadResp, err := fileSvc.CreateAttachment(ctx2, "policies", file_uploads.ResourceIDTypeID, fmt.Sprintf("%d", policyID), tmpFile, false)
	require.NoError(t, err, "CreateAttachment should not return an error")
	require.NotNil(t, uploadResp)
	assert.Contains(t, []int{200, 201}, uploadResp.StatusCode)
	acc.LogTestSuccess(t, "File uploaded successfully to policy ID=%d", policyID)
}

// =============================================================================
// TestAcceptance_FileUploads_validation_errors validates error handling.
// =============================================================================

func TestAcceptance_FileUploads_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicFileUploads
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(tmpFile, []byte("test"), 0644)

	t.Run("InvalidResource", func(t *testing.T) {
		_, err := svc.CreateAttachment(context.Background(), "invalidresource", file_uploads.ResourceIDTypeID, "1", tmpFile, false)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid resource type")
	})

	t.Run("PeripheralsWithName", func(t *testing.T) {
		_, err := svc.CreateAttachment(context.Background(), "peripherals", file_uploads.ResourceIDTypeName, "somename", tmpFile, false)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "peripherals resource only supports ID type")
	})

	t.Run("EmptyIdentifier", func(t *testing.T) {
		_, err := svc.CreateAttachment(context.Background(), "policies", file_uploads.ResourceIDTypeID, "", tmpFile, false)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "identifier cannot be empty")
	})

	t.Run("EmptyFilePath", func(t *testing.T) {
		_, err := svc.CreateAttachment(context.Background(), "policies", file_uploads.ResourceIDTypeID, "1", "", false)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "file path cannot be empty")
	})
}
