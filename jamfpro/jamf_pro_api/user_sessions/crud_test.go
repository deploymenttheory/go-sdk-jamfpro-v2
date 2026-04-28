package user_sessions

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/user_sessions/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*UserSessions, *mocks.UserSessionsMock) {
	t.Helper()
	mock := mocks.NewUserSessionsMock()
	return NewUserSessions(mock), mock
}

// Test GetActiveV1 success
func TestUnit_UserSessions_GetActiveV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetActiveSessionsMock()

	result, resp, err := svc.GetActiveV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
}

// Test GetCountV1 success
func TestUnit_UserSessions_GetCountV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCountMock()

	result, resp, err := svc.GetCountV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.Count)
}

// Error tests (no mock registered)
func TestUnit_UserSessions_GetActiveV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetActiveV1(context.Background())
	require.Error(t, err)
}

func TestUnit_UserSessions_GetCountV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetCountV1(context.Background())
	require.Error(t, err)
}
