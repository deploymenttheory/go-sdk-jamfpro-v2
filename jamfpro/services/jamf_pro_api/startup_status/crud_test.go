package startup_status

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/startup_status/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.StartupStatusMock) {
	t.Helper()
	mock := mocks.NewStartupStatusMock()
	return NewService(mock), mock
}

func TestUnit_StartupStatus_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetStartupStatusMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Database", result.Step)
	assert.Equal(t, "DB_READY", result.StepCode)
	assert.Equal(t, 100, result.Percentage)
}

func TestUnit_StartupStatus_GetV1_Error_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)
	// Do not register any mock - client will return error

	result, resp, err := svc.GetV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	assert.Contains(t, err.Error(), "StartupStatusMock: no response")
}
