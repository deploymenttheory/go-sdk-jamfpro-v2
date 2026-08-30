package environment_type

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/environment_type/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*EnvironmentType, *mocks.EnvironmentTypeMock) {
	t.Helper()
	mock := mocks.NewEnvironmentTypeMock()
	mock.RegisterMocks()
	return NewEnvironmentType(mock), mock
}

func TestUnit_EnvironmentType_GetV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, EnvironmentProduction, result.Environment)
}

func TestUnit_EnvironmentType_GetV2_ClientError(t *testing.T) {
	mock := mocks.NewEnvironmentTypeMock()
	mock.RegisterGetErrorMock()
	svc := NewEnvironmentType(mock)

	result, resp, err := svc.GetV2(context.Background())

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mock client error")
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Nil(t, result)
}

func TestUnit_EnvironmentType_GetV2_NoMockRegistered(t *testing.T) {
	mock := mocks.NewEnvironmentTypeMock()
	mock.RegisterGetNoResponseErrorMock()
	svc := NewEnvironmentType(mock)

	result, resp, err := svc.GetV2(context.Background())

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
}

func TestUnit_EnvironmentType_NewService(t *testing.T) {
	mock := mocks.NewEnvironmentTypeMock()
	svc := NewEnvironmentType(mock)
	require.NotNil(t, svc)
	assert.NotNil(t, svc.client)
}
