package last_login

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/last_login/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*LastLogin, *mocks.LastLoginMock) {
	t.Helper()
	mock := mocks.NewLastLoginMock()
	return NewLastLogin(mock), mock
}

// Test GetV1 success
func TestUnit_LastLogin_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLastLoginMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.LastLogin)
	assert.Equal(t, "2024-04-28T09:15:32Z", result.LastLogin)
}

// Test GetV1 error (no mock registered)
func TestUnit_LastLogin_GetV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetV1(context.Background())
	require.Error(t, err)
}
