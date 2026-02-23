package self_service_branding_upload

import (
	"context"
	"strings"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_branding_upload/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SelfServiceBrandingUploadMock) {
	t.Helper()
	mock := mocks.NewSelfServiceBrandingUploadMock()
	mock.RegisterUploadMock()
	return NewService(mock), mock
}

func TestUnitUpload_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	r := strings.NewReader("fake png bytes")
	result, resp, err := svc.Upload(context.Background(), r, 14, "branding.png")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode)
	require.NotEmpty(t, result.URL)
	require.Contains(t, result.URL, "uploaded-branding.png")
}

func TestUnitUpload_DefaultFileName(t *testing.T) {
	svc, _ := setupMockService(t)
	r := strings.NewReader("fake png bytes")
	result, resp, err := svc.Upload(context.Background(), r, 14, "")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.NotEmpty(t, result.URL)
}

func TestUnitUploadFromFile_FileNotFound(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UploadFromFile(context.Background(), "/nonexistent/path/branding.png")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "open branding image file")
}
