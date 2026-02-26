package branding

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/branding/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.BrandingMock) {
	t.Helper()
	mock := mocks.NewBrandingMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_Branding_DownloadBrandingImageV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	body, resp, err := svc.DownloadBrandingImageV1(context.Background(), "test-id")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.NotNil(t, body)
	require.Greater(t, len(body), 0)
}
