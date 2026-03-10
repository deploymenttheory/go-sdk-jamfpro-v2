package jamf_pro_version

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_version/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*JamfProVersion, *mocks.JamfProVersionMock) {
	t.Helper()
	mock := mocks.NewJamfProVersionMock()
	mock.RegisterMocks()
	return NewJamfProVersion(mock), mock
}

func TestUnit_JamfProVersion_GetV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.NotNil(t, result.Version)
	assert.Equal(t, "11.0.0", *result.Version)
}

func TestUnit_JamfProVersion_GetV1_ClientError(t *testing.T) {
	mock := mocks.NewJamfProVersionMock()
	mock.RegisterGetErrorMock()
	svc := NewJamfProVersion(mock)

	result, resp, err := svc.GetV1(context.Background())

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mock client error")
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Nil(t, result)
}

func TestUnit_JamfProVersion_GetV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewJamfProVersionMock()
	mock.RegisterGetNoResponseErrorMock()
	svc := NewJamfProVersion(mock)

	result, resp, err := svc.GetV1(context.Background())

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

func TestUnit_JamfProVersion_NewService(t *testing.T) {
	mock := mocks.NewJamfProVersionMock()
	svc := NewJamfProVersion(mock)
	require.NotNil(t, svc)
}
