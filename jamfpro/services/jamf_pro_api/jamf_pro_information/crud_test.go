package jamf_pro_information

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_information/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.JamfProInformationMock) {
	t.Helper()
	mock := mocks.NewJamfProInformationMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitGetV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.NotNil(t, result.VppTokenEnabled)
	require.True(t, *result.VppTokenEnabled)
}
