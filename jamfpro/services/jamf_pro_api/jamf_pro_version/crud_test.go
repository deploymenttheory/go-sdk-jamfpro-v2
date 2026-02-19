package jamf_pro_version

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_version/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.JamfProVersionMock) {
	t.Helper()
	mock := mocks.NewJamfProVersionMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitGetV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.NotNil(t, result.Version)
	require.Equal(t, "11.0.0", *result.Version)
}
