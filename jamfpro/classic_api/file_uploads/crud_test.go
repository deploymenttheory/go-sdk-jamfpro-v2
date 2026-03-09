package file_uploads_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/file_uploads"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/file_uploads/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_FileUploads_CreateAttachment(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	err := os.WriteFile(tmpFile, []byte("test content"), 0644)
	require.NoError(t, err)

	mockClient := mocks.NewFileUploadsMock()
	mockClient.RegisterCreateAttachmentMock()
	svc := file_uploads.NewFileUploads(mockClient)

	resp, err := svc.CreateAttachment(context.Background(), "policies", file_uploads.ResourceIDTypeID, "1", tmpFile, false)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_FileUploads_CreateAttachment_WithForceIpaUpload(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "app.ipa")
	err := os.WriteFile(tmpFile, []byte("fake ipa"), 0644)
	require.NoError(t, err)

	mockClient := mocks.NewFileUploadsMock()
	mockClient.RegisterCreateAttachmentMockForPath("/JSSResource/fileuploads/mobiledeviceapplicationsipa/id/1?FORCE_IPA_UPLOAD=true")
	svc := file_uploads.NewFileUploads(mockClient)

	resp, err := svc.CreateAttachment(context.Background(), "mobiledeviceapplicationsipa", file_uploads.ResourceIDTypeID, "1", tmpFile, true)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_FileUploads_CreateAttachment_ByName(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.pdf")
	err := os.WriteFile(tmpFile, []byte("test"), 0644)
	require.NoError(t, err)

	mockClient := mocks.NewFileUploadsMock()
	// Service constructs path without URL encoding: /JSSResource/fileuploads/policies/name/My Policy
	mockClient.RegisterCreateAttachmentMockForPath("/JSSResource/fileuploads/policies/name/My Policy")
	svc := file_uploads.NewFileUploads(mockClient)

	resp, err := svc.CreateAttachment(context.Background(), "policies", file_uploads.ResourceIDTypeName, "My Policy", tmpFile, false)

	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestUnit_FileUploads_CreateAttachment_InvalidResource(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(tmpFile, []byte("test"), 0644)

	mockClient := mocks.NewFileUploadsMock()
	svc := file_uploads.NewFileUploads(mockClient)

	_, err := svc.CreateAttachment(context.Background(), "invalidresource", file_uploads.ResourceIDTypeID, "1", tmpFile, false)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid resource type")
}

func TestUnit_FileUploads_CreateAttachment_InvalidIDType(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(tmpFile, []byte("test"), 0644)

	mockClient := mocks.NewFileUploadsMock()
	svc := file_uploads.NewFileUploads(mockClient)

	_, err := svc.CreateAttachment(context.Background(), "policies", file_uploads.ResourceIDType("invalid"), "1", tmpFile, false)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid ID type")
}

func TestUnit_FileUploads_CreateAttachment_PeripheralsWithName(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(tmpFile, []byte("test"), 0644)

	mockClient := mocks.NewFileUploadsMock()
	svc := file_uploads.NewFileUploads(mockClient)

	_, err := svc.CreateAttachment(context.Background(), "peripherals", file_uploads.ResourceIDTypeName, "somename", tmpFile, false)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "peripherals resource only supports ID type")
}

func TestUnit_FileUploads_CreateAttachment_EmptyIdentifier(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	_ = os.WriteFile(tmpFile, []byte("test"), 0644)

	mockClient := mocks.NewFileUploadsMock()
	svc := file_uploads.NewFileUploads(mockClient)

	_, err := svc.CreateAttachment(context.Background(), "policies", file_uploads.ResourceIDTypeID, "", tmpFile, false)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "identifier cannot be empty")
}

func TestUnit_FileUploads_CreateAttachment_EmptyFilePath(t *testing.T) {
	mockClient := mocks.NewFileUploadsMock()
	svc := file_uploads.NewFileUploads(mockClient)

	_, err := svc.CreateAttachment(context.Background(), "policies", file_uploads.ResourceIDTypeID, "1", "", false)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "file path cannot be empty")
}

func TestUnit_FileUploads_CreateAttachment_FileNotFound(t *testing.T) {
	mockClient := mocks.NewFileUploadsMock()
	svc := file_uploads.NewFileUploads(mockClient)

	_, err := svc.CreateAttachment(context.Background(), "policies", file_uploads.ResourceIDTypeID, "1", "/nonexistent/path/file.txt", false)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "open file")
}

func TestUnit_FileUploads_CreateAttachment_DirectoryNotFile(t *testing.T) {
	tmpDir := t.TempDir()

	mockClient := mocks.NewFileUploadsMock()
	svc := file_uploads.NewFileUploads(mockClient)

	_, err := svc.CreateAttachment(context.Background(), "policies", file_uploads.ResourceIDTypeID, "1", tmpDir, false)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "must point to a file")
}

func TestUnit_FileUploads_ValidResources(t *testing.T) {
	// Ensure all documented resources are present
	expected := []string{
		"computers", "mobiledevices", "enrollmentprofiles", "printers",
		"peripherals", "policies", "ebooks", "mobiledeviceapplications",
		"icon", "mobiledeviceapplicationsipa", "diskencryptionconfigurations",
	}
	assert.ElementsMatch(t, expected, file_uploads.ValidFileUploadResources)
}
