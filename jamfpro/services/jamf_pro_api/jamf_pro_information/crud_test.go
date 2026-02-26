package jamf_pro_information

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_information/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.JamfProInformationMock) {
	t.Helper()
	mock := mocks.NewJamfProInformationMock()
	return NewService(mock), mock
}

func TestUnit_JamfProInformation_GetV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterMocks()

	result, resp, err := svc.GetV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	require.NotNil(t, result.VppTokenEnabled)
	require.True(t, *result.VppTokenEnabled)
}

func TestUnit_JamfProInformation_GetV2_ClientError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV2ErrorMock()

	result, resp, err := svc.GetV2(context.Background())

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mock client error")
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
	assert.Nil(t, result)
}

func TestUnit_JamfProInformation_GetV2_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered - dispatch returns (nil, err)

	result, resp, err := svc.GetV2(context.Background())

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no response for")
	assert.Nil(t, resp)
	assert.Nil(t, result)
}
