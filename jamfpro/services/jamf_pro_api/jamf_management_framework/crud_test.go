package jamf_management_framework

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_management_framework/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.JamfManagementFrameworkMock) {
	t.Helper()
	mock := mocks.NewJamfManagementFrameworkMock()
	return NewService(mock), mock
}

func TestUnit_JamfManagementFramework_NewService(t *testing.T) {
	mock := mocks.NewJamfManagementFrameworkMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
}

func TestUnit_JamfManagementFramework_RedeployV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterRedeployMock("123")

	result, resp, err := svc.RedeployV1(context.Background(), "123")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "123", result.DeviceID)
	assert.Equal(t, "abc-123-uuid", result.CommandUUID)
}

func TestUnit_JamfManagementFramework_RedeployV1_Success_WithWhitespaceTrimmed(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterRedeployMock("123")

	result, resp, err := svc.RedeployV1(context.Background(), "  123  ")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "123", result.DeviceID)
	assert.Equal(t, "abc-123-uuid", result.CommandUUID)
}

func TestUnit_JamfManagementFramework_RedeployV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.RedeployV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "computer ID is required")
}

func TestUnit_JamfManagementFramework_RedeployV1_WhitespaceOnlyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.RedeployV1(context.Background(), "   ")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "computer ID is required")
}

func TestUnit_JamfManagementFramework_RedeployV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock("999")

	result, resp, err := svc.RedeployV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Contains(t, err.Error(), "failed to redeploy jamf management framework")
}

func TestUnit_JamfManagementFramework_RedeployV1_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.RedeployV1(context.Background(), "123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to redeploy jamf management framework")
}
