package icons

import (
	"context"
	"strings"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/icons/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.IconsMock) {
	t.Helper()
	mock := mocks.NewIconsMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitGetByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 1, result.ID)
	require.Equal(t, "icon.png", result.Name)
}

func TestUnitUploadV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	r := strings.NewReader("fake png bytes")
	result, resp, err := svc.UploadV1(context.Background(), r, 14, "test.png")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode)
	require.Equal(t, 2, result.ID)
	require.Equal(t, "uploaded.png", result.Name)
}
