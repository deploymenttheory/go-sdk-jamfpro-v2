package icon

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/icon/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.IconsMock) {
	t.Helper()
	mock := mocks.NewIconsMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_Icon_GetByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 1, result.ID)
	require.Equal(t, "icon.png", result.Name)
}

func TestUnit_Icon_UploadV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	r := strings.NewReader("fake png bytes")
	result, resp, err := svc.UploadV1(context.Background(), r, 14, "test.png")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode)
	require.Equal(t, 2, result.ID)
	require.Equal(t, "uploaded.png", result.Name)
}

func TestUnit_Icon_DownloadV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	body, resp, err := svc.DownloadV1(context.Background(), 1, "original", "0")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.NotNil(t, body)
	require.Greater(t, len(body), 0)
}

func TestUnit_Icon_DownloadV1_Defaults(t *testing.T) {
	svc, _ := setupMockService(t)
	body, resp, err := svc.DownloadV1(context.Background(), 1, "", "")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.NotNil(t, body)
}

func TestUnit_Icons_GetByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), 999)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_Icons_UploadV1_Error(t *testing.T) {
	mock := mocks.NewIconsMock()
	svc := NewService(mock)
	r := strings.NewReader("fake png bytes")
	result, resp, err := svc.UploadV1(context.Background(), r, 14, "test.png")
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_Icons_UploadV1FromFile_FileNotFound(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UploadV1FromFile(context.Background(), "/nonexistent/path/icon.png")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "open icon file")
}

func TestUnit_Icons_UploadV1FromFile_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	f, err := os.CreateTemp(t.TempDir(), "icon-*.png")
	require.NoError(t, err)
	_, err = f.WriteString("fake png bytes")
	require.NoError(t, err)
	require.NoError(t, f.Close())
	result, resp, err := svc.UploadV1FromFile(context.Background(), f.Name())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode)
}

func TestUnit_Icons_UploadV1_DefaultFileName(t *testing.T) {
	svc, _ := setupMockService(t)
	r := strings.NewReader("fake png bytes")
	result, resp, err := svc.UploadV1(context.Background(), r, 14, "")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode)
}
