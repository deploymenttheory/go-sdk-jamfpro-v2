package cloud_information

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_information/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.CloudInformationMock) {
	t.Helper()
	mock := mocks.NewCloudInformationMock()
	return NewService(mock), mock
}

func TestUnit_CloudInformation_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCloudInformationMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.CloudInstance)
	assert.False(t, result.RampInstance)
	assert.False(t, result.GovCloudInstance)
	assert.True(t, result.ManagedServiceProviderInstance)
}
